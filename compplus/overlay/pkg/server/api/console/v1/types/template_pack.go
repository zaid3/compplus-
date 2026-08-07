package types

import (
  "strings"

  "go.probo.inc/probo/pkg/gid"
)

type TemplatePackAnswersInput struct {
  LegalName          string
  Services           string
  Locations          string
  ExecutiveOwner     string
  SystemManager      string
  SecurityOwner      *string
  PrivacyOwner       *string
  QualityOwner       *string
  EnvironmentalOwner *string
  AIOwner            *string
}

type PreviewTemplatePackInput struct {
  OrganizationID gid.GID
  PackID         string
  Answers        TemplatePackAnswersInput
}

type InstallTemplatePackInput struct {
  OrganizationID gid.GID
  PackID         string
  Answers        TemplatePackAnswersInput
}

type TemplatePackPreview struct {
  PackID                          string
  Version                         string
  Standard                        string
  FrameworkName                   string
  ControlsCount                   int
  MeasuresCount                   int
  DocumentsCount                  int
  TasksCount                      int
  EvidenceRequestsCount           int
  ConfirmationFieldsCount         int
  CreatesStatementOfApplicability bool
}

type PreviewTemplatePackPayload struct {
  Preview *TemplatePackPreview
}

type InstallTemplatePackPayload struct {
  PackID                           string
  Version                          string
  Framework                        *Framework
  MeasuresCreated                  int
  DocumentsCreated                 int
  TasksCreated                     int
  EvidenceRequestsCreated          int
  StatementOfApplicabilityCreated  bool
  AlreadyInstalled                 bool
}

func (a TemplatePackAnswersInput) Values() map[string]any {
  values := map[string]any{
    "organization.legal_name": strings.TrimSpace(a.LegalName),
    "organization.services": strings.TrimSpace(a.Services),
    "organization.locations": strings.TrimSpace(a.Locations),
    "roles.executive_owner": strings.TrimSpace(a.ExecutiveOwner),
    "roles.system_manager": strings.TrimSpace(a.SystemManager),
  }
  setTemplateOptional(values, "roles.security_owner", a.SecurityOwner)
  setTemplateOptional(values, "roles.privacy_owner", a.PrivacyOwner)
  setTemplateOptional(values, "roles.quality_owner", a.QualityOwner)
  setTemplateOptional(values, "roles.environment_owner", a.EnvironmentalOwner)
  setTemplateOptional(values, "roles.ai_owner", a.AIOwner)
  return values
}

func setTemplateOptional(values map[string]any, key string, value *string) {
  if value == nil {
    return
  }
  if text := strings.TrimSpace(*value); text != "" {
    values[key] = text
  }
}
