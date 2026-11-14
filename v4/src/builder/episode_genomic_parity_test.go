package builder_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestEpisodeOfCareMatchesPHPPayload(t *testing.T) {
	rank := 1
	akhir := "2026-02-05"

	b := builder.NewEpisodeOfCareBuilder()
	b.SetId("eoc-1")
	b.SetStatus("active")
	b.SetPatient(&datatype.Reference{Reference: "Patient/1000000001"})
	b.SetManagingOrganization(&datatype.Reference{Reference: "Organization/1000000001"})
	b.SetCareManager(&datatype.Reference{Reference: "Practitioner/1000000002"})
	b.SetPeriod(&datatype.Period{Start: ptr("2026-01-05"), End: &akhir})
	b.AddType(konsep("http://terminology.hl7.org/CodeSystem/episodeofcare-type", "hacc", "Home and Community Care"))
	b.AddStatusHistory("active", "2026-01-05", nil)
	b.AddDiagnosis(
		&datatype.Reference{Reference: "Condition/cond-1"},
		konsep("http://terminology.hl7.org/CodeSystem/diagnosis-role", "AD", "Admission diagnosis"),
		&rank,
	)
	b.AddReferralRequest(&datatype.Reference{Reference: "ServiceRequest/sr-1"})
	b.AddTeam(&datatype.Reference{Reference: "CareTeam/ct-1"})
	b.AddAccount(&datatype.Reference{Reference: "Account/1500000001"})

	if err := b.Err(); err != nil {
		t.Fatalf("status sah ditolak: %v", err)
	}
	samakan(t, bacaGolden(t, "episode_of_care"), b)
}

func TestEpisodeOfCareRefusesAnUnknownStatus(t *testing.T) {
	b := builder.NewEpisodeOfCareBuilder()
	b.SetId("eoc-1")
	b.AddStatusHistory("aktif", "2026-01-05", nil)

	if b.Err() == nil {
		t.Fatal("status tidak dikenal seharusnya dicatat sebagai galat")
	}
	if _, ada := b.Build()["statusHistory"]; ada {
		t.Fatal("entri dengan status tidak dikenal tidak boleh ditulis")
	}
}

func TestGenomicStudyMatchesPHPPayload(t *testing.T) {
	b := builder.NewGenomicStudyBuilder()
	b.SetId("gs-1")
	b.SetStatus("available")
	b.SetSubject("Patient/1000000001")
	b.SetMetaProfile("http://example.org/StructureDefinition/genomic-study")
	b.AddType(konsep("http://snomed.info/sct", "443910001", "Genetic testing"))
	b.SetStartDate("2026-01-05")
	b.AddBasedOn(&datatype.Reference{Reference: "ServiceRequest/sr-2"})
	b.SetReferrer(&datatype.Reference{Reference: "Practitioner/1000000002"})
	b.AddInterpreter(&datatype.Reference{Reference: "Practitioner/1000000003"})
	b.AddReason(konsep("http://snomed.info/sct", "399068003", "Routine health examination"))
	b.AddPerformer(&datatype.Reference{Reference: "Practitioner/1000000002"}, "")
	b.AddAnalysis(map[string]interface{}{
		"identifier": []interface{}{map[string]interface{}{"system": "http://example.org/an", "value": "AN-1"}},
	})
	samakan(t, bacaGolden(t, "genomic_study"), b)
}
