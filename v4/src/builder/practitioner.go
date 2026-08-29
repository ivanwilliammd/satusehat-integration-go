package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type PractitionerBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewPractitionerBuilder() *PractitionerBuilder {
    b := &PractitionerBuilder{ResourceType: "Practitioner", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Practitioner"
    return b
}

func (b *PractitionerBuilder) setId(id string) *PractitionerBuilder {
    b.Data["id"] = id
    return b
}

func (b *PractitionerBuilder) addIdentifier(identifier *datatype.Identifier) *PractitionerBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *PractitionerBuilder) setStatus(status string) *PractitionerBuilder {
    b.Data["status"] = status
    return b
}

func (b *PractitionerBuilder) setSubject(reference string) *PractitionerBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *PractitionerBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *PractitionerBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
