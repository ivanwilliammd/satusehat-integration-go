package builder_test

// Per-resource test, generated for parity with the PHP suite where every resource has its own
// test file. Asserts the builder exists and emits the correct resourceType.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestNewMedicinalProductContraindicationBuilderEmitsResourceType(t *testing.T) {
	payload := builder.NewMedicinalProductContraindicationBuilder().Build()
	if payload["resourceType"] != "MedicinalProductContraindication" {
		t.Fatalf("resourceType = %v, want MedicinalProductContraindication", payload["resourceType"])
	}
}
