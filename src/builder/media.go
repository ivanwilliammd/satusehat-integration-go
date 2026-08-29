package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MediaBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMediaBuilder() *MediaBuilder {
    b := &MediaBuilder{ResourceType: "Media", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Media"
    return b
}

func (b *MediaBuilder) setId(id string) *MediaBuilder {
    b.Data["id"] = id
    return b
}

func (b *MediaBuilder) addIdentifier(identifier *datatype.Identifier) *MediaBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *MediaBuilder) setStatus(status string) *MediaBuilder {
    b.Data["status"] = status
    return b
}

func (b *MediaBuilder) setSubject(reference string) *MediaBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *MediaBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *MediaBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
