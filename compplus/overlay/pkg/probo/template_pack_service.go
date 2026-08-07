package probo

import (
  "context"
  "encoding/json"
  "errors"
  "fmt"
  "strings"
  "time"

  "go.gearno.de/kit/pg"
  compplustemplates "go.probo.inc/probo/compplus/templates"
  "go.probo.inc/probo/pkg/coredata"
  "go.probo.inc/probo/pkg/gid"
  "go.probo.inc/probo/pkg/page"
)

type TemplatePackService struct {
  svc *Service
}

type CompileTemplatePackRequest struct {
  PackID  string
  Answers map[string]any
  Now     time.Time
}

type InstallTemplatePackRequest struct {
  OrganizationID gid.GID
  PackID         string
  Answers        map[string]any
  Now            time.Time
}

type InstallTemplatePackResult struct {
  PackID                   string
  Version                  string
  Framework                *coredata.Framework
  Measures                 coredata.Measures
  Documents                coredata.Documents
  StatementOfApplicability *coredata.StatementOfApplicability
  AlreadyInstalled         bool
}

func NewTemplatePackService(svc *Service) *TemplatePackService {
  return &TemplatePackService{svc: svc}
}

func (s *TemplatePackService) Compile(_ context.Context, req CompileTemplatePackRequest) (*compplustemplates.CompiledPack, error) {
  compiled, err := compplustemplates.Compile(req.PackID, compplustemplates.CompileOptions{
    Answers: req.Answers,
    Now: req.Now,
  })
  if err != nil {
    return nil, fmt.Errorf("cannot compile template pack: %w", err)
  }
  return compiled, nil
}

// Install is intentionally idempotent at pack level: a framework reference ID
// uniquely identifies the installed Comp Plus+ pack in an organisation. Repeated
// clicks return the existing framework rather than creating duplicate documents.
func (s *TemplatePackService) Install(ctx context.Context, scope coredata.Scoper, req InstallTemplatePackRequest) (*InstallTemplatePackResult, error) {
  compiled, err := s.Compile(ctx, CompileTemplatePackRequest{PackID: req.PackID, Answers: req.Answers, Now: req.Now})
  if err != nil {
    return nil, err
  }

  existing, err := s.findInstalledFramework(ctx, scope, req.OrganizationID, compiled.Framework.ID)
  if err != nil {
    return nil, fmt.Errorf("cannot check template pack installation: %w", err)
  }
  if existing != nil {
    return &InstallTemplatePackResult{
      PackID: compiled.PackID,
      Version: compiled.Version,
      Framework: existing,
      AlreadyInstalled: true,
    }, nil
  }

  frameworkRequest, err := frameworkImportRequest(compiled.Framework)
  if err != nil {
    return nil, fmt.Errorf("cannot prepare framework import: %w", err)
  }
  framework, err := s.svc.Frameworks.Import(ctx, scope, req.OrganizationID, frameworkRequest)
  if err != nil {
    return nil, fmt.Errorf("cannot install template pack framework: %w", err)
  }

  measureRequest, err := measureImportRequest(compiled.Measures)
  if err != nil {
    return nil, fmt.Errorf("cannot prepare measure import: %w", err)
  }
  measurePage, err := s.svc.Measures.Import(ctx, scope, req.OrganizationID, measureRequest)
  if err != nil {
    return nil, fmt.Errorf("cannot install template pack measures: %w", err)
  }

  measureByReference := make(map[string]*coredata.Measure, len(measurePage.Data))
  for _, measure := range measurePage.Data {
    measureByReference[measure.ReferenceID] = measure
  }

  documents := make(coredata.Documents, 0, len(compiled.Documents))
  for _, compiledDocument := range compiled.Documents {
    measure, ok := measureByReference[compiledDocument.MeasureReferenceID]
    if !ok {
      return nil, fmt.Errorf("cannot install document %q: measure %q was not imported", compiledDocument.Title, compiledDocument.MeasureReferenceID)
    }

    document, _, err := s.svc.Documents.Create(ctx, scope, CreateDocumentRequest{
      OrganizationID: req.OrganizationID,
      Title: compiledDocument.Title,
      Content: compiledDocument.Content,
      Classification: coredata.DocumentClassification(compiledDocument.Classification),
      DocumentType: coredata.DocumentType(compiledDocument.DocumentType),
    })
    if err != nil {
      return nil, fmt.Errorf("cannot create template document %q: %w", compiledDocument.Title, err)
    }

    if _, _, err := s.svc.Measures.CreateDocumentMapping(ctx, scope, measure.ID, document.ID); err != nil {
      return nil, fmt.Errorf("cannot link template document %q to measure %q: %w", compiledDocument.Title, measure.Name, err)
    }
    documents = append(documents, document)
  }

  var soa *coredata.StatementOfApplicability
  if compiled.CreateSOA {
    soa, err = s.installISO27001SOA(ctx, scope, req.OrganizationID, framework)
    if err != nil {
      return nil, fmt.Errorf("cannot create ISO 27001 Statement of Applicability starter: %w", err)
    }
  }

  return &InstallTemplatePackResult{
    PackID: compiled.PackID,
    Version: compiled.Version,
    Framework: framework,
    Measures: measurePage.Data,
    Documents: documents,
    StatementOfApplicability: soa,
  }, nil
}

