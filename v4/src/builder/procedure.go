package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// ProcedureBuilder builds a FHIR R4 Procedure resource.
//
// Method-for-method counterpart of PayloadBuilderProcedure in the PHP SDK: the whole performed[x]
// family, performer with optional function and onBehalfOf, focalDevice, usedReference/usedCode,
// bodySite, report, followUp and extension.
//
// @link https://www.hl7.org/fhir/procedure.html
type ProcedureBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewProcedureBuilder creates an empty Procedure with its resourceType set.
func NewProcedureBuilder() *ProcedureBuilder {
	b := &ProcedureBuilder{ResourceType: "Procedure", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *ProcedureBuilder) set(k string, v interface{}) *ProcedureBuilder {
	b.Data[k] = v
	return b
}

func (b *ProcedureBuilder) push(k string, v interface{}) *ProcedureBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *ProcedureBuilder) SetId(id string) *ProcedureBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *ProcedureBuilder) AddIdentifier(identifier *datatype.Identifier) *ProcedureBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetStatus sets status (preparation | in-progress | not-done | on-hold | stopped | completed | entered-in-error | unknown).
func (b *ProcedureBuilder) SetStatus(status string) *ProcedureBuilder { return b.set("status", status) }

// SetCategory sets category.
func (b *ProcedureBuilder) SetCategory(category *datatype.CodeableConcept) *ProcedureBuilder {
	return b.set("category", category.ToArray())
}

// SetCode sets code.
func (b *ProcedureBuilder) SetCode(code *datatype.CodeableConcept) *ProcedureBuilder {
	return b.set("code", code.ToArray())
}

// SetSubject sets subject.
func (b *ProcedureBuilder) SetSubject(subject *datatype.Reference) *ProcedureBuilder {
	return b.set("subject", subject.ToArray())
}

// SetEncounter sets encounter.
func (b *ProcedureBuilder) SetEncounter(encounter *datatype.Reference) *ProcedureBuilder {
	return b.set("encounter", encounter.ToArray())
}

// AddBasedOn appends a basedOn reference.
func (b *ProcedureBuilder) AddBasedOn(basedOn *datatype.Reference) *ProcedureBuilder {
	return b.push("basedOn", basedOn.ToArray())
}

// AddReasonCode appends a reasonCode.
func (b *ProcedureBuilder) AddReasonCode(reason *datatype.CodeableConcept) *ProcedureBuilder {
	return b.push("reasonCode", reason.ToArray())
}

// SetPerformedDateTime sets performedDateTime.
func (b *ProcedureBuilder) SetPerformedDateTime(v string) *ProcedureBuilder {
	return b.set("performedDateTime", v)
}

// SetPerformedPeriod sets performedPeriod.
func (b *ProcedureBuilder) SetPerformedPeriod(v *datatype.Period) *ProcedureBuilder {
	return b.set("performedPeriod", v.ToArray())
}

// SetPerformedString sets performedString.
func (b *ProcedureBuilder) SetPerformedString(v string) *ProcedureBuilder {
	return b.set("performedString", v)
}

// SetPerformedAge sets performedAge. The PHP builder types this as Range; kept for parity.
func (b *ProcedureBuilder) SetPerformedAge(v *datatype.Range) *ProcedureBuilder {
	return b.set("performedAge", v.ToArray())
}

// SetPerformedRange sets performedRange.
func (b *ProcedureBuilder) SetPerformedRange(v *datatype.Range) *ProcedureBuilder {
	return b.set("performedRange", v.ToArray())
}

// AddPerformer appends a performer. function and onBehalfOf are optional.
func (b *ProcedureBuilder) AddPerformer(actor *datatype.Reference, function *datatype.CodeableConcept, onBehalfOf *datatype.Reference) *ProcedureBuilder {
	performer := map[string]interface{}{"actor": actor.ToArray()}
	if function != nil {
		performer["function"] = function.ToArray()
	}
	if onBehalfOf != nil {
		performer["onBehalfOf"] = onBehalfOf.ToArray()
	}
	return b.push("performer", performer)
}

// SetOutcome sets outcome.
func (b *ProcedureBuilder) SetOutcome(outcome *datatype.CodeableConcept) *ProcedureBuilder {
	return b.set("outcome", outcome.ToArray())
}

// AddReport appends a report reference.
func (b *ProcedureBuilder) AddReport(report *datatype.Reference) *ProcedureBuilder {
	return b.push("report", report.ToArray())
}

// AddFollowUp appends a followUp.
func (b *ProcedureBuilder) AddFollowUp(followUp *datatype.CodeableConcept) *ProcedureBuilder {
	return b.push("followUp", followUp.ToArray())
}

// AddNote appends a note, either as a typed Annotation or as plain text.
func (b *ProcedureBuilder) AddNote(note interface{}) *ProcedureBuilder {
	switch v := note.(type) {
	case *datatype.Annotation:
		return b.push("note", v.ToArray())
	case string:
		return b.push("note", map[string]interface{}{"text": v})
	}
	return b
}

// AddFocalDevice appends a focalDevice. manufactureItem and device are optional.
func (b *ProcedureBuilder) AddFocalDevice(action *datatype.CodeableConcept, manufactureItem, device *datatype.Reference) *ProcedureBuilder {
	focal := map[string]interface{}{"action": action.ToArray()}
	if device != nil {
		focal["device"] = device.ToArray()
	}
	if manufactureItem != nil {
		focal["manufactureItem"] = manufactureItem.ToArray()
	}
	return b.push("focalDevice", focal)
}

// AddUsedReference appends a usedReference. type is optional.
func (b *ProcedureBuilder) AddUsedReference(reference *datatype.Reference, typ *datatype.CodeableConcept) *ProcedureBuilder {
	used := reference.ToArray()
	if typ != nil {
		used["type"] = typ.ToArray()
	}
	return b.push("usedReference", used)
}

// AddUsedCode appends a usedCode.
func (b *ProcedureBuilder) AddUsedCode(usedCode *datatype.CodeableConcept) *ProcedureBuilder {
	return b.push("usedCode", usedCode.ToArray())
}

// AddBodySite appends a bodySite.
func (b *ProcedureBuilder) AddBodySite(bodySite *datatype.CodeableConcept) *ProcedureBuilder {
	return b.push("bodySite", bodySite.ToArray())
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *ProcedureBuilder) AddExtension(url string, value interface{}, valueType string) *ProcedureBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *ProcedureBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *ProcedureBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
