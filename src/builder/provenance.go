package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ProvenanceBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewProvenanceBuilder() *ProvenanceBuilder {
    b := &ProvenanceBuilder{ResourceType: "Provenance", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Provenance"
    return b
}

func (b *ProvenanceBuilder) setId(id string) *ProvenanceBuilder {
    b.Data["id"] = id
    return b
}

func (b *ProvenanceBuilder) addIdentifier(identifier *datatype.Identifier) *ProvenanceBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ProvenanceBuilder) setStatus(status string) *ProvenanceBuilder {
    b.Data["status"] = status
    return b
}

func (b *ProvenanceBuilder) setSubject(reference string) *ProvenanceBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *ProvenanceBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ProvenanceBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
