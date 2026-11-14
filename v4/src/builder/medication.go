package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// MedicationBuilder builds a FHIR R4 Medication resource.
//
// Method-for-method counterpart of PayloadBuilderMedication in the PHP SDK, including the
// Kemkes MedicationType extension helper with the same code/display lookup table.
//
// @link https://www.hl7.org/fhir/medication.html
type MedicationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewMedicationBuilder creates an empty Medication with its resourceType set.
func NewMedicationBuilder() *MedicationBuilder {
	b := &MedicationBuilder{ResourceType: "Medication", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	// The PHP constructor sets meta.profile for this resource; match it so payloads are identical.
	b.Data["meta"] = map[string]interface{}{
		"profile": []interface{}{"https://fhir.kemkes.go.id/r4/StructureDefinition/Medication"},
	}
	return b
}

func (b *MedicationBuilder) set(k string, v interface{}) *MedicationBuilder {
	b.Data[k] = v
	return b
}

func (b *MedicationBuilder) push(k string, v interface{}) *MedicationBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *MedicationBuilder) SetId(id string) *MedicationBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *MedicationBuilder) AddIdentifier(identifier *datatype.Identifier) *MedicationBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetCode sets code.
func (b *MedicationBuilder) SetCode(code *datatype.CodeableConcept) *MedicationBuilder {
	return b.set("code", code.ToArray())
}

// SetStatus sets status (active | inactive | entered-in-error).
func (b *MedicationBuilder) SetStatus(status string) *MedicationBuilder {
	return b.set("status", status)
}

// SetManufacturer sets manufacturer.
func (b *MedicationBuilder) SetManufacturer(manufacturer *datatype.Reference) *MedicationBuilder {
	return b.set("manufacturer", manufacturer.ToArray())
}

// SetForm sets form.
func (b *MedicationBuilder) SetForm(form *datatype.CodeableConcept) *MedicationBuilder {
	return b.set("form", form.ToArray())
}

// AddIngredient appends an ingredient. strength is optional.
func (b *MedicationBuilder) AddIngredient(item *datatype.CodeableConcept, isActive bool, strength *datatype.Quantity) *MedicationBuilder {
	ingredient := map[string]interface{}{
		"itemCodeableConcept": item.ToArray(),
		"isActive":            isActive,
	}
	if strength != nil {
		ingredient["strength"] = strength.ToArray()
	}
	return b.push("ingredient", ingredient)
}

// SetBatch sets batch (lotNumber and expirationDate).
func (b *MedicationBuilder) SetBatch(lotNumber, expirationDate string) *MedicationBuilder {
	return b.set("batch", map[string]interface{}{
		"lotNumber":      lotNumber,
		"expirationDate": expirationDate,
	})
}

// AddMedicationType appends the Kemkes MedicationType extension. The display is looked up from the
// same table the PHP builder uses; an unknown key is written as-is.
func (b *MedicationBuilder) AddMedicationType(code, display string) *MedicationBuilder {
	opsi := map[string]string{
		"NC": "Non-compound",
		"SD": "Gives of such doses",
		"EP": "Divide into equal parts",
	}
	if d, ok := opsi[display]; ok {
		display = d
	}
	return b.push("extension", map[string]interface{}{
		"url": "https://fhir.kemkes.go.id/r4/StructureDefinition/MedicationType",
		"valueCodeableConcept": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{
				"system":  "http://terminology.kemkes.go.id/CodeSystem/medication-type",
				"code":    code,
				"display": display,
			}},
		},
	})
}

// Build returns the payload.
func (b *MedicationBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *MedicationBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
