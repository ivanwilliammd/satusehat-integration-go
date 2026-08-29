package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// MedicationBuilder builds FHIR Medication payload
type MedicationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMedicationBuilder() *MedicationBuilder {
	b := &MedicationBuilder{ResourceType: "Medication", Data: make(map[string]interface{})}
	return b
}

func (b *MedicationBuilder) setId(id string) *MedicationBuilder {
	b.Data["id"] = id
	return b
}

func (b *MedicationBuilder) addIdentifier(identifier *datatype.Identifier) *MedicationBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *MedicationBuilder) setCode(code *datatype.CodeableConcept) *MedicationBuilder {
	b.Data["code"] = code.ToArray()
	return b
}

func (b *MedicationBuilder) setStatus(status string) *MedicationBuilder {
	b.Data["status"] = status
	return b
}

func (b *MedicationBuilder) setManufacturer(manufacturer *datatype.Reference) *MedicationBuilder {
	b.Data["manufacturer"] = manufacturer.ToArray()
	return b
}

func (b *MedicationBuilder) addExtension(url string, value interface{}, valueType string) *MedicationBuilder {
	ext := map[string]interface{}{"url": url}
	if valueType != "" {
		capitalized := strings.ToUpper(valueType[:1]) + valueType[1:]
		ext["value"+capitalized] = value
	} else {
		ext["valueString"] = value
	}
	if _, ok := b.Data["extension"]; !ok {
		b.Data["extension"] = make([]interface{}, 0)
	}
	b.Data["extension"] = append(b.Data["extension"].([]interface{}), ext)
	return b
}

func (b *MedicationBuilder) Build() map[string]interface{} {
	return b.Data
}
