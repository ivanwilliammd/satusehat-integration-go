package builder_test

// Claim lines: the item carries its service coding and the optional sequence links, which PHP writes as
// single-element arrays rather than bare integers.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestClaimItemsMatchPHPPayload(t *testing.T) {
	diagnosis, procedure := 1, 1
	b := builder.NewClaimBuilder()
	b.SetId("claim-1")
	b.SetPatient("Patient/1000000001", "Budi Santoso")
	b.SetEnterer("Practitioner/1000000002")
	b.SetProvider("Organization/1000000001")
	b.AddItem(1, "http://hl7.org/fhir/sid/icd-9-cm", "88.72", "Diagnostic ultrasound of heart", nil, &diagnosis, &procedure, nil)
	samakan(t, bacaGolden(t, "claim_items"), b)
}
