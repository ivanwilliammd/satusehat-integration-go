// Package conformance memeriksa payload terhadap definisi FHIR R4 resmi.
//
// Sebagian builder bisa menulis elemen yang tidak ada di resource-nya (mis. `onsetDateTime` pada
// `Goal`). Payload seperti itu sah secara JSON, hanya bukan R4 — server menolaknya tanpa pesan yang
// berguna. Daftar elemen sah berasal dari data/r4_elements.json, dihasilkan repo PHP dari
// profiles-resources.json HL7 R4 resmi.
package conformance

import (
	_ "embed"
	"encoding/json"
	"os"
	"sort"
)

//go:embed data/r4_elements.json
var elemenMentah []byte

// alias: nama resource berbeda antara kelas builder dan definisi FHIR.
var alias = map[string]string{"ListResource": "List"}

var peta map[string][]string

func muat() map[string][]string {
	if peta != nil {
		return peta
	}
	peta = map[string][]string{}
	if err := json.Unmarshal(elemenMentah, &peta); err != nil {
		// Data yang tidak terbaca berarti penjaga tidak punya dasar memeriksa. Mengembalikan
		// peta kosong akan membuat setiap resource "tidak diperiksa" — gagal senyap, bukan pilihan.
		panic("data/r4_elements.json tidak terbaca: " + err.Error())
	}
	return peta
}

// UnknownElements mengembalikan elemen tingkat atas yang tidak ada di R4 untuk payload ini.
func UnknownElements(payload map[string]any) []string {
	tipe, ok := payload["resourceType"].(string)
	if !ok || tipe == "" {
		return []string{"resourceType"}
	}
	nama := tipe
	if a, ada := alias[tipe]; ada {
		nama = a
	}
	sah := muat()[nama]
	if len(sah) == 0 {
		// Resource di luar definisi R4: tidak ada dasar untuk memeriksa.
		return nil
	}
	diizinkan := make(map[string]struct{}, len(sah))
	for _, e := range sah {
		diizinkan[e] = struct{}{}
	}
	asing := []string{}
	for k := range payload {
		if k == "resourceType" {
			continue
		}
		if _, ada := diizinkan[k]; !ada {
			asing = append(asing, k)
		}
	}
	sort.Strings(asing)
	return asing
}

// AssertR4 mengembalikan error bila ada elemen asing.
func AssertR4(payload map[string]any) error {
	asing := UnknownElements(payload)
	if len(asing) == 0 {
		return nil
	}
	tipe, _ := payload["resourceType"].(string)
	if tipe == "" {
		tipe = "(tanpa resourceType)"
	}
	return &ElementError{Resource: tipe, Elements: asing}
}

// ElementError menyebut resource dan elemen yang melanggar, supaya bisa diperbaiki tanpa menebak.
type ElementError struct {
	Resource string
	Elements []string
}

func (e *ElementError) Error() string {
	p := ""
	for i, el := range e.Elements {
		if i > 0 {
			p += ", "
		}
		p += el
	}
	return e.Resource + " menulis elemen yang tidak ada di FHIR R4: " + p
}

// StrictR4Enabled melaporkan apakah SATUSEHAT_STRICT_R4 menyalakan pemeriksaan.
func StrictR4Enabled() bool {
	switch os.Getenv("SATUSEHAT_STRICT_R4") {
	case "", "0", "false", "off", "no", "FALSE", "OFF", "NO":
		return false
	}
	return true
}
