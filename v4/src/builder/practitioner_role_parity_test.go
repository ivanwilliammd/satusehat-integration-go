package builder_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestPractitionerRoleMatchesPHPPayload(t *testing.T) {
	sistemTelepon, nilaiTelepon := datatype.ContactPointSystem("phone"), "+628123456789"
	penggunaan := datatype.ContactPointUse("work")

	b := builder.NewPractitionerRoleBuilder()
	b.SetId("pr-1")
	b.SetActive(true)
	b.SetPractitioner("Practitioner/1000000002")
	b.SetOrganization("Organization/1000000001")
	b.AddCode(konsep("http://terminology.hl7.org/CodeSystem/practitioner-role", "doctor", "Doctor"))
	b.AddSpecialty(konsep("http://snomed.info/sct", "394814009", "General practice"))
	b.AddLocation("Location/3000000001")
	b.AddTelecom(&datatype.ContactPoint{System: &sistemTelepon, Value: &nilaiTelepon, Use: &penggunaan})
	b.AddEndpoint("Endpoint/4000000001")
	b.SetAvailabilityExceptions("Tidak praktik hari Minggu.")
	b.AddExtension("http://example.org/str", "1234567890", "String")
	samakan(t, bacaGolden(t, "practitioner_role"), b)
}
