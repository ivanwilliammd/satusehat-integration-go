package builder_test

// MedicationDispenseBuilder must produce the same payload as PayloadBuilderMedicationDispense in the
// PHP SDK. That builder takes plain ids and builds the Reference itself, and normalises timestamps
// through date('c'), so the test pins both behaviours.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestMedicationDispenseMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "medication_dispense_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewMedicationDispenseBuilder()
	b.SetId("8000000001")
	b.AddIdentifier("http://sys-ids.kemkes.go.id/medicationdispense/1000000001", "8000000001")
	b.SetStatus("completed")
	b.SetMedicationReference("Medication/6000000001", "Paracetamol 500 mg")
	b.SetSubject("1000000001", "Budi Santoso")
	b.SetContext("2000000001")
	b.AddPerformer("1000000002", "Apt. Rina")
	b.SetLocation("1000000003", "Depo Farmasi")
	b.AddAuthorizingPrescription("7000000001")
	b.SetCategory("outpatient", "Outpatient", "")
	b.SetQuantity(10, "tablet", "http://unitsofmeasure.org", "TAB")
	b.SetDaysSupply(3, "day", "http://unitsofmeasure.org", "d")
	b.SetWhenPrepared("2026-01-05T10:15:00+07:00")
	b.SetWhenHandedOver("2026-01-05T10:30:00+07:00")
	b.AddDosageInstruction(1, "3 x 1 tablet sesudah makan", 1, "tablet", "O", "Oral", "BID", "Twice a day", "PC", "Sesudah makan")
	b.SetSubstitution(true, "G", "generic")

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
