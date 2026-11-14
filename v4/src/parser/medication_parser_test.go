package parser_test

// MedicationParser is checked against the official example payload, element by element: every accessor
// must return exactly what the payload holds.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/parser"
)

func TestMedicationParserMatchesOfficialExample(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "Medication.json"))
	if err != nil {
		t.Fatalf("fixture hilang: %v", err)
	}

	var data map[string]any
	if err := json.Unmarshal(isi, &data); err != nil {
		t.Fatalf("fixture tidak sah: %v", err)
	}
	if data["resourceType"] != "Medication" {
		t.Fatalf("resourceType = %v, want Medication", data["resourceType"])
	}

	pr := parser.NewMedicationParser(data)
	if pr.GetRaw() == nil {
		t.Fatal("GetRaw() nil")
	}
	if !reflect.DeepEqual(pr.GetResourceType(), data["resourceType"]) {
		t.Fatalf("GetResourceType() = %v, want %v", pr.GetResourceType(), data["resourceType"])
	}
	if !reflect.DeepEqual(pr.GetMeta(), data["meta"]) {
		t.Fatalf("GetMeta() = %v, want %v", pr.GetMeta(), data["meta"])
	}
	if !reflect.DeepEqual(pr.GetIdentifier(), data["identifier"]) {
		t.Fatalf("GetIdentifier() = %v, want %v", pr.GetIdentifier(), data["identifier"])
	}
	if !reflect.DeepEqual(pr.GetCode(), data["code"]) {
		t.Fatalf("GetCode() = %v, want %v", pr.GetCode(), data["code"])
	}
	if !reflect.DeepEqual(pr.GetStatus(), data["status"]) {
		t.Fatalf("GetStatus() = %v, want %v", pr.GetStatus(), data["status"])
	}
	if !reflect.DeepEqual(pr.GetManufacturer(), data["manufacturer"]) {
		t.Fatalf("GetManufacturer() = %v, want %v", pr.GetManufacturer(), data["manufacturer"])
	}
	if !reflect.DeepEqual(pr.GetForm(), data["form"]) {
		t.Fatalf("GetForm() = %v, want %v", pr.GetForm(), data["form"])
	}
	if !reflect.DeepEqual(pr.GetIngredient(), data["ingredient"]) {
		t.Fatalf("GetIngredient() = %v, want %v", pr.GetIngredient(), data["ingredient"])
	}
	if !reflect.DeepEqual(pr.GetBatch(), data["batch"]) {
		t.Fatalf("GetBatch() = %v, want %v", pr.GetBatch(), data["batch"])
	}
	if !reflect.DeepEqual(pr.GetExtension(), data["extension"]) {
		t.Fatalf("GetExtension() = %v, want %v", pr.GetExtension(), data["extension"])
	}
}
