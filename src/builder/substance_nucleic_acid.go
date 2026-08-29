package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type SubstanceNucleicAcidBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewSubstanceNucleicAcidBuilder() *SubstanceNucleicAcidBuilder {
    b := &SubstanceNucleicAcidBuilder{ResourceType: "SubstanceNucleicAcid", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "SubstanceNucleicAcid"
    return b
}

func (b *SubstanceNucleicAcidBuilder) setId(id string) *SubstanceNucleicAcidBuilder {
    b.Data["id"] = id
    return b
}

func (b *SubstanceNucleicAcidBuilder) addIdentifier(identifier *datatype.Identifier) *SubstanceNucleicAcidBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *SubstanceNucleicAcidBuilder) setStatus(status string) *SubstanceNucleicAcidBuilder {
    b.Data["status"] = status
    return b
}

func (b *SubstanceNucleicAcidBuilder) setSubject(reference string) *SubstanceNucleicAcidBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *SubstanceNucleicAcidBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *SubstanceNucleicAcidBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
