package terminology_test

// Uji aturan wajib SATUSEHAT. Sumbernya data/satusehat/required-rules.json yang disertakan di paket,
// bukan ingatan: uji terakhir memastikan konstanta kode tidak menyimpang dari berkas itu.

import (
	"strings"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/terminology"
)

func encounter(system, status string) map[string]any {
	return map[string]any{
		"resourceType": "Encounter",
		"status":       status,
		"class":        map[string]any{"system": system, "code": "AMB"},
	}
}

func TestEncounterDenganActCodeLolos(t *testing.T) {
	if masalah := terminology.Violations(encounter(terminology.ActCode, "finished")); len(masalah) != 0 {
		t.Fatalf("harus lolos, dapat %v", masalah)
	}
}

func TestEncounterDenganSistemLainDitolak(t *testing.T) {
	masalah := terminology.Violations(encounter("http://contoh.local/class", "finished"))
	if len(masalah) != 1 || !contains(masalah[0], "v3-ActCode") {
		t.Fatalf("harus menolak sistem lain, dapat %v", masalah)
	}
}

func TestStatusDiLuarDaftarDitolak(t *testing.T) {
	masalah := terminology.Violations(encounter(terminology.ActCode, "selesai"))
	if len(masalah) != 1 || !contains(masalah[0], "bukan kode yang sah") {
		t.Fatalf("harus menolak status asing, dapat %v", masalah)
	}
}

func TestObservationTanpaEncounterMenyebutAlasanR4(t *testing.T) {
	masalah := terminology.Violations(map[string]any{
		"resourceType": "Observation",
		"status":       "final",
		"code":         map[string]any{"coding": []any{map[string]any{"system": "http://loinc.org"}}},
		"subject":      map[string]any{"reference": "Patient/1"},
	})
	if len(masalah) != 1 || !contains(masalah[0], "opsional di R4") {
		t.Fatalf("harus menyebut alasan R4, dapat %v", masalah)
	}
}

func TestSubjectBukanPatientDitolak(t *testing.T) {
	masalah := terminology.Violations(map[string]any{
		"resourceType": "Condition",
		"code":         map[string]any{"coding": []any{map[string]any{"system": "http://hl7.org/fhir/sid/icd-10"}}},
		"subject":      map[string]any{"reference": "Group/1"},
		"encounter":    map[string]any{"reference": "Encounter/2"},
	})
	if len(masalah) != 1 || !contains(masalah[0], "Patient") {
		t.Fatalf("harus menolak subject bukan Patient, dapat %v", masalah)
	}
}

func TestPatientWajibPunyaNama(t *testing.T) {
	masalah := terminology.Violations(map[string]any{
		"resourceType": "Patient",
		"identifier":   []any{map[string]any{"system": "https://fhir.kemkes.go.id/id/nik", "value": "1"}},
	})
	if len(masalah) != 1 || !contains(masalah[0], "name") {
		t.Fatalf("harus menuntut name, dapat %v", masalah)
	}
}

func TestImmunizationMenerimaOccurrenceDateTimeDanString(t *testing.T) {
	dasar := map[string]any{
		"resourceType": "Immunization",
		"status":       "completed",
		"vaccineCode":  map[string]any{"coding": []any{map[string]any{"system": "http://hl7.org/fhir/sid/cvx"}}},
		"patient":      map[string]any{"reference": "Patient/1"},
	}
	for _, kunci := range []string{"occurrenceDateTime", "occurrenceString"} {
		salin := map[string]any{}
		for k, v := range dasar {
			salin[k] = v
		}
		salin[kunci] = "2026-11-16"
		if masalah := terminology.Violations(salin); len(masalah) != 0 {
			t.Fatalf("%s harus diterima, dapat %v", kunci, masalah)
		}
	}
	if masalah := terminology.Violations(dasar); len(masalah) != 1 || !contains(masalah[0], "occurrenceDateTime") {
		t.Fatalf("tanpa occurrence harus ditolak, dapat %v", masalah)
	}
}

