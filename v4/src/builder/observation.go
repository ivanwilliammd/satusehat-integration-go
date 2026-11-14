package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/terminology"
)

// ObservationBuilder builds a FHIR R4 Observation resource.
//
// Method-for-method counterpart of PayloadBuilderObservation in the PHP SDK. Strings passed to
// AddCategory and SetCode go through the offline terminology resolver, exactly as the PHP builder
// does, so no terminology server is needed.
//
// @link https://www.hl7.org/fhir/observation.html
type ObservationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewObservationBuilder creates an empty Observation with its resourceType set.
func NewObservationBuilder() *ObservationBuilder {
	b := &ObservationBuilder{ResourceType: "Observation", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *ObservationBuilder) set(k string, v interface{}) *ObservationBuilder {
	b.Data[k] = v
	return b
}

func (b *ObservationBuilder) push(k string, v interface{}) *ObservationBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *ObservationBuilder) SetId(id string) *ObservationBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *ObservationBuilder) AddIdentifier(identifier *datatype.Identifier) *ObservationBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetStatus sets status (registered | preliminary | final | amended | corrected | cancelled | entered-in-error | unknown).
func (b *ObservationBuilder) SetStatus(status string) *ObservationBuilder {
	return b.set("status", status)
}

// AddCategory appends a category. A string is resolved through the offline terminology catalogs.
func (b *ObservationBuilder) AddCategory(category interface{}) *ObservationBuilder {
	if s, ok := category.(string); ok {
		return b.push("category", terminology.Resolve(s))
	}
	return b.push("category", category.(*datatype.CodeableConcept).ToArray())
}

// AddCategoryConcept appends a category given as a typed concept.
func (b *ObservationBuilder) AddCategoryConcept(category *datatype.CodeableConcept) *ObservationBuilder {
	return b.push("category", category.ToArray())
}

// SetCode sets code, either as a catalog key or as a typed concept.
func (b *ObservationBuilder) SetCode(code interface{}) *ObservationBuilder {
	switch v := code.(type) {
	case string:
		return b.set("code", terminology.Resolve(v))
	case *datatype.CodeableConcept:
		return b.set("code", v.ToArray())
	}
	return b
}

// SetSubject sets subject.
func (b *ObservationBuilder) SetSubject(subject *datatype.Reference) *ObservationBuilder {
	return b.set("subject", subject.ToArray())
}

// SetEncounter sets encounter.
func (b *ObservationBuilder) SetEncounter(encounter *datatype.Reference) *ObservationBuilder {
	return b.set("encounter", encounter.ToArray())
}

// SetEffectiveDateTime sets effectiveDateTime.
func (b *ObservationBuilder) SetEffectiveDateTime(dateTime string) *ObservationBuilder {
	return b.set("effectiveDateTime", dateTime)
}

// SetEffectivePeriod sets effectivePeriod.
func (b *ObservationBuilder) SetEffectivePeriod(period *datatype.Period) *ObservationBuilder {
	return b.set("effectivePeriod", period.ToArray())
}

// SetEffectiveInstant sets effectiveInstant.
func (b *ObservationBuilder) SetEffectiveInstant(instant string) *ObservationBuilder {
	return b.set("effectiveInstant", instant)
}

// SetEffectiveTiming sets effectiveTiming.
func (b *ObservationBuilder) SetEffectiveTiming(timing *datatype.Timing) *ObservationBuilder {
	return b.set("effectiveTiming", timing.ToArray())
}

// SetEffectivePeriodStart sets effectivePeriod.start, keeping an existing end.
func (b *ObservationBuilder) SetEffectivePeriodStart(start string) *ObservationBuilder {
	existing, _ := b.Data["effectivePeriod"].(map[string]interface{})
	if existing == nil {
		existing = map[string]interface{}{}
	}
	existing["start"] = start
	return b.set("effectivePeriod", existing)
}

// SetEffectivePeriodEnd sets effectivePeriod.end, keeping an existing start.
func (b *ObservationBuilder) SetEffectivePeriodEnd(end string) *ObservationBuilder {
	existing, _ := b.Data["effectivePeriod"].(map[string]interface{})
	if existing == nil {
		existing = map[string]interface{}{}
	}
	existing["end"] = end
	return b.set("effectivePeriod", existing)
}

