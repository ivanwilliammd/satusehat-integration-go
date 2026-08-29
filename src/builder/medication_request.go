package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// MedicationRequestBuilder builds FHIR MedicationRequest payload
type MedicationRequestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMedicationRequestBuilder() *MedicationRequestBuilder {
	b := &MedicationRequestBuilder{ResourceType: "MedicationRequest", Data: make(map[string]interface{})}
	return b
}

func (b *MedicationRequestBuilder) setId(id string) *MedicationRequestBuilder {
	b.Data["id"] = id
	return b
}

func (b *MedicationRequestBuilder) addIdentifier(identifier *datatype.Identifier) *MedicationRequestBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *MedicationRequestBuilder) setStatus(status string) *MedicationRequestBuilder {
	b.Data["status"] = status
	return b
}

func (b *MedicationRequestBuilder) setIntent(intent string) *MedicationRequestBuilder {
	b.Data["intent"] = intent
	return b
}

func (b *MedicationRequestBuilder) setCategory(category *datatype.CodeableConcept) *MedicationRequestBuilder {
	b.Data["category"] = category.ToArray()
	return b
}

func (b *MedicationRequestBuilder) setPriority(priority string) *MedicationRequestBuilder {
	b.Data["priority"] = priority
	return b
}

func (b *MedicationRequestBuilder) setMedicationCodeableConcept(medication *datatype.CodeableConcept) *MedicationRequestBuilder {
	b.Data["medicationCodeableConcept"] = medication.ToArray()
	return b
}

func (b *MedicationRequestBuilder) setMedicationReference(medication *datatype.Reference) *MedicationRequestBuilder {
	b.Data["medicationReference"] = medication.ToArray()
	return b
}

func (b *MedicationRequestBuilder) setSubject(subject *datatype.Reference) *MedicationRequestBuilder {
	b.Data["subject"] = subject.ToArray()
	return b
}

func (b *MedicationRequestBuilder) setEncounter(encounter *datatype.Reference) *MedicationRequestBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

func (b *MedicationRequestBuilder) setAuthoredOn(authoredOn string) *MedicationRequestBuilder {
	b.Data["authoredOn"] = authoredOn
	return b
}

func (b *MedicationRequestBuilder) setRequester(requester *datatype.Reference) *MedicationRequestBuilder {
	b.Data["requester"] = requester.ToArray()
	return b
}

func (b *MedicationRequestBuilder) setRecorder(recorder *datatype.Reference) *MedicationRequestBuilder {
	b.Data["recorder"] = recorder.ToArray()
	return b
}

func (b *MedicationRequestBuilder) setReasonCode(reasonCode *datatype.CodeableConcept) *MedicationRequestBuilder {
	b.Data["reasonCode"] = reasonCode.ToArray()
	return b
}

func (b *MedicationRequestBuilder) setReasonReference(reasonRef *datatype.Reference) *MedicationRequestBuilder {
	b.Data["reasonReference"] = reasonRef.ToArray()
	return b
}

func (b *MedicationRequestBuilder) addNote(note *datatype.Annotation) *MedicationRequestBuilder {
	if _, ok := b.Data["note"]; !ok {
		b.Data["note"] = make([]interface{}, 0)
	}
	b.Data["note"] = append(b.Data["note"].([]interface{}), note.ToArray())
	return b
}

func (b *MedicationRequestBuilder) addDosageInstruction(
	sequence int,
	text string,
	timing *datatype.Timing,
	site *datatype.CodeableConcept,
	route *datatype.CodeableConcept,
	method *datatype.CodeableConcept,
	doseAndRate []map[string]interface{},
) *MedicationRequestBuilder {
	dosage := map[string]interface{}{}
	if sequence != 0 {
		dosage["sequence"] = sequence
	}
	if text != "" {
		dosage["text"] = text
	}
	if timing != nil {
		dosage["timing"] = timing.ToArray()
	}
	if site != nil {
		dosage["site"] = site.ToArray()
	}
	if route != nil {
		dosage["route"] = route.ToArray()
	}
	if method != nil {
		dosage["method"] = method.ToArray()
	}
	if len(doseAndRate) > 0 {
		dosage["doseAndRate"] = doseAndRate
	}
	if _, ok := b.Data["dosageInstruction"]; !ok {
		b.Data["dosageInstruction"] = make([]interface{}, 0)
	}
	b.Data["dosageInstruction"] = append(b.Data["dosageInstruction"].([]interface{}), dosage)
	return b
}

func (b *MedicationRequestBuilder) addExtension(url string, value interface{}, valueType string) *MedicationRequestBuilder {
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

func (b *MedicationRequestBuilder) Build() map[string]interface{} {
	return b.Data
}
