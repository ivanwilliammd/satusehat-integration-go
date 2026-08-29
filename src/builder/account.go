package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type AccountBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewAccountBuilder() *AccountBuilder {
    b := &AccountBuilder{ResourceType: "Account", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Account"
    return b
}

func (b *AccountBuilder) setId(id string) *AccountBuilder {
    b.Data["id"] = id
    return b
}

func (b *AccountBuilder) addIdentifier(identifier *datatype.Identifier) *AccountBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *AccountBuilder) setStatus(status string) *AccountBuilder {
    b.Data["status"] = status
    return b
}

func (b *AccountBuilder) setSubject(reference string) *AccountBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *AccountBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *AccountBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
