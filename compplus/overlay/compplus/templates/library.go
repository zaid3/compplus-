package templates

import "strings"

func allPacks() []Pack {
  return []Pack{
    corePack(),
    iso27001Pack(),
    iso9001Pack(),
    ukGDPRPack(),
    iso14001Pack(),
    iso42001Pack(),
  }
}

func makePack(id, name, version, standard, description string, templates []TemplateDefinition, createSOA bool, extraControls []ControlDefinition) Pack {
  seen := map[string]bool{}
  controls := make([]ControlDefinition, 0, len(templates)+len(extraControls))
  for _, tmpl := range templates {
    if !seen[tmpl.ControlID] {
      controls = append(controls, ControlDefinition{ID: tmpl.ControlID, Name: tmpl.Title, Description: tmpl.Purpose})
      seen[tmpl.ControlID] = true
    }
  }
  for _, control := range extraControls {
    if !seen[control.ID] {
      controls = append(controls, control)
      seen[control.ID] = true
    }
  }
  return Pack{
    ID: id,
    Name: name,
    Version: version,
    Standard: standard,
    Description: description,
    CreateSOA: createSOA,
    Framework: FrameworkDefinition{
      ID: "COMPPLUS-" + strings.ToUpper(strings.ReplaceAll(id, "-", "_")) + "-2026",
      Name: name,
      Description: description,
      Controls: controls,
    },
    Templates: templates,
  }
}

func tmpl(id, controlID, title, kind, category, documentType, ownerKey, purpose string, requirements, evidence, references, confirmations, starter []string) TemplateDefinition {
  return TemplateDefinition{
    ID: id,
    ControlID: controlID,
    Title: title,
    Kind: kind,
    Category: category,
    DocumentType: documentType,
    Classification: "INTERNAL",
    OwnerKey: ownerKey,
    Purpose: purpose,
    Requirements: requirements,
    Evidence: evidence,
    References: references,
    Confirmations: confirmations,
    StarterContent: starter,
  }
}

func confidential(t TemplateDefinition) TemplateDefinition {
  t.Classification = "CONFIDENTIAL"
  return t
}
