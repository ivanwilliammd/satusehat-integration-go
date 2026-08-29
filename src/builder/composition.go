package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type CompositionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewCompositionBuilder() *CompositionBuilder {
    b := &CompositionBuilder{ResourceType: "Composition", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Composition"
    return b
}

func (b *CompositionBuilder) setId(id string) *CompositionBuilder {
    b.Data["id"] = id
    return b
}

func (b *CompositionBuilder) addIdentifier(identifier *datatype.Identifier) *CompositionBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *CompositionBuilder) setStatus(status string) *CompositionBuilder {
    b.Data["status"] = status
    return b
}

func (b *CompositionBuilder) setSubject(reference string) *CompositionBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *CompositionBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *CompositionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
