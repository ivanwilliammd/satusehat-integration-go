package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type PaymentNoticeBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewPaymentNoticeBuilder() *PaymentNoticeBuilder {
	b := &PaymentNoticeBuilder{ResourceType: "PaymentNotice", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "PaymentNotice"
	return b
}

func (b *PaymentNoticeBuilder) SetId(id string) *PaymentNoticeBuilder {
	b.Data["id"] = id
	return b
}

func (b *PaymentNoticeBuilder) AddIdentifier(identifier *datatype.Identifier) *PaymentNoticeBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *PaymentNoticeBuilder) SetStatus(status string) *PaymentNoticeBuilder {
	b.Data["status"] = status
	return b
}

func (b *PaymentNoticeBuilder) SetSubject(reference string, display ...string) *PaymentNoticeBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *PaymentNoticeBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *PaymentNoticeBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetInsurer sets insurer by reference, as the PHP builder does.
func (b *PaymentNoticeBuilder) SetInsurer(reference string) *PaymentNoticeBuilder {
	b.Data["insurer"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetCreated sets created, as the PHP builder does.
func (b *PaymentNoticeBuilder) SetCreated(dateTime string) *PaymentNoticeBuilder {
	b.Data["created"] = dateTime
	return b
}

// SetRequest sets request by reference, as the PHP builder does.
func (b *PaymentNoticeBuilder) SetRequest(reference string) *PaymentNoticeBuilder {
	b.Data["request"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}
