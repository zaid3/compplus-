package templates

import (
  "encoding/json"
  "fmt"
  "sort"
  "strings"
  "time"

  "go.probo.inc/probo/pkg/prosemirror"
)

type Pack struct {
  ID          string
  Name        string
  Version     string
  Standard    string
  Description string
  Framework   FrameworkDefinition
  CreateSOA   bool
  Templates   []TemplateDefinition
}

type FrameworkDefinition struct {
  ID          string              `json:"id"`
  Name        string              `json:"name"`
  Description string              `json:"description,omitempty"`
  Controls    []ControlDefinition `json:"controls"`
}

type ControlDefinition struct {
  ID          string `json:"id"`
  Name        string `json:"name"`
  Description string `json:"description"`
}

type TemplateDefinition struct {
  ID             string
  ControlID      string
  Title          string
  Kind           string
  Category       string
  DocumentType   string
  Classification string
  OwnerKey       string
  Purpose        string
  Requirements   []string
  Evidence       []string
  References     []string
  Confirmations  []string
  StarterContent []string
}

type CompileOptions struct {
  Answers map[string]any
  Now     time.Time
}

type CompiledPack struct {
  PackID    string              `json:"pack_id"`
  Version   string              `json:"version"`
  Standard  string              `json:"standard"`
  CreateSOA bool                `json:"create_soa"`
  Framework FrameworkDefinition `json:"framework"`
  Measures  []CompiledMeasure   `json:"measures"`
  Documents []CompiledDocument  `json:"documents"`
}

type CompiledMeasure struct {
  Name        string                    `json:"name"`
  Description string                    `json:"description"`
  Category    string                    `json:"category"`
  ReferenceID string                    `json:"reference-id"`
  Standards   []CompiledStandardMapping `json:"standards"`
  Tasks       []CompiledTask            `json:"tasks"`
}

type CompiledStandardMapping struct {
  Framework string `json:"framework"`
  Control   string `json:"control"`
}

type CompiledTask struct {
  Name               string             `json:"name"`
  Description        string             `json:"description"`
  ReferenceID        string             `json:"reference-id"`
  RequestedEvidences []CompiledEvidence `json:"requested-evidences"`
}

type CompiledEvidence struct {
  ReferenceID string `json:"reference-id"`
  Type        string `json:"type"`
  Name        string `json:"name"`
}

type CompiledDocument struct {
  ReferenceID        string `json:"reference_id"`
  MeasureReferenceID string `json:"measure_reference_id"`
  Title              string `json:"title"`
  Content            string `json:"content"`
  DocumentType       string `json:"document_type"`
  Classification     string `json:"classification"`
}

func AvailablePacks() []Pack {
  packs := allPacks()
  sort.Slice(packs, func(i, j int) bool { return packs[i].ID < packs[j].ID })
  return packs
}

func GetPack(id string) (Pack, bool) {
  for _, pack := range allPacks() {
    if pack.ID == id {
      return pack, true
    }
  }
  return Pack{}, false
}

