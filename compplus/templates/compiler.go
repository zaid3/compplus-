// Copyright (c) 2026 CompPlus.
// Use of this source code is governed by the MIT license in the repository root.

package templates

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"regexp"
	"strings"
	"time"

	"go.probo.inc/probo/pkg/prosemirror"
)

// Files contains the versioned CompPlus template library.
//
//go:embed catalog.json core schema
var Files embed.FS

var placeholderPattern = regexp.MustCompile(`\{\{\s*([^}|]+?)\s*(?:\|\s*([a-z_]+)\s*)?\}\}`)

type (
	InstallManifest struct {
		PackID    string              `json:"pack_id"`
		Version   string              `json:"version"`
		Framework FrameworkDefinition `json:"framework"`
		Items     []InstallItem       `json:"items"`
	}

	FrameworkDefinition struct {
		ID          string              `json:"id"`
		Name        string              `json:"name"`
		Description string              `json:"description"`
		Controls    []ControlDefinition `json:"controls"`
	}

	ControlDefinition struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	InstallItem struct {
		ID       string             `json:"id"`
		Control  string             `json:"control"`
		Category string             `json:"category"`
		Document DocumentDefinition `json:"document"`
		Measure  MeasureDefinition  `json:"measure"`
		Task     TaskDefinition     `json:"task"`
		Evidence []EvidenceDefinition `json:"evidence"`
	}

	DocumentDefinition struct {
		Title          string `json:"title"`
		ContentFile    string `json:"content_file"`
		Type           string `json:"type"`
		Classification string `json:"classification"`
	}

	MeasureDefinition struct {
		ReferenceID string `json:"reference_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	TaskDefinition struct {
		ReferenceID string `json:"reference_id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	EvidenceDefinition struct {
		ReferenceID string `json:"reference_id"`
		Type        string `json:"type"`
		Name        string `json:"name"`
	}

	CompileOptions struct {
		Answers map[string]any
		Now     time.Time
	}

	CompiledPack struct {
		PackID    string                    `json:"pack_id"`
		Version   string                    `json:"version"`
		Framework FrameworkDefinition       `json:"framework"`
		Measures  []CompiledMeasure         `json:"measures"`
		Documents []CompiledDocument        `json:"documents"`
	}

	CompiledMeasure struct {
		Name        string                    `json:"name"`
		Description string                    `json:"description"`
		Category    string                    `json:"category"`
		ReferenceID string                    `json:"reference-id"`
		Standards   []CompiledStandardMapping `json:"standards"`
		Tasks       []CompiledTask            `json:"tasks"`
	}

	CompiledStandardMapping struct {
		Framework string `json:"framework"`
		Control   string `json:"control"`
	}

	CompiledTask struct {
		Name               string             `json:"name"`
		Description        string             `json:"description"`
		ReferenceID        string             `json:"reference-id"`
		RequestedEvidences []CompiledEvidence `json:"requested-evidences"`
	}

	CompiledEvidence struct {
		ReferenceID string `json:"reference-id"`
		Type        string `json:"type"`
		Name        string `json:"name"`
	}

	CompiledDocument struct {
		ReferenceID       string `json:"reference_id"`
		MeasureReferenceID string `json:"measure_reference_id"`
		Title             string `json:"title"`
		Content           string `json:"content"`
		DocumentType      string `json:"document_type"`
		Classification    string `json:"classification"`
	}
)

func LoadInstallManifest(packID string) (*InstallManifest, error) {
	if !safePackID(packID) {
		return nil, fmt.Errorf("invalid template pack id %q", packID)
	}

	path := fmt.Sprintf("%s/install.json", packID)
	data, err := fs.ReadFile(Files, path)
	if err != nil {
		return nil, fmt.Errorf("cannot read template pack manifest %q: %w", packID, err)
	}

	manifest := &InstallManifest{}
	if err := json.Unmarshal(data, manifest); err != nil {
		return nil, fmt.Errorf("cannot decode template pack manifest %q: %w", packID, err)
	}

	if err := manifest.Validate(); err != nil {
		return nil, fmt.Errorf("invalid template pack manifest %q: %w", packID, err)
	}

	return manifest, nil
}

