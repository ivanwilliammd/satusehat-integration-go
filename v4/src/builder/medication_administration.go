package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// MedicationAdministrationBuilder builds a FHIR R4 MedicationAdministration resource.
//
// Method-for-method counterpart of PayloadBuilderMedicationAdministration in the PHP SDK. The
// reference setters accept either a plain string or a *datatype.Reference, matching refArray() in
// PHP, and write the element unchanged — no type prefixing happens here.
//
// @link https://www.hl7.org/fhir/medicationadministration.html
type MedicationAdministrationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewMedicationAdministrationBuilder creates an empty MedicationAdministration with its resourceType set.
func NewMedicationAdministrationBuilder() *MedicationAdministrationBuilder {
	b := &MedicationAdministrationBuilder{ResourceType: "MedicationAdministration", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *MedicationAdministrationBuilder) set(k string, v interface{}) *MedicationAdministrationBuilder {
	b.Data[k] = v
	return b
}

func (b *MedicationAdministrationBuilder) push(k string, v interface{}) *MedicationAdministrationBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// refArray meniru refArray() PHP: Reference apa adanya, string menjadi {"reference": s}.
func refArray(v interface{}) map[string]interface{} {
	switch r := v.(type) {
	case *datatype.Reference:
		return r.ToArray()
	case string:
		return map[string]interface{}{"reference": r}
	}
	return map[string]interface{}{}
}

// SetId sets the logical id.
func (b *MedicationAdministrationBuilder) SetId(id string) *MedicationAdministrationBuilder {
	return b.set("id", id)
}

// AddIdentifier appends an identifier.
func (b *MedicationAdministrationBuilder) AddIdentifier(identifier *datatype.Identifier) *MedicationAdministrationBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetStatus sets status (in-progress | not-done | on-hold | completed | entered-in-error | stopped | unknown).
func (b *MedicationAdministrationBuilder) SetStatus(status string) *MedicationAdministrationBuilder {
	return b.set("status", status)
}

// SetMedicationReference sets medicationReference from a string or a Reference.
func (b *MedicationAdministrationBuilder) SetMedicationReference(reference interface{}) *MedicationAdministrationBuilder {
	return b.set("medicationReference", refArray(reference))
}

// SetSubject sets subject from a string or a Reference.
func (b *MedicationAdministrationBuilder) SetSubject(subject interface{}) *MedicationAdministrationBuilder {
	return b.set("subject", refArray(subject))
}

// SetContext sets context from a string or a Reference.
func (b *MedicationAdministrationBuilder) SetContext(context interface{}) *MedicationAdministrationBuilder {
	return b.set("context", refArray(context))
}

// SetEffectivePeriod sets effectivePeriod.
func (b *MedicationAdministrationBuilder) SetEffectivePeriod(period *datatype.Period) *MedicationAdministrationBuilder {
	return b.set("effectivePeriod", period.ToArray())
}

// AddPerformer appends a performer.
func (b *MedicationAdministrationBuilder) AddPerformer(actor *datatype.Reference) *MedicationAdministrationBuilder {
	return b.push("performer", map[string]interface{}{"actor": actor.ToArray()})
}

// AddReasonCode appends a reasonCode.
func (b *MedicationAdministrationBuilder) AddReasonCode(reason *datatype.CodeableConcept) *MedicationAdministrationBuilder {
	return b.push("reasonCode", reason.ToArray())
}

// SetRequest sets request.
func (b *MedicationAdministrationBuilder) SetRequest(request *datatype.Reference) *MedicationAdministrationBuilder {
	return b.set("request", request.ToArray())
}

// SetDosage sets dosage. All three arguments are required by the PHP signature.
func (b *MedicationAdministrationBuilder) SetDosage(dose *datatype.Quantity, route, site *datatype.CodeableConcept) *MedicationAdministrationBuilder {
	return b.set("dosage", map[string]interface{}{
		"dose":  dose.ToArray(),
		"route": route.ToArray(),
		"site":  site.ToArray(),
	})
}

// AddContained appends a contained resource.
func (b *MedicationAdministrationBuilder) AddContained(resource map[string]interface{}) *MedicationAdministrationBuilder {
	return b.push("contained", resource)
}

// Build returns the payload.
func (b *MedicationAdministrationBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *MedicationAdministrationBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
