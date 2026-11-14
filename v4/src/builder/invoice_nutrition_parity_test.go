package builder_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestInvoiceFullSurfaceMatchesPHPPayload(t *testing.T) {
	jumlah, satuan, faktor := 1.0, 250000.0, 1.0

	b := builder.NewInvoiceBuilder()
	b.SetId("inv-1")
	b.SetStatus("issued")
	b.SetCustomer("Patient/1000000001")
	b.SetIssuer("Organization/1000000001")
	b.SetAccount("Account/1500000001")
	b.SetParticipant("http://terminology.hl7.org/CodeSystem/v3-ParticipationType", "PCP", "Practitioner/1000000002", "dr. Budi")
	b.AddLineItem(1, "http://terminology.kemkes.go.id/CodeSystem/invoice-item", "VISIT", "Kunjungan", &jumlah, &satuan, &faktor)
	b.SetTotalGross(250000, "IDR")
	b.SetTotalNet(225000, "IDR")
	b.SetTotalVAT(25000, "IDR")
	b.SetPaymentTerms("Bayar dalam 14 hari.")
	b.SetCancelledReason("")
	samakan(t, bacaGolden(t, "invoice_full"), b)
}

func TestNutritionOrderMatchesPHPPayload(t *testing.T) {
	timing := (&datatype.Timing{}).SetRepeat(&datatype.TimingRepeat{Bounds: map[string]interface{}{"start": "2026-01-05"}})

	b := builder.NewNutritionOrderBuilder()
	b.SetId("no-1")
	b.SetStatus("active")
	b.SetIntent("order")
	b.SetPatient(&datatype.Reference{Reference: "Patient/1000000001"})
	b.SetDateTime("2026-01-05T09:00:00+07:00")
	b.SetOrderer(&datatype.Reference{Reference: "Practitioner/1000000002"})
	b.AddAllergyType(konsep("http://snomed.info/sct", "91936005", "Allergy to penicillin"))
	b.AddFoodPreferenceModifier(konsep("http://terminology.hl7.org/CodeSystem/diet", "vegetarian", "Vegetarian"))
	b.AddExcludeFoodModifier(konsep("http://snomed.info/sct", "227493005", "Cashew nut"))
	b.SetOralDietType(konsep("http://snomed.info/sct", "151000004", "Low salt diet"))
	b.AddOralDietSchedule(timing)
	b.AddSupplement(map[string]interface{}{
		"type": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": "http://snomed.info/sct", "code": "442901000124106"}},
		},
		"productName": "Suplemen A",
	})
	b.SetEnteralFormula(map[string]interface{}{
		"baseFormulaType": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": "http://snomed.info/sct", "code": "443031000124106"}},
		},
	})
	samakan(t, bacaGolden(t, "nutrition_order"), b)
}
