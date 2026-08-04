// Copyright (c) 2026 CompPlus.
// Use of this source code is governed by the MIT license in the repository root.

package types

import (
	"strings"

	"go.probo.inc/probo/pkg/gid"
)

type (
	TemplatePackAnswersInput struct {
		LegalName             string
		TradingName           *string
		Address               string
		Services              string
		Locations             string
		EmployeeCount         int
		ExecutiveOwner        string
		SystemManager         string
		SecurityOwner         *string
		PrivacyOwner          *string
		QualityOwner          *string
		EnvironmentalOwner    *string
		AIOwner               *string
		CertificationTarget   *string
		ReviewMonth           string
		UsesSuppliers         bool
		UsesAI                bool
		ProcessesPersonalData bool
		EnvironmentalImpacts  bool
		SelectedStandards     []string
	}

	PreviewTemplatePackInput struct {
		OrganizationID gid.GID
		PackID         string
		Answers        TemplatePackAnswersInput
	}

	InstallTemplatePackInput struct {
		OrganizationID gid.GID
		PackID         string
		Answers        TemplatePackAnswersInput
	}

	TemplatePackPreview struct {
		PackID                  string
		Version                 string
		FrameworkName           string
		ControlsCount           int
		MeasuresCount           int
		DocumentsCount          int
		TasksCount              int
		EvidenceRequestsCount   int
		ConfirmationFieldsCount int
	}

	PreviewTemplatePackPayload struct {
		Preview *TemplatePackPreview
	}

	InstallTemplatePackPayload struct {
		PackID                  string
		Version                 string
		Framework               *Framework
		MeasuresCreated         int
		DocumentsCreated        int
		TasksCreated            int
		EvidenceRequestsCreated int
	}
)

func (a TemplatePackAnswersInput) Values() map[string]any {
	values := map[string]any{
		"organization.legal_name":             strings.TrimSpace(a.LegalName),
		"organization.address":                strings.TrimSpace(a.Address),
		"organization.services":               strings.TrimSpace(a.Services),
		"organization.locations":               strings.TrimSpace(a.Locations),
		"organization.employee_count":          a.EmployeeCount,
		"organization.review_month":            strings.TrimSpace(a.ReviewMonth),
		"organization.uses_suppliers":           a.UsesSuppliers,
		"organization.uses_ai":                  a.UsesAI,
		"organization.processes_personal_data": a.ProcessesPersonalData,
		"organization.environmental_impacts":   a.EnvironmentalImpacts,
		"organization.selected_standards":      strings.Join(a.SelectedStandards, ", "),
		"roles.executive_owner":                strings.TrimSpace(a.ExecutiveOwner),
		"roles.system_manager":                 strings.TrimSpace(a.SystemManager),
	}

	setOptional(values, "organization.trading_name", a.TradingName)
	setOptional(values, "roles.security_owner", a.SecurityOwner)
	setOptional(values, "roles.privacy_owner", a.PrivacyOwner)
	setOptional(values, "roles.quality_owner", a.QualityOwner)
	setOptional(values, "roles.environment_owner", a.EnvironmentalOwner)
	setOptional(values, "roles.ai_owner", a.AIOwner)
	setOptional(values, "organization.certification_target", a.CertificationTarget)

	return values
}

func setOptional(values map[string]any, key string, value *string) {
	if value == nil {
		return
	}

	trimmed := strings.TrimSpace(*value)
	if trimmed != "" {
		values[key] = trimmed
	}
}
