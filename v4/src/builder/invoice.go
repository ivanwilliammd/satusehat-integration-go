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

// AddLineItem appends an invoice line. A unit price becomes a base price component and a price factor a
// separate factor component, which is the shape the canonical builder writes.
func (b *InvoiceBuilder) AddLineItem(sequence int, serviceSystem, serviceCode, serviceDisplay string, quantity, unitPrice, priceFactor *float64) *InvoiceBuilder {
	baris := map[string]interface{}{
		"sequence": sequence,
		"service": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{
				"system": serviceSystem, "code": serviceCode, "display": serviceDisplay,
			}},
		},
	}
	if quantity != nil {
		baris["quantity"] = map[string]interface{}{"value": *quantity}
	}
	if unitPrice != nil {
		baris["priceComponent"] = append(komponenHarga(baris), map[string]interface{}{
			"type": "base", "amount": map[string]interface{}{"value": *unitPrice},
		})
	}
	if priceFactor != nil {
		baris["priceComponent"] = append(komponenHarga(baris), map[string]interface{}{
			"type": "factor", "factor": *priceFactor,
		})
	}
	pushPath(b.Data, "lineItem", baris)
	return b
}

func komponenHarga(baris map[string]interface{}) []interface{} {
	if ada, ok := baris["priceComponent"].([]interface{}); ok {
		return ada
	}
	return []interface{}{}
}

// SetParticipant adds a participant with its role coding and actor.
func (b *InvoiceBuilder) SetParticipant(roleSystem, roleCode, reference, display string) *InvoiceBuilder {
	aktor := map[string]interface{}{"reference": reference}
	if display != "" {
		aktor["display"] = display
	}
	pushPath(b.Data, "participant", map[string]interface{}{
		"role": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": roleSystem, "code": roleCode}},
		},
		"actor": aktor,
	})
	return b
}

// SetIssuer sets issuer by reference.
func (b *InvoiceBuilder) SetIssuer(reference string) *InvoiceBuilder {
	b.Data["issuer"] = map[string]interface{}{"reference": reference}
	return b
}

// SetCustomer sets the customer. The canonical builder writes it as "subject" although the method is
// called setCustomer, and so does this.
func (b *InvoiceBuilder) SetCustomer(reference string) *InvoiceBuilder {
	b.Data["subject"] = map[string]interface{}{"reference": reference}
	return b
}

// SetAccount sets account by reference.
func (b *InvoiceBuilder) SetAccount(reference string) *InvoiceBuilder {
	b.Data["account"] = map[string]interface{}{"reference": reference}
	return b
}

// SetTotalNet writes totalNet through nested paths.
func (b *InvoiceBuilder) SetTotalNet(value float64, code string) *InvoiceBuilder {
	setPath(b.Data, "totalNet/value", value)
	setPath(b.Data, "totalNet/currency", code)
	return b
}

// SetTotalGross writes totalGross through nested paths.
func (b *InvoiceBuilder) SetTotalGross(value float64, code string) *InvoiceBuilder {
	setPath(b.Data, "totalGross/value", value)
	setPath(b.Data, "totalGross/currency", code)
	return b
}

// SetTotalVAT writes totalVAT through nested paths.
func (b *InvoiceBuilder) SetTotalVAT(value float64, code string) *InvoiceBuilder {
	setPath(b.Data, "totalVAT/value", value)
	setPath(b.Data, "totalVAT/currency", code)
	return b
}

// SetPaymentTerms sets paymentTerms.
func (b *InvoiceBuilder) SetPaymentTerms(text string) *InvoiceBuilder {
	b.Data["paymentTerms"] = text
	return b
}

// SetCancelledReason sets cancelledReason.
func (b *InvoiceBuilder) SetCancelledReason(reason string) *InvoiceBuilder {
	b.Data["cancelledReason"] = reason
	return b
}
