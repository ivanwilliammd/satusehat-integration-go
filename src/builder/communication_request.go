package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type CommunicationRequestBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewCommunicationRequestBuilder() *CommunicationRequestBuilder {
    b := &CommunicationRequestBuilder{ResourceType: "CommunicationRequest", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "CommunicationRequest"
    return b
}

func (b *CommunicationRequestBuilder) setId(id string) *CommunicationRequestBuilder {
    b.Data["id"] = id
    return b
}

func (b *CommunicationRequestBuilder) addIdentifier(identifier *datatype.Identifier) *CommunicationRequestBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *CommunicationRequestBuilder) setStatus(status string) *CommunicationRequestBuilder {
    b.Data["status"] = status
    return b
}

func (b *CommunicationRequestBuilder) setSubject(reference string) *CommunicationRequestBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *CommunicationRequestBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *CommunicationRequestBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
