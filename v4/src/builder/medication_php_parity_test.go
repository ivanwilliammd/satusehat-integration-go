package builder_test

// MedicationBuilder must produce the same payload as PayloadBuilderMedication in the PHP SDK,
// including the meta.profile element the PHP constructor sets.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestMedicationMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "medication_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewMedicationBuilder()
	b.SetId("6000000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/medication/1000000001", Value: "6000000001"})
	b.SetCode(concept("http://terminology.kemkes.go.id/CodeSystem/medication-code", "910002021", "Paracetamol 500 mg"))
	b.SetStatus("active")
	b.SetManufacturer(&datatype.Reference{Reference: "Organization/1000000001", Display: "PT Kimia Farma"})
	b.SetForm(concept("http://terminology.kemkes.go.id/CodeSystem/medication-form", "BS006", "Tablet"))
	b.AddIngredient(
		concept("http://terminology.kemkes.go.id/CodeSystem/active-substance", "AS0001", "Paracetamol"),
		true,
		qty(500, "mg", "http://unitsofmeasure.org", "mg"),
	)
	b.SetBatch("LOT-2026-001", "2028-01-31")
	b.AddMedicationType("NC", "Non-compound")

	mentah, err := b.BuildJSON()
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var dapat map[string]interface{}
	if err := json.Unmarshal(mentah, &dapat); err != nil {
		t.Fatalf("hasil bukan JSON sah: %v", err)
	}

	if !reflect.DeepEqual(harapan, dapat) {
		h, _ := json.MarshalIndent(harapan, "", "  ")
		d, _ := json.MarshalIndent(dapat, "", "  ")
		t.Fatalf("payload Go berbeda dari PHP\n--- PHP ---\n%s\n--- Go ---\n%s", h, d)
	}
}
