package builder_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestExplanationOfBenefitMatchesPHPPayload(t *testing.T) {
	b := builder.NewExplanationOfBenefitBuilder()
	b.SetId("eob-1")
	b.AddIdentifier("http://sys.id/ni", "NI-7")
	b.SetStatus("active")
	b.SetType("http://terminology.hl7.org/CodeSystem/claim-type", "institutional", "Institutional")
	b.SetUse("http://terminology.hl7.org/CodeSystem/claim-use", "claim", "Claim")
	b.SetPatient("Patient/1000000001", "Budi Santoso")
	b.SetCreated("2026-01-05")
	b.SetProvider("Organization/1000000001", "RS Sehat")
	samakan(t, bacaGolden(t, "explanation_of_benefit"), b)
}

func TestInsurancePlanMatchesPHPPayload(t *testing.T) {
	b := builder.NewInsurancePlanBuilder()
	b.SetId("ip-1")
	b.AddIdentifier("http://sys.id/ni", "NI-8")
	b.SetStatus("active")
	b.SetName("BPJS Kesehatan")
	b.SetType("http://terminology.hl7.org/CodeSystem/insurance-plan-type", "medical", "Medical")
	b.SetAdministeredBy("Organization/1000000001", "BPJS")
	b.SetCoverage("http://terminology.hl7.org/CodeSystem/coverage-type", "institutional", "Institutional")
	samakan(t, bacaGolden(t, "insurance_plan"), b)
}

func TestMedicationKnowledgeMatchesPHPPayload(t *testing.T) {
	b := builder.NewMedicationKnowledgeBuilder()
	b.SetId("mk-1")
	b.AddIdentifier("http://sys.id/ni", "NI-9")
	b.SetStatus("active")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/medication", "MED-1", "Amoksisilin")
	b.SetManufacturer("Organization/1000000001", "Kimia Farma")
	b.SetDoseForm("http://snomed.info/sct", "385055001", "Tablet")
	samakan(t, bacaGolden(t, "medication_knowledge"), b)
}

func TestObservationDefinitionMatchesPHPPayload(t *testing.T) {
	b := builder.NewObservationDefinitionBuilder()
	b.SetId("od-1")
	b.AddIdentifier("http://sys.id/ni", "NI-10")
	b.SetStatus("active")
	b.SetCode("http://loinc.org", "8867-4", "Heart rate")
	b.SetCategory("http://terminology.hl7.org/CodeSystem/observation-category", "vital-signs", "Vital Signs")
	b.SetMethod("http://snomed.info/sct", "37931006", "Auscultation")
	samakan(t, bacaGolden(t, "observation_definition"), b)
}

func TestOrganizationAffiliationMatchesPHPPayload(t *testing.T) {
	b := builder.NewOrganizationAffiliationBuilder()
	b.SetId("oa-1")
	b.AddIdentifier("http://sys.id/ni", "NI-11")
	b.SetStatus("active")
	b.SetOrganization("Organization/1000000001", "RS Sehat")
	b.SetParticipatingOrganization("Organization/1000000002", "Klinik Sehat")
	b.SetNetwork("http://terminology.hl7.org/CodeSystem/network-type", "nt1", "Network 1")
	b.SetCode("http://terminology.hl7.org/CodeSystem/organization-role", "provider", "Provider")
	samakan(t, bacaGolden(t, "organization_affiliation"), b)
}

func TestResearchStudyMatchesPHPPayload(t *testing.T) {
	b := builder.NewResearchStudyBuilder()
	b.SetId("rs-1")
	b.AddIdentifier("http://sys.id/ni", "NI-12")
	b.SetStatus("active")
	b.SetTitle("Studi Hipertensi")
	b.SetProtocol("PlanDefinition/pd-1", "Protokol")
	b.SetSponsor("Organization/1000000001", "FKUI")
	b.SetPhase("http://terminology.hl7.org/CodeSystem/research-study-phase", "phase-2", "Phase 2")
	samakan(t, bacaGolden(t, "research_study"), b)
}

func TestInsurancePlanKeepsEveryType(t *testing.T) {
	b := builder.NewInsurancePlanBuilder()
	b.SetId("ip-2")
	b.SetType("http://terminology.hl7.org/CodeSystem/insurance-plan-type", "medical", "Medical")
	b.SetType("http://terminology.hl7.org/CodeSystem/insurance-plan-type", "dental", "Dental")
	samakan(t, bacaGolden(t, "insurance_plan_two_types"), b)
}
