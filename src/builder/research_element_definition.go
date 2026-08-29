package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ResearchElementDefinitionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewResearchElementDefinitionBuilder() *ResearchElementDefinitionBuilder {
    b := &ResearchElementDefinitionBuilder{ResourceType: "ResearchElementDefinition", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ResearchElementDefinition"
    return b
}

func (b *ResearchElementDefinitionBuilder) setId(id string) *ResearchElementDefinitionBuilder {
    b.Data["id"] = id
    return b
}

func (b *ResearchElementDefinitionBuilder) addIdentifier(identifier *datatype.Identifier) *ResearchElementDefinitionBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ResearchElementDefinitionBuilder) setStatus(status string) *ResearchElementDefinitionBuilder {
    b.Data["status"] = status
    return b
}

func (b *ResearchElementDefinitionBuilder) setSubject(reference string) *ResearchElementDefinitionBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ResearchElementDefinitionBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ResearchElementDefinitionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
