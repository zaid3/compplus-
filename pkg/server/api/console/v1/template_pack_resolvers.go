// Copyright (c) 2026 CompPlus.
// Use of this source code is governed by the MIT license in the repository root.

package console_v1

import (
	"context"
	"errors"
	"strings"

	"go.gearno.de/kit/log"
	"go.probo.inc/probo/pkg/probo"
	"go.probo.inc/probo/pkg/server/api/console/v1/types"
	"go.probo.inc/probo/pkg/server/gqlutils"
)

// PreviewTemplatePack is the resolver for the previewTemplatePack field.
func (r *mutationResolver) PreviewTemplatePack(
	ctx context.Context,
	input types.PreviewTemplatePackInput,
) (*types.PreviewTemplatePackPayload, error) {
	if _, err := r.authorize(ctx, input.OrganizationID, probo.ActionFrameworkList); err != nil {
		return nil, err
	}

	compiled, err := probo.NewTemplatePackService(r.probo).Compile(
		ctx,
		probo.CompileTemplatePackRequest{
			PackID:  input.PackID,
			Answers: input.Answers.Values(),
		},
	)
	if err != nil {
		return nil, gqlutils.Invalid(ctx, err)
	}

	tasksCount, evidenceRequestsCount := compiledWorkCounts(compiled)
	confirmationFieldsCount := 0
	for _, document := range compiled.Documents {
		confirmationFieldsCount += strings.Count(document.Content, "CONFIRM")
		confirmationFieldsCount += strings.Count(document.Content, "[ADD:")
	}

	return &types.PreviewTemplatePackPayload{
		Preview: &types.TemplatePackPreview{
			PackID:                  compiled.PackID,
			Version:                 compiled.Version,
			FrameworkName:           compiled.Framework.Name,
			ControlsCount:           len(compiled.Framework.Controls),
			MeasuresCount:           len(compiled.Measures),
			DocumentsCount:          len(compiled.Documents),
			TasksCount:              tasksCount,
			EvidenceRequestsCount:   evidenceRequestsCount,
			ConfirmationFieldsCount: confirmationFieldsCount,
		},
	}, nil
}

// InstallTemplatePack is the resolver for the installTemplatePack field.
func (r *mutationResolver) InstallTemplatePack(
	ctx context.Context,
	input types.InstallTemplatePackInput,
) (*types.InstallTemplatePackPayload, error) {
	scope, err := r.authorize(ctx, input.OrganizationID, probo.ActionFrameworkImport)
	if err != nil {
		return nil, err
	}

	for _, action := range []string{
		probo.ActionMeasureImport,
		probo.ActionDocumentCreate,
		probo.ActionMeasureDocumentMappingCreate,
	} {
		if _, err := r.authorize(ctx, input.OrganizationID, action); err != nil {
			return nil, err
		}
	}

	templatePackService := probo.NewTemplatePackService(r.probo)
	compiled, err := templatePackService.Compile(
		ctx,
		probo.CompileTemplatePackRequest{
			PackID:  input.PackID,
			Answers: input.Answers.Values(),
		},
	)
	if err != nil {
		return nil, gqlutils.Invalid(ctx, err)
	}

	result, err := templatePackService.Install(
		ctx,
		scope,
		probo.InstallTemplatePackRequest{
			OrganizationID: input.OrganizationID,
			PackID:         input.PackID,
			Answers:        input.Answers.Values(),
		},
	)
	if err != nil {
		if _, ok := errors.AsType[*probo.ErrTemplatePackAlreadyInstalled](err); ok {
			return nil, gqlutils.Conflict(ctx, err)
		}

		r.logger.ErrorCtx(
			ctx,
			"cannot install CompPlus template pack",
			log.Error(err),
			log.String("pack_id", input.PackID),
			log.String("organization_id", input.OrganizationID.String()),
		)

		return nil, gqlutils.Internal(ctx)
	}

	tasksCount, evidenceRequestsCount := compiledWorkCounts(compiled)

	return &types.InstallTemplatePackPayload{
		PackID:                  result.PackID,
		Version:                 result.Version,
		Framework:               types.NewFramework(result.Framework),
		MeasuresCreated:         len(result.Measures),
		DocumentsCreated:        len(result.Documents),
		TasksCreated:            tasksCount,
		EvidenceRequestsCreated: evidenceRequestsCount,
	}, nil
}

func compiledWorkCounts(compiled interface {
	GetMeasures() []proboTemplateMeasure
}) (int, int) {
	return 0, 0
}