func (s *TemplatePackService) installISO27001SOA(ctx context.Context, scope coredata.Scoper, organizationID gid.GID, framework *coredata.Framework) (*coredata.StatementOfApplicability, error) {
  soa, err := s.svc.StatementsOfApplicability.Create(ctx, scope, CreateStatementOfApplicabilityRequest{
    OrganizationID: organizationID,
    Name: "Comp Plus+ ISO/IEC 27001 Statement of Applicability",
  })
  if err != nil {
    return nil, err
  }

  controlsPage, err := s.svc.Controls.ListForFrameworkID(
    ctx,
    scope,
    framework.ID,
    page.NewCursor(
      500,
      nil,
      page.Head,
      page.OrderBy[coredata.ControlOrderField]{
        Field: coredata.ControlOrderFieldSectionTitle,
        Direction: page.OrderDirectionAsc,
      },
    ),
    coredata.NewControlFilter(nil),
  )
  if err != nil {
    return nil, err
  }

  starterJustification := "Starter status: provisionally applicable pending review against the organisation's information-security risks, legal/regulatory/contractual obligations and business context. Confirm final applicability, rationale and implementation evidence before approving the SoA."
  for _, control := range controlsPage.Data {
    if !strings.HasPrefix(control.SectionTitle, "A.") {
      continue
    }
    if _, err := s.svc.StatementsOfApplicability.CreateApplicabilityStatement(
      ctx,
      scope,
      soa.ID,
      control.ID,
      true,
      &starterJustification,
    ); err != nil {
      return nil, fmt.Errorf("cannot add SoA control %s: %w", control.SectionTitle, err)
    }
  }

  return soa, nil
}

func (s *TemplatePackService) findInstalledFramework(ctx context.Context, scope coredata.Scoper, organizationID gid.GID, referenceID string) (*coredata.Framework, error) {
  framework := &coredata.Framework{}
  err := s.svc.pg.WithConn(ctx, func(ctx context.Context, conn pg.Querier) error {
    return framework.LoadByReferenceID(ctx, conn, scope, referenceID)
  })
  if errors.Is(err, coredata.ErrResourceNotFound) {
    return nil, nil
  }
  if err != nil {
    return nil, err
  }
  if framework.OrganizationID != organizationID {
    return nil, nil
  }
  return framework, nil
}

func frameworkImportRequest(definition compplustemplates.FrameworkDefinition) (ImportFrameworkRequest, error) {
  payload, err := json.Marshal(definition)
  if err != nil {
    return ImportFrameworkRequest{}, err
  }
  request := ImportFrameworkRequest{}
  if err := json.Unmarshal(payload, &request.Framework); err != nil {
    return ImportFrameworkRequest{}, err
  }
  return request, nil
}

func measureImportRequest(measures []compplustemplates.CompiledMeasure) (ImportMeasureRequest, error) {
  payload, err := json.Marshal(struct {
    Measures []compplustemplates.CompiledMeasure `json:"measures"`
  }{Measures: measures})
  if err != nil {
    return ImportMeasureRequest{}, err
  }
  request := ImportMeasureRequest{}
  if err := json.Unmarshal(payload, &request); err != nil {
    return ImportMeasureRequest{}, err
  }
  for _, measure := range request.Measures {
    for _, task := range measure.Tasks {
      for _, evidence := range task.RequestedEvidences {
        if !evidence.Type.IsValid() {
          return ImportMeasureRequest{}, fmt.Errorf("invalid evidence type %q", evidence.Type)
        }
      }
    }
  }
  return request, nil
}
