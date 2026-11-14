package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/terminology"
)

// ConditionBuilder builds a FHIR R4 Condition resource.
//
// Method-for-method counterpart of PayloadBuilderCondition in the PHP SDK. SetSeverity and SetCode
// accept a string and resolve it through the offline terminology catalogs, as PHP does.
//
// @link https://www.hl7.org/fhir/condition.html
type ConditionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewConditionBuilder creates an empty Condition with its resourceType set.
func NewConditionBuilder() *ConditionBuilder {
	b := &ConditionBuilder{ResourceType: "Condition", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *ConditionBuilder) set(k string, v interface{}) *ConditionBuilder {
	b.Data[k] = v
	return b
}

func (b *ConditionBuilder) push(k string, v interface{}) *ConditionBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

func (b *ConditionBuilder) setCodeable(key string, value interface{}) *ConditionBuilder {
	switch v := value.(type) {
	case string:
		return b.set(key, terminology.Resolve(v))
	case *datatype.CodeableConcept:
		return b.set(key, v.ToArray())
	}
	return b
}

// SetId sets the logical id.
func (b *ConditionBuilder) SetId(id string) *ConditionBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *ConditionBuilder) AddIdentifier(identifier *datatype.Identifier) *ConditionBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetClinicalStatus sets clinicalStatus.
func (b *ConditionBuilder) SetClinicalStatus(status *datatype.CodeableConcept) *ConditionBuilder {
	return b.set("clinicalStatus", status.ToArray())
}

// SetVerificationStatus sets verificationStatus.
func (b *ConditionBuilder) SetVerificationStatus(status *datatype.CodeableConcept) *ConditionBuilder {
	return b.set("verificationStatus", status.ToArray())
}

// AddCategory appends a category.
func (b *ConditionBuilder) AddCategory(category *datatype.CodeableConcept) *ConditionBuilder {
	return b.push("category", category.ToArray())
}

// SetSeverity sets severity, either as a catalog key or as a typed concept.
func (b *ConditionBuilder) SetSeverity(severity interface{}) *ConditionBuilder {
	return b.setCodeable("severity", severity)
}

// SetCode sets code, either as a catalog key or as a typed concept.
func (b *ConditionBuilder) SetCode(code interface{}) *ConditionBuilder {
	return b.setCodeable("code", code)
}

// SetSubject sets subject.
func (b *ConditionBuilder) SetSubject(subject *datatype.Reference) *ConditionBuilder {
	return b.set("subject", subject.ToArray())
}

// SetEncounter sets encounter.
func (b *ConditionBuilder) SetEncounter(encounter *datatype.Reference) *ConditionBuilder {
	return b.set("encounter", encounter.ToArray())
}

// SetOnsetDateTime sets onsetDateTime.
func (b *ConditionBuilder) SetOnsetDateTime(v string) *ConditionBuilder {
	return b.set("onsetDateTime", v)
}

// SetOnsetAge sets onsetAge.
func (b *ConditionBuilder) SetOnsetAge(v *datatype.Range) *ConditionBuilder {
	return b.set("onsetAge", v.ToArray())
}

// SetOnsetPeriod sets onsetPeriod.
func (b *ConditionBuilder) SetOnsetPeriod(v *datatype.Period) *ConditionBuilder {
	return b.set("onsetPeriod", v.ToArray())
}

// SetOnsetRange sets onsetRange.
func (b *ConditionBuilder) SetOnsetRange(v *datatype.Range) *ConditionBuilder {
	return b.set("onsetRange", v.ToArray())
}

// SetOnsetString sets onsetString.
func (b *ConditionBuilder) SetOnsetString(v string) *ConditionBuilder { return b.set("onsetString", v) }

// SetAbatementDateTime sets abatementDateTime.
func (b *ConditionBuilder) SetAbatementDateTime(v string) *ConditionBuilder {
	return b.set("abatementDateTime", v)
}

// SetAbatementAge sets abatementAge.
func (b *ConditionBuilder) SetAbatementAge(v *datatype.Range) *ConditionBuilder {
	return b.set("abatementAge", v.ToArray())
}

// SetAbatementPeriod sets abatementPeriod.
func (b *ConditionBuilder) SetAbatementPeriod(v *datatype.Period) *ConditionBuilder {
	return b.set("abatementPeriod", v.ToArray())
}

// SetAbatementRange sets abatementRange.
func (b *ConditionBuilder) SetAbatementRange(v *datatype.Range) *ConditionBuilder {
	return b.set("abatementRange", v.ToArray())
}

// SetAbatementString sets abatementString.
func (b *ConditionBuilder) SetAbatementString(v string) *ConditionBuilder {
	return b.set("abatementString", v)
}

// SetRecordedDate sets recordedDate.
func (b *ConditionBuilder) SetRecordedDate(v string) *ConditionBuilder {
	return b.set("recordedDate", v)
}

// SetRecorder sets recorder.
func (b *ConditionBuilder) SetRecorder(v *datatype.Reference) *ConditionBuilder {
	return b.set("recorder", v.ToArray())
}

// SetAsserter sets asserter.
func (b *ConditionBuilder) SetAsserter(v *datatype.Reference) *ConditionBuilder {
	return b.set("asserter", v.ToArray())
}

// AddStage appends a stage. assessment is optional.
func (b *ConditionBuilder) AddStage(summary *datatype.CodeableConcept, assessment *datatype.Reference) *ConditionBuilder {
	stage := map[string]interface{}{"summary": summary.ToArray()}
	if assessment != nil {
		stage["assessment"] = []interface{}{assessment.ToArray()}
	}
	return b.push("stage", stage)
}

// AddEvidence appends evidence. detail is optional.
func (b *ConditionBuilder) AddEvidence(code *datatype.CodeableConcept, detail *datatype.Reference) *ConditionBuilder {
	evidence := map[string]interface{}{"code": []interface{}{code.ToArray()}}
	if detail != nil {
		evidence["detail"] = []interface{}{detail.ToArray()}
	}
	return b.push("evidence", evidence)
}

// AddNote appends a note, either as a typed Annotation or as plain text.
func (b *ConditionBuilder) AddNote(note interface{}) *ConditionBuilder {
	switch v := note.(type) {
	case *datatype.Annotation:
		return b.push("note", v.ToArray())
	case string:
		return b.push("note", map[string]interface{}{"text": v})
	}
	return b
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *ConditionBuilder) AddExtension(url string, value interface{}, valueType string) *ConditionBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *ConditionBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *ConditionBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
