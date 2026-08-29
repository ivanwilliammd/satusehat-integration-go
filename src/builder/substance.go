package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type SubstanceBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewSubstanceBuilder() *SubstanceBuilder {
    b := &SubstanceBuilder{ResourceType: "Substance", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Substance"
    return b
}

func (b *SubstanceBuilder) setId(id string) *SubstanceBuilder {
    b.Data["id"] = id
    return b
}

func (b *SubstanceBuilder) addIdentifier(identifier *datatype.Identifier) *SubstanceBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *SubstanceBuilder) setStatus(status string) *SubstanceBuilder {
    b.Data["status"] = status
    return b
}

func (b *SubstanceBuilder) setSubject(reference string) *SubstanceBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *SubstanceBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *SubstanceBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
