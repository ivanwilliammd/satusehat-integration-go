package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type BasicBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewBasicBuilder() *BasicBuilder {
    b := &BasicBuilder{ResourceType: "Basic", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Basic"
    return b
}

func (b *BasicBuilder) setId(id string) *BasicBuilder {
    b.Data["id"] = id
    return b
}

func (b *BasicBuilder) addIdentifier(identifier *datatype.Identifier) *BasicBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *BasicBuilder) setStatus(status string) *BasicBuilder {
    b.Data["status"] = status
    return b
}

func (b *BasicBuilder) setSubject(reference string) *BasicBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *BasicBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *BasicBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