func Compile(packID string, options CompileOptions) (*CompiledPack, error) {
  pack, ok := GetPack(packID)
  if !ok {
    return nil, fmt.Errorf("unknown Comp Plus+ template pack %q", packID)
  }
  if err := validatePack(pack); err != nil {
    return nil, err
  }

  now := options.Now
  if now.IsZero() {
    now = time.Now()
  }
  answers := withDefaults(options.Answers, now)

  compiled := &CompiledPack{
    PackID: pack.ID,
    Version: pack.Version,
    Standard: pack.Standard,
    CreateSOA: pack.CreateSOA,
    Framework: pack.Framework,
    Measures: make([]CompiledMeasure, 0, len(pack.Templates)),
    Documents: make([]CompiledDocument, 0, len(pack.Templates)),
  }

  for _, tmpl := range pack.Templates {
    markdown := renderTemplate(pack, tmpl, answers)
    node, err := prosemirror.ParseMarkdown(markdown)
    if err != nil {
      return nil, fmt.Errorf("cannot convert template %q to editor content: %w", tmpl.Title, err)
    }
    content, err := json.Marshal(node)
    if err != nil {
      return nil, fmt.Errorf("cannot encode template %q: %w", tmpl.Title, err)
    }

    measureRef := fmt.Sprintf("compplus-%s-%s", pack.ID, tmpl.ID)
    evidence := make([]CompiledEvidence, 0, len(tmpl.Evidence))
    for i, name := range tmpl.Evidence {
      evidence = append(evidence, CompiledEvidence{
        ReferenceID: fmt.Sprintf("%s-evidence-%02d", measureRef, i+1),
        Type: "FILE",
        Name: name,
      })
    }

    compiled.Measures = append(compiled.Measures, CompiledMeasure{
      Name: tmpl.Title + " is implemented and maintained",
      Description: tmpl.Purpose,
      Category: tmpl.Category,
      ReferenceID: measureRef,
      Standards: []CompiledStandardMapping{{Framework: pack.Framework.ID, Control: tmpl.ControlID}},
      Tasks: []CompiledTask{{
        Name: "Review and approve " + tmpl.Title,
        Description: "Review the prepared wording, complete highlighted organisation-specific confirmations, attach evidence where requested, and approve the document when accurate.",
        ReferenceID: measureRef + "-review",
        RequestedEvidences: evidence,
      }},
    })

    compiled.Documents = append(compiled.Documents, CompiledDocument{
      ReferenceID: "compplus-" + pack.ID + "-document-" + tmpl.ID,
      MeasureReferenceID: measureRef,
      Title: tmpl.Title,
      Content: string(content),
      DocumentType: tmpl.DocumentType,
      Classification: tmpl.Classification,
    })
  }

  return compiled, nil
}

func renderTemplate(pack Pack, tmpl TemplateDefinition, values map[string]any) string {
  owner := valueOr(values, tmpl.OwnerKey, valueOr(values, "roles.system_manager", "[CONFIRM: document owner]"))
  org := valueOr(values, "organization.legal_name", "[CONFIRM: legal organisation name]")
  services := valueOr(values, "organization.services", "[CONFIRM: products and services]")
  locations := valueOr(values, "organization.locations", "[CONFIRM: locations and remote operations]")
  effective := valueOr(values, "document.effective_date", "[CONFIRM: effective date]")
  review := valueOr(values, "document.next_review_date", "[CONFIRM: next review date]")

  var b strings.Builder
  fmt.Fprintf(&b, "# %s\n\n", tmpl.Title)
  fmt.Fprintf(&b, "**Organisation:** %s  \n", org)
  fmt.Fprintf(&b, "**Document owner:** %s  \n", owner)
  fmt.Fprintf(&b, "**Effective date:** %s  \n", effective)
  fmt.Fprintf(&b, "**Next review:** %s  \n", review)
  fmt.Fprintf(&b, "**Status:** Draft prepared by Comp Plus+ — review confirmations before approval.\n\n")

  fmt.Fprintf(&b, "## Purpose\n\n%s\n\n", tmpl.Purpose)
  fmt.Fprintf(&b, "## Scope\n\nThis %s applies to %s and the people, systems, suppliers and locations used to deliver %s across %s, except where a documented and approved exclusion is recorded.\n\n", strings.ToLower(tmpl.Kind), org, services, locations)

  fmt.Fprintf(&b, "## How %s will work\n\n", org)
  for _, item := range tmpl.Requirements {
    fmt.Fprintf(&b, "- %s\n", renderText(item, values))
  }
  b.WriteString("\n")

  if len(tmpl.StarterContent) > 0 {
    b.WriteString("## Starter content\n\n")
    for _, item := range tmpl.StarterContent {
      fmt.Fprintf(&b, "- %s\n", renderText(item, values))
    }
    b.WriteString("\n")
  }

  b.WriteString("## Responsibilities\n\n")
  fmt.Fprintf(&b, "- **Owner:** %s keeps this document accurate, coordinates actions and retains evidence.\n", owner)
  b.WriteString("- **Senior management:** provides resources, resolves escalated issues and approves material changes.\n")
  b.WriteString("- **Workers and relevant third parties:** follow the approved process and report issues promptly.\n\n")

  if len(tmpl.Evidence) > 0 {
    b.WriteString("## Records and evidence\n\n")
    for _, item := range tmpl.Evidence {
      fmt.Fprintf(&b, "- %s\n", item)
    }
    b.WriteString("\n")
  }

  if len(tmpl.Confirmations) > 0 {
    b.WriteString("## Organisation-specific confirmations\n\n")
    b.WriteString("Complete only these highlighted points; the rest of the template is already prepared.\n\n")
    for _, item := range tmpl.Confirmations {
      fmt.Fprintf(&b, "- [CONFIRM: %s]\n", item)
    }
    b.WriteString("\n")
  }

  if len(tmpl.References) > 0 {
    b.WriteString("## Compliance mapping\n\n")
    fmt.Fprintf(&b, "This Comp Plus+ implementation template supports: %s. It is implementation guidance, not a reproduction of the standard text.\n\n", strings.Join(tmpl.References, ", "))
  }

  b.WriteString("## Review and approval\n\n")
  b.WriteString("Review after significant legal, regulatory, organisational, technology, supplier or incident-related change and at least on the scheduled review date. Approval confirms that the wording reflects actual practice; it does not by itself establish certification or legal compliance.\n")
  return b.String()
}

