package parser_test

// MedicationDispenseParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestMedicationDispenseParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "MedicationDispense.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "MedicationDispense" {
		t.Fatalf("resourceType = %v, want MedicationDispense", data["resourceType"])
	}

	pr := parser.NewMedicationDispenseParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetContained(), data["contained"]) {
		t.Fatalf("GetContained() = %v, want %v", pr.GetContained(), data["contained"])
	}
	if !reflect.DeepEqual(pr.GetIdentifier(), data["identifier"]) {
		t.Fatalf("GetIdentifier() = %v, want %v", pr.GetIdentifier(), data["identifier"])
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
	if !reflect.DeepEqual(pr.GetPerformer(), data["performer"]) {
		t.Fatalf("GetPerformer() = %v, want %v", pr.GetPerformer(), data["performer"])
	}
	if !reflect.DeepEqual(pr.GetLocation(), data["location"]) {
		t.Fatalf("GetLocation() = %v, want %v", pr.GetLocation(), data["location"])
	}
	if !reflect.DeepEqual(pr.GetAuthorizingPrescription(), data["authorizingPrescription"]) {
		t.Fatalf("GetAuthorizingPrescription() = %v, want %v", pr.GetAuthorizingPrescription(), data["authorizingPrescription"])
	}
	if !reflect.DeepEqual(pr.GetQuantity(), data["quantity"]) {
		t.Fatalf("GetQuantity() = %v, want %v", pr.GetQuantity(), data["quantity"])
	}
	if !reflect.DeepEqual(pr.GetDaysSupply(), data["daysSupply"]) {
		t.Fatalf("GetDaysSupply() = %v, want %v", pr.GetDaysSupply(), data["daysSupply"])
	}
	if !reflect.DeepEqual(pr.GetWhenPrepared(), data["whenPrepared"]) {
		t.Fatalf("GetWhenPrepared() = %v, want %v", pr.GetWhenPrepared(), data["whenPrepared"])
	}
	if !reflect.DeepEqual(pr.GetWhenHandedOver(), data["whenHandedOver"]) {
		t.Fatalf("GetWhenHandedOver() = %v, want %v", pr.GetWhenHandedOver(), data["whenHandedOver"])
	}
	if !reflect.DeepEqual(pr.GetDosageInstruction(), data["dosageInstruction"]) {
		t.Fatalf("GetDosageInstruction() = %v, want %v", pr.GetDosageInstruction(), data["dosageInstruction"])
	}
}
