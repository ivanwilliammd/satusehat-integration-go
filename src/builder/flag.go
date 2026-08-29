package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type FlagBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewFlagBuilder() *FlagBuilder {
    b := &FlagBuilder{ResourceType: "Flag", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Flag"
    return b
}

func (b *FlagBuilder) setId(id string) *FlagBuilder {
    b.Data["id"] = id
    return b
}

func (b *FlagBuilder) addIdentifier(identifier *datatype.Identifier) *FlagBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *FlagBuilder) setStatus(status string) *FlagBuilder {
    b.Data["status"] = status
    return b
}

func (b *FlagBuilder) setSubject(reference string) *FlagBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *FlagBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *FlagBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
