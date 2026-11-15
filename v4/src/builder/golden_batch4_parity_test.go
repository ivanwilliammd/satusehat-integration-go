package builder_test

// Six builders share one generic surface: id, identifier, status, subject, code, encounter. They had no
// goldens, and the shared shape is precisely why one test file settles all six.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func ident(system, value string) *datatype.Identifier {
	return &datatype.Identifier{System: system, Value: value}
}

func TestAppointmentMatchesPHPPayload(t *testing.T) {
	b := builder.NewAppointmentBuilder()
	b.SetId("appointment-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-30"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "appointment"), b)
}

func TestAppointmentResponseMatchesPHPPayload(t *testing.T) {
	b := builder.NewAppointmentResponseBuilder()
	b.SetId("appointmentresponse-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-31"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "appointment_response"), b)
}

func TestCommunicationMatchesPHPPayload(t *testing.T) {
	b := builder.NewCommunicationBuilder()
	b.SetId("communication-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-32"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "communication"), b)
}

func TestCommunicationRequestMatchesPHPPayload(t *testing.T) {
	b := builder.NewCommunicationRequestBuilder()
	b.SetId("communicationrequest-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-33"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "communication_request"), b)
}

func TestDetectedIssueMatchesPHPPayload(t *testing.T) {
	b := builder.NewDetectedIssueBuilder()
	b.SetId("detectedissue-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-34"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "detected_issue"), b)
}

func TestConsentMatchesPHPPayload(t *testing.T) {
	b := builder.NewConsentBuilder()
	b.SetId("consent-1")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-35"))
	b.SetStatus("active")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetCode("http://terminology.kemkes.go.id/CodeSystem/activity", "ACT-1", "Konsultasi")
	b.SetEncounter("Encounter/enc-1")
	samakan(t, bacaGolden(t, "consent"), b)
}
