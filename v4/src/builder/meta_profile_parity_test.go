package builder_test

// Two profiles must survive. PHP pushes "meta/profile" so the list grows; a setter that writes the list
// outright keeps only the last profile and nothing complains.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestMetaProfileAccumulates(t *testing.T) {
	b := builder.NewPatientBuilder()
	b.SetId("Patient/1000000001")
	b.SetMetaProfile("http://example.org/p1")
	b.SetMetaProfile("http://example.org/p2")
	samakan(t, bacaGolden(t, "meta_profile_two"), b)
}

func TestGroupMetaProfileAccumulates(t *testing.T) {
	b := builder.NewGroupBuilder()
	b.SetId("g-1")
	b.SetMetaProfile("http://example.org/p1")
	b.SetMetaProfile("http://example.org/p2")

	meta, _ := b.Build()["meta"].(map[string]interface{})
	profil, _ := meta["profile"].([]interface{})
	if len(profil) != 2 {
		t.Fatalf("profile tertimpa: dapat %d entri, harap 2", len(profil))
	}
}
