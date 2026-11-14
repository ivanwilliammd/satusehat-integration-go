package builder_test

// Date, note, description, outcome and request setters — the shapes the canonical builder writes, checked
// against payloads produced by running it.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestPaymentReconciliationOutcomeMatchesPHPPayload(t *testing.T) {
	b := builder.NewPaymentReconciliationBuilder()
	b.SetId("pr-1")
	b.SetOutcome("http://terminology.hl7.org/CodeSystem/remittance-outcome", "complete", "Complete")
	b.SetRequest("Claim/claim-1")
	samakan(t, bacaGolden(t, "payment_reconciliation"), b)
}

func TestClinicalImpressionDateAndNoteMatchPHPPayload(t *testing.T) {
	b := builder.NewClinicalImpressionBuilder()
	b.SetId("ci-1")
	b.SetDate("2026-01-05")
	b.AddNote("Pasien stabil.")
	samakan(t, bacaGolden(t, "clinical_impression"), b)
}

func TestInvoiceDateAndNoteMatchPHPPayload(t *testing.T) {
	b := builder.NewInvoiceBuilder()
	b.SetId("inv-1")
	b.SetDate("2026-01-05T11:00:00+07:00")
	b.AddNote("Tagihan rawat jalan.")
	samakan(t, bacaGolden(t, "invoice"), b)
}
