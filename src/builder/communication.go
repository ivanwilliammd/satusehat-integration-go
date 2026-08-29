package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type CommunicationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewCommunicationBuilder() *CommunicationBuilder {
    b := &CommunicationBuilder{ResourceType: "Communication", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Communication"
    return b
}

func (b *CommunicationBuilder) setId(id string) *CommunicationBuilder {
    b.Data["id"] = id
    return b
}

func (b *CommunicationBuilder) addIdentifier(identifier *datatype.Identifier) *CommunicationBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *CommunicationBuilder) setStatus(status string) *CommunicationBuilder {
    b.Data["status"] = status
    return b
}

func (b *CommunicationBuilder) setSubject(reference string) *CommunicationBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *CommunicationBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *CommunicationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
