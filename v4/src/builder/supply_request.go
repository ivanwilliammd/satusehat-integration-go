package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SupplyRequestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSupplyRequestBuilder() *SupplyRequestBuilder {
	b := &SupplyRequestBuilder{ResourceType: "SupplyRequest", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "SupplyRequest"
	return b
}

func (b *SupplyRequestBuilder) SetId(id string) *SupplyRequestBuilder {
	b.Data["id"] = id
	return b
}

func (b *SupplyRequestBuilder) AddIdentifier(identifier *datatype.Identifier) *SupplyRequestBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *SupplyRequestBuilder) SetStatus(status string) *SupplyRequestBuilder {
	b.Data["status"] = status
	return b
}

func (b *SupplyRequestBuilder) SetSubject(reference string, display ...string) *SupplyRequestBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *SupplyRequestBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SupplyRequestBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *SupplyRequestBuilder) SetCode(system, code, display string) *SupplyRequestBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *SupplyRequestBuilder) SetEncounter(reference string) *SupplyRequestBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
