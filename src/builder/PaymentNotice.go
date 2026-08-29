package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type PaymentNoticeBuilder struct {
	Data map[string]interface{}
}

func NewPaymentNoticeBuilder() *PaymentNoticeBuilder {
	return &PaymentNoticeBuilder{Data: make(map[string]interface{})}
}

func (b *PaymentNoticeBuilder) SetId(id string) *PaymentNoticeBuilder {
	b.Data["id"] = id
	return b
}

func (b *PaymentNoticeBuilder) SetStatus(status string) *PaymentNoticeBuilder {
	b.Data["status"] = status
	return b
}

func (b *PaymentNoticeBuilder) SetRequest(ref string) *PaymentNoticeBuilder {
	b.Data["request"] = map[string]string{"reference": ref}
	return b
}

func (b *PaymentNoticeBuilder) SetResponse(ref string) *PaymentNoticeBuilder {
	b.Data["response"] = map[string]string{"reference": ref}
	return b
}

func (b *PaymentNoticeBuilder) SetCreated(dt string) *PaymentNoticeBuilder {
	b.Data["created"] = dt
	return b
}

func (b *PaymentNoticeBuilder) SetProvider(ref *datatype.Reference) *PaymentNoticeBuilder {
	b.Data["provider"] = ref.ToArray()
	return b
}

func (b *PaymentNoticeBuilder) SetAmount(money *datatype.Money) *PaymentNoticeBuilder {
	b.Data["amount"] = money.ToArray()
	return b
}

func (b *PaymentNoticeBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "PaymentNotice"
	return dt
}
