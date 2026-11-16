package terminology

// Aturan wajib SATUSEHAT yang bisa diputuskan dari payload saja.
//
// Audit R4 menjawab "apakah elemen ini ada di R4"; ia tidak menjawab "apakah isinya bentuk yang
// diterima SATUSEHAT". Aturannya TIDAK ditulis di sini: data/satusehat/required-rules.json (disertakan
// lewat go:embed) memuat tiap butir bersama URL dokumentasi sumbernya, dan berkas ini hanya menjalankan
// butir itu. Menambah resource berarti menambah data, bukan kode.
//
// Yang tidak diklaim: aturan yang butuh data di luar payload (mis. identifier organisasi harus
// terdaftar di master sarana).

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
)

// ActCode adalah sistem wajib untuk Encounter.class menurut dokumentasi SATUSEHAT.
const ActCode = "http://terminology.hl7.org/CodeSystem/v3-ActCode"

var (
	EncounterStatus = []string{
		"planned", "arrived", "triaged", "in-progress", "onleave",
		"finished", "cancelled", "entered-in-error", "unknown",
	}

	ObservationStatus = []string{
		"registered", "preliminary", "final", "amended", "corrected",
		"cancelled", "entered-in-error", "unknown",
	}

	ProcedureStatus = []string{
		"preparation", "in-progress", "not-done", "on-hold", "stopped",
		"completed", "entered-in-error", "unknown",
	}
)

//go:embed data/satusehat/required-rules.json
var aturanMentah []byte

// Butir adalah satu aturan wajib untuk satu elemen.
type Butir struct {
	Elemen  string   `json:"elemen"`
	Jenis   string   `json:"jenis"`
	Kode    []string `json:"kode"`
	Awalan  string   `json:"awalan"`
	Sistem  string   `json:"sistem"`
	Pilihan []string `json:"pilihan"`
	Sumber  string   `json:"sumber"`
}

var (
	sekali   sync.Once
	aturan   map[string][]Butir
	galatIsi error
)

// AturanWajib mengembalikan berkas aturan, dibaca sekali.
func AturanWajib() (map[string][]Butir, error) {
	sekali.Do(func() {
		var isi map[string]json.RawMessage
		if err := json.Unmarshal(aturanMentah, &isi); err != nil {
			galatIsi = fmt.Errorf("berkas aturan tidak terbaca: %w", err)
			return
		}
		aturan = make(map[string][]Butir, len(isi))
		for kunci, mentah := range isi {
			if strings.HasPrefix(kunci, "_") {
				continue
			}
			var butir []Butir
			if err := json.Unmarshal(mentah, &butir); err != nil {
				galatIsi = fmt.Errorf("aturan %s tidak terbaca: %w", kunci, err)
				return
			}
			aturan[kunci] = butir
		}
	})
	return aturan, galatIsi
}

// Violations mengembalikan daftar pelanggaran; kosong berarti lolos.
func Violations(payload map[string]any) []string {
	resourceType, _ := payload["resourceType"].(string)
	if resourceType == "" {
		return []string{"resourceType tidak ada"}
	}

	semua, err := AturanWajib()
	if err != nil {
		// Aturan yang tidak terbaca bukan alasan menyatakan payload bersih.
		return []string{err.Error()}
	}

	masalah := []string{}
	for _, butir := range semua[resourceType] {
		if pesan := periksa(payload, resourceType, butir); pesan != "" {
			masalah = append(masalah, pesan)
		}
	}
	return masalah
}

func periksa(payload map[string]any, resourceType string, butir Butir) string {
	nilai := ambil(payload, strings.ReplaceAll(butir.Elemen, "[x]", ""))
	sebut := resourceType + "." + butir.Elemen
	// SATUSEHAT mewajibkan encounter pada beberapa resource yang di R4 opsional. Alasannya ditulis di
	// pesan supaya tidak dikira pustaka ini salah membaca spesifikasi.
	catatanR4 := ""
	if butir.Elemen == "encounter" {
		catatanR4 = " (SATUSEHAT, meski opsional di R4)"
	}

	switch butir.Jenis {
	case "kode":
		if teks, ok := nilai.(string); ok && berisi(butir.Kode, teks) {
			return ""
		}
		if teks, ok := nilai.(string); ok {
			return fmt.Sprintf("%s bukan kode yang sah: %s", sebut, teks)
		}
		return fmt.Sprintf("%s bukan kode yang sah: (kosong)", sebut)
	case "isi":
		if kosong(nilai) {
			return sebut + " wajib diisi"
		}
		return ""
	case "ada":
		if nilai == nil {
			return sebut + " wajib ada" + catatanR4
		}
		return ""
	case "referensi":
		if acuan, ok := nilai.(map[string]any)["reference"].(string); ok && strings.HasPrefix(acuan, butir.Awalan) {
			return ""
		}
		return fmt.Sprintf("%s wajib merujuk %s", sebut, butir.Awalan)
	case "sistem":
		if sistem(nilai) == butir.Sistem {
			return ""
		}
		return fmt.Sprintf("%s wajib system %s", sebut, butir.Sistem)
	case "salah_satu":
		for _, pilihan := range butir.Pilihan {
			if _, ada := payload[pilihan]; ada {
				return ""
			}
		}
		return fmt.Sprintf("%s wajib salah satu: %s", sebut, strings.Join(butir.Pilihan, ", "))
	default:
		return fmt.Sprintf("jenis aturan tidak dikenal: %s", butir.Jenis)
	}
}

// sistem mengambil `system` langsung (Coding) atau dari coding[0] (CodeableConcept). Encounter.class
// bertipe Coding, sedangkan elemen lain yang diperiksa di sini bertipe CodeableConcept.
func sistem(nilai any) string {
	objek, ok := nilai.(map[string]any)
	if !ok {
		return ""
	}
	if langsung, ok := objek["system"].(string); ok {
		return langsung
	}
	coding, ok := objek["coding"].([]any)
	if !ok || len(coding) == 0 {
		return ""
	}
	pertama, ok := coding[0].(map[string]any)
	if !ok {
		return ""
	}
	teks, _ := pertama["system"].(string)
	return teks
}

func ambil(payload map[string]any, jalur string) any {
	var simpul any = payload
	for _, kunci := range strings.Split(jalur, ".") {
		objek, ok := simpul.(map[string]any)
		if !ok {
			return nil
		}
		simpul, ok = objek[kunci]
		if !ok {
			return nil
		}
	}
	return simpul
}

func kosong(nilai any) bool {
	switch v := nilai.(type) {
	case nil:
		return true
	case string:
		return v == ""
	case []any:
		return len(v) == 0
	case map[string]any:
		return len(v) == 0
	}
	return false
}

func berisi(daftar []string, teks string) bool {
	for _, satu := range daftar {
		if satu == teks {
			return true
		}
	}
	return false
}
