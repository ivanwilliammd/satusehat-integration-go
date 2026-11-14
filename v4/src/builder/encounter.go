package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// EncounterBuilder builds a FHIR R4 Encounter resource.
//
// Method-for-method counterpart of PayloadBuilderEncounter in the PHP SDK. Optional arguments are
// pointers: pass nil to omit the element, which is what the PHP nullable parameters do.
//
// @link https://www.hl7.org/fhir/encounter.html
type EncounterBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewEncounterBuilder creates an empty Encounter with its resourceType set.
func NewEncounterBuilder() *EncounterBuilder {
	b := &EncounterBuilder{ResourceType: "Encounter", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *EncounterBuilder) set(k string, v interface{}) *EncounterBuilder {
	b.Data[k] = v
	return b
}

func (b *EncounterBuilder) push(k string, v interface{}) *EncounterBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *EncounterBuilder) SetId(id string) *EncounterBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *EncounterBuilder) AddIdentifier(identifier *datatype.Identifier) *EncounterBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetStatus sets status (planned | arrived | triaged | in-progress | onleave | finished | cancelled).
func (b *EncounterBuilder) SetStatus(status string) *EncounterBuilder { return b.set("status", status) }

// SetClass sets class (a Coding: ambulatory | emergency | inpatient | home | virtual …).
func (b *EncounterBuilder) SetClass(class *datatype.Coding) *EncounterBuilder {
	return b.set("class", class.ToArray())
}

// AddType appends a type.
func (b *EncounterBuilder) AddType(t *datatype.CodeableConcept) *EncounterBuilder {
	return b.push("type", t.ToArray())
}

// SetSubject sets subject (usually a Patient reference).
func (b *EncounterBuilder) SetSubject(subject *datatype.Reference) *EncounterBuilder {
	return b.set("subject", subject.ToArray())
}

// AddParticipant appends a participant. type and period are optional.
func (b *EncounterBuilder) AddParticipant(individual *datatype.Reference, t *datatype.CodeableConcept, period *datatype.Period) *EncounterBuilder {
	participant := map[string]interface{}{"individual": individual.ToArray()}
	if t != nil {
		// participant.type is 0..* in FHIR, so the PHP builder wraps it in a list; match that.
		participant["type"] = []interface{}{t.ToArray()}
	}
	if period != nil {
		participant["period"] = period.ToArray()
	}
	return b.push("participant", participant)
}

// AddLocation appends a location. status and physicalType are optional.
func (b *EncounterBuilder) AddLocation(location *datatype.Reference, status string, physicalType *datatype.CodeableConcept) *EncounterBuilder {
	entry := map[string]interface{}{"location": location.ToArray()}
	if status != "" {
		entry["status"] = status
	}
	if physicalType != nil {
		entry["physicalType"] = physicalType.ToArray()
	}
	return b.push("location", entry)
}

// SetPeriod sets period.
func (b *EncounterBuilder) SetPeriod(period *datatype.Period) *EncounterBuilder {
	return b.set("period", period.ToArray())
}

// SetServiceProvider sets serviceProvider.
func (b *EncounterBuilder) SetServiceProvider(provider *datatype.Reference) *EncounterBuilder {
	return b.set("serviceProvider", provider.ToArray())
}

// AddDiagnosis appends a diagnosis. rank, use and role are optional.
func (b *EncounterBuilder) AddDiagnosis(condition *datatype.Reference, rank int, use *datatype.CodeableConcept, role *datatype.CodeableConcept) *EncounterBuilder {
	diagnosis := map[string]interface{}{"condition": condition.ToArray()}
	if rank > 0 {
		diagnosis["rank"] = rank
	}
	if use != nil {
		diagnosis["use"] = use.ToArray()
	}
	if role != nil {
		diagnosis["role"] = role.ToArray()
	}
	return b.push("diagnosis", diagnosis)
}

// AddReasonCode appends a reasonCode.
func (b *EncounterBuilder) AddReasonCode(reason *datatype.CodeableConcept) *EncounterBuilder {
	return b.push("reasonCode", reason.ToArray())
}

// AddReasonReference appends a reasonReference.
func (b *EncounterBuilder) AddReasonReference(reason *datatype.Reference) *EncounterBuilder {
	return b.push("reasonReference", reason.ToArray())
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *EncounterBuilder) AddExtension(url string, value interface{}, valueType string) *EncounterBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *EncounterBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *EncounterBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
