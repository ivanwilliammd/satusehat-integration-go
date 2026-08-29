package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type InvoiceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewInvoiceBuilder() *InvoiceBuilder {
	b := &InvoiceBuilder{ResourceType: "Invoice", Data: make(map[string]interface{})}
	return b
}

func (b *InvoiceBuilder) SetId(id string) *InvoiceBuilder { b.Data["id"] = id; return b }
func (b *InvoiceBuilder) AddIdentifier(id *datatype.Identifier) *InvoiceBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}
func (b *InvoiceBuilder) SetStatus(status string) *InvoiceBuilder { b.Data["status"] = status; return b }
func (b *InvoiceBuilder) SetType(code *datatype.CodeableConcept) *InvoiceBuilder { b.Data["type"] = code.ToArray(); return b }
func (b *InvoiceBuilder) SetDate(date string) *InvoiceBuilder { b.Data["date"] = date; return b }
func (b *InvoiceBuilder) SetParticipant(participantRef string, role string) *InvoiceBuilder {
	if _, ok := b.Data["participant"]; !ok { b.Data["participant"] = make([]interface{}, 0) }
	b.Data["participant"] = append(b.Data["participant"].([]interface{}), map[string]interface{}{
		"actor":     map[string]interface{}{"reference": participantRef},
		"role":      []interface{}{map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": role}}}},
	})
	return b
}
func (b *InvoiceBuilder) SetIssuer(issuerRef string) *InvoiceBuilder { b.Data["issuer"] = map[string]interface{}{"reference": issuerRef}; return b }
func (b *InvoiceBuilder) SetAccount(accountRef string) *InvoiceBuilder { b.Data["account"] = map[string]interface{}{"reference": accountRef}; return b }
func (b *InvoiceBuilder) SetTotalPrice(amount float64, currency string) *InvoiceBuilder {
	b.Data["totalPriceComponent"] = []interface{}{
		map[string]interface{}{"type": "base", "amount": map[string]interface{}{"value": amount, "currency": currency}},
	}
	return b
}
func (b *InvoiceBuilder) SetTotalNet(amount float64, currency string) *InvoiceBuilder {
	b.Data["totalNet"] = map[string]interface{}{"value": amount, "currency": currency}
	return b
}
func (b *InvoiceBuilder) AddLineItem(description string, priceAmount float64, currency string, quantity int) *InvoiceBuilder {
	if _, ok := b.Data["lineItem"]; !ok { b.Data["lineItem"] = make([]interface{}, 0) }
	b.Data["lineItem"] = append(b.Data["lineItem"].([]interface{}), map[string]interface{}{
		"priceComponent": []interface{}{
			map[string]interface{}{"type": "base", "amount": map[string]interface{}{"value": priceAmount, "currency": currency}},
		},
	})
	return b
}
func (b *InvoiceBuilder) Build() map[string]interface{} { return b.Data }
