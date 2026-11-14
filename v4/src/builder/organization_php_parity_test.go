package builder_test

// OrganizationBuilder must produce the same payload as PayloadBuilderOrganization in the PHP SDK,
// including the list-shaped type element and the contact wrapper ({telecom: […], purpose: {text}}).

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestOrganizationMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "organization_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	tel := &datatype.ContactPoint{System: useCP("phone"), Value: str("0254000000")}
	addr := &datatype.Address{Line: []string{"Jl. Rumah Sakit No. 1"}, City: str("Serang")}

	b := builder.NewOrganizationBuilder()
	b.SetId("1000000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/organization/1000000001", Value: "1000000001"})
	b.SetActive(true)
	b.SetName("RS Contoh")
	b.AddAlias("RSC")
	b.SetType(concept("http://terminology.hl7.org/CodeSystem/organization-type", "prov", "Healthcare Provider"))
	b.AddTelecom(tel)
	b.AddAddress(addr)
	b.SetPartOf(&datatype.Reference{Reference: "Organization/1000000009"})
	b.AddContact(tel, "ADMIN", "Bagian Admisi", addr)
	b.AddEndpoint(&datatype.Reference{Reference: "Endpoint/end-1"})

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
