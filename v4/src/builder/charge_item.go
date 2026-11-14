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

func (b *ChargeItemBuilder) SetSubject(reference string, display ...string) *ChargeItemBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ChargeItemBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ChargeItemBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ChargeItemBuilder) SetCode(system, code, display string) *ChargeItemBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ChargeItemBuilder) SetEncounter(reference string) *ChargeItemBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}

// SetEnterer sets enterer by reference, as the PHP builder does.
func (b *ChargeItemBuilder) SetEnterer(reference string) *ChargeItemBuilder {
	b.Data["enterer"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}
