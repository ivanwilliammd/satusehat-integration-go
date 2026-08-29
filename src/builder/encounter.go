package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type EncounterBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewEncounterBuilder() *EncounterBuilder {
    b := &EncounterBuilder{ResourceType: "Encounter", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Encounter"
    return b
}

func (b *EncounterBuilder) setId(id string) *EncounterBuilder {
    b.Data["id"] = id
    return b
}

func (b *EncounterBuilder) addIdentifier(identifier *datatype.Identifier) *EncounterBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *EncounterBuilder) setStatus(status string) *EncounterBuilder {
    b.Data["status"] = status
    return b
}

func (b *EncounterBuilder) setSubject(reference string) *EncounterBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *EncounterBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *EncounterBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
