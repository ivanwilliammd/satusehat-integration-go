package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type EvidenceBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewEvidenceBuilder() *EvidenceBuilder {
    b := &EvidenceBuilder{ResourceType: "Evidence", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Evidence"
    return b
}

func (b *EvidenceBuilder) setId(id string) *EvidenceBuilder {
    b.Data["id"] = id
    return b
}

func (b *EvidenceBuilder) addIdentifier(identifier *datatype.Identifier) *EvidenceBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *EvidenceBuilder) setStatus(status string) *EvidenceBuilder {
    b.Data["status"] = status
    return b
}

func (b *EvidenceBuilder) setSubject(reference string) *EvidenceBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *EvidenceBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *EvidenceBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
