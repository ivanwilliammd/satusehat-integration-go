package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type ChargeItemResponseBuilder struct {
	Data map[string]interface{}
}

func NewChargeItemResponseBuilder() *ChargeItemResponseBuilder {
	return &ChargeItemResponseBuilder{Data: make(map[string]interface{})}
}

func (b *ChargeItemResponseBuilder) SetId(id string) *ChargeItemResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *ChargeItemResponseBuilder) AddIdentifier(system, value string) *ChargeItemResponseBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), map[string]string{"system": system, "value": value})
	return b
}

func (b *ChargeItemResponseBuilder) SetStatus(status string) *ChargeItemResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *ChargeItemResponseBuilder) SetChargeItem(ref string) *ChargeItemResponseBuilder {
	b.Data["chargeItem"] = map[string]string{"reference": ref}
	return b
}

func (b *ChargeItemResponseBuilder) SetRequest(ref string) *ChargeItemResponseBuilder {
	b.Data["request"] = map[string]string{"reference": ref}
	return b
}

func (b *ChargeItemResponseBuilder) SetOutcome(cc *datatype.CodeableConcept) *ChargeItemResponseBuilder {
	b.Data["outcome"] = cc.ToArray()
	return b
}

func (b *ChargeItemResponseBuilder) SetDescription(desc string) *ChargeItemResponseBuilder {
	b.Data["description"] = desc
	return b
}

func (b *ChargeItemResponseBuilder) SetCreated(dt string) *ChargeItemResponseBuilder {
	b.Data["created"] = dt
	return b
}

func (b *ChargeItemResponseBuilder) SetRequestor(ref string) *ChargeItemResponseBuilder {
	b.Data["requestor"] = map[string]string{"reference": ref}
	return b
}

func (b *ChargeItemResponseBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "ChargeItemResponse"
	return dt
}
