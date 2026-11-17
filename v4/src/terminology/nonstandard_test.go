package terminology_test

import (
	"strings"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/terminology"
)

func TestSemuaURLDomainKemkes(t *testing.T) {
	for _, pair := range terminology.NonStandardAll {
		if !strings.HasPrefix(pair[1], "https://fhir.kemkes.go.id/") {
			t.Fatalf("%s bukan domain Kemkes: %s", pair[0], pair[1])
		}
	}
}

func TestSebelasObjek(t *testing.T) {
	if len(terminology.NonStandardAll) != 11 {
		t.Fatalf("harus 11, dapat %d", len(terminology.NonStandardAll))
	}
}

func TestTidakAdaURLGanda(t *testing.T) {
	lihat := map[string]bool{}
	for _, pair := range terminology.NonStandardAll {
		if lihat[pair[1]] {
			t.Fatalf("URL ganda: %s", pair[1])
		}
		lihat[pair[1]] = true
	}
}

func TestIdentifierSelaluOfficial(t *testing.T) {
	id := terminology.OfficialIdentifier(terminology.NIKSystem, "3171234567890001")
	if id["use"] != "official" || id["system"] != terminology.NIKSystem {
		t.Fatalf("identifier salah: %v", id)
	}
}

func TestAdministrativeCodeIrisan(t *testing.T) {
	ext := terminology.AdministrativeCode([][2]string{{"province", "31"}, {"city", "3171"}})
	if ext["url"] != terminology.Administrative {
		t.Fatalf("url salah: %v", ext["url"])
	}
	if n := len(ext["extension"].([]any)); n != 2 {
		t.Fatalf("irisan harus 2, dapat %d", n)
	}
}
