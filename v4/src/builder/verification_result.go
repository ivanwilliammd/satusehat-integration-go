package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type VerificationResultBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewVerificationResultBuilder() *VerificationResultBuilder {
    b := &VerificationResultBuilder{ResourceType: "VerificationResult", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "VerificationResult"
    return b
}

func (b *VerificationResultBuilder) setId(id string) *VerificationResultBuilder {
    b.Data["id"] = id
    return b
}

func (b *VerificationResultBuilder) addIdentifier(identifier *datatype.Identifier) *VerificationResultBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *VerificationResultBuilder) setStatus(status string) *VerificationResultBuilder {
    b.Data["status"] = status
    return b
}

func (b *VerificationResultBuilder) setSubject(reference string) *VerificationResultBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *VerificationResultBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *VerificationResultBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
