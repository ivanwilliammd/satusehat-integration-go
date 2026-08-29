package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ConsentBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewConsentBuilder() *ConsentBuilder {
    b := &ConsentBuilder{ResourceType: "Consent", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Consent"
    return b
}

func (b *ConsentBuilder) setId(id string) *ConsentBuilder {
    b.Data["id"] = id
    return b
}

func (b *ConsentBuilder) addIdentifier(identifier *datatype.Identifier) *ConsentBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ConsentBuilder) setStatus(status string) *ConsentBuilder {
    b.Data["status"] = status
    return b
}

func (b *ConsentBuilder) setSubject(reference string) *ConsentBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ConsentBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ConsentBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
