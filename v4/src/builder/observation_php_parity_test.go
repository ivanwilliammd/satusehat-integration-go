package builder_test

// ObservationBuilder must produce the same payload as PayloadBuilderObservation in the PHP SDK.
// testdata/observation_php_golden.json is the PHP builder's output for a fully populated
// Observation, including a component and a referenceRange.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func qty(value float64, unit, system, code string) *datatype.Quantity {
	return &datatype.Quantity{Value: ptr(value), Unit: str(unit), System: str(system), Code: str(code)}
}

func concept(system, code, display string) *datatype.CodeableConcept {
	return &datatype.CodeableConcept{Coding: []datatype.Coding{{System: system, Code: code, Display: display}}}
}

func TestObservationMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "observation_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewObservationBuilder()
	b.SetId("3000000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/observation/1000000001", Value: "3000000001"})
	b.SetStatus("final")
	b.AddCategoryConcept(concept("http://terminology.hl7.org/CodeSystem/observation-category", "laboratory", "Laboratory"))
	b.SetCode(concept("http://loinc.org", "8867-4", "Heart rate"))
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetEncounter(&datatype.Reference{Reference: "Encounter/2000000001"})
	b.SetEffectiveDateTime("2026-01-05T08:30:00+07:00")
	b.SetValueQuantity(qty(88, "beats/minute", "http://unitsofmeasure.org", "/min"))
	b.AddInterpretation(concept("http://terminology.hl7.org/CodeSystem/v3-ObservationInterpretation", "N", "Normal"))
	b.AddNote("Diukur saat pasien duduk.")
	b.AddBodySite(concept("http://snomed.info/sct", "7569003", "Finger structure"))
	b.SetSpecimen(&datatype.Reference{Reference: "Specimen/4000000001"})
	b.AddReferenceRange(
		qty(60, "beats/minute", "http://unitsofmeasure.org", "/min"),
		qty(100, "beats/minute", "http://unitsofmeasure.org", "/min"),
		nil,
		"Dewasa saat istirahat",
	)
	b.AddComponent(
		concept("http://loinc.org", "8480-6", "Systolic blood pressure"),
		qty(120, "mmHg", "http://unitsofmeasure.org", "mm[Hg]"),
	)

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
