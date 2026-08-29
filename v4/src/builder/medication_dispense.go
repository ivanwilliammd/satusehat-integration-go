package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MedicationDispenseBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicationDispenseBuilder() *MedicationDispenseBuilder {
    b := &MedicationDispenseBuilder{ResourceType: "MedicationDispense", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicationDispense"
    return b
}

func (b *MedicationDispenseBuilder) setId(id string) *MedicationDispenseBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicationDispenseBuilder) addIdentifier(identifier *datatype.Identifier) *MedicationDispenseBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *MedicationDispenseBuilder) setStatus(status string) *MedicationDispenseBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicationDispenseBuilder) setSubject(reference string) *MedicationDispenseBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *MedicationDispenseBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *MedicationDispenseBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
