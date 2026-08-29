package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ServiceRequestBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewServiceRequestBuilder() *ServiceRequestBuilder {
    b := &ServiceRequestBuilder{ResourceType: "ServiceRequest", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ServiceRequest"
    return b
}

func (b *ServiceRequestBuilder) setId(id string) *ServiceRequestBuilder {
    b.Data["id"] = id
    return b
}

func (b *ServiceRequestBuilder) addIdentifier(identifier *datatype.Identifier) *ServiceRequestBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ServiceRequestBuilder) setStatus(status string) *ServiceRequestBuilder {
    b.Data["status"] = status
    return b
}

func (b *ServiceRequestBuilder) setSubject(reference string) *ServiceRequestBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ServiceRequestBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ServiceRequestBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
