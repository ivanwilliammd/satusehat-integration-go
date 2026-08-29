package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ImmunizationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewImmunizationBuilder() *ImmunizationBuilder {
    b := &ImmunizationBuilder{ResourceType: "Immunization", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Immunization"
    return b
}

func (b *ImmunizationBuilder) setId(id string) *ImmunizationBuilder {
    b.Data["id"] = id
    return b
}

func (b *ImmunizationBuilder) addIdentifier(identifier *datatype.Identifier) *ImmunizationBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ImmunizationBuilder) setStatus(status string) *ImmunizationBuilder {
    b.Data["status"] = status
    return b
}

func (b *ImmunizationBuilder) setSubject(reference string) *ImmunizationBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ImmunizationBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ImmunizationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
