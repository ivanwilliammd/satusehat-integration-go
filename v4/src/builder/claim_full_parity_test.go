package builder_test

// The whole Claim and ClaimResponse surface as the PHP builder writes it: care team, diagnosis, procedure,
// nested insurance account, billable period, and on the response side the adjudication list that lives under
// "item", the payment block and the nested totals.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestClaimFullSurfaceMatchesPHPPayload(t *testing.T) {
	peranSistem, peranKode := "http://terminology.hl7.org/CodeSystem/v3-ParticipationType", "ATND"
	peranDisplay := "attender"
	diagnosisTipe, diagnosisKode := "http://terminology.hl7.org/CodeSystem/ex-diagnosistype", "principal"
	b := builder.NewClaimBuilder()
	b.SetId("claim-1")
	b.SetUse("claim")
	b.SetStatus("active")
	b.SetPatient("Patient/1000000001", "Budi Santoso")
	b.SetType("http://terminology.hl7.org/CodeSystem/claim-type", "institutional", "Institutional")
	b.SetSubType("http://terminology.hl7.org/CodeSystem/ex-claimsubtype", "ip", "Inpatient")
	b.SetPriority("http://terminology.hl7.org/CodeSystem/processpriority", "normal", "Normal")
	b.SetFacility("Organization/1000000001", "RS Sehat")
	b.SetPrescriber("Practitioner/1000000002")
	b.SetPriorityPayor("Organization/1000000003")
	b.AddCareTeam(1, "Practitioner/1000000002", &peranSistem, &peranKode, &peranDisplay)
	b.AddDiagnosis(1, "http://hl7.org/fhir/sid/icd-10", "I10", &diagnosisTipe, &diagnosisKode)
	tanggal := "2026-01-05"
	b.AddProcedure(1, "http://hl7.org/fhir/sid/icd-9-cm", "88.72", &tanggal, nil, nil)
	b.AddInsuranceAccount("Account/1500000001")
	b.SetBillablePeriod("2026-01-05", "2026-01-06")
	b.SetTotal(250000, "IDR")
	samakan(t, bacaGolden(t, "claim_full"), b)
}

func TestClaimResponseFullSurfaceMatchesPHPPayload(t *testing.T) {
	referensi := "Payment/1"
	nilai := 250000.0

	b := builder.NewClaimResponseBuilder()
	b.SetId("cr-1")
	b.SetUse("claim")
	b.SetOutcome("complete")
	b.SetDisposition("Selesai")
	b.SetPayeeType("http://terminology.hl7.org/CodeSystem/payeetype", "provider", "Provider")
	b.SetRequestor("Organization/1000000001")
	b.SetPreAuthRef("PA-1")
	b.SetPreAuthPeriod("2026-01-01", "2026-01-31")
	b.SetTotal(250000, "http://terminology.hl7.org/CodeSystem/adjudication", "submitted")
	b.AddItemAdjudication(1, "http://terminology.hl7.org/CodeSystem/adjudication", "eligible", &nilai, nil, nil)
	b.SetPayment(250000, "http://terminology.hl7.org/CodeSystem/payment-type", "payment", "2026-01-10", &referensi)
	b.AddAddItem(1, "http://hl7.org/fhir/sid/icd-10", "I10", "Hipertensi")
	b.AddError(1, "http://terminology.hl7.org/CodeSystem/adjudication-error", "16", "Kode tidak valid")
	samakan(t, bacaGolden(t, "claim_response"), b)
}
