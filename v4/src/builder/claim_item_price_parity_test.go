package builder_test

// Item-level amounts travel through nested paths (item/0/quantity/value). The golden pins that they land
// inside the list element, not under a literal "0" key.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestClaimItemPricesMatchPHPPayload(t *testing.T) {
	b := builder.NewClaimBuilder()
	b.SetId("claim-1")
	b.AddItem(1, "http://hl7.org/fhir/sid/icd-9-cm", "88.72", "Diagnostic ultrasound of heart", nil, nil, nil, nil)
	b.AddItemQuantity(0, 2, "IDR")
	b.AddItemUnitPrice(0, 250000, "IDR")
	b.AddItemServicedPeriod(0, "2026-01-05", "2026-01-06")
	samakan(t, bacaGolden(t, "claim_item_prices"), b)
}
