package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MessageDefinitionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMessageDefinitionBuilder() *MessageDefinitionBuilder {
    b := &MessageDefinitionBuilder{ResourceType: "MessageDefinition", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MessageDefinition"
    return b
}

func (b *MessageDefinitionBuilder) setId(id string) *MessageDefinitionBuilder {
    b.Data["id"] = id
    return b
}

func (b *MessageDefinitionBuilder) addIdentifier(identifier *datatype.Identifier) *MessageDefinitionBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *MessageDefinitionBuilder) setStatus(status string) *MessageDefinitionBuilder {
    b.Data["status"] = status
    return b
}

func (b *MessageDefinitionBuilder) setSubject(reference string) *MessageDefinitionBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *MessageDefinitionBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *MessageDefinitionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
