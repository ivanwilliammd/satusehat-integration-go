package parser_test

// MedicationStatementParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestMedicationStatementParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "MedicationStatement.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "MedicationStatement" {
		t.Fatalf("resourceType = %v, want MedicationStatement", data["resourceType"])
	}

	pr := parser.NewMedicationStatementParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetContained(), data["contained"]) {
		t.Fatalf("GetContained() = %v, want %v", pr.GetContained(), data["contained"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetCategory(), data["category"]) {
		t.Fatalf("GetCategory() = %v, want %v", pr.GetCategory(), data["category"])
	}
	if !reflect.DeepEqual(pr.GetMedicationReference(), data["medicationReference"]) {
		t.Fatalf("GetMedicationReference() = %v, want %v", pr.GetMedicationReference(), data["medicationReference"])
	}
	if !reflect.DeepEqual(pr.GetSubject(), data["subject"]) {
		t.Fatalf("GetSubject() = %v, want %v", pr.GetSubject(), data["subject"])
	}
	if !reflect.DeepEqual(pr.GetDosage(), data["dosage"]) {
		t.Fatalf("GetDosage() = %v, want %v", pr.GetDosage(), data["dosage"])
	}
	if !reflect.DeepEqual(pr.GetEffectiveDateTime(), data["effectiveDateTime"]) {
		t.Fatalf("GetEffectiveDateTime() = %v, want %v", pr.GetEffectiveDateTime(), data["effectiveDateTime"])
	}
	if !reflect.DeepEqual(pr.GetDateAsserted(), data["dateAsserted"]) {
		t.Fatalf("GetDateAsserted() = %v, want %v", pr.GetDateAsserted(), data["dateAsserted"])
	}
	if !reflect.DeepEqual(pr.GetInformationSource(), data["informationSource"]) {
		t.Fatalf("GetInformationSource() = %v, want %v", pr.GetInformationSource(), data["informationSource"])
	}
	if !reflect.DeepEqual(pr.GetContext(), data["context"]) {
		t.Fatalf("GetContext() = %v, want %v", pr.GetContext(), data["context"])
	}
}
