package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type GroupBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewGroupBuilder() *GroupBuilder {
    b := &GroupBuilder{ResourceType: "Group", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Group"
    return b
}

func (b *GroupBuilder) setId(id string) *GroupBuilder {
    b.Data["id"] = id
    return b
}

func (b *GroupBuilder) addIdentifier(identifier *datatype.Identifier) *GroupBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *GroupBuilder) setStatus(status string) *GroupBuilder {
    b.Data["status"] = status
    return b
}

func (b *GroupBuilder) setSubject(reference string) *GroupBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *GroupBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *GroupBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
