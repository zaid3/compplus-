package console_v1

import (
  "context"
  "strings"

  "go.gearno.de/kit/log"
  "go.probo.inc/probo/pkg/iam"
  "go.probo.inc/probo/pkg/probo"
  "go.probo.inc/probo/pkg/server/api/console/v1/types"
  "go.probo.inc/probo/pkg/server/gqlutils"
)

func (r *mutationResolver) PreviewTemplatePack(ctx context.Context, input types.PreviewTemplatePackInput) (*types.PreviewTemplatePackPayload, error) {
  if _, err := r.authorize(ctx, input.OrganizationID, probo.ActionDocumentList); err != nil {
    return nil, err
  }

  compiled, err := probo.NewTemplatePackService(r.probo).Compile(ctx, probo.CompileTemplatePackRequest{
    PackID: input.PackID,
    Answers: input.Answers.Values(),
  })
  if err != nil {
    return nil, gqlutils.Invalid(ctx, err)
  }

  tasks, evidenceRequests := compiledTemplateWorkCounts(compiled)
  confirmations := 0
  for _, document := range compiled.Documents {
    confirmations += strings.Count(document.Content, "CONFIRM")
  }

  return &types.PreviewTemplatePackPayload{
    Preview: &types.TemplatePackPreview{
      PackID: compiled.PackID,
      Version: compiled.Version,
      Standard: compiled.Standard,
      FrameworkName: compiled.Framework.Name,
      ControlsCount: len(compiled.Framework.Controls),
      MeasuresCount: len(compiled.Measures),
      DocumentsCount: len(compiled.Documents),
      TasksCount: tasks,
      EvidenceRequestsCount: evidenceRequests,
      ConfirmationFieldsCount: confirmations,
      CreatesStatementOfApplicability: compiled.CreateSOA,
    },
  }, nil
}

func (r *mutationResolver) InstallTemplatePack(ctx context.Context, input types.InstallTemplatePackInput) (*types.InstallTemplatePackPayload, error) {
  scope, err := r.authorize(ctx, input.OrganizationID, probo.ActionFrameworkImport)
  if err != nil {
    return nil, err
  }

  actions := []iam.Action{
    probo.ActionMeasureImport,
    probo.ActionDocumentCreate,
    probo.ActionMeasureDocumentMappingCreate,
  }
  if input.PackID == "iso27001" {
    actions = append(actions,
      probo.ActionStatementOfApplicabilityCreate,
      probo.ActionApplicabilityStatementCreate,
      probo.ActionControlList,
    )
  }
  for _, action := range actions {
    if _, err := r.authorize(ctx, input.OrganizationID, action); err != nil {
      return nil, err
    }
  }

  templateService := probo.NewTemplatePackService(r.probo)
  compiled, err := templateService.Compile(ctx, probo.CompileTemplatePackRequest{
    PackID: input.PackID,
    Answers: input.Answers.Values(),
  })
  if err != nil {
    return nil, gqlutils.Invalid(ctx, err)
  }

  result, err := templateService.Install(ctx, scope, probo.InstallTemplatePackRequest{
    OrganizationID: input.OrganizationID,
    PackID: input.PackID,
    Answers: input.Answers.Values(),
  })
  if err != nil {
    r.logger.ErrorCtx(ctx, "cannot install Comp Plus+ template pack", log.Error(err), log.String("pack_id", input.PackID), log.String("organization_id", input.OrganizationID.String()))
    return nil, gqlutils.Internal(ctx)
  }

  tasks, evidenceRequests := compiledTemplateWorkCounts(compiled)
  if result.AlreadyInstalled {
    tasks = 0
    evidenceRequests = 0
  }

  return &types.InstallTemplatePackPayload{
    PackID: result.PackID,
    Version: result.Version,
    Framework: types.NewFramework(result.Framework),
    MeasuresCreated: len(result.Measures),
    DocumentsCreated: len(result.Documents),
    TasksCreated: tasks,
    EvidenceRequestsCreated: evidenceRequests,
    StatementOfApplicabilityCreated: result.StatementOfApplicability != nil,
    AlreadyInstalled: result.AlreadyInstalled,
  }, nil
}
