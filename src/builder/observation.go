package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ObservationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewObservationBuilder() *ObservationBuilder {
    b := &ObservationBuilder{ResourceType: "Observation", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Observation"
    return b
}

func (b *ObservationBuilder) setId(id string) *ObservationBuilder {
    b.Data["id"] = id
    return b
}

func (b *ObservationBuilder) addIdentifier(identifier *datatype.Identifier) *ObservationBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ObservationBuilder) setStatus(status string) *ObservationBuilder {
    b.Data["status"] = status
    return b
}

func (b *ObservationBuilder) setSubject(reference string) *ObservationBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ObservationBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ObservationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
