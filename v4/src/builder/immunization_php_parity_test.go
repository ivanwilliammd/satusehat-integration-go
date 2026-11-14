package builder_test

// ImmunizationBuilder must produce the same payload as PayloadBuilderImmunization in the PHP SDK,
// and must reject the same values that builder throws on.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestImmunizationMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "immunization_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewImmunizationBuilder()
	b.SetMetaProfile("https://fhir.kemkes.go.id/r4/StructureDefinition/Immunization")
	b.SetId("9200000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/immunization/1000000001", Value: "9200000001"})
	b.SetStatus("completed")
	b.SetVaccineCodeFromCode("93001282", "Vaksin COVID-19", "")
	b.SetPatient(&datatype.Reference{Reference: "Patient/1000000001"}, "Budi Santoso")
	b.SetOccurrenceDateTime("2026-01-05T09:00:00+07:00")
	b.AddPerformer(
		&datatype.Reference{Reference: "Practitioner/1000000002"},
		concept("http://terminology.hl7.org/CodeSystem/v2-0443", "AP", "Administering Provider"),
	)
	b.SetDoseQuantity(72, "ug", "", "")
	b.SetLocation(&datatype.Reference{Reference: "Location/1000000003"}, "Poli Umum")
	b.SetLotNumber("LOT-2026-001")
	b.SetRecorded("2026-01-05T09:05:00+07:00")
	b.SetPrimarySource(true)
	b.AddProtocolApplied(2, concept("http://snomed.info/sct", "1234", "Primary series"))
	b.SetRouteFromCode("IM", "Intramuscular", "")
	b.AddReasonCode(concept("http://snomed.info/sct", "429060002", "Prophylaxis"))
	b.AddExtension("https://fhir.kemkes.go.id/r4/StructureDefinition/immunization-vaccine-holder", "PT Bio Farma")

	if err := b.Err(); err != nil {
		t.Fatalf("tidak diharapkan error: %v", err)
	}

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

// PHP throws FHIRInvalidPropertyValue for these two cases; Go records the error instead of
// panicking, and must not write the element.
func TestImmunizationRejectsInvalidValues(t *testing.T) {
	b := builder.NewImmunizationBuilder()
	b.SetStatus("done")

	if b.Err() == nil {
		t.Fatal("status tidak sah seharusnya menghasilkan error")
	}
	if _, ada := b.Build()["status"]; ada {
		t.Fatal("status tidak sah tidak boleh ditulis ke payload")
	}

	c := builder.NewImmunizationBuilder()
	c.SetVaccineCodeFromCode("12345", "Bukan vaksin resmi", "")
	if c.Err() == nil {
		t.Fatal("kode vaksin di luar allowlist seharusnya menghasilkan error")
	}
	if _, ada := c.Build()["vaccineCode"]; ada {
		t.Fatal("kode vaksin tidak sah tidak boleh ditulis ke payload")
	}
}
