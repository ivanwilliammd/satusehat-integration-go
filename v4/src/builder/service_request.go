package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// ServiceRequestBuilder builds a FHIR R4 ServiceRequest resource.
//
// Method-for-method counterpart of PayloadBuilderServiceRequest in the PHP SDK: requisition,
// intent, priority, doNotPerform, quantityQuantity, occurrenceDateTime, authoredOn, requester,
// performer, reasonCode, supportingInfo, specimen, note, patientInstruction, relevantHistory.
//
// @link https://www.hl7.org/fhir/servicerequest.html
type ServiceRequestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewServiceRequestBuilder creates an empty ServiceRequest with its resourceType set.
func NewServiceRequestBuilder() *ServiceRequestBuilder {
	b := &ServiceRequestBuilder{ResourceType: "ServiceRequest", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *ServiceRequestBuilder) set(k string, v interface{}) *ServiceRequestBuilder {
	b.Data[k] = v
	return b
}

func (b *ServiceRequestBuilder) push(k string, v interface{}) *ServiceRequestBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *ServiceRequestBuilder) SetId(id string) *ServiceRequestBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *ServiceRequestBuilder) AddIdentifier(identifier *datatype.Identifier) *ServiceRequestBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetRequisition sets requisition.
func (b *ServiceRequestBuilder) SetRequisition(identifier *datatype.Identifier) *ServiceRequestBuilder {
	return b.set("requisition", identifier.ToArray())
}

// SetStatus sets status (draft | active | on-hold | revoked | completed | entered-in-error | unknown).
func (b *ServiceRequestBuilder) SetStatus(status string) *ServiceRequestBuilder {
	return b.set("status", status)
}

// SetIntent sets intent (proposal | plan | directive | order | original-order | reflex-order | filler-order | instance-order | option).
func (b *ServiceRequestBuilder) SetIntent(intent string) *ServiceRequestBuilder {
	return b.set("intent", intent)
}

// AddCategory appends a category.
func (b *ServiceRequestBuilder) AddCategory(category *datatype.CodeableConcept) *ServiceRequestBuilder {
	return b.push("category", category.ToArray())
}

// SetPriority sets priority (routine | urgent | asap | stat).
func (b *ServiceRequestBuilder) SetPriority(priority string) *ServiceRequestBuilder {
	return b.set("priority", priority)
}

// SetDoNotPerform sets doNotPerform.
func (b *ServiceRequestBuilder) SetDoNotPerform(doNotPerform bool) *ServiceRequestBuilder {
	return b.set("doNotPerform", doNotPerform)
}

// SetCode sets code.
func (b *ServiceRequestBuilder) SetCode(code *datatype.CodeableConcept) *ServiceRequestBuilder {
	return b.set("code", code.ToArray())
}

// SetQuantityQuantity sets quantityQuantity.
func (b *ServiceRequestBuilder) SetQuantityQuantity(quantity *datatype.Quantity) *ServiceRequestBuilder {
	return b.set("quantityQuantity", quantity.ToArray())
}

// SetSubject sets subject.
func (b *ServiceRequestBuilder) SetSubject(subject *datatype.Reference) *ServiceRequestBuilder {
	return b.set("subject", subject.ToArray())
}

// SetEncounter sets encounter.
func (b *ServiceRequestBuilder) SetEncounter(encounter *datatype.Reference) *ServiceRequestBuilder {
	return b.set("encounter", encounter.ToArray())
}

// SetOccurrenceDateTime sets occurrenceDateTime.
func (b *ServiceRequestBuilder) SetOccurrenceDateTime(dateTime string) *ServiceRequestBuilder {
	return b.set("occurrenceDateTime", dateTime)
}

// SetAuthoredOn sets authoredOn.
func (b *ServiceRequestBuilder) SetAuthoredOn(dateTime string) *ServiceRequestBuilder {
	return b.set("authoredOn", dateTime)
}

// SetRequester sets requester.
func (b *ServiceRequestBuilder) SetRequester(requester *datatype.Reference) *ServiceRequestBuilder {
	return b.set("requester", requester.ToArray())
}

// AddPerformer appends a performer.
func (b *ServiceRequestBuilder) AddPerformer(performer *datatype.Reference) *ServiceRequestBuilder {
	return b.push("performer", performer.ToArray())
}

// AddReasonCode appends a reasonCode.
func (b *ServiceRequestBuilder) AddReasonCode(reason *datatype.CodeableConcept) *ServiceRequestBuilder {
	return b.push("reasonCode", reason.ToArray())
}

// AddSupportingInfo appends a supportingInfo reference.
func (b *ServiceRequestBuilder) AddSupportingInfo(info *datatype.Reference) *ServiceRequestBuilder {
	return b.push("supportingInfo", info.ToArray())
}

// AddSpecimen appends a specimen reference.
func (b *ServiceRequestBuilder) AddSpecimen(specimen *datatype.Reference) *ServiceRequestBuilder {
	return b.push("specimen", specimen.ToArray())
}

// AddNote appends a note.
func (b *ServiceRequestBuilder) AddNote(text string) *ServiceRequestBuilder {
	return b.push("note", map[string]interface{}{"text": text})
}

// SetPatientInstruction sets patientInstruction.
func (b *ServiceRequestBuilder) SetPatientInstruction(instruction string) *ServiceRequestBuilder {
	return b.set("patientInstruction", instruction)
}

// AddRelevantHistory appends a relevantHistory reference.
func (b *ServiceRequestBuilder) AddRelevantHistory(history *datatype.Reference) *ServiceRequestBuilder {
	return b.push("relevantHistory", history.ToArray())
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *ServiceRequestBuilder) AddExtension(url string, value interface{}, valueType string) *ServiceRequestBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *ServiceRequestBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *ServiceRequestBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
