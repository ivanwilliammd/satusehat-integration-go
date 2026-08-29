package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type BodyStructureBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewBodyStructureBuilder() *BodyStructureBuilder {
    b := &BodyStructureBuilder{ResourceType: "BodyStructure", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "BodyStructure"
    return b
}

func (b *BodyStructureBuilder) setId(id string) *BodyStructureBuilder {
    b.Data["id"] = id
    return b
}

func (b *BodyStructureBuilder) addIdentifier(identifier *datatype.Identifier) *BodyStructureBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *BodyStructureBuilder) setStatus(status string) *BodyStructureBuilder {
    b.Data["status"] = status
    return b
}

func (b *BodyStructureBuilder) setSubject(reference string) *BodyStructureBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *BodyStructureBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *BodyStructureBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
