package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type OperationOutcomeBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewOperationOutcomeBuilder() *OperationOutcomeBuilder {
    b := &OperationOutcomeBuilder{ResourceType: "OperationOutcome", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "OperationOutcome"
    return b
}

func (b *OperationOutcomeBuilder) setId(id string) *OperationOutcomeBuilder {
    b.Data["id"] = id
    return b
}

func (b *OperationOutcomeBuilder) addIdentifier(identifier *datatype.Identifier) *OperationOutcomeBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *OperationOutcomeBuilder) setStatus(status string) *OperationOutcomeBuilder {
    b.Data["status"] = status
    return b
}

func (b *OperationOutcomeBuilder) setSubject(reference string) *OperationOutcomeBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *OperationOutcomeBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *OperationOutcomeBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
