package builder_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestChargeItemMatchesPHPPayload(t *testing.T) {
	sistem, kode := "http://terminology.hl7.org/CodeSystem/v3-ParticipationType", "ATND"
	b := builder.NewChargeItemBuilder()
	b.SetId("ci-1")
	b.SetStatus("billable")
	b.SetSubject("Patient/1000000001")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/charge-item", "CI-001", "Tindakan")
	b.SetEncounter("Encounter/2000000001")
	b.SetPerformedDate("2026-01-05T09:00:00+07:00")
	b.SetPerformerActor("Practitioner/1000000002", &sistem, &kode)
	b.SetOrganization("Organization/1000000001")
	b.SetQuantity(1)
	b.SetPriceOverride(250000, "IDR")
	b.SetEnteredDate("2026-01-05T09:30:00+07:00")
	b.SetAccount("Account/1500000001")
	samakan(t, bacaGolden(t, "charge_item"), b)
}
