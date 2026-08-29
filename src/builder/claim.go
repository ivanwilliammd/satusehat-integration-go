package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ClaimBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewClaimBuilder() *ClaimBuilder {
    b := &ClaimBuilder{ResourceType: "Claim", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Claim"
    return b
}

func (b *ClaimBuilder) setId(id string) *ClaimBuilder {
    b.Data["id"] = id
    return b
}

func (b *ClaimBuilder) addIdentifier(identifier *datatype.Identifier) *ClaimBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ClaimBuilder) setStatus(status string) *ClaimBuilder {
    b.Data["status"] = status
    return b
}

func (b *ClaimBuilder) setSubject(reference string) *ClaimBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ClaimBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ClaimBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
