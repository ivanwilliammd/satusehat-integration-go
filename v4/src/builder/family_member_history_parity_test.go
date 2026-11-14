package builder_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func konsep(system, code, display string) *datatype.CodeableConcept {
	cc := &datatype.CodeableConcept{}
	cc.Coding = []datatype.Coding{{System: system, Code: code, Display: display}}
	return cc
}

func TestFamilyMemberHistoryMatchesPHPPayload(t *testing.T) {
	b := builder.NewFamilyMemberHistoryBuilder()
	b.SetId("fmh-1")
	b.SetStatus("completed")
	b.SetPatient(&datatype.Reference{Reference: "Patient/1000000001"})
	b.SetDate("2026-01-05")
	b.SetName("Ibu kandung")
	b.SetRelationship(konsep("http://terminology.hl7.org/CodeSystem/v3-RoleCode", "MTH", "mother"))
	b.SetSex(konsep("http://hl7.org/fhir/administrative-gender", "female", "Female"))
	b.SetBornDate("1960-03-01")
	b.SetAgeAge(65, "a")
	b.SetDeceasedBoolean(false)
	b.SetNote("Riwayat hipertensi.")
	samakan(t, bacaGolden(t, "family_member_history"), b)
}
