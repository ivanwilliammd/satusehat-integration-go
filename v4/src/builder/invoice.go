package builder

import (
	"encoding/json"
	"fmt"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type InvoiceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewInvoiceBuilder() *InvoiceBuilder {
	b := &InvoiceBuilder{ResourceType: "Invoice", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Invoice"
	return b
}

func (b *InvoiceBuilder) SetId(id string) *InvoiceBuilder {
	b.Data["id"] = id
	return b
}

func (b *InvoiceBuilder) AddIdentifier(identifier *datatype.Identifier) *InvoiceBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *InvoiceBuilder) SetStatus(status string) *InvoiceBuilder {
	b.Data["status"] = status
	return b
}

func (b *InvoiceBuilder) SetSubject(reference string, display ...string) *InvoiceBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *InvoiceBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *InvoiceBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetType sets type with a single coding, as the PHP builder does.
func (b *InvoiceBuilder) SetType(system, code, display string) *InvoiceBuilder {
	b.Data["type"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetDate sets date, as the PHP builder does.
func (b *InvoiceBuilder) SetDate(date string) *InvoiceBuilder {
	b.Data["date"] = date
	return b
}

// AddNote adds a note with plain text, as the PHP builder does.
func (b *InvoiceBuilder) AddNote(text string) *InvoiceBuilder {
	daftar, _ := b.Data["note"].([]interface{})
	b.Data["note"] = append(daftar, map[string]interface{}{"text": text})
	return b
}

// AddLineItemPriceComponent adds a price component to an existing line item.
func (b *InvoiceBuilder) AddLineItemPriceComponent(itemIdx int, kind string, amount float64, code *string) *InvoiceBuilder {
	komponen := map[string]interface{}{"type": kind, "amount": map[string]interface{}{"value": amount}}
	if code != nil {
		komponen["code"] = *code
	}
	pushPath(b.Data, fmt.Sprintf("lineItem/%d/priceComponent", itemIdx), komponen)
	return b
}
