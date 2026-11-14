package builder

import (
	"encoding/json"
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
