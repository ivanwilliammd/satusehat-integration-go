package builder_test

// These builders share the generic surface — id, identifier, status, subject, code, encounter — and none had
// a golden. One generated test per resource, one script that produced every golden by running the PHP
// builder. The repetition is deliberate: these are the resources nobody looks at twice.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestAuditEventMatchesPHPPayload(t *testing.T) {
	b := builder.NewAuditEventBuilder()
	b.SetId("audit-event-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-100"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "audit_event"), b)
}

func TestBasicMatchesPHPPayload(t *testing.T) {
	b := builder.NewBasicBuilder()
	b.SetId("basic-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-101"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "basic"), b)
}

func TestBinaryMatchesPHPPayload(t *testing.T) {
	b := builder.NewBinaryBuilder()
	b.SetId("binary-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-102"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "binary"), b)
}

func TestBiologicallyDerivedProductMatchesPHPPayload(t *testing.T) {
	b := builder.NewBiologicallyDerivedProductBuilder()
	b.SetId("biologically-derived-product-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-103"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "biologically_derived_product"), b)
}

func TestBodyStructureMatchesPHPPayload(t *testing.T) {
	b := builder.NewBodyStructureBuilder()
	b.SetId("body-structure-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-104"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "body_structure"), b)
}

func TestContractMatchesPHPPayload(t *testing.T) {
	b := builder.NewContractBuilder()
	b.SetId("contract-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-105"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "contract"), b)
}

func TestDeviceDefinitionMatchesPHPPayload(t *testing.T) {
	b := builder.NewDeviceDefinitionBuilder()
	b.SetId("device-definition-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-106"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "device_definition"), b)
}

func TestDeviceRequestMatchesPHPPayload(t *testing.T) {
	b := builder.NewDeviceRequestBuilder()
	b.SetId("device-request-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-107"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "device_request"), b)
}

func TestDeviceUseStatementMatchesPHPPayload(t *testing.T) {
	b := builder.NewDeviceUseStatementBuilder()
	b.SetId("device-use-statement-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-108"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "device_use_statement"), b)
}

func TestEffectEvidenceSynthesisMatchesPHPPayload(t *testing.T) {
	b := builder.NewEffectEvidenceSynthesisBuilder()
	b.SetId("effect-evidence-synthesis-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-109"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "effect_evidence_synthesis"), b)
}

func TestEnrollmentRequestMatchesPHPPayload(t *testing.T) {
	b := builder.NewEnrollmentRequestBuilder()
	b.SetId("enrollment-request-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-110"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "enrollment_request"), b)
}

func TestEventDefinitionMatchesPHPPayload(t *testing.T) {
	b := builder.NewEventDefinitionBuilder()
	b.SetId("event-definition-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-111"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "event_definition"), b)
}

func TestEvidenceMatchesPHPPayload(t *testing.T) {
	b := builder.NewEvidenceBuilder()
	b.SetId("evidence-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-112"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "evidence"), b)
}

func TestEvidenceVariableMatchesPHPPayload(t *testing.T) {
	b := builder.NewEvidenceVariableBuilder()
	b.SetId("evidence-variable-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-113"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "evidence_variable"), b)
}

func TestExampleScenarioMatchesPHPPayload(t *testing.T) {
	b := builder.NewExampleScenarioBuilder()
	b.SetId("example-scenario-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-114"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "example_scenario"), b)
}

func TestFlagMatchesPHPPayload(t *testing.T) {
	b := builder.NewFlagBuilder()
	b.SetId("flag-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-115"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "flag"), b)
}

func TestGuidanceResponseMatchesPHPPayload(t *testing.T) {
	b := builder.NewGuidanceResponseBuilder()
	b.SetId("guidance-response-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-116"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "guidance_response"), b)
}

func TestImagingManifestMatchesPHPPayload(t *testing.T) {
	b := builder.NewImagingManifestBuilder()
	b.SetId("imaging-manifest-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-117"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "imaging_manifest"), b)
}

func TestImagingSelectionMatchesPHPPayload(t *testing.T) {
	b := builder.NewImagingSelectionBuilder()
	b.SetId("imaging-selection-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-118"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "imaging_selection"), b)
}

func TestImmunizationEvaluationMatchesPHPPayload(t *testing.T) {
	b := builder.NewImmunizationEvaluationBuilder()
	b.SetId("immunization-evaluation-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-119"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "immunization_evaluation"), b)
}

func TestImmunizationRecommendationMatchesPHPPayload(t *testing.T) {
	b := builder.NewImmunizationRecommendationBuilder()
	b.SetId("immunization-recommendation-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-120"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "immunization_recommendation"), b)
}

func TestLibraryMatchesPHPPayload(t *testing.T) {
	b := builder.NewLibraryBuilder()
	b.SetId("library-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-121"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "library"), b)
}

func TestLinkageMatchesPHPPayload(t *testing.T) {
	b := builder.NewLinkageBuilder()
	b.SetId("linkage-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-122"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "linkage"), b)
}

func TestMeasureMatchesPHPPayload(t *testing.T) {
	b := builder.NewMeasureBuilder()
	b.SetId("measure-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-123"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "measure"), b)
}

func TestMeasureReportMatchesPHPPayload(t *testing.T) {
	b := builder.NewMeasureReportBuilder()
	b.SetId("measure-report-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-124"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "measure_report"), b)
}

func TestMediaMatchesPHPPayload(t *testing.T) {
	b := builder.NewMediaBuilder()
	b.SetId("media-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-125"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "media"), b)
}

