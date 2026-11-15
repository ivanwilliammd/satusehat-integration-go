package builder_test

// CoverageEligibilityRequest and QuestionnaireResponse are the first two of the eleven builders with methods
// of their own, and both were proven only by tests that predate the goldens.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestCoverageEligibilityRequestMatchesPHPPayload(t *testing.T) {
	b := builder.NewCoverageEligibilityRequestBuilder()
	b.SetId("cer-2")
	b.AddIdentifier(ident("http://sys.id/ni", "NI-60"))
	b.SetStatus("active")
	b.SetPatient("Patient/1000000001", "Budi Santoso")
	b.SetInsurer("Organization/1000000001")
	b.SetCreated("2026-01-05")
	b.AddCoverage("Coverage/cov-1", "PRE-1")
	b.SetEnterer("Practitioner/1000000002")
	b.SetProvider("Organization/1000000001")
	b.AddItem("http://terminology.kemkes.go.id/CodeSystem/lab", "LAB-1", "Darah Lengkap",
		"http://terminology.hl7.org/CodeSystem/ex-benefitcategory", "1")
	b.SetPriority("http://terminology.hl7.org/CodeSystem/processpriority", "normal", "Normal")
	b.SetPurpose([]string{"validation", "benefits"})
	b.SetServicedDate("2026-01-05")
	samakan(t, bacaGolden(t, "coverage_eligibility_request"), b)
}

func TestQuestionnaireResponseMatchesPHPPayload(t *testing.T) {
	cc := &datatype.CodeableConcept{
		Coding: []datatype.Coding{{System: "http://snomed.info/sct", Code: "25064002", Display: "Headache"}},
	}
	b := builder.NewQuestionnaireResponseBuilder()
	b.SetId("qr-1")
	b.SetStatus("completed")
	b.SetQuestionnaire("Questionnaire/q-1")
	b.SetAuthored("2026-01-05")
	b.SetAuthor(&datatype.Reference{Reference: "Practitioner/1000000002"})
	b.SetSource(&datatype.Reference{Reference: "Patient/1000000001"})
	b.SetSubject("Patient/1000000001")
	b.SetEncounter(&datatype.Reference{Reference: "Encounter/enc-1"})
	b.AddItem("1", "Keluhan utama", cc)
	samakan(t, bacaGolden(t, "questionnaire_response"), b)
}

func TestQuestionnaireResponseDefaultsStatusLikePHP(t *testing.T) {
	b := builder.NewQuestionnaireResponseBuilder()
	b.SetId("qr-2")
	b.SetQuestionnaire("Questionnaire/q-1")
	samakan(t, bacaGolden(t, "questionnaire_response_default_status"), b)
}

func TestQuestionnaireResponseRefusesUnknownStatus(t *testing.T) {
	b := builder.NewQuestionnaireResponseBuilder()
	b.SetId("qr-3")
	b.SetStatus("finished")
	if b.Err() == nil {
		t.Fatal("status finished diterima; kanon menolak")
	}
	if _, ada := b.Build()["status"]; ada {
		t.Fatal("status ditulis walau ditolak")
	}
}
