package builder_test

// PractitionerBuilder must produce the same payload as PayloadBuilderPractitioner in the PHP SDK.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// dua helper kecil agar test tidak bergantung pada nama konstanta enum ContactPoint.
func useCP(v string) *datatype.ContactPointSystem {
	s := datatype.ContactPointSystem(v)
	return &s
}

func useCPU(v string) *datatype.ContactPointUse {
	u := datatype.ContactPointUse(v)
	return &u
}

func TestPractitionerMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "practitioner_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewPractitionerBuilder()
	b.SetId("1000000002")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/practitioner/1000000002", Value: "1000000002"})
	b.SetActive(true)
	b.AddName(&datatype.HumanName{Text: "dr. Ani Lestari", Family: "Lestari", Given: []string{"Ani"}})
	b.AddTelecom(&datatype.ContactPoint{
		System: useCP("phone"), Value: str("081200000000"), Use: useCPU("work"),
	})
	b.AddAddress(&datatype.Address{
		Use: str("work"), Line: []string{"Jl. Contoh No. 1"}, City: str("Serang"), PostalCode: str("42111"),
	})
	b.SetGender("female")
	b.SetBirthDate("1985-04-12")
	b.AddPhoto("https://example.org/ani.jpg", "image/jpeg")
	b.AddQualification(
		&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/qualification/1000000002", Value: "STR-1"},
		concept("http://terminology.hl7.org/CodeSystem/v2-0360", "MD", "Doctor of Medicine"),
		"2020-01-01",
		&datatype.Reference{Reference: "Organization/1000000001"},
	)
	ya := true
	b.AddCommunication(concept("urn:ietf:bcp:47", "id-ID", "Indonesian"), &ya)

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