func renderText(s string, values map[string]any) string {
  for key, value := range values {
    s = strings.ReplaceAll(s, "{{"+key+"}}", fmt.Sprint(value))
  }
  return s
}

func valueOr(values map[string]any, key, fallback string) string {
  if value, ok := values[key]; ok {
    text := strings.TrimSpace(fmt.Sprint(value))
    if text != "" {
      return text
    }
  }
  return fallback
}

func withDefaults(input map[string]any, now time.Time) map[string]any {
  values := make(map[string]any, len(input)+16)
  for k, v := range input {
    values[k] = v
  }
  if valueOr(values, "roles.security_owner", "") == "" {
    values["roles.security_owner"] = valueOr(values, "roles.system_manager", "[CONFIRM: security owner]")
  }
  if valueOr(values, "roles.privacy_owner", "") == "" {
    values["roles.privacy_owner"] = valueOr(values, "roles.system_manager", "[CONFIRM: privacy owner]")
  }
  if valueOr(values, "roles.quality_owner", "") == "" {
    values["roles.quality_owner"] = valueOr(values, "roles.system_manager", "[CONFIRM: quality owner]")
  }
  if valueOr(values, "roles.environment_owner", "") == "" {
    values["roles.environment_owner"] = valueOr(values, "roles.system_manager", "[CONFIRM: environmental owner]")
  }
  if valueOr(values, "roles.ai_owner", "") == "" {
    values["roles.ai_owner"] = valueOr(values, "roles.system_manager", "[CONFIRM: AI governance owner]")
  }
  values["document.effective_date"] = now.Format(time.DateOnly)
  values["document.next_review_date"] = now.AddDate(1, 0, 0).Format(time.DateOnly)
  values["document.year"] = now.Year()
  return values
}

func validatePack(pack Pack) error {
  if pack.ID == "" || pack.Version == "" || pack.Framework.ID == "" || pack.Framework.Name == "" {
    return fmt.Errorf("invalid template pack metadata")
  }
  controls := map[string]bool{}
  for _, control := range pack.Framework.Controls {
    if control.ID == "" || control.Name == "" {
      return fmt.Errorf("pack %q contains an invalid control", pack.ID)
    }
    controls[control.ID] = true
  }
  for _, tmpl := range pack.Templates {
    if tmpl.ID == "" || tmpl.Title == "" || !controls[tmpl.ControlID] {
      return fmt.Errorf("pack %q template %q refers to an unknown control %q", pack.ID, tmpl.ID, tmpl.ControlID)
    }
    if tmpl.DocumentType == "" || tmpl.Classification == "" {
      return fmt.Errorf("pack %q template %q is missing document metadata", pack.ID, tmpl.ID)
    }
  }
  return nil
}
