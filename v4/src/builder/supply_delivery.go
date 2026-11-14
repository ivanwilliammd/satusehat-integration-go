package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SupplyDeliveryBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSupplyDeliveryBuilder() *SupplyDeliveryBuilder {
	b := &SupplyDeliveryBuilder{ResourceType: "SupplyDelivery", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "SupplyDelivery"
	return b
}

func (b *SupplyDeliveryBuilder) SetId(id string) *SupplyDeliveryBuilder {
	b.Data["id"] = id
	return b
}

func (b *SupplyDeliveryBuilder) AddIdentifier(identifier *datatype.Identifier) *SupplyDeliveryBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *SupplyDeliveryBuilder) SetStatus(status string) *SupplyDeliveryBuilder {
	b.Data["status"] = status
	return b
}

func (b *SupplyDeliveryBuilder) SetSubject(reference string) *SupplyDeliveryBuilder {
	b.Data["subject"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

func (b *SupplyDeliveryBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SupplyDeliveryBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
