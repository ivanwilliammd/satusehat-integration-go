package conformance_test

import (
	"errors"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/conformance"
)

func TestPayloadSahTidakAdaTemuan(t *testing.T) {
	p := map[string]any{"resourceType": "Patient", "gender": "male"}
	if asing := conformance.UnknownElements(p); len(asing) != 0 {
		t.Fatalf("temuan tak terduga: %v", asing)
	}
}

func TestElemenBukanR4Ditangkap(t *testing.T) {
	p := map[string]any{"resourceType": "Goal", "onsetDateTime": "2026-01-01"}
	asing := conformance.UnknownElements(p)
	if len(asing) != 1 || asing[0] != "onsetDateTime" {
		t.Fatalf("harus onsetDateTime, dapat %v", asing)
	}
}

func TestAliasListResource(t *testing.T) {
	p := map[string]any{"resourceType": "ListResource", "entry": []any{}}
	if asing := conformance.UnknownElements(p); len(asing) != 0 {
		t.Fatalf("temuan tak terduga: %v", asing)
	}
}

func TestResourceDiLuarR4TidakDiperiksa(t *testing.T) {
	p := map[string]any{"resourceType": "BillingStatus", "apa": 1}
	if asing := conformance.UnknownElements(p); len(asing) != 0 {
		t.Fatalf("temuan tak terduga: %v", asing)
	}
}

func TestAssertMenyebutResourceDanElemen(t *testing.T) {
	err := conformance.AssertR4(map[string]any{"resourceType": "Goal", "onsetDateTime": "x"})
	var el *conformance.ElementError
	if !errors.As(err, &el) {
		t.Fatalf("error harus ElementError, dapat %v", err)
	}
	if el.Resource != "Goal" || len(el.Elements) != 1 || el.Elements[0] != "onsetDateTime" {
		t.Fatalf("isi error salah: %+v", el)
	}
}

func TestTanpaResourceTypeDitolak(t *testing.T) {
	if asing := conformance.UnknownElements(map[string]any{"nama": "x"}); len(asing) != 1 || asing[0] != "resourceType" {
		t.Fatalf("harus resourceType, dapat %v", asing)
	}
}
