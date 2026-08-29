package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type PlanDefinitionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewPlanDefinitionBuilder() *PlanDefinitionBuilder {
    b := &PlanDefinitionBuilder{ResourceType: "PlanDefinition", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "PlanDefinition"
    return b
}

func (b *PlanDefinitionBuilder) setId(id string) *PlanDefinitionBuilder {
    b.Data["id"] = id
    return b
}

func (b *PlanDefinitionBuilder) addIdentifier(identifier *datatype.Identifier) *PlanDefinitionBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *PlanDefinitionBuilder) setStatus(status string) *PlanDefinitionBuilder {
    b.Data["status"] = status
    return b
}

func (b *PlanDefinitionBuilder) setSubject(reference string) *PlanDefinitionBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *PlanDefinitionBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *PlanDefinitionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
