package builder_test

// LocationBuilder must produce the same payload as PayloadBuilderLocation in the PHP SDK, including
// the position object and the list-shaped type element.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestLocationMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "location_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	addr := &datatype.Address{Line: []string{"Jl. Rumah Sakit No. 1"}, City: str("Serang")}
	tel := &datatype.ContactPoint{System: useCP("phone"), Value: str("0254000001")}

	b := builder.NewLocationBuilder()
	b.SetId("1000000003")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/location/1000000003", Value: "1000000003"})
	b.SetStatus("active")
	b.SetOperationalStatus(concept("http://terminology.hl7.org/CodeSystem/v2-0116", "O", "Occupied"))
	b.SetName("Poli Umum")
	b.AddAlias("POLI-01")
	b.SetDescription("Ruang pemeriksaan umum lantai 1")
	b.SetType(concept("http://terminology.hl7.org/CodeSystem/v3-RoleCode", "OUTPHARM", "outpatient pharmacy"))
	b.AddTelecom(tel)
	b.AddAddress(addr)
	b.SetPhysicalType(concept("http://terminology.hl7.org/CodeSystem/location-physical-type", "ro", "Room"))
	lat, lon, alt := -6.1234, 106.1234, 12.5
	b.SetPosition(&lat, &lon, &alt)
	b.SetManagingOrganization(&datatype.Reference{Reference: "Organization/1000000001"})
	b.SetPartOf(&datatype.Reference{Reference: "Location/1000000009"})

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
