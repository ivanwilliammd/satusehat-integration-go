package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type LibraryBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewLibraryBuilder() *LibraryBuilder {
    b := &LibraryBuilder{ResourceType: "Library", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Library"
    return b
}

func (b *LibraryBuilder) setId(id string) *LibraryBuilder {
    b.Data["id"] = id
    return b
}

func (b *LibraryBuilder) addIdentifier(identifier *datatype.Identifier) *LibraryBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *LibraryBuilder) setStatus(status string) *LibraryBuilder {
    b.Data["status"] = status
    return b
}

func (b *LibraryBuilder) setSubject(reference string) *LibraryBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *LibraryBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *LibraryBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
