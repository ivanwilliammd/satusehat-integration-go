package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type RequestGroupBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewRequestGroupBuilder() *RequestGroupBuilder {
    b := &RequestGroupBuilder{ResourceType: "RequestGroup", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "RequestGroup"
    return b
}

func (b *RequestGroupBuilder) setId(id string) *RequestGroupBuilder {
    b.Data["id"] = id
    return b
}

func (b *RequestGroupBuilder) addIdentifier(identifier *datatype.Identifier) *RequestGroupBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *RequestGroupBuilder) setStatus(status string) *RequestGroupBuilder {
    b.Data["status"] = status
    return b
}

func (b *RequestGroupBuilder) setSubject(reference string) *RequestGroupBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *RequestGroupBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *RequestGroupBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
