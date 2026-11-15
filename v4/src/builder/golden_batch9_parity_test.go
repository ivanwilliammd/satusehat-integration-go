package builder_test

// The last three builders without a golden. With these, every one of the 141 builders has a payload compared
// against the canonical PHP builder.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestBillingStatusMatchesPHPPayload(t *testing.T) {
	b := builder.NewBillingStatusBuilder()
	b.SetId("bs-1")
	b.AddIdentifier("http://sys.id/ni", "NI-90")
	b.SetStatus("active")
	b.SetInsurer("Organization/1000000001", "BPJS")
	b.SetRecipient("Organization/1000000002", "RS Sehat")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetRequest("CoverageEligibilityRequest/cer-1", "Klaim")
	samakan(t, bacaGolden(t, "billing_status"), b)
}

func TestPurificationDecisionMatchesPHPPayload(t *testing.T) {
	b := builder.NewPurificationDecisionBuilder()
	b.SetId("pd-1")
	b.AddIdentifier("http://sys.id/ni", "NI-91")
	b.SetStatus("approved", "Approved", "http://terminology.hl7.org/CodeSystem/decision")
	b.SetInsurer("Organization/1000000001", "BPJS")
	b.SetProvider("Organization/1000000002", "RS Sehat")
	b.SetClaimResponse("ClaimResponse/cr-1", "Respons")
	b.SetCreated("2026-01-05")
	samakan(t, bacaGolden(t, "purification_decision"), b)
}

func TestResourceGuideMatchesPHPPayload(t *testing.T) {
	b := builder.NewResourceGuideBuilder()
	b.SetId("rg-1")
	b.AddIdentifier("http://sys.id/ni", "NI-92", "official", "MC", "Medical Record Number")
	b.SetStatus("active")
	b.SetName("Panduan")
	b.SetDescription("Panduan implementasi")
	b.SetVersion("1.0.0")
	b.SetPublisher("Kemenkes")
	samakan(t, bacaGolden(t, "resource_guide"), b)
}
