package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ConditionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewConditionBuilder() *ConditionBuilder {
    b := &ConditionBuilder{ResourceType: "Condition", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Condition"
    return b
}

func (b *ConditionBuilder) setId(id string) *ConditionBuilder {
    b.Data["id"] = id
    return b
}

func (b *ConditionBuilder) addIdentifier(identifier *datatype.Identifier) *ConditionBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ConditionBuilder) setStatus(status string) *ConditionBuilder {
    b.Data["status"] = status
    return b
}

func (b *ConditionBuilder) setSubject(reference string) *ConditionBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *ConditionBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ConditionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
