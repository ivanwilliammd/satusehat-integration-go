package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// AllergyIntoleranceBuilder builds FHIR AllergyIntolerance payload
type AllergyIntoleranceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewAllergyIntoleranceBuilder() *AllergyIntoleranceBuilder {
	b := &AllergyIntoleranceBuilder{ResourceType: "AllergyIntolerance", Data: make(map[string]interface{})}
	return b
}

func (b *AllergyIntoleranceBuilder) setId(id string) *AllergyIntoleranceBuilder {
	b.Data["id"] = id
	return b
}

func (b *AllergyIntoleranceBuilder) addIdentifier(identifier *datatype.Identifier) *AllergyIntoleranceBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *AllergyIntoleranceBuilder) setClinicalStatus(clinicalStatus *datatype.CodeableConcept) *AllergyIntoleranceBuilder {
	b.Data["clinicalStatus"] = clinicalStatus.ToArray()
	return b
}

func (b *AllergyIntoleranceBuilder) setVerificationStatus(verificationStatus *datatype.CodeableConcept) *AllergyIntoleranceBuilder {
	b.Data["verificationStatus"] = verificationStatus.ToArray()
	return b
}

func (b *AllergyIntoleranceBuilder) setType(typ string) *AllergyIntoleranceBuilder {
	b.Data["type"] = typ
	return b
}

func (b *AllergyIntoleranceBuilder) setCategory(category string) *AllergyIntoleranceBuilder {
	b.Data["category"] = category
	return b
}

func (b *AllergyIntoleranceBuilder) setCriticality(criticality string) *AllergyIntoleranceBuilder {
	b.Data["criticality"] = criticality
	return b
}

func (b *AllergyIntoleranceBuilder) setCode(code *datatype.CodeableConcept) *AllergyIntoleranceBuilder {
	b.Data["code"] = code.ToArray()
	return b
}

func (b *AllergyIntoleranceBuilder) setPatient(patient *datatype.Reference) *AllergyIntoleranceBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}

func (b *AllergyIntoleranceBuilder) setEncounter(encounter *datatype.Reference) *AllergyIntoleranceBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

func (b *AllergyIntoleranceBuilder) setOnsetDateTime(dateTime string) *AllergyIntoleranceBuilder {
	b.Data["onsetDateTime"] = dateTime
	return b
}

func (b *AllergyIntoleranceBuilder) setOnsetAge(age *datatype.Age) *AllergyIntoleranceBuilder {
	b.Data["onsetAge"] = age.ToArray()
	return b
}

func (b *AllergyIntoleranceBuilder) setOnsetPeriod(period *datatype.Period) *AllergyIntoleranceBuilder {
	b.Data["onsetPeriod"] = period.ToArray()
	return b
}

func (b *AllergyIntoleranceBuilder) setOnsetRange(rng *datatype.Range) *AllergyIntoleranceBuilder {
	b.Data["onsetRange"] = rng.ToArray()
	return b
}

func (b *AllergyIntoleranceBuilder) setOnsetString(onsetString string) *AllergyIntoleranceBuilder {
	b.Data["onsetString"] = onsetString
	return b
}

func (b *AllergyIntoleranceBuilder) setRecordedDate(recordedDate string) *AllergyIntoleranceBuilder {
	b.Data["recordedDate"] = recordedDate
	return b
}

func (b *AllergyIntoleranceBuilder) setRecorder(recorder *datatype.Reference) *AllergyIntoleranceBuilder {
	b.Data["recorder"] = recorder.ToArray()
	return b
}

func (b *AllergyIntoleranceBuilder) setAsserter(asserter *datatype.Reference) *AllergyIntoleranceBuilder {
	b.Data["asserter"] = asserter.ToArray()
	return b
}

func (b *AllergyIntoleranceBuilder) setLastOccurrence(lastOccurrence string) *AllergyIntoleranceBuilder {
	b.Data["lastOccurrence"] = lastOccurrence
	return b
}

func (b *AllergyIntoleranceBuilder) addNote(note *datatype.Annotation) *AllergyIntoleranceBuilder {
	if _, ok := b.Data["note"]; !ok {
		b.Data["note"] = make([]interface{}, 0)
	}
	b.Data["note"] = append(b.Data["note"].([]interface{}), note.ToArray())
	return b
}

func (b *AllergyIntoleranceBuilder) addReaction(
	substance *datatype.CodeableConcept,
	manifestation []datatype.CodeableConcept,
	description string,
	onset string,
	severity string,
	note *datatype.Annotation,
) *AllergyIntoleranceBuilder {
	reaction := map[string]interface{}{}
	if substance != nil {
		reaction["substance"] = substance.ToArray()
	}
	if len(manifestation) > 0 {
		mans := make([]interface{}, len(manifestation))
		for i, m := range manifestation {
			mans[i] = m.ToArray()
		}
		reaction["manifestation"] = mans
	}
	if description != "" {
		reaction["description"] = description
	}
	if onset != "" {
		reaction["onset"] = onset
	}
	if severity != "" {
		reaction["severity"] = severity
	}
	if note != nil {
		reaction["note"] = []interface{}{note.ToArray()}
	}
	if _, ok := b.Data["reaction"]; !ok {
		b.Data["reaction"] = make([]interface{}, 0)
	}
	b.Data["reaction"] = append(b.Data["reaction"].([]interface{}), reaction)
	return b
}

func (b *AllergyIntoleranceBuilder) addExtension(url string, value interface{}, valueType string) *AllergyIntoleranceBuilder {
	ext := map[string]interface{}{"url": url}
	if valueType != "" {
		capitalized := strings.ToUpper(valueType[:1]) + valueType[1:]
		ext["value"+capitalized] = value
	} else {
		ext["valueString"] = value
	}
	if _, ok := b.Data["extension"]; !ok {
		b.Data["extension"] = make([]interface{}, 0)
	}
	b.Data["extension"] = append(b.Data["extension"].([]interface{}), ext)
	return b
}

func (b *AllergyIntoleranceBuilder) Build() map[string]interface{} {
	return b.Data
}
