package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type MedicationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicationBuilder() *MedicationBuilder {
    b := &MedicationBuilder{ResourceType: "Medication", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Medication"
    return b
}

func (b *MedicationBuilder) setId(id string) *MedicationBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicationBuilder) addIdentifier(identifier *datatype.Identifier) *MedicationBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *MedicationBuilder) setStatus(status string) *MedicationBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicationBuilder) setSubject(reference string) *MedicationBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *MedicationBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *MedicationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
