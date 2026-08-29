package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type SubstanceSpecificationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewSubstanceSpecificationBuilder() *SubstanceSpecificationBuilder {
    b := &SubstanceSpecificationBuilder{ResourceType: "SubstanceSpecification", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "SubstanceSpecification"
    return b
}

func (b *SubstanceSpecificationBuilder) setId(id string) *SubstanceSpecificationBuilder {
    b.Data["id"] = id
    return b
}

func (b *SubstanceSpecificationBuilder) addIdentifier(identifier *datatype.Identifier) *SubstanceSpecificationBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *SubstanceSpecificationBuilder) setStatus(status string) *SubstanceSpecificationBuilder {
    b.Data["status"] = status
    return b
}

func (b *SubstanceSpecificationBuilder) setSubject(reference string) *SubstanceSpecificationBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *SubstanceSpecificationBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *SubstanceSpecificationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
