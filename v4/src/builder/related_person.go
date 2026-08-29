package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type RelatedPersonBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewRelatedPersonBuilder() *RelatedPersonBuilder {
    b := &RelatedPersonBuilder{ResourceType: "RelatedPerson", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "RelatedPerson"
    return b
}

func (b *RelatedPersonBuilder) setId(id string) *RelatedPersonBuilder {
    b.Data["id"] = id
    return b
}

func (b *RelatedPersonBuilder) addIdentifier(identifier *datatype.Identifier) *RelatedPersonBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *RelatedPersonBuilder) setStatus(status string) *RelatedPersonBuilder {
    b.Data["status"] = status
    return b
}

func (b *RelatedPersonBuilder) setSubject(reference string) *RelatedPersonBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *RelatedPersonBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *RelatedPersonBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
