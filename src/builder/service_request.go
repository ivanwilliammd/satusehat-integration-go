package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// ServiceRequestBuilder builds FHIR ServiceRequest payload
type ServiceRequestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewServiceRequestBuilder() *ServiceRequestBuilder {
	b := &ServiceRequestBuilder{ResourceType: "ServiceRequest", Data: make(map[string]interface{})}
	return b
}

func (b *ServiceRequestBuilder) setId(id string) *ServiceRequestBuilder {
	b.Data["id"] = id
	return b
}

func (b *ServiceRequestBuilder) addIdentifier(identifier *datatype.Identifier) *ServiceRequestBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *ServiceRequestBuilder) setRequisition(identifier *datatype.Identifier) *ServiceRequestBuilder {
	b.Data["requisition"] = identifier.ToArray()
	return b
}

func (b *ServiceRequestBuilder) setStatus(status string) *ServiceRequestBuilder {
	b.Data["status"] = status
	return b
}

func (b *ServiceRequestBuilder) setIntent(intent string) *ServiceRequestBuilder {
	b.Data["intent"] = intent
	return b
}

func (b *ServiceRequestBuilder) addCategory(category *datatype.CodeableConcept) *ServiceRequestBuilder {
	if _, ok := b.Data["category"]; !ok {
		b.Data["category"] = make([]interface{}, 0)
	}
	b.Data["category"] = append(b.Data["category"].([]interface{}), category.ToArray())
	return b
}

func (b *ServiceRequestBuilder) setPriority(priority string) *ServiceRequestBuilder {
	b.Data["priority"] = priority
	return b
}

func (b *ServiceRequestBuilder) setDoNotPerform(doNotPerform bool) *ServiceRequestBuilder {
	b.Data["doNotPerform"] = doNotPerform
	return b
}

func (b *ServiceRequestBuilder) setCode(code *datatype.CodeableConcept) *ServiceRequestBuilder {
	b.Data["code"] = code.ToArray()
	return b
}

func (b *ServiceRequestBuilder) setQuantityQuantity(quantity *datatype.Quantity) *ServiceRequestBuilder {
	b.Data["quantityQuantity"] = quantity.ToArray()
	return b
}

func (b *ServiceRequestBuilder) setSubject(subject *datatype.Reference) *ServiceRequestBuilder {
	b.Data["subject"] = subject.ToArray()
	return b
}

func (b *ServiceRequestBuilder) setEncounter(encounter *datatype.Reference) *ServiceRequestBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

func (b *ServiceRequestBuilder) setOccurrenceDateTime(dateTime string) *ServiceRequestBuilder {
	b.Data["occurrenceDateTime"] = dateTime
	return b
}

func (b *ServiceRequestBuilder) setAuthoredOn(authoredOn string) *ServiceRequestBuilder {
	b.Data["authoredOn"] = authoredOn
	return b
}

func (b *ServiceRequestBuilder) setRequester(requester *datatype.Reference) *ServiceRequestBuilder {
	b.Data["requester"] = requester.ToArray()
	return b
}

func (b *ServiceRequestBuilder) addPerformer(performer *datatype.Reference) *ServiceRequestBuilder {
	if _, ok := b.Data["performer"]; !ok {
		b.Data["performer"] = make([]interface{}, 0)
	}
	b.Data["performer"] = append(b.Data["performer"].([]interface{}), performer.ToArray())
	return b
}

func (b *ServiceRequestBuilder) addReasonCode(reasonCode *datatype.CodeableConcept) *ServiceRequestBuilder {
	if _, ok := b.Data["reasonCode"]; !ok {
		b.Data["reasonCode"] = make([]interface{}, 0)
	}
	b.Data["reasonCode"] = append(b.Data["reasonCode"].([]interface{}), reasonCode.ToArray())
	return b
}

func (b *ServiceRequestBuilder) addSupportingInfo(supportingInfo *datatype.Reference) *ServiceRequestBuilder {
	if _, ok := b.Data["supportingInfo"]; !ok {
		b.Data["supportingInfo"] = make([]interface{}, 0)
	}
	b.Data["supportingInfo"] = append(b.Data["supportingInfo"].([]interface{}), supportingInfo.ToArray())
	return b
}

func (b *ServiceRequestBuilder) addSpecimen(specimen *datatype.Reference) *ServiceRequestBuilder {
	if _, ok := b.Data["specimen"]; !ok {
		b.Data["specimen"] = make([]interface{}, 0)
	}
	b.Data["specimen"] = append(b.Data["specimen"].([]interface{}), specimen.ToArray())
	return b
}

func (b *ServiceRequestBuilder) addNote(text string) *ServiceRequestBuilder {
	if _, ok := b.Data["note"]; !ok {
		b.Data["note"] = make([]interface{}, 0)
	}
	b.Data["note"] = append(b.Data["note"].([]interface{}), map[string]interface{}{"text": text})
	return b
}

func (b *ServiceRequestBuilder) setPatientInstruction(patientInstruction string) *ServiceRequestBuilder {
	b.Data["patientInstruction"] = patientInstruction
	return b
}

func (b *ServiceRequestBuilder) addRelevantHistory(relevantHistory *datatype.Reference) *ServiceRequestBuilder {
	if _, ok := b.Data["relevantHistory"]; !ok {
		b.Data["relevantHistory"] = make([]interface{}, 0)
	}
	b.Data["relevantHistory"] = append(b.Data["relevantHistory"].([]interface{}), relevantHistory.ToArray())
	return b
}

func (b *ServiceRequestBuilder) addExtension(url string, value interface{}, valueType string) *ServiceRequestBuilder {
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

func (b *ServiceRequestBuilder) Build() map[string]interface{} {
	return b.Data
}
