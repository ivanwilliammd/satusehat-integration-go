package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// EncounterBuilder builds FHIR Encounter payload
type EncounterBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewEncounterBuilder() *EncounterBuilder {
	b := &EncounterBuilder{ResourceType: "Encounter", Data: make(map[string]interface{})}
	return b
}

func (b *EncounterBuilder) setId(id string) *EncounterBuilder {
	b.Data["id"] = id
	return b
}

func (b *EncounterBuilder) addIdentifier(identifier *datatype.Identifier) *EncounterBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *EncounterBuilder) setStatus(status string) *EncounterBuilder {
	b.Data["status"] = status
	return b
}

func (b *EncounterBuilder) setClass(class *datatype.Coding) *EncounterBuilder {
	b.Data["class"] = class.ToArray()
	return b
}

func (b *EncounterBuilder) addType(typeCode *datatype.CodeableConcept) *EncounterBuilder {
	if _, ok := b.Data["type"]; !ok {
		b.Data["type"] = make([]interface{}, 0)
	}
	b.Data["type"] = append(b.Data["type"].([]interface{}), typeCode.ToArray())
	return b
}

func (b *EncounterBuilder) setSubject(subject *datatype.Reference) *EncounterBuilder {
	b.Data["subject"] = subject.ToArray()
	return b
}

func (b *EncounterBuilder) addParticipant(
	individual *datatype.Reference,
	ptype *datatype.CodeableConcept,
	period *datatype.Period,
) *EncounterBuilder {
	participant := map[string]interface{}{"individual": individual.ToArray()}
	if ptype != nil {
		participant["type"] = []interface{}{ptype.ToArray()}
	}
	if period != nil {
		participant["period"] = period.ToArray()
	}
	if _, ok := b.Data["participant"]; !ok {
		b.Data["participant"] = make([]interface{}, 0)
	}
	b.Data["participant"] = append(b.Data["participant"].([]interface{}), participant)
	return b
}

func (b *EncounterBuilder) addLocation(
	location *datatype.Reference,
	status *string,
	physicalType *datatype.CodeableConcept,
) *EncounterBuilder {
	loc := map[string]interface{}{"location": location.ToArray()}
	if status != nil {
		loc["status"] = *status
	}
	if physicalType != nil {
		loc["physicalType"] = physicalType.ToArray()
	}
	if _, ok := b.Data["location"]; !ok {
		b.Data["location"] = make([]interface{}, 0)
	}
	b.Data["location"] = append(b.Data["location"].([]interface{}), loc)
	return b
}

func (b *EncounterBuilder) setPeriod(period *datatype.Period) *EncounterBuilder {
	b.Data["period"] = period.ToArray()
	return b
}

func (b *EncounterBuilder) setServiceProvider(serviceProvider *datatype.Reference) *EncounterBuilder {
	b.Data["serviceProvider"] = serviceProvider.ToArray()
	return b
}

func (b *EncounterBuilder) addDiagnosis(
	condition *datatype.Reference,
	rank *int,
	use *datatype.CodeableConcept,
	role *datatype.CodeableConcept,
) *EncounterBuilder {
	diagnosis := map[string]interface{}{"condition": condition.ToArray()}
	if rank != nil {
		diagnosis["rank"] = *rank
	}
	if use != nil {
		diagnosis["use"] = use.ToArray()
	}
	if role != nil {
		diagnosis["role"] = role.ToArray()
	}
	if _, ok := b.Data["diagnosis"]; !ok {
		b.Data["diagnosis"] = make([]interface{}, 0)
	}
	b.Data["diagnosis"] = append(b.Data["diagnosis"].([]interface{}), diagnosis)
	return b
}

func (b *EncounterBuilder) addReasonCode(reasonCode *datatype.CodeableConcept) *EncounterBuilder {
	if _, ok := b.Data["reasonCode"]; !ok {
		b.Data["reasonCode"] = make([]interface{}, 0)
	}
	b.Data["reasonCode"] = append(b.Data["reasonCode"].([]interface{}), reasonCode.ToArray())
	return b
}

func (b *EncounterBuilder) addReasonReference(reasonReference *datatype.Reference) *EncounterBuilder {
	if _, ok := b.Data["reasonReference"]; !ok {
		b.Data["reasonReference"] = make([]interface{}, 0)
	}
	b.Data["reasonReference"] = append(b.Data["reasonReference"].([]interface{}), reasonReference.ToArray())
	return b
}

func (b *EncounterBuilder) addExtension(url string, value interface{}, valueType string) *EncounterBuilder {
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

func (b *EncounterBuilder) Build() map[string]interface{} {
	return b.Data
}
