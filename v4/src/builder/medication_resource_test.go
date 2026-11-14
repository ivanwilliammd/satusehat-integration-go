package builder_test

// Per-resource test, generated for parity with the PHP suite where every resource has its own
// test file. Asserts the builder exists and emits the correct resourceType.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestNewMedicationBuilderEmitsResourceType(t *testing.T) {
	payload := builder.NewMedicationBuilder().Build()
	if payload["resourceType"] != "Medication" {
		t.Fatalf("resourceType = %v, want Medication", payload["resourceType"])
	}
}