func (m *InstallManifest) Validate() error {
	if m.PackID == "" {
		return fmt.Errorf("pack_id is required")
	}
	if m.Version == "" {
		return fmt.Errorf("version is required")
	}
	if m.Framework.ID == "" || m.Framework.Name == "" {
		return fmt.Errorf("framework id and name are required")
	}
	if len(m.Framework.Controls) == 0 {
		return fmt.Errorf("at least one framework control is required")
	}
	if len(m.Items) == 0 {
		return fmt.Errorf("at least one install item is required")
	}

	controlIDs := make(map[string]struct{}, len(m.Framework.Controls))
	for _, control := range m.Framework.Controls {
		if control.ID == "" || control.Name == "" {
			return fmt.Errorf("every control requires an id and name")
		}
		if _, exists := controlIDs[control.ID]; exists {
			return fmt.Errorf("duplicate control id %q", control.ID)
		}
		controlIDs[control.ID] = struct{}{}
	}

	itemIDs := make(map[string]struct{}, len(m.Items))
	measureIDs := make(map[string]struct{}, len(m.Items))
	for _, item := range m.Items {
		if item.ID == "" {
			return fmt.Errorf("every install item requires an id")
		}
		if _, exists := itemIDs[item.ID]; exists {
			return fmt.Errorf("duplicate install item id %q", item.ID)
		}
		itemIDs[item.ID] = struct{}{}
		if _, exists := controlIDs[item.Control]; !exists {
			return fmt.Errorf("item %q refers to unknown control %q", item.ID, item.Control)
		}
		if item.Document.Title == "" || item.Document.ContentFile == "" {
			return fmt.Errorf("item %q requires a document title and content file", item.ID)
		}
		if !validDocumentType(item.Document.Type) {
			return fmt.Errorf("item %q has invalid document type %q", item.ID, item.Document.Type)
		}
		if !validClassification(item.Document.Classification) {
			return fmt.Errorf("item %q has invalid classification %q", item.ID, item.Document.Classification)
		}
		if item.Measure.ReferenceID == "" || item.Measure.Name == "" {
			return fmt.Errorf("item %q requires a measure reference id and name", item.ID)
		}
		if _, exists := measureIDs[item.Measure.ReferenceID]; exists {
			return fmt.Errorf("duplicate measure reference id %q", item.Measure.ReferenceID)
		}
		measureIDs[item.Measure.ReferenceID] = struct{}{}
		if item.Task.ReferenceID == "" || item.Task.Name == "" {
			return fmt.Errorf("item %q requires a task reference id and name", item.ID)
		}
	}

	return nil
}

func Compile(packID string, options CompileOptions) (*CompiledPack, error) {
	manifest, err := LoadInstallManifest(packID)
	if err != nil {
		return nil, err
	}

	now := options.Now
	if now.IsZero() {
		now = time.Now()
	}
	answers := withDefaults(options.Answers, now)

	compiled := &CompiledPack{
		PackID:    manifest.PackID,
		Version:   manifest.Version,
		Framework: manifest.Framework,
		Measures:  make([]CompiledMeasure, 0, len(manifest.Items)),
		Documents: make([]CompiledDocument, 0, len(manifest.Items)),
	}

	for _, item := range manifest.Items {
		contentPath := fmt.Sprintf("%s/%s", packID, item.Document.ContentFile)
		markdownData, err := fs.ReadFile(Files, contentPath)
		if err != nil {
			return nil, fmt.Errorf("cannot read document template %q: %w", contentPath, err)
		}

		rendered := render(string(markdownData), answers)
		if placeholderPattern.MatchString(rendered) {
			return nil, fmt.Errorf("document template %q contains unresolved placeholders", contentPath)
		}

		node, err := prosemirror.ParseMarkdown(rendered)
		if err != nil {
			return nil, fmt.Errorf("cannot convert document template %q to editor content: %w", contentPath, err)
		}
		content, err := json.Marshal(node)
		if err != nil {
			return nil, fmt.Errorf("cannot encode document template %q: %w", contentPath, err)
		}

		evidences := make([]CompiledEvidence, 0, len(item.Evidence))
		for _, evidence := range item.Evidence {
			evidences = append(evidences, CompiledEvidence{
				ReferenceID: evidence.ReferenceID,
				Type:        evidence.Type,
				Name:        evidence.Name,
			})
		}

		compiled.Measures = append(compiled.Measures, CompiledMeasure{
			Name:        item.Measure.Name,
			Description: item.Measure.Description,
			Category:    item.Category,
			ReferenceID: item.Measure.ReferenceID,
			Standards: []CompiledStandardMapping{{
				Framework: manifest.Framework.ID,
				Control:   item.Control,
			}},
			Tasks: []CompiledTask{{
				Name:               item.Task.Name,
				Description:        item.Task.Description,
				ReferenceID:        item.Task.ReferenceID,
				RequestedEvidences: evidences,
			}},
		})

		compiled.Documents = append(compiled.Documents, CompiledDocument{
			ReferenceID:        "compplus-" + packID + "-document-" + item.ID,
			MeasureReferenceID: item.Measure.ReferenceID,
			Title:              item.Document.Title,
			Content:            string(content),
			DocumentType:       item.Document.Type,
			Classification:     item.Document.Classification,
		})
	}

	return compiled, nil
}

