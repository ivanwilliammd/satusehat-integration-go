package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// ConditionBuilder builds FHIR Condition payload
type ConditionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewConditionBuilder() *ConditionBuilder {
	b := &ConditionBuilder{ResourceType: "Condition", Data: make(map[string]interface{})}
	return b
}

func (b *ConditionBuilder) setId(id string) *ConditionBuilder {
	b.Data["id"] = id
	return b
}

func (b *ConditionBuilder) addIdentifier(identifier *datatype.Identifier) *ConditionBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *ConditionBuilder) setClinicalStatus(clinicalStatus *datatype.CodeableConcept) *ConditionBuilder {
	b.Data["clinicalStatus"] = clinicalStatus.ToArray()
	return b
}

func (b *ConditionBuilder) setVerificationStatus(verificationStatus *datatype.CodeableConcept) *ConditionBuilder {
	b.Data["verificationStatus"] = verificationStatus.ToArray()
	return b
}

func (b *ConditionBuilder) addCategory(category *datatype.CodeableConcept) *ConditionBuilder {
	if _, ok := b.Data["category"]; !ok {
		b.Data["category"] = make([]interface{}, 0)
	}
	b.Data["category"] = append(b.Data["category"].([]interface{}), category.ToArray())
	return b
}

func (b *ConditionBuilder) setSeverity(severity *datatype.CodeableConcept) *ConditionBuilder {
	b.Data["severity"] = severity.ToArray()
	return b
}

func (b *ConditionBuilder) setCode(code *datatype.CodeableConcept) *ConditionBuilder {
	b.Data["code"] = code.ToArray()
	return b
}

func (b *ConditionBuilder) setSubject(subject *datatype.Reference) *ConditionBuilder {
	b.Data["subject"] = subject.ToArray()
	return b
}

func (b *ConditionBuilder) setEncounter(encounter *datatype.Reference) *ConditionBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

func (b *ConditionBuilder) setOnsetDateTime(dateTime string) *ConditionBuilder {
	b.Data["onsetDateTime"] = dateTime
	return b
}

func (b *ConditionBuilder) setOnsetAge(age *datatype.Range) *ConditionBuilder {
	b.Data["onsetAge"] = age.ToArray()
	return b
}

func (b *ConditionBuilder) setOnsetPeriod(period *datatype.Period) *ConditionBuilder {
	b.Data["onsetPeriod"] = period.ToArray()
	return b
}

func (b *ConditionBuilder) setOnsetRange(rng *datatype.Range) *ConditionBuilder {
	b.Data["onsetRange"] = rng.ToArray()
	return b
}

func (b *ConditionBuilder) setOnsetString(onsetString string) *ConditionBuilder {
	b.Data["onsetString"] = onsetString
	return b
}

func (b *ConditionBuilder) setAbatementDateTime(dateTime string) *ConditionBuilder {
	b.Data["abatementDateTime"] = dateTime
	return b
}

func (b *ConditionBuilder) setAbatementAge(age *datatype.Range) *ConditionBuilder {
	b.Data["abatementAge"] = age.ToArray()
	return b
}

func (b *ConditionBuilder) setAbatementPeriod(period *datatype.Period) *ConditionBuilder {
	b.Data["abatementPeriod"] = period.ToArray()
	return b
}

func (b *ConditionBuilder) setAbatementRange(rng *datatype.Range) *ConditionBuilder {
	b.Data["abatementRange"] = rng.ToArray()
	return b
}

func (b *ConditionBuilder) setAbatementString(abatementString string) *ConditionBuilder {
	b.Data["abatementString"] = abatementString
	return b
}

func (b *ConditionBuilder) setRecordedDate(recordedDate string) *ConditionBuilder {
	b.Data["recordedDate"] = recordedDate
	return b
}

func (b *ConditionBuilder) setRecorder(recorder *datatype.Reference) *ConditionBuilder {
	b.Data["recorder"] = recorder.ToArray()
	return b
}

func (b *ConditionBuilder) setAsserter(asserter *datatype.Reference) *ConditionBuilder {
	b.Data["asserter"] = asserter.ToArray()
	return b
}

func (b *ConditionBuilder) addStage(summary *datatype.CodeableConcept, assessment *datatype.Reference) *ConditionBuilder {
	stage := map[string]interface{}{"summary": summary.ToArray()}
	if assessment != nil {
		stage["assessment"] = []interface{}{assessment.ToArray()}
	}
	if _, ok := b.Data["stage"]; !ok {
		b.Data["stage"] = make([]interface{}, 0)
	}
	b.Data["stage"] = append(b.Data["stage"].([]interface{}), stage)
	return b
}

func (b *ConditionBuilder) addEvidence(code *datatype.CodeableConcept, detail *datatype.Reference) *ConditionBuilder {
	evidence := map[string]interface{}{"code": []interface{}{code.ToArray()}}
	if detail != nil {
		evidence["detail"] = []interface{}{detail.ToArray()}
	}
	if _, ok := b.Data["evidence"]; !ok {
		b.Data["evidence"] = make([]interface{}, 0)
	}
	b.Data["evidence"] = append(b.Data["evidence"].([]interface{}), evidence)
	return b
}

func (b *ConditionBuilder) addNote(note *datatype.Annotation) *ConditionBuilder {
	if _, ok := b.Data["note"]; !ok {
		b.Data["note"] = make([]interface{}, 0)
	}
	b.Data["note"] = append(b.Data["note"].([]interface{}), note.ToArray())
	return b
}

func (b *ConditionBuilder) addExtension(url string, value interface{}, valueType string) *ConditionBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType != "" {
		capitalized := strings.ToUpper(valueType[:1]) + valueType[1:]
		extension["value"+capitalized] = value
	} else {
		extension["valueString"] = value
	}
	if _, ok := b.Data["extension"]; !ok {
		b.Data["extension"] = make([]interface{}, 0)
	}
	b.Data["extension"] = append(b.Data["extension"].([]interface{}), extension)
	return b
}

func (b *ConditionBuilder) Build() map[string]interface{} {
	return b.Data
}
