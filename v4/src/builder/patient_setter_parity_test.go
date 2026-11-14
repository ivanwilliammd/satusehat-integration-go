package builder_test

// The patient setter exists in two shapes across the canonical builders: a reference string with an
// optional display, and a whole Reference. Both are checked against the PHP payload rather than against
// what the Go signature implies.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func bacaGolden(t *testing.T, nama string) map[string]interface{} {
	t.Helper()
	isi, err := os.ReadFile(filepath.Join("testdata", nama+"_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}
	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}
	return harapan
}

func samakan(t *testing.T, harapan map[string]interface{}, b interface{ Build() map[string]interface{} }) {
	t.Helper()
	dapat := b.Build()
	// Bandingkan lewat JSON supaya tipe Go dan PHP dibandingkan sebagai data, bukan sebagai tipe.
	h, err := json.Marshal(harapan)
	if err != nil {
		t.Fatalf("marshal harapan: %v", err)
	}
	d, err := json.Marshal(dapat)
	if err != nil {
		t.Fatalf("marshal dapat: %v", err)
	}
	var hn, dn interface{}
	_ = json.Unmarshal(h, &hn)
	_ = json.Unmarshal(d, &dn)
	if !reflect.DeepEqual(hn, dn) {
		hi, _ := json.MarshalIndent(hn, "", "  ")
		di, _ := json.MarshalIndent(dn, "", "  ")
		t.Fatalf("payload berbeda\n--- PHP ---\n%s\n--- Go ---\n%s", hi, di)
	}
}

func TestClaimPatientMatchesPHPPayload(t *testing.T) {
	b := builder.NewClaimBuilder()
	b.SetId("claim-1")
	b.SetPatient("Patient/1000000001", "Budi Santoso")
	samakan(t, bacaGolden(t, "claim_minimal"), b)
}

func TestDevicePatientMatchesPHPPayload(t *testing.T) {
	b := builder.NewDeviceBuilder()
	b.SetId("dev-1")
	b.SetPatient(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	samakan(t, bacaGolden(t, "device"), b)
}

func TestClaimInsuranceFieldsMatchPHPPayload(t *testing.T) {
	focal := true
	sistem, kode := "http://terminology.hl7.org/CodeSystem/coverage-priority", "P1"
	b := builder.NewClaimBuilder()
	b.SetId("claim-1")
	b.SetPatient("Patient/1000000001", "Budi Santoso")
	b.SetInsurer("Organization/1000000001")
	b.SetCreated("2026-01-05T10:00:00+07:00")
	b.SetType("http://terminology.hl7.org/CodeSystem/claim-type", "institutional", "Institutional")
	b.AddCoverage("Coverage/1400000001", &focal, &sistem, &kode)
	samakan(t, bacaGolden(t, "claim"), b)
}

func TestAccountTypeAndCoverageMatchPHPPayload(t *testing.T) {
	b := builder.NewAccountBuilder()
	b.SetId("acc-1")
	b.SetType("http://terminology.hl7.org/CodeSystem/v3-ActCode", "PBI", "Patient Billing", "Tagihan pasien")
	b.AddCoverage("Coverage/1400000001", 1)
	samakan(t, bacaGolden(t, "account"), b)
}
