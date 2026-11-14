package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// MedicationRequestBuilder builds a FHIR R4 MedicationRequest resource.
//
// Method-for-method counterpart of PayloadBuilderMedicationRequest in the PHP SDK, including the
// nested dispenseRequest element and its incremental setters.
//
// @link https://www.hl7.org/fhir/medicationrequest.html
type MedicationRequestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewMedicationRequestBuilder creates an empty MedicationRequest with its resourceType set.
func NewMedicationRequestBuilder() *MedicationRequestBuilder {
	b := &MedicationRequestBuilder{ResourceType: "MedicationRequest", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *MedicationRequestBuilder) set(k string, v interface{}) *MedicationRequestBuilder {
	b.Data[k] = v
	return b
}

func (b *MedicationRequestBuilder) push(k string, v interface{}) *MedicationRequestBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

func (b *MedicationRequestBuilder) dispenseRequest() map[string]interface{} {
	if dr, ok := b.Data["dispenseRequest"].(map[string]interface{}); ok {
		return dr
	}
	return map[string]interface{}{}
}

// SetId sets the logical id.
func (b *MedicationRequestBuilder) SetId(id string) *MedicationRequestBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *MedicationRequestBuilder) AddIdentifier(identifier *datatype.Identifier) *MedicationRequestBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetStatus sets status (active | on-hold | cancelled | completed | entered-in-error | stopped | draft | unknown).
func (b *MedicationRequestBuilder) SetStatus(status string) *MedicationRequestBuilder {
	return b.set("status", status)
}

// SetIntent sets intent (proposal | plan | order | original-order | reflex-order | filler-order | instance-order | option).
func (b *MedicationRequestBuilder) SetIntent(intent string) *MedicationRequestBuilder {
	return b.set("intent", intent)
}

// AddCategory appends a category.
func (b *MedicationRequestBuilder) AddCategory(category *datatype.CodeableConcept) *MedicationRequestBuilder {
	return b.push("category", category.ToArray())
}

// SetMedicationCodeableConcept sets medicationCodeableConcept.
func (b *MedicationRequestBuilder) SetMedicationCodeableConcept(medication *datatype.CodeableConcept) *MedicationRequestBuilder {
	return b.set("medicationCodeableConcept", medication.ToArray())
}

// SetMedicationReference sets medicationReference.
func (b *MedicationRequestBuilder) SetMedicationReference(medication *datatype.Reference) *MedicationRequestBuilder {
	return b.set("medicationReference", medication.ToArray())
}

// SetSubject sets subject.
func (b *MedicationRequestBuilder) SetSubject(subject *datatype.Reference) *MedicationRequestBuilder {
	return b.set("subject", subject.ToArray())
}

// SetEncounter sets encounter.
func (b *MedicationRequestBuilder) SetEncounter(encounter *datatype.Reference) *MedicationRequestBuilder {
	return b.set("encounter", encounter.ToArray())
}

// SetAuthoredOn sets authoredOn.
func (b *MedicationRequestBuilder) SetAuthoredOn(authoredOn string) *MedicationRequestBuilder {
	return b.set("authoredOn", authoredOn)
}

// SetRequester sets requester.
func (b *MedicationRequestBuilder) SetRequester(requester *datatype.Reference) *MedicationRequestBuilder {
	return b.set("requester", requester.ToArray())
}

// SetRecorder sets recorder.
func (b *MedicationRequestBuilder) SetRecorder(recorder *datatype.Reference) *MedicationRequestBuilder {
	return b.set("recorder", recorder.ToArray())
}

// AddReasonCode appends a reasonCode.
func (b *MedicationRequestBuilder) AddReasonCode(reason *datatype.CodeableConcept) *MedicationRequestBuilder {
	return b.push("reasonCode", reason.ToArray())
}

// AddReasonReference appends a reasonReference.
func (b *MedicationRequestBuilder) AddReasonReference(reason *datatype.Reference) *MedicationRequestBuilder {
	return b.push("reasonReference", reason.ToArray())
}

// AddDosageInstruction appends a dosageInstruction.
func (b *MedicationRequestBuilder) AddDosageInstruction(dosage *datatype.Dosage) *MedicationRequestBuilder {
	return b.push("dosageInstruction", dosage.ToArray())
}

// SetDispenseRequest replaces dispenseRequest. Every argument is optional; pass nil (or 0) to omit.
func (b *MedicationRequestBuilder) SetDispenseRequest(validityPeriod *datatype.Reference, numberOfRepeatsAllowed int, quantity *datatype.Quantity, expectedSupplyDuration *datatype.Duration, performer *datatype.CodeableConcept) *MedicationRequestBuilder {
	dr := map[string]interface{}{}
	if validityPeriod != nil {
		dr["validityPeriod"] = validityPeriod.ToArray()
	}
	if numberOfRepeatsAllowed != 0 {
		dr["numberOfRepeatsAllowed"] = numberOfRepeatsAllowed
	}
	if quantity != nil {
		dr["quantity"] = quantity.ToArray()
	}
	if expectedSupplyDuration != nil {
		dr["expectedSupplyDuration"] = expectedSupplyDuration.ToArray()
	}
	if performer != nil {
		dr["performer"] = performer.ToArray()
	}
	return b.set("dispenseRequest", dr)
}

// SetDispenseInterval sets dispenseRequest.dispenseInterval, keeping the rest of the element.
func (b *MedicationRequestBuilder) SetDispenseInterval(interval *datatype.Duration) *MedicationRequestBuilder {
	dr := b.dispenseRequest()
	dr["dispenseInterval"] = interval.ToArray()
	return b.set("dispenseRequest", dr)
}

// SetInitialFill sets dispenseRequest.initialFill, keeping the rest of the element.
func (b *MedicationRequestBuilder) SetInitialFill(initialFill *datatype.Quantity) *MedicationRequestBuilder {
	dr := b.dispenseRequest()
	dr["initialFill"] = initialFill.ToArray()
	return b.set("dispenseRequest", dr)
}

// SetSubstitution sets substitution. allowed is optional; nil omits allowedBoolean.
func (b *MedicationRequestBuilder) SetSubstitution(substitution *datatype.CodeableConcept, allowed *bool) *MedicationRequestBuilder {
	data := map[string]interface{}{"code": substitution.ToArray()}
	if allowed != nil {
		data["allowedBoolean"] = *allowed
	}
	return b.set("substitution", data)
}

// AddPriorPrescription appends a priorPrescription.
func (b *MedicationRequestBuilder) AddPriorPrescription(prior *datatype.Reference) *MedicationRequestBuilder {
	return b.push("priorPrescription", prior.ToArray())
}

// AddDetectedIssue appends a detectedIssue.
func (b *MedicationRequestBuilder) AddDetectedIssue(issue *datatype.Reference) *MedicationRequestBuilder {
	return b.push("detectedIssue", issue.ToArray())
}

// AddEventHistory appends an eventHistory entry.
func (b *MedicationRequestBuilder) AddEventHistory(event *datatype.Reference) *MedicationRequestBuilder {
	return b.push("eventHistory", event.ToArray())
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *MedicationRequestBuilder) AddExtension(url string, value interface{}, valueType string) *MedicationRequestBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *MedicationRequestBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *MedicationRequestBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