// SetValueQuantity sets valueQuantity.
func (b *ObservationBuilder) SetValueQuantity(v *datatype.Quantity) *ObservationBuilder {
	return b.set("valueQuantity", v.ToArray())
}

// SetValueCodeableConcept sets valueCodeableConcept.
func (b *ObservationBuilder) SetValueCodeableConcept(v *datatype.CodeableConcept) *ObservationBuilder {
	return b.set("valueCodeableConcept", v.ToArray())
}

// SetValueString sets valueString.
func (b *ObservationBuilder) SetValueString(v string) *ObservationBuilder {
	return b.set("valueString", v)
}

// SetValueBoolean sets valueBoolean.
func (b *ObservationBuilder) SetValueBoolean(v bool) *ObservationBuilder {
	return b.set("valueBoolean", v)
}

// SetValueInteger sets valueInteger.
func (b *ObservationBuilder) SetValueInteger(v int) *ObservationBuilder {
	return b.set("valueInteger", v)
}

// SetValueRange sets valueRange.
func (b *ObservationBuilder) SetValueRange(v *datatype.Range) *ObservationBuilder {
	return b.set("valueRange", v.ToArray())
}

// SetValueRatio sets valueRatio.
func (b *ObservationBuilder) SetValueRatio(v *datatype.Ratio) *ObservationBuilder {
	return b.set("valueRatio", v.ToArray())
}

// SetValueTime sets valueTime.
func (b *ObservationBuilder) SetValueTime(v string) *ObservationBuilder { return b.set("valueTime", v) }

// SetValueDateTime sets valueDateTime.
func (b *ObservationBuilder) SetValueDateTime(v string) *ObservationBuilder {
	return b.set("valueDateTime", v)
}

// SetValuePeriod sets valuePeriod.
func (b *ObservationBuilder) SetValuePeriod(v *datatype.Period) *ObservationBuilder {
	return b.set("valuePeriod", v.ToArray())
}

// AddInterpretation appends an interpretation.
func (b *ObservationBuilder) AddInterpretation(v *datatype.CodeableConcept) *ObservationBuilder {
	return b.push("interpretation", v.ToArray())
}

// AddNote appends a note, either as a typed Annotation or as plain text.
func (b *ObservationBuilder) AddNote(note interface{}) *ObservationBuilder {
	switch v := note.(type) {
	case *datatype.Annotation:
		return b.push("note", v.ToArray())
	case string:
		return b.push("note", map[string]interface{}{"text": v})
	}
	return b
}

// AddBodySite appends a bodySite.
func (b *ObservationBuilder) AddBodySite(v *datatype.CodeableConcept) *ObservationBuilder {
	return b.push("bodySite", v.ToArray())
}

// SetMethod sets method.
func (b *ObservationBuilder) SetMethod(v *datatype.CodeableConcept) *ObservationBuilder {
	return b.set("method", v.ToArray())
}

// SetSpecimen sets specimen.
func (b *ObservationBuilder) SetSpecimen(v *datatype.Reference) *ObservationBuilder {
	return b.set("specimen", v.ToArray())
}

// SetDevice sets device.
func (b *ObservationBuilder) SetDevice(v *datatype.Reference) *ObservationBuilder {
	return b.set("device", v.ToArray())
}

// AddReferenceRange appends a referenceRange. low, high, type and text are optional.
func (b *ObservationBuilder) AddReferenceRange(low, high *datatype.Quantity, t *datatype.CodeableConcept, text string) *ObservationBuilder {
	refRange := map[string]interface{}{}
	if low != nil || high != nil {
		if low != nil {
			refRange["low"] = low.ToArray()
		}
		if high != nil {
			refRange["high"] = high.ToArray()
		}
	}
	if t != nil {
		refRange["type"] = t.ToArray()
	}
	if text != "" {
		refRange["text"] = text
	}
	return b.push("referenceRange", refRange)
}

// AddComponent appends a component. The value may be a Quantity, CodeableConcept, Range, Ratio,
// SampledData, string, int or bool; the matching value[x] element is written, as in PHP.
func (b *ObservationBuilder) AddComponent(code *datatype.CodeableConcept, value interface{}) *ObservationBuilder {
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
	return b.push("component", component)
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *ObservationBuilder) AddExtension(url string, value interface{}, valueType string) *ObservationBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *ObservationBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *ObservationBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
