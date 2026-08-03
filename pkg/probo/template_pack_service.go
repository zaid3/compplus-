// Copyright (c) 2026 CompPlus.
// Use of this source code is governed by the MIT license in the repository root.

package probo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go.gearno.de/kit/pg"
	compplustemplates "go.probo.inc/probo/compplus/templates"
	"go.probo.inc/probo/pkg/coredata"
	"go.probo.inc/probo/pkg/gid"
)

type (
	TemplatePackService struct {
		svc *Service
	}

	CompileTemplatePackRequest struct {
		PackID  string
		Answers map[string]any
		Now     time.Time
	}

	InstallTemplatePackRequest struct {
		OrganizationID gid.GID
		PackID         string
		Answers        map[string]any
		Now            time.Time
	}

	InstallTemplatePackResult struct {
		PackID    string
		Version   string
		Framework *coredata.Framework
		Measures  coredata.Measures
		Documents coredata.Documents
	}

	ErrTemplatePackAlreadyInstalled struct {
		PackID         string
		OrganizationID gid.GID
	}
)

func (e *ErrTemplatePackAlreadyInstalled) Error() string {
	return fmt.Sprintf("template pack %q is already installed for organization %q", e.PackID, e.OrganizationID)
}

// Compile prepares a complete pack without changing organization data. It is
// used by the setup wizard to preview exactly what will be created.
func (s *TemplatePackService) Compile(
	_ context.Context,
	req CompileTemplatePackRequest,
) (*compplustemplates.CompiledPack, error) {
	compiled, err := compplustemplates.Compile(req.PackID, compplustemplates.CompileOptions{
		Answers: req.Answers,
		Now:     req.Now,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot compile template pack: %w", err)
	}

	return compiled, nil
}

// Install creates the framework, controls, measures, tasks, evidence requests
// and editable draft documents for one organization. It reuses the existing
// application services so normal validation and tenant scoping remain in force.
func (s *TemplatePackService) Install(
	ctx context.Context,
	scope coredata.Scoper,
	req InstallTemplatePackRequest,
) (*InstallTemplatePackResult, error) {
	compiled, err := s.Compile(ctx, CompileTemplatePackRequest{
		PackID:  req.PackID,
		Answers: req.Answers,
		Now:     req.Now,
	})
	if err != nil {
		return nil, err
	}

	alreadyInstalled, err := s.isInstalled(ctx, scope, req.OrganizationID, compiled.Framework.ID)
	if err != nil {
		return nil, fmt.Errorf("cannot check template pack installation: %w", err)
	}
	if alreadyInstalled {
		return nil, &ErrTemplatePackAlreadyInstalled{
			PackID:         req.PackID,
			OrganizationID: req.OrganizationID,
		}
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
			return nil, fmt.Errorf(
				"cannot install document %q: measure %q was not imported",
				compiledDocument.Title,
				compiledDocument.MeasureReferenceID,
			)
		}

		document, _, err := s.svc.Documents.Create(ctx, scope, CreateDocumentRequest{
			OrganizationID: req.OrganizationID,
			Title:          compiledDocument.Title,
			Content:        compiledDocument.Content,
			Classification: coredata.DocumentClassification(compiledDocument.Classification),
			DocumentType:   coredata.DocumentType(compiledDocument.DocumentType),
		})
		if err != nil {
			return nil, fmt.Errorf("cannot create template document %q: %w", compiledDocument.Title, err)
		}

		if _, _, err := s.svc.Measures.CreateDocumentMapping(ctx, scope, measure.ID, document.ID); err != nil {
			return nil, fmt.Errorf(
				"cannot link template document %q to measure %q: %w",
				compiledDocument.Title,
				measure.Name,
				err,
			)
		}

		documents = append(documents, document)
	}

	return &InstallTemplatePackResult{
		PackID:    compiled.PackID,
		Version:   compiled.Version,
		Framework: framework,
		Measures:  measurePage.Data,
		Documents: documents,
	}, nil
}

func (s *TemplatePackService) isInstalled(
	ctx context.Context,
	scope coredata.Scoper,
	organizationID gid.GID,
	frameworkReferenceID string,
) (bool, error) {
	framework := &coredata.Framework{}
	err := s.svc.pg.WithConn(ctx, func(ctx context.Context, conn pg.Querier) error {
		return framework.LoadByReferenceID(ctx, conn, scope, frameworkReferenceID)
	})
	if errors.Is(err, coredata.ErrResourceNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return framework.OrganizationID == organizationID, nil
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