func render(content string, values map[string]any) string {
	return placeholderPattern.ReplaceAllStringFunc(content, func(raw string) string {
		parts := placeholderPattern.FindStringSubmatch(raw)
		key := strings.TrimSpace(parts[1])
		filter := strings.TrimSpace(parts[2])

		if value, ok := lookup(values, key); ok {
			text := strings.TrimSpace(fmt.Sprint(value))
			if text != "" {
				return text
			}
		}

		label := humanize(key)
		switch filter {
		case "confirm":
			return "[CONFIRM: " + label + "]"
		case "conditional":
			return "[CONFIRM APPLICABILITY: " + label + "]"
		default:
			return "[ADD: " + label + "]"
		}
	})
}

func withDefaults(input map[string]any, now time.Time) map[string]any {
	values := make(map[string]any, len(input)+8)
	for key, value := range input {
		values[key] = value
	}

	defaults := map[string]any{
		"document.effective_date":   now.Format(time.DateOnly),
		"document.next_review_date": now.AddDate(1, 0, 0).Format(time.DateOnly),
		"objectives.reporting_year": now.Year(),
		"objectives.year_end":       time.Date(now.Year(), time.December, 31, 0, 0, 0, 0, now.Location()).Format(time.DateOnly),
		"register.last_reviewed":    now.Format(time.DateOnly),
		"register.next_review":      now.AddDate(1, 0, 0).Format(time.DateOnly),
	}
	for key, value := range defaults {
		if _, exists := values[key]; !exists {
			values[key] = value
		}
	}

	return values
}

func lookup(values map[string]any, key string) (any, bool) {
	if value, ok := values[key]; ok {
		return value, true
	}

	parts := strings.Split(key, ".")
	var current any = values
	for _, part := range parts {
		switch typed := current.(type) {
		case map[string]any:
			value, ok := typed[part]
			if !ok {
				return nil, false
			}
			current = value
		case map[string]string:
			value, ok := typed[part]
			if !ok {
				return nil, false
			}
			current = value
		default:
			return nil, false
		}
	}

	return current, true
}

func humanize(key string) string {
	parts := strings.FieldsFunc(key, func(r rune) bool {
		return r == '.' || r == '_' || r == '-'
	})
	for i, part := range parts {
		if part == "" {
			continue
		}
		parts[i] = strings.ToUpper(part[:1]) + part[1:]
	}
	return strings.Join(parts, " ")
}

func safePackID(packID string) bool {
	if packID == "" {
		return false
	}
	for _, r := range packID {
		if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
			return false
		}
	}
	return true
}

func validDocumentType(value string) bool {
	switch value {
	case "OTHER", "GOVERNANCE", "POLICY", "PROCEDURE", "PLAN", "REGISTER", "RECORD", "REPORT", "TEMPLATE":
		return true
	default:
		return false
	}
}

func validClassification(value string) bool {
	switch value {
	case "PUBLIC", "INTERNAL", "CONFIDENTIAL", "SECRET":
		return true
	default:
		return false
	}
}