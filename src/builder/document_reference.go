package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type DocumentReferenceBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewDocumentReferenceBuilder() *DocumentReferenceBuilder {
    b := &DocumentReferenceBuilder{ResourceType: "DocumentReference", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "DocumentReference"
    return b
}

func (b *DocumentReferenceBuilder) setId(id string) *DocumentReferenceBuilder {
    b.Data["id"] = id
    return b
}

func (b *DocumentReferenceBuilder) addIdentifier(identifier *datatype.Identifier) *DocumentReferenceBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *DocumentReferenceBuilder) setStatus(status string) *DocumentReferenceBuilder {
    b.Data["status"] = status
    return b
}

func (b *DocumentReferenceBuilder) setSubject(reference string) *DocumentReferenceBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *DocumentReferenceBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *DocumentReferenceBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
