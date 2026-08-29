package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ResearchDefinitionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewResearchDefinitionBuilder() *ResearchDefinitionBuilder {
    b := &ResearchDefinitionBuilder{ResourceType: "ResearchDefinition", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ResearchDefinition"
    return b
}

func (b *ResearchDefinitionBuilder) setId(id string) *ResearchDefinitionBuilder {
    b.Data["id"] = id
    return b
}

func (b *ResearchDefinitionBuilder) addIdentifier(identifier *datatype.Identifier) *ResearchDefinitionBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ResearchDefinitionBuilder) setStatus(status string) *ResearchDefinitionBuilder {
    b.Data["status"] = status
    return b
}

func (b *ResearchDefinitionBuilder) setSubject(reference string) *ResearchDefinitionBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ResearchDefinitionBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ResearchDefinitionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
