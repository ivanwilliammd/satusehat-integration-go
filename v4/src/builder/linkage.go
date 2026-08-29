package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type LinkageBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewLinkageBuilder() *LinkageBuilder {
    b := &LinkageBuilder{ResourceType: "Linkage", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Linkage"
    return b
}

func (b *LinkageBuilder) setId(id string) *LinkageBuilder {
    b.Data["id"] = id
    return b
}

func (b *LinkageBuilder) addIdentifier(identifier *datatype.Identifier) *LinkageBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *LinkageBuilder) setStatus(status string) *LinkageBuilder {
    b.Data["status"] = status
    return b
}

func (b *LinkageBuilder) setSubject(reference string) *LinkageBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *LinkageBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *LinkageBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
