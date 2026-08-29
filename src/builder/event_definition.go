package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type EventDefinitionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewEventDefinitionBuilder() *EventDefinitionBuilder {
    b := &EventDefinitionBuilder{ResourceType: "EventDefinition", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "EventDefinition"
    return b
}

func (b *EventDefinitionBuilder) setId(id string) *EventDefinitionBuilder {
    b.Data["id"] = id
    return b
}

func (b *EventDefinitionBuilder) addIdentifier(identifier *datatype.Identifier) *EventDefinitionBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *EventDefinitionBuilder) setStatus(status string) *EventDefinitionBuilder {
    b.Data["status"] = status
    return b
}

func (b *EventDefinitionBuilder) setSubject(reference string) *EventDefinitionBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *EventDefinitionBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *EventDefinitionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
