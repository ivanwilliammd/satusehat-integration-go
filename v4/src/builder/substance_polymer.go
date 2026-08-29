package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SubstancePolymerBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewSubstancePolymerBuilder() *SubstancePolymerBuilder {
    b := &SubstancePolymerBuilder{ResourceType: "SubstancePolymer", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "SubstancePolymer"
    return b
}

func (b *SubstancePolymerBuilder) setId(id string) *SubstancePolymerBuilder {
    b.Data["id"] = id
    return b
}

func (b *SubstancePolymerBuilder) addIdentifier(identifier *datatype.Identifier) *SubstancePolymerBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *SubstancePolymerBuilder) setStatus(status string) *SubstancePolymerBuilder {
    b.Data["status"] = status
    return b
}

func (b *SubstancePolymerBuilder) setSubject(reference string) *SubstancePolymerBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *SubstancePolymerBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *SubstancePolymerBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
