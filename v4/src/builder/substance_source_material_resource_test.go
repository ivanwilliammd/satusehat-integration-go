package builder_test

// Per-resource test, generated for parity with the PHP suite where every resource has its own
// test file. Asserts the builder exists and emits the correct resourceType.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestNewSubstanceSourceMaterialBuilderEmitsResourceType(t *testing.T) {
	payload := builder.NewSubstanceSourceMaterialBuilder().Build()
	if payload["resourceType"] != "SubstanceSourceMaterial" {
		t.Fatalf("resourceType = %v, want SubstanceSourceMaterial", payload["resourceType"])
	}
}
