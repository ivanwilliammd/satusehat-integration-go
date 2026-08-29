package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MedicationStatementBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMedicationStatementBuilder() *MedicationStatementBuilder {
    b := &MedicationStatementBuilder{ResourceType: "MedicationStatement", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MedicationStatement"
    return b
}

func (b *MedicationStatementBuilder) setId(id string) *MedicationStatementBuilder {
    b.Data["id"] = id
    return b
}

func (b *MedicationStatementBuilder) addIdentifier(identifier *datatype.Identifier) *MedicationStatementBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *MedicationStatementBuilder) setStatus(status string) *MedicationStatementBuilder {
    b.Data["status"] = status
    return b
}

func (b *MedicationStatementBuilder) setSubject(reference string) *MedicationStatementBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *MedicationStatementBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *MedicationStatementBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
