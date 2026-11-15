package builder_test

// Golden coverage for six builders that had none: the payload shape is proven by running the PHP builder,
// not by inspection.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
)

func TestActivityDefinitionMatchesPHPPayload(t *testing.T) {
	b := builder.NewActivityDefinitionBuilder()
	b.SetId("ad-1")
	b.AddIdentifier("http://sys.id/ni", "NI-1")
	b.SetStatus("active")
	b.SetDescription("Edukasi gizi")
	b.SetKind("http://terminology.hl7.org/CodeSystem/activity-definition-kind", "ServiceRequest", "Service Request")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "EDU-1", "Edukasi")
	b.SetAuthor("Practitioner/1000000002", "dr. Budi")
	b.SetTiming("2026-01-05")
	samakan(t, bacaGolden(t, "activity_definition"), b)
}

func TestCapabilityStatementMatchesPHPPayload(t *testing.T) {
	b := builder.NewCapabilityStatementBuilder()
	b.SetId("cs-1")
	b.AddIdentifier("http://sys.id/ni", "NI-2")
	b.SetStatus("active")
	b.SetDate("2026-01-05")
	b.SetKind("http://hl7.org/fhir/capability-statement-kind", "instance", "Instance")
	b.SetFhirVersion("http://hl7.org/fhir/FHIR-version", "4.0.1", "4.0.1")
	b.SetFormat("http://hl7.org/fhir/restful-format", "json", "JSON")
	b.SetRest("http://hl7.org/fhir/restful-mode", "server", "Server")
	samakan(t, bacaGolden(t, "capability_statement"), b)
}

func TestCatalogEntryMatchesPHPPayload(t *testing.T) {
	b := builder.NewCatalogEntryBuilder()
	b.SetId("ce-1")
	b.AddIdentifier("http://sys.id/ni", "NI-3")
	b.SetStatus("active")
	b.SetType("http://terminology.hl7.org/CodeSystem/catalogType", "medication", "Medication")
	b.SetReference("Medication/med-1", "Amoksisilin")
	b.SetValidityPeriod("2026-01-05/2026-12-31")
	samakan(t, bacaGolden(t, "catalog_entry"), b)
}

func TestDeviceMetricMatchesPHPPayload(t *testing.T) {
	b := builder.NewDeviceMetricBuilder()
	b.SetId("dm-1")
	b.AddIdentifier("http://sys.id/ni", "NI-4")
	b.SetStatus("active")
	b.SetType("http://terminology.hl7.org/CodeSystem/device-metric-type", "heart-rate", "Heart rate")
	b.SetCategory("http://terminology.hl7.org/CodeSystem/metric-category", "measurement", "Measurement")
	b.SetMeasurementPeriod("2026-01-05")
	samakan(t, bacaGolden(t, "device_metric"), b)
}

func TestDocumentManifestMatchesPHPPayload(t *testing.T) {
	b := builder.NewDocumentManifestBuilder()
	b.SetId("dmf-1")
	b.AddIdentifier("http://sys.id/ni", "NI-5")
	b.SetStatus("current")
	b.SetType("http://loinc.org", "34133-9", "Summary of episode note")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCreated("2026-01-05")
	b.SetSource("Practitioner/1000000002", "dr. Budi")
	samakan(t, bacaGolden(t, "document_manifest"), b)
}

func TestEnrollmentResponseMatchesPHPPayload(t *testing.T) {
	b := builder.NewEnrollmentResponseBuilder()
	b.SetId("er-1")
	b.AddIdentifier("http://sys.id/ni", "NI-6")
	b.SetStatus("active")
	b.SetRequest("EnrollmentRequest/erq-1", "Pendaftaran")
	b.SetOutcome("http://terminology.hl7.org/CodeSystem/remittance-outcome", "complete", "Complete")
	b.SetDisposition("Disetujui")
	b.SetCreated("2026-01-05")
	samakan(t, bacaGolden(t, "enrollment_response"), b)
}
