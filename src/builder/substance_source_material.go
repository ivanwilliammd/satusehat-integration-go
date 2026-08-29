package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SubstanceSourceMaterialBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewSubstanceSourceMaterialBuilder() *SubstanceSourceMaterialBuilder {
    b := &SubstanceSourceMaterialBuilder{ResourceType: "SubstanceSourceMaterial", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "SubstanceSourceMaterial"
    return b
}

func (b *SubstanceSourceMaterialBuilder) setId(id string) *SubstanceSourceMaterialBuilder {
    b.Data["id"] = id
    return b
}

func (b *SubstanceSourceMaterialBuilder) addIdentifier(identifier *datatype.Identifier) *SubstanceSourceMaterialBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *SubstanceSourceMaterialBuilder) setStatus(status string) *SubstanceSourceMaterialBuilder {
    b.Data["status"] = status
    return b
}

func (b *SubstanceSourceMaterialBuilder) setSubject(reference string) *SubstanceSourceMaterialBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *SubstanceSourceMaterialBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *SubstanceSourceMaterialBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
