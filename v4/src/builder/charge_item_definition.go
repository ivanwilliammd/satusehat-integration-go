package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ChargeItemDefinitionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewChargeItemDefinitionBuilder() *ChargeItemDefinitionBuilder {
	b := &ChargeItemDefinitionBuilder{ResourceType: "ChargeItemDefinition", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ChargeItemDefinition"
	return b
}

func (b *ChargeItemDefinitionBuilder) SetId(id string) *ChargeItemDefinitionBuilder {
	b.Data["id"] = id
	return b
}

func (b *ChargeItemDefinitionBuilder) AddIdentifier(identifier *datatype.Identifier) *ChargeItemDefinitionBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ChargeItemDefinitionBuilder) SetStatus(status string) *ChargeItemDefinitionBuilder {
	b.Data["status"] = status
	return b
}

func (b *ChargeItemDefinitionBuilder) SetSubject(reference string) *ChargeItemDefinitionBuilder {
	b.Data["subject"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

func (b *ChargeItemDefinitionBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ChargeItemDefinitionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
