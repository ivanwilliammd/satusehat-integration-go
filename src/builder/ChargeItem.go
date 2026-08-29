package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type ChargeItemBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewChargeItemBuilder() *ChargeItemBuilder {
	b := &ChargeItemBuilder{ResourceType: "ChargeItem", Data: make(map[string]interface{})}
	return b
}

func (b *ChargeItemBuilder) SetId(id string) *ChargeItemBuilder { b.Data["id"] = id; return b }
func (b *ChargeItemBuilder) AddIdentifier(id *datatype.Identifier) *ChargeItemBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}
func (b *ChargeItemBuilder) SetStatus(status string) *ChargeItemBuilder { b.Data["status"] = status; return b }
func (b *ChargeItemBuilder) SetCode(code *datatype.CodeableConcept) *ChargeItemBuilder { b.Data["code"] = code.ToArray(); return b }
func (b *ChargeItemBuilder) SetSubject(subjectRef string) *ChargeItemBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *ChargeItemBuilder) SetEncounter(encRef string) *ChargeItemBuilder { b.Data["context"] = map[string]interface{}{"reference": encRef}; return b }
func (b *ChargeItemBuilder) SetOccurrence(occurrence string) *ChargeItemBuilder { b.Data["occurrenceDateTime"] = occurrence; return b }
func (b *ChargeItemBuilder) SetPerformer(performerRef string, function string) *ChargeItemBuilder {
	if _, ok := b.Data["performer"]; !ok { b.Data["performer"] = make([]interface{}, 0) }
	b.Data["performer"] = append(b.Data["performer"].([]interface{}), map[string]interface{}{
		"actor":    map[string]interface{}{"reference": performerRef},
		"function": map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": function}}},
	})
	return b
}
func (b *ChargeItemBuilder) SetPerformingOrganization(orgRef string) *ChargeItemBuilder { b.Data["performingOrganization"] = map[string]interface{}{"reference": orgRef}; return b }
func (b *ChargeItemBuilder) SetRequestingOrganization(orgRef string) *ChargeItemBuilder { b.Data["requestingOrganization"] = map[string]interface{}{"reference": orgRef}; return b }
func (b *ChargeItemBuilder) SetQuantity(qty int) *ChargeItemBuilder { b.Data["quantity"] = qty; return b }
func (b *ChargeItemBuilder) SetBodysite(bodySite *datatype.CodeableConcept) *ChargeItemBuilder { b.Data["bodysite"] = []interface{}{bodySite.ToArray()}; return b }
func (b *ChargeItemBuilder) SetNetCost(amount float64, currency string) *ChargeItemBuilder {
	b.Data["priceOverride"] = map[string]interface{}{"value": amount, "currency": currency}
	return b
}
func (b *ChargeItemBuilder) SetOverrideReason(reason string) *ChargeItemBuilder { b.Data["overrideReason"] = reason; return b }
func (b *ChargeItemBuilder) Build() map[string]interface{} { return b.Data }
