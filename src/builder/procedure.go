package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// ProcedureBuilder builds FHIR Procedure payload
type ProcedureBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewProcedureBuilder() *ProcedureBuilder {
	b := &ProcedureBuilder{ResourceType: "Procedure", Data: make(map[string]interface{})}
	return b
}

func (b *ProcedureBuilder) setId(id string) *ProcedureBuilder {
	b.Data["id"] = id
	return b
}

func (b *ProcedureBuilder) addIdentifier(identifier *datatype.Identifier) *ProcedureBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *ProcedureBuilder) setStatus(status string) *ProcedureBuilder {
	b.Data["status"] = status
	return b
}

func (b *ProcedureBuilder) setCategory(category *datatype.CodeableConcept) *ProcedureBuilder {
	b.Data["category"] = category.ToArray()
	return b
}

func (b *ProcedureBuilder) setCode(code *datatype.CodeableConcept) *ProcedureBuilder {
	b.Data["code"] = code.ToArray()
	return b
}

func (b *ProcedureBuilder) setSubject(subject *datatype.Reference) *ProcedureBuilder {
	b.Data["subject"] = subject.ToArray()
	return b
}

func (b *ProcedureBuilder) setEncounter(encounter *datatype.Reference) *ProcedureBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

func (b *ProcedureBuilder) setPerformedDateTime(dateTime string) *ProcedureBuilder {
	b.Data["performedDateTime"] = dateTime
	return b
}

func (b *ProcedureBuilder) setPerformedPeriod(period *datatype.Period) *ProcedureBuilder {
	b.Data["performedPeriod"] = period.ToArray()
	return b
}

func (b *ProcedureBuilder) setPerformedString(performedString string) *ProcedureBuilder {
	b.Data["performedString"] = performedString
	return b
}

func (b *ProcedureBuilder) setPerformedAge(age *datatype.Range) *ProcedureBuilder {
	b.Data["performedAge"] = age.ToArray()
	return b
}

func (b *ProcedureBuilder) setPerformedRange(rng *datatype.Range) *ProcedureBuilder {
	b.Data["performedRange"] = rng.ToArray()
	return b
}

func (b *ProcedureBuilder) addPerformer(
	actor *datatype.Reference,
	fn *datatype.CodeableConcept,
	onBehalfOf *datatype.Reference,
) *ProcedureBuilder {
	performer := map[string]interface{}{"actor": actor.ToArray()}
	if fn != nil {
		performer["function"] = fn.ToArray()
	}
	if onBehalfOf != nil {
		performer["onBehalfOf"] = onBehalfOf.ToArray()
	}
	if _, ok := b.Data["performer"]; !ok {
		b.Data["performer"] = make([]interface{}, 0)
	}
	b.Data["performer"] = append(b.Data["performer"].([]interface{}), performer)
	return b
}

func (b *ProcedureBuilder) setOutcome(outcome *datatype.CodeableConcept) *ProcedureBuilder {
	b.Data["outcome"] = outcome.ToArray()
	return b
}

func (b *ProcedureBuilder) addReport(report *datatype.Reference) *ProcedureBuilder {
	if _, ok := b.Data["report"]; !ok {
		b.Data["report"] = make([]interface{}, 0)
	}
	b.Data["report"] = append(b.Data["report"].([]interface{}), report.ToArray())
	return b
}

func (b *ProcedureBuilder) addFollowUp(followUp *datatype.CodeableConcept) *ProcedureBuilder {
	if _, ok := b.Data["followUp"]; !ok {
		b.Data["followUp"] = make([]interface{}, 0)
	}
	b.Data["followUp"] = append(b.Data["followUp"].([]interface{}), followUp.ToArray())
	return b
}

func (b *ProcedureBuilder) addNote(note *datatype.Annotation) *ProcedureBuilder {
	if _, ok := b.Data["note"]; !ok {
		b.Data["note"] = make([]interface{}, 0)
	}
	b.Data["note"] = append(b.Data["note"].([]interface{}), note.ToArray())
	return b
}

func (b *ProcedureBuilder) addFocalDevice(
	action *datatype.CodeableConcept,
	device *datatype.Reference,
) *ProcedureBuilder {
	focalDevice := map[string]interface{}{"action": action.ToArray()}
	if device != nil {
		focalDevice["device"] = device.ToArray()
	}
	if _, ok := b.Data["focalDevice"]; !ok {
		b.Data["focalDevice"] = make([]interface{}, 0)
	}
	b.Data["focalDevice"] = append(b.Data["focalDevice"].([]interface{}), focalDevice)
	return b
}

func (b *ProcedureBuilder) addUsedReference(reference *datatype.Reference, typ *datatype.CodeableConcept) *ProcedureBuilder {
	used := reference.ToArray()
	if typ != nil {
		used["type"] = typ.ToArray()
	}
	if _, ok := b.Data["usedReference"]; !ok {
		b.Data["usedReference"] = make([]interface{}, 0)
	}
	b.Data["usedReference"] = append(b.Data["usedReference"].([]interface{}), used)
	return b
}

func (b *ProcedureBuilder) addUsedCode(usedCode *datatype.CodeableConcept) *ProcedureBuilder {
	if _, ok := b.Data["usedCode"]; !ok {
		b.Data["usedCode"] = make([]interface{}, 0)
	}
	b.Data["usedCode"] = append(b.Data["usedCode"].([]interface{}), usedCode.ToArray())
	return b
}

func (b *ProcedureBuilder) addBodySite(bodySite *datatype.CodeableConcept) *ProcedureBuilder {
	if _, ok := b.Data["bodySite"]; !ok {
		b.Data["bodySite"] = make([]interface{}, 0)
	}
	b.Data["bodySite"] = append(b.Data["bodySite"].([]interface{}), bodySite.ToArray())
	return b
}

func (b *ProcedureBuilder) addExtension(url string, value interface{}, valueType string) *ProcedureBuilder {
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

func (b *ProcedureBuilder) Build() map[string]interface{} {
	return b.Data
}
