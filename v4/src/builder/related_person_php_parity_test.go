package builder_test

// RelatedPersonBuilder must produce the same payload as PayloadBuilderRelatedPerson in the PHP SDK.
// The golden includes a relationship given as a bare code without display, because PHP writes
// "display": null in that case — a detail a cleaner implementation would silently drop.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestRelatedPersonMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "related_person_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewRelatedPersonBuilder()
	b.SetId("rp-1")
	b.AddIdentifier("http://sys-ids.kemkes.go.id/relatedperson/1000000001", str("rp-1"))
	b.SetActive(true)
	b.SetPatient("1000000001", str("Budi Santoso"))
	b.AddRelationship("SPS", str("Spouse"))
	b.AddRelationship("CHILD", nil) // PHP menulis display: null
	b.AddName(&datatype.HumanName{Text: "Yessike", Family: "Yessike"})
	b.AddTelecom("phone", str("081211112222"), "")
	b.AddAddress(&datatype.Address{Line: []string{"Jl. Contoh No. 2"}, City: str("Serang")})
	ya := true
	b.AddCommunication("id-ID", &ya)
	b.SetGender("female")
	b.SetBirthDate("1990-02-02")
	b.AddExtension("https://fhir.kemkes.go.id/r4/StructureDefinition/relatedperson-is-guardian", true)

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
