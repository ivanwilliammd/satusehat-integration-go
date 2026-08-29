package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type AdverseEventBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewAdverseEventBuilder() *AdverseEventBuilder {
    b := &AdverseEventBuilder{ResourceType: "AdverseEvent", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "AdverseEvent"
    return b
}

func (b *AdverseEventBuilder) setId(id string) *AdverseEventBuilder {
    b.Data["id"] = id
    return b
}

func (b *AdverseEventBuilder) addIdentifier(identifier *datatype.Identifier) *AdverseEventBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *AdverseEventBuilder) setStatus(status string) *AdverseEventBuilder {
    b.Data["status"] = status
    return b
}

func (b *AdverseEventBuilder) setSubject(reference string) *AdverseEventBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *AdverseEventBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *AdverseEventBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
