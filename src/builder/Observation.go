package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// ObservationBuilder builds FHIR Observation payload
type ObservationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewObservationBuilder() *ObservationBuilder {
	b := &ObservationBuilder{ResourceType: "Observation", Data: make(map[string]interface{})}
	return b
}

func (b *ObservationBuilder) setId(id string) *ObservationBuilder {
	b.Data["id"] = id
	return b
}

func (b *ObservationBuilder) addIdentifier(identifier *datatype.Identifier) *ObservationBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *ObservationBuilder) setStatus(status string) *ObservationBuilder {
	b.Data["status"] = status
	return b
}

func (b *ObservationBuilder) addCategory(category *datatype.CodeableConcept) *ObservationBuilder {
	if _, ok := b.Data["category"]; !ok {
		b.Data["category"] = make([]interface{}, 0)
	}
	b.Data["category"] = append(b.Data["category"].([]interface{}), category.ToArray())
	return b
}

func (b *ObservationBuilder) setCode(code *datatype.CodeableConcept) *ObservationBuilder {
	b.Data["code"] = code.ToArray()
	return b
}

func (b *ObservationBuilder) setSubject(subject *datatype.Reference) *ObservationBuilder {
	b.Data["subject"] = subject.ToArray()
	return b
}

func (b *ObservationBuilder) setEncounter(encounter *datatype.Reference) *ObservationBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

func (b *ObservationBuilder) setEffectiveDateTime(dateTime string) *ObservationBuilder {
	b.Data["effectiveDateTime"] = dateTime
	return b
}

func (b *ObservationBuilder) setEffectivePeriod(period *datatype.Period) *ObservationBuilder {
	b.Data["effectivePeriod"] = period.ToArray()
	return b
}

func (b *ObservationBuilder) setEffectiveInstant(instant string) *ObservationBuilder {
	b.Data["effectiveInstant"] = instant
	return b
}

func (b *ObservationBuilder) setEffectiveTiming(timing *datatype.Timing) *ObservationBuilder {
	b.Data["effectiveTiming"] = timing.ToArray()
	return b
}

func (b *ObservationBuilder) setValueQuantity(value *datatype.Quantity) *ObservationBuilder {
	b.Data["valueQuantity"] = value.ToArray()
	return b
}

func (b *ObservationBuilder) setValueCodeableConcept(value *datatype.CodeableConcept) *ObservationBuilder {
	b.Data["valueCodeableConcept"] = value.ToArray()
	return b
}

func (b *ObservationBuilder) setValueString(value string) *ObservationBuilder {
	b.Data["valueString"] = value
	return b
}

func (b *ObservationBuilder) setValueBoolean(value bool) *ObservationBuilder {
	b.Data["valueBoolean"] = value
	return b
}

func (b *ObservationBuilder) setValueInteger(value int) *ObservationBuilder {
	b.Data["valueInteger"] = value
	return b
}

func (b *ObservationBuilder) setValueRange(value *datatype.Range) *ObservationBuilder {
	b.Data["valueRange"] = value.ToArray()
	return b
}

func (b *ObservationBuilder) setValueRatio(value *datatype.Ratio) *ObservationBuilder {
	b.Data["valueRatio"] = value.ToArray()
	return b
}

func (b *ObservationBuilder) setValueTime(value string) *ObservationBuilder {
	b.Data["valueTime"] = value
	return b
}

func (b *ObservationBuilder) setValueDateTime(value string) *ObservationBuilder {
	b.Data["valueDateTime"] = value
	return b
}

func (b *ObservationBuilder) setValuePeriod(value *datatype.Period) *ObservationBuilder {
	b.Data["valuePeriod"] = value.ToArray()
	return b
}

func (b *ObservationBuilder) addInterpretation(interpretation *datatype.CodeableConcept) *ObservationBuilder {
	if _, ok := b.Data["interpretation"]; !ok {
		b.Data["interpretation"] = make([]interface{}, 0)
	}
	b.Data["interpretation"] = append(b.Data["interpretation"].([]interface{}), interpretation.ToArray())
	return b
}

func (b *ObservationBuilder) addNote(note *datatype.Annotation) *ObservationBuilder {
	if _, ok := b.Data["note"]; !ok {
		b.Data["note"] = make([]interface{}, 0)
	}
	b.Data["note"] = append(b.Data["note"].([]interface{}), note.ToArray())
	return b
}

func (b *ObservationBuilder) addBodySite(bodySite *datatype.CodeableConcept) *ObservationBuilder {
	if _, ok := b.Data["bodySite"]; !ok {
		b.Data["bodySite"] = make([]interface{}, 0)
	}
	b.Data["bodySite"] = append(b.Data["bodySite"].([]interface{}), bodySite.ToArray())
	return b
}

func (b *ObservationBuilder) setMethod(method *datatype.CodeableConcept) *ObservationBuilder {
	b.Data["method"] = method.ToArray()
	return b
}

func (b *ObservationBuilder) setSpecimen(specimen *datatype.Reference) *ObservationBuilder {
	b.Data["specimen"] = specimen.ToArray()
	return b
}

func (b *ObservationBuilder) setDevice(device *datatype.Reference) *ObservationBuilder {
	b.Data["device"] = device.ToArray()
	return b
}

func (b *ObservationBuilder) addReferenceRange(
	low, high *datatype.Quantity,
	typeCode *datatype.CodeableConcept,
	text *string,
) *ObservationBuilder {
	refRange := make(map[string]interface{})
	if low != nil || high != nil {
		if low != nil {
			refRange["low"] = low.ToArray()
		}
		if high != nil {
			refRange["high"] = high.ToArray()
		}
	}
	if typeCode != nil {
		refRange["type"] = typeCode.ToArray()
	}
	if text != nil {
		refRange["text"] = *text
	}
	if _, ok := b.Data["referenceRange"]; !ok {
		b.Data["referenceRange"] = make([]interface{}, 0)
	}
	b.Data["referenceRange"] = append(b.Data["referenceRange"].([]interface{}), refRange)
	return b
}

func (b *ObservationBuilder) addComponent(code *datatype.CodeableConcept, value interface{}) *ObservationBuilder {
	component := map[string]interface{}{"code": code.ToArray()}
	switch v := value.(type) {
	case *datatype.Quantity:
		component["valueQuantity"] = v.ToArray()
	case *datatype.CodeableConcept:
		component["valueCodeableConcept"] = v.ToArray()
	case *datatype.Range:
		component["valueRange"] = v.ToArray()
	case *datatype.Ratio:
		component["valueRatio"] = v.ToArray()
	case *datatype.SampledData:
		component["valueSampledData"] = v.ToArray()
	case string:
		component["valueString"] = v
	case int:
		component["valueInteger"] = v
	case bool:
		component["valueBoolean"] = v
	}
	if _, ok := b.Data["component"]; !ok {
		b.Data["component"] = make([]interface{}, 0)
	}
	b.Data["component"] = append(b.Data["component"].([]interface{}), component)
	return b
}

func (b *ObservationBuilder) addExtension(url string, value interface{}, valueType string) *ObservationBuilder {
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

func (b *ObservationBuilder) Build() map[string]interface{} {
	return b.Data
}
