package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ChargeItemResponseBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewChargeItemResponseBuilder() *ChargeItemResponseBuilder {
    b := &ChargeItemResponseBuilder{ResourceType: "ChargeItemResponse", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ChargeItemResponse"
    return b
}

func (b *ChargeItemResponseBuilder) setId(id string) *ChargeItemResponseBuilder {
    b.Data["id"] = id
    return b
}

func (b *ChargeItemResponseBuilder) addIdentifier(identifier *datatype.Identifier) *ChargeItemResponseBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ChargeItemResponseBuilder) setStatus(status string) *ChargeItemResponseBuilder {
    b.Data["status"] = status
    return b
}

func (b *ChargeItemResponseBuilder) setSubject(reference string) *ChargeItemResponseBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *ChargeItemResponseBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ChargeItemResponseBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
