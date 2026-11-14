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

func (b *SupplyDeliveryBuilder) SetSubject(reference string, display ...string) *SupplyDeliveryBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *SupplyDeliveryBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SupplyDeliveryBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *SupplyDeliveryBuilder) SetCode(system, code, display string) *SupplyDeliveryBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *SupplyDeliveryBuilder) SetEncounter(reference string) *SupplyDeliveryBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
