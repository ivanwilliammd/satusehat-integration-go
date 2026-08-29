package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type MessageHeaderBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMessageHeaderBuilder() *MessageHeaderBuilder {
    b := &MessageHeaderBuilder{ResourceType: "MessageHeader", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MessageHeader"
    return b
}

func (b *MessageHeaderBuilder) setId(id string) *MessageHeaderBuilder {
    b.Data["id"] = id
    return b
}

func (b *MessageHeaderBuilder) addIdentifier(identifier *datatype.Identifier) *MessageHeaderBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *MessageHeaderBuilder) setStatus(status string) *MessageHeaderBuilder {
    b.Data["status"] = status
    return b
}

func (b *MessageHeaderBuilder) setSubject(reference string) *MessageHeaderBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *MessageHeaderBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *MessageHeaderBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
