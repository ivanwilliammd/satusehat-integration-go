package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ChargeItemBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewChargeItemBuilder() *ChargeItemBuilder {
	b := &ChargeItemBuilder{ResourceType: "ChargeItem", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ChargeItem"
	return b
}

func (b *ChargeItemBuilder) SetId(id string) *ChargeItemBuilder {
	b.Data["id"] = id
	return b
}

func (b *ChargeItemBuilder) AddIdentifier(identifier *datatype.Identifier) *ChargeItemBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ChargeItemBuilder) SetStatus(status string) *ChargeItemBuilder {
	b.Data["status"] = status
	return b
}

func (b *ChargeItemBuilder) SetSubject(reference string) *ChargeItemBuilder {
	b.Data["subject"] = datatype.Reference{}.ToArray()
	return b
}

func (b *ChargeItemBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ChargeItemBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
