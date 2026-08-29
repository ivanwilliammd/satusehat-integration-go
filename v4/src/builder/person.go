package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type PersonBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewPersonBuilder() *PersonBuilder {
    b := &PersonBuilder{ResourceType: "Person", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Person"
    return b
}

func (b *PersonBuilder) setId(id string) *PersonBuilder {
    b.Data["id"] = id
    return b
}

func (b *PersonBuilder) addIdentifier(identifier *datatype.Identifier) *PersonBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *PersonBuilder) setStatus(status string) *PersonBuilder {
    b.Data["status"] = status
    return b
}

func (b *PersonBuilder) setSubject(reference string) *PersonBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *PersonBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *PersonBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
