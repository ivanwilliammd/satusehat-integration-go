package terminology_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/terminology"
)

// Counts come from the PHP catalog dump. If a catalog is regenerated, these numbers move
// together with the data; a mismatch means the generated file is out of date.
func TestCatalogCountsMatchPhp(t *testing.T) {
	kasus := []struct {
		nama    string
		katalog map[string]any
		kunci   string
		jumlah  int
	}{
		{"FamilyRelationship.relationshipCodes", terminology.FamilyRelationshipCatalogs, "relationshipCodes", 105},
		{"Occupation.performer_role", terminology.OccupationCatalogs, "performer_role", 512},
		{"ImmunizationTerminology.vaccine_map", terminology.ImmunizationTerminologyCatalogs, "vaccine_map", 39},
		{"CompositionTerminology.category", terminology.CompositionTerminologyCatalogs, "category", 125},
		{"CompositionTerminology.section", terminology.CompositionTerminologyCatalogs, "section", 84},
		{"CompositionTerminology.type", terminology.CompositionTerminologyCatalogs, "type", 2},
	}
	for _, k := range kasus {
		isi, ok := k.katalog[k.kunci].(map[string]any)
		if !ok {
			t.Fatalf("%s: kunci %q tidak ada", k.nama, k.kunci)
		}
		if len(isi) != k.jumlah {
			t.Fatalf("%s: %d entri, harap %d", k.nama, len(isi), k.jumlah)
		}
	}
}

func TestMedicationCatalogHasFourteenGroups(t *testing.T) {
	if len(terminology.MedicationTerminologyCatalogs) != 14 {
		t.Fatalf("MedicationTerminology: %d grup, harap 14", len(terminology.MedicationTerminologyCatalogs))
	}
}
