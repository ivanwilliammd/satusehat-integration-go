package builder_test

// Per-resource test, generated for parity with the PHP suite where every resource has its own
// test file. Asserts the builder exists and emits the correct resourceType.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestNewEvidenceBuilderEmitsResourceType(t *testing.T) {
	payload := builder.NewEvidenceBuilder().Build()
	if payload["resourceType"] != "Evidence" {
		t.Fatalf("resourceType = %v, want Evidence", payload["resourceType"])
	}
}
