package parser_test

// ClinicalImpressionParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestClinicalImpressionParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "ClinicalImpression.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "ClinicalImpression" {
		t.Fatalf("resourceType = %v, want ClinicalImpression", data["resourceType"])
	}

	pr := parser.NewClinicalImpressionParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetIdentifier(), data["identifier"]) {
		t.Fatalf("GetIdentifier() = %v, want %v", pr.GetIdentifier(), data["identifier"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetCode(), data["code"]) {
		t.Fatalf("GetCode() = %v, want %v", pr.GetCode(), data["code"])
	}
	if !reflect.DeepEqual(pr.GetDescription(), data["description"]) {
		t.Fatalf("GetDescription() = %v, want %v", pr.GetDescription(), data["description"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetEncounter(), data["encounter"]) {
		t.Fatalf("GetEncounter() = %v, want %v", pr.GetEncounter(), data["encounter"])
	}
	if !reflect.DeepEqual(pr.GetEffectiveDateTime(), data["effectiveDateTime"]) {
		t.Fatalf("GetEffectiveDateTime() = %v, want %v", pr.GetEffectiveDateTime(), data["effectiveDateTime"])
	}
	if !reflect.DeepEqual(pr.GetDate(), data["date"]) {
		t.Fatalf("GetDate() = %v, want %v", pr.GetDate(), data["date"])
	}
	if !reflect.DeepEqual(pr.GetAssessor(), data["assessor"]) {
		t.Fatalf("GetAssessor() = %v, want %v", pr.GetAssessor(), data["assessor"])
	}
	if !reflect.DeepEqual(pr.GetProblem(), data["problem"]) {
		t.Fatalf("GetProblem() = %v, want %v", pr.GetProblem(), data["problem"])
	}
	if !reflect.DeepEqual(pr.GetInvestigation(), data["investigation"]) {
		t.Fatalf("GetInvestigation() = %v, want %v", pr.GetInvestigation(), data["investigation"])
	}
	if !reflect.DeepEqual(pr.GetSummary(), data["summary"]) {
		t.Fatalf("GetSummary() = %v, want %v", pr.GetSummary(), data["summary"])
	}
	if !reflect.DeepEqual(pr.GetFinding(), data["finding"]) {
		t.Fatalf("GetFinding() = %v, want %v", pr.GetFinding(), data["finding"])
	}
	if !reflect.DeepEqual(pr.GetPrognosisCodeableConcept(), data["prognosisCodeableConcept"]) {
		t.Fatalf("GetPrognosisCodeableConcept() = %v, want %v", pr.GetPrognosisCodeableConcept(), data["prognosisCodeableConcept"])
	}
}
