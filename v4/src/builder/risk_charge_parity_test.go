package builder_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestRiskAssessmentMatchesPHPPayload(t *testing.T) {
	relatif := 1.5

	b := builder.NewRiskAssessmentBuilder()
	b.SetId("ra-1")
	b.SetStatus("final")
	b.SetSubject("Patient/1000000001")
	b.SetOccurrenceDateTime("2026-01-05T09:00:00+07:00")
	b.SetCondition(&datatype.Reference{Reference: "Condition/cond-1"})
	b.SetPerformer(&datatype.Reference{Reference: "Practitioner/1000000002"})
	b.AddReasonReference(&datatype.Reference{Reference: "Observation/obs-1"})
	b.AddBasis(&datatype.Reference{Reference: "Observation/obs-2"})
	b.AddPrediction(
		konsep("http://snomed.info/sct", "52101004", "Present"),
		0.7,
		konsep("http://terminology.hl7.org/CodeSystem/risk-probability", "moderate", "Moderate"),
		&relatif,
	)
	b.SetMitigation("Pemantauan tekanan darah tiap bulan.")
	b.AddExtension("http://example.org/risk", 3, "Integer")
	samakan(t, bacaGolden(t, "risk_assessment"), b)
}

func TestChargeItemDefinitionMatchesPHPPayload(t *testing.T) {
	referensi := "Organization/1000000001"
	tampil := "RS Sehat"
	waktu := "2026-01-05T09:00:00+07:00"

	b := builder.NewChargeItemDefinitionBuilder()
	b.SetId("cid-1")
	b.SetStatus("active")
	b.SetUrl("http://example.org/ChargeItemDefinition/cid-1")
	b.SetVersion("1.0.0")
	b.SetTitle("Tarif kunjungan")
	b.SetPublisher("RS Sehat")
	b.AddInstance(&referensi, &tampil, &waktu)
	b.AddProperty("http://example.org/prop", "kategori", "Rawat jalan")
	b.AddApplies("Location", "Location/3000000001")
	samakan(t, bacaGolden(t, "charge_item_definition"), b)
}

func TestChargeItemResponseMatchesPHPPayload(t *testing.T) {
	b := builder.NewChargeItemResponseBuilder()
	b.SetId("cir-1")
	b.SetStatus("active")
	b.SetChargeItem("ChargeItem/ci-1")
	b.SetPartOf("ChargeItem/ci-0")
	b.SetRequestor("Practitioner/1000000002")
	b.SetResponseDate("2026-01-05")
	b.AddOutcomeInfo("http://terminology.hl7.org/CodeSystem/chargeitem-status", "billable", "Billable")
	b.AddError("http://terminology.hl7.org/CodeSystem/chargeitem-error", "E1", "Tarif tidak ditemukan")
	b.AddCostCenterChargeItemResponse("Organization/1000000001", "RS Sehat")
	samakan(t, bacaGolden(t, "charge_item_response"), b)
}
