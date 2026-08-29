package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ListBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewListBuilder() *ListBuilder {
    b := &ListBuilder{ResourceType: "List", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "List"
    return b
}

func (b *ListBuilder) setId(id string) *ListBuilder {
    b.Data["id"] = id
    return b
}

func (b *ListBuilder) addIdentifier(identifier *datatype.Identifier) *ListBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ListBuilder) setStatus(status string) *ListBuilder {
    b.Data["status"] = status
    return b
}

func (b *ListBuilder) setSubject(reference string) *ListBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ListBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ListBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
