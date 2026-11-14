package builder_test

// Per-resource test, generated for parity with the PHP suite where every resource has its own
// test file. Asserts the builder exists and emits the correct resourceType.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestNewSupplyRequestBuilderEmitsResourceType(t *testing.T) {
	payload := builder.NewSupplyRequestBuilder().Build()
	if payload["resourceType"] != "SupplyRequest" {
		t.Fatalf("resourceType = %v, want SupplyRequest", payload["resourceType"])
	}
}