func TestMessageDefinitionMatchesPHPPayload(t *testing.T) {
	b := builder.NewMessageDefinitionBuilder()
	b.SetId("message-definition-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-126"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "message_definition"), b)
}

func TestMessageHeaderMatchesPHPPayload(t *testing.T) {
	b := builder.NewMessageHeaderBuilder()
	b.SetId("message-header-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-127"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "message_header"), b)
}

func TestOperationOutcomeMatchesPHPPayload(t *testing.T) {
	b := builder.NewOperationOutcomeBuilder()
	b.SetId("operation-outcome-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-128"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "operation_outcome"), b)
}

func TestPersonMatchesPHPPayload(t *testing.T) {
	b := builder.NewPersonBuilder()
	b.SetId("person-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-129"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "person"), b)
}

func TestPlanDefinitionMatchesPHPPayload(t *testing.T) {
	b := builder.NewPlanDefinitionBuilder()
	b.SetId("plan-definition-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-130"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "plan_definition"), b)
}

func TestProvenanceMatchesPHPPayload(t *testing.T) {
	b := builder.NewProvenanceBuilder()
	b.SetId("provenance-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-131"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "provenance"), b)
}

func TestQuestionnaireMatchesPHPPayload(t *testing.T) {
	b := builder.NewQuestionnaireBuilder()
	b.SetId("questionnaire-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-132"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "questionnaire"), b)
}

func TestRequestGroupMatchesPHPPayload(t *testing.T) {
	b := builder.NewRequestGroupBuilder()
	b.SetId("request-group-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-133"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "request_group"), b)
}

func TestResearchDefinitionMatchesPHPPayload(t *testing.T) {
	b := builder.NewResearchDefinitionBuilder()
	b.SetId("research-definition-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-134"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "research_definition"), b)
}

func TestResearchElementDefinitionMatchesPHPPayload(t *testing.T) {
	b := builder.NewResearchElementDefinitionBuilder()
	b.SetId("research-element-definition-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-135"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "research_element_definition"), b)
}

func TestResearchSubjectMatchesPHPPayload(t *testing.T) {
	b := builder.NewResearchSubjectBuilder()
	b.SetId("research-subject-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-136"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "research_subject"), b)
}

func TestRiskEvidenceSynthesisMatchesPHPPayload(t *testing.T) {
	b := builder.NewRiskEvidenceSynthesisBuilder()
	b.SetId("risk-evidence-synthesis-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-137"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "risk_evidence_synthesis"), b)
}

func TestScheduleMatchesPHPPayload(t *testing.T) {
	b := builder.NewScheduleBuilder()
	b.SetId("schedule-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-138"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "schedule"), b)
}

func TestSearchParameterMatchesPHPPayload(t *testing.T) {
	b := builder.NewSearchParameterBuilder()
	b.SetId("search-parameter-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-139"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "search_parameter"), b)
}

func TestSlotMatchesPHPPayload(t *testing.T) {
	b := builder.NewSlotBuilder()
	b.SetId("slot-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-140"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "slot"), b)
}

func TestSubscriptionMatchesPHPPayload(t *testing.T) {
	b := builder.NewSubscriptionBuilder()
	b.SetId("subscription-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-141"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "subscription"), b)
}

func TestSubstanceNucleicAcidMatchesPHPPayload(t *testing.T) {
	b := builder.NewSubstanceNucleicAcidBuilder()
	b.SetId("substance-nucleic-acid-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-142"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "substance_nucleic_acid"), b)
}

func TestSubstancePolymerMatchesPHPPayload(t *testing.T) {
	b := builder.NewSubstancePolymerBuilder()
	b.SetId("substance-polymer-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-143"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "substance_polymer"), b)
}

func TestSubstanceProteinMatchesPHPPayload(t *testing.T) {
	b := builder.NewSubstanceProteinBuilder()
	b.SetId("substance-protein-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-144"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "substance_protein"), b)
}

func TestSubstanceSourceMaterialMatchesPHPPayload(t *testing.T) {
	b := builder.NewSubstanceSourceMaterialBuilder()
	b.SetId("substance-source-material-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-145"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "substance_source_material"), b)
}

func TestSubstanceSpecificationMatchesPHPPayload(t *testing.T) {
	b := builder.NewSubstanceSpecificationBuilder()
	b.SetId("substance-specification-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-146"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "substance_specification"), b)
}

func TestSupplyDeliveryMatchesPHPPayload(t *testing.T) {
	b := builder.NewSupplyDeliveryBuilder()
	b.SetId("supply-delivery-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-147"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "supply_delivery"), b)
}

func TestSupplyRequestMatchesPHPPayload(t *testing.T) {
	b := builder.NewSupplyRequestBuilder()
	b.SetId("supply-request-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-148"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "supply_request"), b)
}

func TestTestReportMatchesPHPPayload(t *testing.T) {
	b := builder.NewTestReportBuilder()
	b.SetId("test-report-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-149"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "test_report"), b)
}

func TestTestScriptMatchesPHPPayload(t *testing.T) {
	b := builder.NewTestScriptBuilder()
	b.SetId("test-script-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-150"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "test_script"), b)
}

func TestVerificationResultMatchesPHPPayload(t *testing.T) {
	b := builder.NewVerificationResultBuilder()
	b.SetId("verification-result-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-151"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "verification_result"), b)
}

func TestVisionPrescriptionMatchesPHPPayload(t *testing.T) {
	b := builder.NewVisionPrescriptionBuilder()
	b.SetId("vision-prescription-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-152"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "vision_prescription"), b)
}
