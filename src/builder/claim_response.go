package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ClaimResponseBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewClaimResponseBuilder() *ClaimResponseBuilder {
    b := &ClaimResponseBuilder{ResourceType: "ClaimResponse", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ClaimResponse"
    return b
}

func (b *ClaimResponseBuilder) setId(id string) *ClaimResponseBuilder {
    b.Data["id"] = id
    return b
}

func (b *ClaimResponseBuilder) addIdentifier(identifier *datatype.Identifier) *ClaimResponseBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ClaimResponseBuilder) setStatus(status string) *ClaimResponseBuilder {
    b.Data["status"] = status
    return b
}

func (b *ClaimResponseBuilder) setSubject(reference string) *ClaimResponseBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ClaimResponseBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ClaimResponseBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
