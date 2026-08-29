package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type PaymentReconciliationBuilder struct {
	Data map[string]interface{}
}

func NewPaymentReconciliationBuilder() *PaymentReconciliationBuilder {
	return &PaymentReconciliationBuilder{Data: make(map[string]interface{})}
}

func (b *PaymentReconciliationBuilder) SetId(id string) *PaymentReconciliationBuilder {
	b.Data["id"] = id
	return b
}

func (b *PaymentReconciliationBuilder) SetStatus(status string) *PaymentReconciliationBuilder {
	b.Data["status"] = status
	return b
}

func (b *PaymentReconciliationBuilder) SetCreated(dt string) *PaymentReconciliationBuilder {
	b.Data["created"] = dt
	return b
}

func (b *PaymentReconciliationBuilder) SetPeriodStart(dt string) *PaymentReconciliationBuilder {
	b.Data["period"] = map[string]string{"start": dt}
	return b
}

func (b *PaymentReconciliationBuilder) SetPeriodEnd(dt string) *PaymentReconciliationBuilder {
	if p, ok := b.Data["period"].(map[string]string); ok {
		p["end"] = dt
	} else {
		b.Data["periodEnd"] = dt
	}
	return b
}

func (b *PaymentReconciliationBuilder) SetRequestor(ref *datatype.Reference) *PaymentReconciliationBuilder {
	b.Data["requestor"] = ref.ToArray()
	return b
}

func (b *PaymentReconciliationBuilder) SetOutcome(outcome string) *PaymentReconciliationBuilder {
	b.Data["outcome"] = outcome
	return b
}

func (b *PaymentReconciliationBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "PaymentReconciliation"
	return dt
}
