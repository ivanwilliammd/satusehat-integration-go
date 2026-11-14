package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type PaymentReconciliationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewPaymentReconciliationBuilder() *PaymentReconciliationBuilder {
	b := &PaymentReconciliationBuilder{ResourceType: "PaymentReconciliation", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "PaymentReconciliation"
	return b
}

func (b *PaymentReconciliationBuilder) SetId(id string) *PaymentReconciliationBuilder {
	b.Data["id"] = id
	return b
}

func (b *PaymentReconciliationBuilder) AddIdentifier(identifier *datatype.Identifier) *PaymentReconciliationBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *PaymentReconciliationBuilder) SetStatus(status string) *PaymentReconciliationBuilder {
	b.Data["status"] = status
	return b
}

func (b *PaymentReconciliationBuilder) SetSubject(reference string, display ...string) *PaymentReconciliationBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *PaymentReconciliationBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *PaymentReconciliationBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetInsurer sets insurer by reference, as the PHP builder does.
func (b *PaymentReconciliationBuilder) SetInsurer(reference string) *PaymentReconciliationBuilder {
	b.Data["insurer"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetCreated sets created, as the PHP builder does.
func (b *PaymentReconciliationBuilder) SetCreated(dateTime string) *PaymentReconciliationBuilder {
	b.Data["created"] = dateTime
	return b
}

// SetType sets type with a single coding, as the PHP builder does.
func (b *PaymentReconciliationBuilder) SetType(system, code, display string) *PaymentReconciliationBuilder {
	b.Data["type"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}
