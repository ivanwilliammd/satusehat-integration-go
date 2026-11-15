package builder_test

// Group, ListResource, SpecimenDefinition, Substance and SubstanceReferenceInformation — the last five
// builders with methods of their own that had no golden.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestGroupMatchesPHPPayload(t *testing.T) {
	b := builder.NewGroupBuilder()
	b.SetId("grp-1")
	b.AddIdentifier("http://sys.id/ni", "NI-80")
	b.SetActive(true)
	b.SetType("person")
	b.SetActual(true)
	b.SetCode(konsep("http://terminology.hl7.org/CodeSystem/group-type", "person", "Person").ToArray())
	b.SetName("Keluarga Santoso")
	b.SetQuantity(3)
	b.SetManagingEntity((&datatype.Reference{Reference: "Organization/1000000001"}).ToArray())
	b.AddMember("Patient/1000000001", "Budi Santoso")
	b.AddMember("Patient/1000000002", map[string]interface{}{"start": "2026-01-05"})
	b.AddMember("Patient/1000000003", nil, true)
	b.AddMemberReference(&datatype.Reference{Reference: "RelatedPerson/rp-1"})
	b.AddExtension("http://example.org/ext", true)
	b.AddExtension("http://example.org/ext2", "teks")
	samakan(t, bacaGolden(t, "group"), b)
}

func TestListResourceMatchesPHPPayload(t *testing.T) {
	b := builder.NewListResourceBuilder()
	b.SetId("list-1")
	b.SetStatus("current")
	b.SetMode("working")
	b.SetTitle("Daftar Pasien")
	b.SetCode("http://loinc.org", "11450-4", "Problem list")
	b.SetSubject("Patient/1000000001", "Budi Santoso")
	b.SetDate("2026-01-05")
	samakan(t, bacaGolden(t, "list_resource"), b)
}

func TestSpecimenDefinitionMatchesPHPPayload(t *testing.T) {
	b := builder.NewSpecimenDefinitionBuilder()
	b.SetId("sd-1")
	b.AddIdentifier("http://sys.id/ni", "NI-81")
	b.SetStatus("active")
	b.SetType("http://snomed.info/sct", "122555007", "Venous blood specimen")
	b.SetCollection("Darah vena")
	b.SetPreparation("http://terminology.hl7.org/CodeSystem/preparation", "spun", "Spun")
	samakan(t, bacaGolden(t, "specimen_definition"), b)
}

func TestSubstanceMatchesPHPPayload(t *testing.T) {
	b := builder.NewSubstanceBuilder()
	b.SetId("sub-1")
	b.SetStatus("active")
	b.SetCode("http://snomed.info/sct", "387517004", "Amoksisilin")
	b.SetText("generated", "<div>Amoksisilin</div>")
	samakan(t, bacaGolden(t, "substance"), b)
}

func TestSubstanceReferenceInformationMatchesPHPPayload(t *testing.T) {
	b := builder.NewSubstanceReferenceInformationBuilder()
	b.SetId("sri-1")
	b.SetStatus("active")
	b.SetComment("Informasi referensi")
	b.SetGene("http://www.genenames.org", "HGNC:5", "A1BG")
	b.SetGeneElement("http://snomed.info/sct", "123", "Elemen")
	samakan(t, bacaGolden(t, "substance_reference_information"), b)
}
