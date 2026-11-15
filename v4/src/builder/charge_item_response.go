package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ChargeItemResponseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewChargeItemResponseBuilder() *ChargeItemResponseBuilder {
	b := &ChargeItemResponseBuilder{ResourceType: "ChargeItemResponse", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ChargeItemResponse"
	return b
}

func (b *ChargeItemResponseBuilder) SetId(id string) *ChargeItemResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *ChargeItemResponseBuilder) AddIdentifier(identifier *datatype.Identifier) *ChargeItemResponseBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ChargeItemResponseBuilder) SetStatus(status string) *ChargeItemResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *ChargeItemResponseBuilder) SetSubject(reference string, display ...string) *ChargeItemResponseBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ChargeItemResponseBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ChargeItemResponseBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCreated sets created, as the PHP builder does.
func (b *ChargeItemResponseBuilder) SetCreated(dateTime string) *ChargeItemResponseBuilder {
	b.Data["created"] = dateTime
	return b
}

// SetDescription sets description, as the PHP builder does.
func (b *ChargeItemResponseBuilder) SetDescription(description string) *ChargeItemResponseBuilder {
	b.Data["description"] = description
	return b
}

// SetOutcome sets outcome with a single coding, as the PHP builder does.
func (b *ChargeItemResponseBuilder) SetOutcome(system, code, display string) *ChargeItemResponseBuilder {
	b.Data["outcome"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetRequest sets request by reference, as the PHP builder does.
func (b *ChargeItemResponseBuilder) SetRequest(reference string) *ChargeItemResponseBuilder {
	b.Data["request"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetChargeItem sets chargeItem by reference.
func (b *ChargeItemResponseBuilder) SetChargeItem(reference string) *ChargeItemResponseBuilder {
	b.Data["chargeItem"] = map[string]interface{}{"reference": reference}
	return b
}

// SetPartOf adds a partOf reference. Named "set" in the canonical builder but it pushes.
func (b *ChargeItemResponseBuilder) SetPartOf(reference string) *ChargeItemResponseBuilder {
	pushPath(b.Data, "partOf", map[string]interface{}{"reference": reference})
	return b
}

// SetRequestor sets requestor by reference.
func (b *ChargeItemResponseBuilder) SetRequestor(reference string) *ChargeItemResponseBuilder {
	b.Data["requestor"] = map[string]interface{}{"reference": reference}
	return b
}

// AddOutcomeInfo appends an outcome coding.
func (b *ChargeItemResponseBuilder) AddOutcomeInfo(system, code, display string) *ChargeItemResponseBuilder {
	pushPath(b.Data, "outcomeInfo", map[string]interface{}{
		"code": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
		},
	})
	return b
}

// AddError appends an error coding.
func (b *ChargeItemResponseBuilder) AddError(system, code, display string) *ChargeItemResponseBuilder {
	pushPath(b.Data, "error", map[string]interface{}{
		"code": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
		},
	})
	return b
}

// AddCostCenterChargeItemResponse appends a cost centre response with an optional display.
func (b *ChargeItemResponseBuilder) AddCostCenterChargeItemResponse(reference, display string) *ChargeItemResponseBuilder {
	pusat := map[string]interface{}{"reference": reference}
	if display != "" {
		pusat["display"] = display
	}
	pushPath(b.Data, "costCenterChargeItemResponse", map[string]interface{}{"costCenter": pusat})
	return b
}

// SetResponseDate sets responseDate.
func (b *ChargeItemResponseBuilder) SetResponseDate(date string) *ChargeItemResponseBuilder {
	b.Data["responseDate"] = date
	return b
}
