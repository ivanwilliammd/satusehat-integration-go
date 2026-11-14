package builder_test

// EncounterBuilder must produce the same payload as PayloadBuilderEncounter in the PHP SDK. The
// golden file is the PHP builder's output for a fully populated Encounter.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestEncounterMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "encounter_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewEncounterBuilder()
	b.SetId("2000000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/encounter/1000000001", Value: "2000000001"})
	b.SetStatus("finished")
	b.SetClass(&datatype.Coding{System: "http://terminology.hl7.org/CodeSystem/v3-ActCode", Code: "AMB", Display: "ambulatory"})
	b.AddType(&datatype.CodeableConcept{Coding: []datatype.Coding{{
		System: "http://terminology.hl7.org/CodeSystem/encounter-type", Code: "AMB", Display: "Ambulatory",
	}}})
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.AddParticipant(
		&datatype.Reference{Reference: "Practitioner/1000000002", Display: "dr. Ani"},
		&datatype.CodeableConcept{Coding: []datatype.Coding{{
			System: "http://terminology.hl7.org/CodeSystem/v3-ParticipationType", Code: "ATND", Display: "attender",
		}}},
		nil,
	)
	b.AddLocation(&datatype.Reference{Reference: "Location/1000000003", Display: "Poli Umum"}, "active", nil)
	b.SetPeriod(&datatype.Period{
		Start: ptr("2026-01-05T08:00:00+07:00"),
		End:   ptr("2026-01-05T09:00:00+07:00"),
	})
	b.SetServiceProvider(&datatype.Reference{Reference: "Organization/1000000001", Display: "RS Contoh"})
	b.AddDiagnosis(&datatype.Reference{Reference: "Condition/1000000004"}, 1, nil, nil)
	b.AddReasonCode(&datatype.CodeableConcept{Coding: []datatype.Coding{{
		System: "http://hl7.org/fhir/sid/icd-10", Code: "A09", Display: "Diarrhoea",
	}}})

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
