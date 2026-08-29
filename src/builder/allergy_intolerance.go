package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type AllergyIntoleranceBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewAllergyIntoleranceBuilder() *AllergyIntoleranceBuilder {
    b := &AllergyIntoleranceBuilder{ResourceType: "AllergyIntolerance", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "AllergyIntolerance"
    return b
}

func (b *AllergyIntoleranceBuilder) setId(id string) *AllergyIntoleranceBuilder {
    b.Data["id"] = id
    return b
}

func (b *AllergyIntoleranceBuilder) addIdentifier(identifier *datatype.Identifier) *AllergyIntoleranceBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *AllergyIntoleranceBuilder) setStatus(status string) *AllergyIntoleranceBuilder {
    b.Data["status"] = status
    return b
}

func (b *AllergyIntoleranceBuilder) setSubject(reference string) *AllergyIntoleranceBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *AllergyIntoleranceBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *AllergyIntoleranceBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
