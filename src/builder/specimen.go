package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type SpecimenBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewSpecimenBuilder() *SpecimenBuilder {
    b := &SpecimenBuilder{ResourceType: "Specimen", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Specimen"
    return b
}

func (b *SpecimenBuilder) setId(id string) *SpecimenBuilder {
    b.Data["id"] = id
    return b
}

func (b *SpecimenBuilder) addIdentifier(identifier *datatype.Identifier) *SpecimenBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *SpecimenBuilder) setStatus(status string) *SpecimenBuilder {
    b.Data["status"] = status
    return b
}

func (b *SpecimenBuilder) setSubject(reference string) *SpecimenBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *SpecimenBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *SpecimenBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
