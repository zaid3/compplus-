package templates

import (
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func TestLoadCoreInstallManifest(t *testing.T) {
	manifest, err := LoadInstallManifest("core")
	if err != nil {
		t.Fatalf("LoadInstallManifest returned an error: %v", err)
	}

	if manifest.Framework.ID != "COMPPLUS-CORE-2026" {
		t.Fatalf("unexpected framework id %q", manifest.Framework.ID)
	}
	if got, want := len(manifest.Framework.Controls), 12; got != want {
		t.Fatalf("control count = %d, want %d", got, want)
	}
	if got, want := len(manifest.Items), 12; got != want {
		t.Fatalf("install item count = %d, want %d", got, want)
	}
}

func TestCompileCorePack(t *testing.T) {
	now := time.Date(2026, time.August, 3, 12, 0, 0, 0, time.UTC)
	compiled, err := Compile("core", CompileOptions{
		Now: now,
		Answers: map[string]any{
			"organization.legal_name":             "Example Education Ltd",
			"organization.services":               "education and student-support services",
			"organization.locations":               "London office and approved remote-working locations",
			"organization.processes_personal_data": true,
			"organization.selected_standards":      "ISO 27001 and UK GDPR",
			"organization.certification_target":    "2027-03-31",
			"roles.executive_owner":                "A. Director",
			"roles.system_manager":                 "C. Manager",
			"roles.security_owner":                 "S. Owner",
			"roles.privacy_owner":                  "P. Owner",
			"roles.quality_owner":                  "Q. Owner",
			"roles.environment_owner":              "E. Owner",
			"roles.ai_owner":                       "AI Owner",
		},
	})
	if err != nil {
		t.Fatalf("Compile returned an error: %v", err)
	}

	if got, want := len(compiled.Measures), 12; got != want {
		t.Fatalf("measure count = %d, want %d", got, want)
	}
	if got, want := len(compiled.Documents), 12; got != want {
		t.Fatalf("document count = %d, want %d", got, want)
	}

	measureReferences := make(map[string]struct{}, len(compiled.Measures))
	for _, measure := range compiled.Measures {
		if _, exists := measureReferences[measure.ReferenceID]; exists {
			t.Fatalf("duplicate compiled measure reference %q", measure.ReferenceID)
		}
		measureReferences[measure.ReferenceID] = struct{}{}
		if len(measure.Standards) != 1 || measure.Standards[0].Framework != compiled.Framework.ID {
			t.Fatalf("measure %q is not mapped to the compiled framework", measure.ReferenceID)
		}
		if len(measure.Tasks) != 1 {
			t.Fatalf("measure %q task count = %d, want 1", measure.ReferenceID, len(measure.Tasks))
		}
	}

	for _, document := range compiled.Documents {
		if strings.Contains(document.Content, "{{") {
			t.Fatalf("document %q contains an unresolved placeholder", document.Title)
		}
		if !strings.Contains(document.Content, "Example Education Ltd") {
			t.Fatalf("document %q was not personalised with the organisation name", document.Title)
		}
		if _, exists := measureReferences[document.MeasureReferenceID]; !exists {
			t.Fatalf("document %q refers to unknown measure %q", document.Title, document.MeasureReferenceID)
		}

		var editorDocument map[string]any
		if err := json.Unmarshal([]byte(document.Content), &editorDocument); err != nil {
			t.Fatalf("document %q contains invalid editor JSON: %v", document.Title, err)
		}
		if editorDocument["type"] != "doc" {
			t.Fatalf("document %q editor root type = %v, want doc", document.Title, editorDocument["type"])
		}
	}
}

func TestCompileMarksMissingAnswersForConfirmation(t *testing.T) {
	compiled, err := Compile("core", CompileOptions{
		Now: time.Date(2026, time.August, 3, 12, 0, 0, 0, time.UTC),
		Answers: map[string]any{
			"organization.legal_name": "Example Ltd",
		},
	})
	if err != nil {
		t.Fatalf("Compile returned an error: %v", err)
	}

	foundConfirmation := false
	for _, document := range compiled.Documents {
		if strings.Contains(document.Content, "CONFIRM") {
			foundConfirmation = true
			break
		}
	}
	if !foundConfirmation {
		t.Fatal("expected missing answers to be converted into confirmation markers")
	}
}