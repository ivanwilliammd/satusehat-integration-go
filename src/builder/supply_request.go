package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type SupplyRequestBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewSupplyRequestBuilder() *SupplyRequestBuilder {
    b := &SupplyRequestBuilder{ResourceType: "SupplyRequest", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "SupplyRequest"
    return b
}

func (b *SupplyRequestBuilder) setId(id string) *SupplyRequestBuilder {
    b.Data["id"] = id
    return b
}

func (b *SupplyRequestBuilder) addIdentifier(identifier *datatype.Identifier) *SupplyRequestBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *SupplyRequestBuilder) setStatus(status string) *SupplyRequestBuilder {
    b.Data["status"] = status
    return b
}

func (b *SupplyRequestBuilder) setSubject(reference string) *SupplyRequestBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *SupplyRequestBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *SupplyRequestBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
