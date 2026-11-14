package builder_test

// SpecimenBuilder must produce the same payload as PayloadBuilderSpecimen in the PHP SDK. The
// interesting part is the nested collection object, which PHP builds from path strings
// ("collection/quantity") and this port builds with setPath.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestSpecimenMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "specimen_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewSpecimenBuilder()
	b.SetId("4000000001")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/specimen/1000000001", Value: "4000000001"})
	b.SetStatus("available")
	b.SetType(concept("http://snomed.info/sct", "119297000", "Blood specimen"))
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetReceivedTime("2026-01-05T08:45:00+07:00")
	b.SetCollectedDateTime("2026-01-05T08:15:00+07:00")
	b.SetCollector(&datatype.Reference{Reference: "Practitioner/1000000002", Display: "Ns. Dwi"})
	b.SetFastingStatusCodeableConcept(concept("http://terminology.hl7.org/CodeSystem/v2-0916", "F", "Patient was fasting"))
	b.SetMethod(concept("http://snomed.info/sct", "129316008", "Venipuncture"))
	b.SetQuantity(&datatype.Quantity{Value: ptr(3.0), Unit: str("mL"), System: str("http://unitsofmeasure.org"), Code: str("mL")})
	b.SetBodySite(concept("http://snomed.info/sct", "368208006", "Left antecubital fossa"))
	b.AddRequest(&datatype.Reference{Reference: "ServiceRequest/9400000001"})
	b.AddCondition("Puasa 8 jam")
	b.AddProcessing("2026-01-05T08:50:00+07:00")
	b.AddTransportedTime("2026-01-05T09:10:00+07:00")
	b.AddTransportedPerson("Kurir Lab", []map[string]interface{}{{"system": "phone", "value": "081200000000"}})
	b.AddReceivedPerson(&datatype.Reference{Reference: "Practitioner/1000000002"})

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
