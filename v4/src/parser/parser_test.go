package parser

import (
	"testing"
)

func TestAll18Parsers(t *testing.T) {
	// 1. DeviceParser
	dev := NewDeviceParser(map[string]any{
		"identifier":   []any{map[string]any{"value": "DEV-1"}},
		"status":       "active",
		"manufacturer": "Acme Med",
		"deviceName":   []any{map[string]any{"name": "Monitor"}},
		"type":         map[string]any{"text": "ECG"},
	})
	if dev.GetStatus() != "active" || dev.GetManufacturer() != "Acme Med" || dev.GetIdentifiers() == nil || dev.GetDeviceNames() == nil || dev.GetType() == nil {
		t.Fatal("DeviceParser failed")
	}

	// 2. DiagnosticReportParser
	dr := NewDiagnosticReportParser(map[string]any{
		"status":            "final",
		"subject":           map[string]any{"reference": "Patient/P1"},
		"encounter":         map[string]any{"reference": "Encounter/E1"},
		"performer":         []any{map[string]any{"reference": "Practitioner/PR1"}},
		"result":            []any{map[string]any{"reference": "Observation/OB1"}},
		"specimen":          []any{map[string]any{"reference": "Specimen/SP1"}},
		"effectiveDateTime": "2024-01-01T00:00:00Z",
		"conclusion":        "Normal",
	})
	if dr.GetSubjectReference() != "P1" || dr.GetEncounterReference() != "E1" ||
		len(dr.GetPerformerReferences()) != 1 || dr.GetPerformerReferences()[0] != "PR1" ||
		len(dr.GetResultReferences()) != 1 || dr.GetResultReferences()[0] != "OB1" ||
		len(dr.GetSpecimenReferences()) != 1 || dr.GetSpecimenReferences()[0] != "SP1" ||
		dr.GetConclusion() != "Normal" {
		t.Fatal("DiagnosticReportParser failed")
	}

	// 3. DocumentReferenceParser
	doc := NewDocumentReferenceParser(map[string]any{
		"status":      "current",
		"docStatus":   "final",
		"subject":     map[string]any{"reference": "Patient/P1"},
		"custodian":   map[string]any{"reference": "Organization/O1"},
		"description": "Discharge Summary",
	})
	if doc.GetStatus() != "current" || doc.GetDocStatus() != "final" || doc.GetSubjectReference() != "Patient/P1" || doc.GetCustodianReference() != "Organization/O1" || doc.GetDescription() != "Discharge Summary" {
		t.Fatal("DocumentReferenceParser failed")
	}

	// 4. EncounterParser
	enc := NewEncounterParser(map[string]any{
		"status":     "finished",
		"identifier": []any{map[string]any{"value": "REG-001"}},
		"subject":    map[string]any{"reference": "Patient/P1", "display": "Budi"},
		"location": []any{
			map[string]any{"location": map[string]any{"reference": "Location/L1", "display": "Poli Umum"}},
		},
		"participant": []any{
			map[string]any{"individual": map[string]any{"reference": "Practitioner/PR1", "display": "dr. Andi"}},
		},
		"diagnosis": []any{
			map[string]any{"condition": map[string]any{"reference": "Condition/C1", "display": "Fever"}},
		},
		"period": map[string]any{"start": "2024-01-01T08:00:00Z", "end": "2024-01-01T09:00:00Z"},
	})
	if enc.GetStatus() != "finished" || enc.GetSubjectReference() != "P1" || enc.GetSubjectDisplay() != "Budi" ||
		enc.GetLocationReferences()[0] != "L1" || enc.GetParticipantReferences()[0] != "PR1" ||
		enc.GetDiagnosisReferences()[0] != "C1" || enc.GetRegistrationIds()[0] != "REG-001" ||
		enc.GetPeriodStart() != "2024-01-01T08:00:00Z" || enc.GetPeriodEnd() != "2024-01-01T09:00:00Z" {
		t.Fatal("EncounterParser failed")
	}

	// 5. GenomicStudyParser
	gs := NewGenomicStudyParser(map[string]any{
		"status":      "registered",
		"subject":     map[string]any{"reference": "Patient/P1"},
		"encounter":   map[string]any{"reference": "Encounter/E1"},
		"referrer":    map[string]any{"reference": "Practitioner/PR1"},
		"startDate":   "2024-01-01",
		"description": "WGS",
	})
	if gs.GetStatus() != "registered" || gs.GetSubjectReference() != "Patient/P1" || gs.GetReferrerReference() != "Practitioner/PR1" || gs.GetDescription() != "WGS" {
		t.Fatal("GenomicStudyParser failed")
	}

	// 6. GroupParser
	grp := NewGroupParser(map[string]any{
		"active":   true,
		"type":     "person",
		"actual":   true,
		"name":     "Cohort A",
		"quantity": 10,
	})
	if grp.GetType() != "person" || grp.GetName() != "Cohort A" || grp.GetQuantity() != 10 {
		t.Fatal("GroupParser failed")
	}

	// 7. LocationParser
	loc := NewLocationParser(map[string]any{
		"name":                 "Ruang ICU",
		"status":               "active",
		"managingOrganization": map[string]any{"reference": "Organization/ORG1"},
		"partOf":               map[string]any{"reference": "Location/LOC0"},
	})
	if loc.GetName() != "Ruang ICU" || loc.GetManagingOrganizationReference() != "ORG1" || loc.GetPartOfReference() != "LOC0" {
		t.Fatal("LocationParser failed")
	}

	// 8. MolecularSequenceParser
	ms := NewMolecularSequenceParser(map[string]any{
		"type":             "dna",
		"coordinateSystem": 1,
		"patient":          map[string]any{"reference": "Patient/P1"},
		"specimen":         map[string]any{"reference": "Specimen/S1"},
		"device":           map[string]any{"reference": "Device/D1"},
		"performer":        map[string]any{"reference": "Organization/O1"},
	})
	if ms.GetType() != "dna" || ms.GetPatientReference() != "P1" || ms.GetSpecimenReference() != "S1" || ms.GetDeviceReference() != "D1" || ms.GetPerformerReference() != "O1" {
		t.Fatal("MolecularSequenceParser failed")
	}

	// 9. ObservationParser
	obs := NewObservationParser(map[string]any{
		"status":      "final",
		"subject":     map[string]any{"reference": "Patient/P1"},
		"encounter":   map[string]any{"reference": "Encounter/E1"},
		"performer":   []any{map[string]any{"reference": "Practitioner/PR1"}},
		"specimen":    map[string]any{"reference": "Specimen/S1"},
		"valueString": "Positive",
	})
	if obs.GetStatus() != "final" || obs.GetSubjectReference() != "P1" || obs.GetEncounterReference() != "E1" ||
		obs.GetPerformerReferences()[0] != "PR1" || obs.GetSpecimenReference() != "S1" || obs.GetValueString() != "Positive" {
		t.Fatal("ObservationParser failed")
	}

	// 10. OrganizationParser
	org := NewOrganizationParser(map[string]any{
		"name":   "RS Harapan",
		"partOf": map[string]any{"reference": "Organization/ORG-ROOT"},
	})
	if org.GetName() != "RS Harapan" || org.GetPartOfReference() != "ORG-ROOT" {
		t.Fatal("OrganizationParser failed")
	}

	// 11. PatientParser
	pat := NewPatientParser(map[string]any{
		"gender":               "male",
		"birthDate":            "1990-05-10",
		"deceasedBoolean":      false,
		"multipleBirthInteger": 2,
	})
	if pat.GetGender() != "male" || pat.GetBirthDate() != "1990-05-10" || pat.GetMultipleBirth() != 2 {
		t.Fatal("PatientParser failed")
	}

	// 12. PractitionerParser
	prac := NewPractitionerParser(map[string]any{
		"id":        "PRAC-100",
		"gender":    "female",
		"birthDate": "1985-03-20",
		"address": []any{
			map[string]any{
				"line": []any{"Jl. Kesehatan No. 1"},
				"extension": []any{
					map[string]any{
						"extension": []any{
							map[string]any{"url": "province", "valueCode": "31"},
							map[string]any{"url": "city", "valueCode": "3171"},
							map[string]any{"url": "district", "valueCode": "317101"},
							map[string]any{"url": "village", "valueCode": "3171011001"},
						},
					},
				},
			},
		},
		"qualification": []any{
			map[string]any{
				"identifier": []any{map[string]any{"value": "STR-999"}},
			},
		},
	})
	if prac.GetId() != "PRAC-100" || prac.GetAddressLine() != "Jl. Kesehatan No. 1" ||
		prac.GetCity() != "3171" || prac.GetVillage() != "3171011001" || prac.GetQualifications() != "STR-999" {
		t.Fatal("PractitionerParser failed")
	}

	// 13. ProcedureParser
	proc := NewProcedureParser(map[string]any{
		"status":    "completed",
		"subject":   map[string]any{"reference": "Patient/P1", "display": "Budi"},
		"encounter": map[string]any{"reference": "Encounter/E1", "display": "Visit 1"},
		"performer": []any{
			map[string]any{"actor": map[string]any{"reference": "Practitioner/PR1", "display": "dr. Andi"}},
		},
		"location": map[string]any{"reference": "Location/L1"},
	})
	if proc.GetStatus() != "completed" || proc.GetSubjectReference() != "P1" || proc.GetEncounterReference() != "E1" ||
		proc.GetPerformerReferences()[0] != "PR1" || proc.GetLocation() != "L1" {
		t.Fatal("ProcedureParser failed")
	}

	// 14. RelatedPersonParser
	rp := NewRelatedPersonParser(map[string]any{
		"active":    true,
		"patient":   map[string]any{"reference": "Patient/P1"},
		"gender":    "female",
		"birthDate": "1965-01-01",
	})
	if rp.GetPatientReference() != "Patient/P1" || rp.GetGender() != "female" {
		t.Fatal("RelatedPersonParser failed")
	}

	// 15. RiskAssessmentParser
	ra := NewRiskAssessmentParser(map[string]any{
		"status":     "final",
		"subject":    map[string]any{"reference": "Patient/P1"},
		"encounter":  map[string]any{"reference": "Encounter/E1"},
		"condition":  map[string]any{"reference": "Condition/C1"},
		"performer":  map[string]any{"reference": "Practitioner/PR1"},
		"mitigation": "Low sodium diet",
	})
	if ra.GetStatus() != "final" || ra.GetSubjectReference() != "P1" || ra.GetEncounterReference() != "E1" ||
		ra.GetConditionReference() != "C1" || ra.GetPerformerReference() != "PR1" || ra.GetMitigation() != "Low sodium diet" {
		t.Fatal("RiskAssessmentParser failed")
	}

	// 16. ServiceRequestParser
	sr := NewServiceRequestParser(map[string]any{
		"status":             "active",
		"intent":             "order",
		"priority":           "routine",
		"subject":            map[string]any{"reference": "Patient/P1", "display": "Budi"},
		"encounter":          map[string]any{"reference": "Encounter/E1", "display": "Enc 1"},
		"requester":          map[string]any{"reference": "Practitioner/PR1", "display": "dr. Andi"},
		"patientInstruction": "Fasting 8 hours",
	})
	if sr.GetStatus() != "active" || sr.GetIntent() != "order" || sr.GetEncounterReference() != "E1" ||
		sr.GetRequesterReference() != "Practitioner/PR1" || sr.GetPatientInstruction() != "Fasting 8 hours" {
		t.Fatal("ServiceRequestParser failed")
	}

	// 17. SpecimenParser
	spec := NewSpecimenParser(map[string]any{
		"status":  "available",
		"subject": map[string]any{"reference": "Patient/P1", "display": "Budi"},
		"collection": map[string]any{
			"collectedDateTime": "2024-01-01T07:00:00Z",
			"collector":         map[string]any{"reference": "Practitioner/PR1", "display": "Nurse A"},
		},
	})
	if spec.GetStatus() != "available" || spec.GetSubjectReference() != "P1" ||
		spec.GetCollectedDateTime() != "2024-01-01T07:00:00Z" || spec.GetCollectorReference() != "PR1" {
		t.Fatal("SpecimenParser failed")
	}

	// 18. TaskParser
	tsk := NewTaskParser(map[string]any{
		"status":          "requested",
		"intent":          "order",
		"priority":        "stat",
		"instantiatesUri": "http://example.org/task",
		"for":             map[string]any{"reference": "Patient/P1"},
		"encounter":       map[string]any{"reference": "Encounter/E1"},
		"requester":       map[string]any{"reference": "Practitioner/PR1"},
	})
	if tsk.GetStatus() != "requested" || tsk.GetForReference() != "P1" ||
		tsk.GetEncounterReference() != "E1" || tsk.GetRequesterReference() != "PR1" {
		t.Fatal("TaskParser failed")
	}
}
