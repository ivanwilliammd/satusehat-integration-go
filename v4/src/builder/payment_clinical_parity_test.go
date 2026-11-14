package builder_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestPaymentReconciliationFullSurfaceMatchesPHPPayload(t *testing.T) {
	tipeCatatan := "display"
	jumlah := 250000.0
	klaim := "Claim/claim-1"

	b := builder.NewPaymentReconciliationBuilder()
	b.SetId("pr-1")
	b.SetPeriod("2026-01-01", "2026-01-31")
	b.SetRequestProvider("Organization/1000000001")
	b.SetRequestMatchDate("2026-01-05")
	b.SetOutcomeCode("http://terminology.hl7.org/CodeSystem/remittance-outcome", "complete", "Complete")
	b.AddRequestor("Organization/1000000002", "RS Sehat")
	b.AddProcessNote("Pembayaran sesuai klaim.", &tipeCatatan)
	b.SetPaymentDate("2026-01-10")
	b.SetPaymentAmount(250000, "IDR")
	b.SetPaymentIdentifier("http://example.org/payment", "PAY-1")
	b.AddDetail("http://terminology.hl7.org/CodeSystem/adjudication", "eligible", &jumlah, &klaim)
	b.SetTotalAmount(250000, "IDR")
	samakan(t, bacaGolden(t, "payment_reconciliation_full"), b)
}

func TestClinicalImpressionFullSurfaceMatchesPHPPayload(t *testing.T) {
	b := builder.NewClinicalImpressionBuilder()
	b.SetId("ci-1")
	b.SetStatus("completed")
	b.SetSubject("Patient/1000000001")
	b.SetEffectiveDateTime("2026-01-05T09:00:00+07:00")
	b.SetAssessor(&datatype.Reference{Reference: "Practitioner/1000000002"})
	b.SetPreviousOpinion(&datatype.Reference{Reference: "ClinicalImpression/ci-0"})
	b.AddFindingCodeableConcept(konsep("http://snomed.info/sct", "38341003", "Hypertension"))
	b.AddFindingReference(&datatype.Reference{Reference: "Observation/obs-1"})
	b.AddPrognosisCodeableConcept(konsep("http://snomed.info/sct", "170968001", "Prognosis good"))
	b.AddSupportingInfo(&datatype.Reference{Reference: "Observation/obs-2"})
	b.AddInvestigation(map[string]interface{}{
		"code": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": "http://terminology.hl7.org/CodeSystem/investigation-sets", "code": "exam"}},
		},
		"item": []interface{}{map[string]interface{}{"display": "Pemeriksaan fisik"}},
	})
	samakan(t, bacaGolden(t, "clinical_impression_full"), b)
}
