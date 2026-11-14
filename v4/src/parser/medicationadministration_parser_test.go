package parser_test

// MedicationAdministrationParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestMedicationAdministrationParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "MedicationAdministration.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "MedicationAdministration" {
		t.Fatalf("resourceType = %v, want MedicationAdministration", data["resourceType"])
	}

	pr := parser.NewMedicationAdministrationParser(data)
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
	if !reflect.DeepEqual(pr.GetContext(), data["context"]) {
		t.Fatalf("GetContext() = %v, want %v", pr.GetContext(), data["context"])
	}
	if !reflect.DeepEqual(pr.GetEffectivePeriod(), data["effectivePeriod"]) {
		t.Fatalf("GetEffectivePeriod() = %v, want %v", pr.GetEffectivePeriod(), data["effectivePeriod"])
	}
	if !reflect.DeepEqual(pr.GetPerformer(), data["performer"]) {
		t.Fatalf("GetPerformer() = %v, want %v", pr.GetPerformer(), data["performer"])
	}
	if !reflect.DeepEqual(pr.GetReasonCode(), data["reasonCode"]) {
		t.Fatalf("GetReasonCode() = %v, want %v", pr.GetReasonCode(), data["reasonCode"])
	}
	if !reflect.DeepEqual(pr.GetRequest(), data["request"]) {
		t.Fatalf("GetRequest() = %v, want %v", pr.GetRequest(), data["request"])
	}
	if !reflect.DeepEqual(pr.GetDosage(), data["dosage"]) {
		t.Fatalf("GetDosage() = %v, want %v", pr.GetDosage(), data["dosage"])
	}
}