func TestMedicationRequestWajibIntent(t *testing.T) {
	masalah := terminology.Violations(map[string]any{
		"resourceType": "MedicationRequest",
		"status":       "active",
		"subject":      map[string]any{"reference": "Patient/1"},
	})
	if len(masalah) != 1 || !contains(masalah[0], "intent") {
		t.Fatalf("harus menuntut intent, dapat %v", masalah)
	}
}

func TestResourceTakDipetakanTidakDiperiksa(t *testing.T) {
	if masalah := terminology.Violations(map[string]any{"resourceType": "AuditEvent"}); len(masalah) != 0 {
		t.Fatalf("resource tak dipetakan harus dilewati, dapat %v", masalah)
	}
}

func TestTanpaResourceTypeDilaporkan(t *testing.T) {
	masalah := terminology.Violations(map[string]any{"status": "finished"})
	if len(masalah) != 1 || masalah[0] != "resourceType tidak ada" {
		t.Fatalf("harus melaporkan resourceType hilang, dapat %v", masalah)
	}
}

func TestSetiapButirMenyebutSumberDanJenisDikenal(t *testing.T) {
	semua, err := terminology.AturanWajib()
	if err != nil {
		t.Fatal(err)
	}
	jenis := map[string]bool{"kode": true, "isi": true, "ada": true, "referensi": true, "sistem": true, "salah_satu": true}
	for resource, butir := range semua {
		if len(butir) == 0 {
			t.Fatalf("%s tanpa aturan", resource)
		}
		for _, satu := range butir {
			if satu.Sumber == "" {
				t.Fatalf("%s tanpa sumber", resource)
			}
			if !jenis[satu.Jenis] {
				t.Fatalf("%s jenis tak dikenal: %s", resource, satu.Jenis)
			}
		}
	}
}

func TestKonstantaSamaDenganBerkasAturan(t *testing.T) {
	semua, err := terminology.AturanWajib()
	if err != nil {
		t.Fatal(err)
	}
	kode := func(resource string) []string {
		for _, butir := range semua[resource] {
			if butir.Elemen == "status" && butir.Jenis == "kode" {
				return butir.Kode
			}
		}
		return nil
	}
	pasangan := map[string][]string{
		"Encounter":   terminology.EncounterStatus,
		"Observation": terminology.ObservationStatus,
		"Procedure":   terminology.ProcedureStatus,
	}
	for resource, harapan := range pasangan {
		dapat := kode(resource)
		if len(dapat) != len(harapan) {
			t.Fatalf("%s: panjang beda %v vs %v", resource, dapat, harapan)
		}
		for i := range harapan {
			if dapat[i] != harapan[i] {
				t.Fatalf("%s: %v vs %v", resource, dapat, harapan)
			}
		}
	}
}

func contains(teks, cari string) bool {
	for i := 0; i+len(cari) <= len(teks); i++ {
		if teks[i:i+len(cari)] == cari {
			return true
		}
	}
	return false
}

// TestSidikJariBerkasAturan memastikan berkas aturan tidak bergeser diam-diam.
// Kelima port memakai berkas yang sama; kalau berkas berubah tanpa uji ikut berubah,
// kelima port memvalidasi hal berbeda tanpa ada yang tahu. Pemeriksa lintas-repo tetap
// scripts/check-rules-drift.sh di repo PHP.
func TestSidikJariBerkasAturan(t *testing.T) {
	// Checkout Windows memakai CRLF; hitung setelah normalisasi.
	mentah := strings.ReplaceAll(string(terminology.BerkasAturan()), "\r\n", "\n")
	if n := len(mentah); n != 5849 {
		t.Fatalf("jumlah byte berkas aturan berubah: %d", n)
	}
	aturan, err := terminology.AturanWajib()
	if err != nil {
		t.Fatalf("berkas aturan gagal dibaca: %v", err)
	}
	harap := map[string]int{
		"Condition":         3,
		"DiagnosticReport":  4,
		"Encounter":         2,
		"Immunization":      4,
		"MedicationRequest": 3,
		"Observation":       4,
		"Patient":           2,
		"Procedure":         4,
	}
	for resource, jumlah := range harap {
		if dapat := len(aturan[resource]); dapat != jumlah {
			t.Fatalf("%s: jumlah butir berubah %d vs %d", resource, dapat, jumlah)
		}
	}
}
