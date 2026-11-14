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

// SetOutcome sets outcome with a single coding, as the PHP builder does.
func (b *PaymentReconciliationBuilder) SetOutcome(system, code, display string) *PaymentReconciliationBuilder {
	b.Data["outcome"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetRequest sets request by reference, as the PHP builder does.
func (b *PaymentReconciliationBuilder) SetRequest(reference string) *PaymentReconciliationBuilder {
	b.Data["request"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetDisposition sets disposition, as the PHP builder does.
func (b *PaymentReconciliationBuilder) SetDisposition(disposition string) *PaymentReconciliationBuilder {
	b.Data["disposition"] = disposition
	return b
}

// SetPeriod writes the reconciliation period through nested paths.
func (b *PaymentReconciliationBuilder) SetPeriod(start, end string) *PaymentReconciliationBuilder {
	setPath(b.Data, "period/start", start)
	if end != "" {
		setPath(b.Data, "period/end", end)
	}
	return b
}

// SetRequestProvider sets requestProvider by reference.
func (b *PaymentReconciliationBuilder) SetRequestProvider(reference string) *PaymentReconciliationBuilder {
	b.Data["requestProvider"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetRequestMatchDate sets requestMatchDate.
func (b *PaymentReconciliationBuilder) SetRequestMatchDate(date string) *PaymentReconciliationBuilder {
	b.Data["requestMatchDate"] = date
	return b
}

// SetOutcomeCode adds a coding to outcomeCode. Named "set" in the canonical builder but it pushes, so
// repeated calls accumulate.
func (b *PaymentReconciliationBuilder) SetOutcomeCode(system, code, display string) *PaymentReconciliationBuilder {
	pushPath(b.Data, "outcomeCode", map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	})
	return b
}

// AddRequestor adds a requestor reference with an optional display.
func (b *PaymentReconciliationBuilder) AddRequestor(reference, display string) *PaymentReconciliationBuilder {
	requestor := map[string]interface{}{"reference": reference}
	if display != "" {
		requestor["display"] = display
	}
	pushPath(b.Data, "requestor", requestor)
	return b
}

// AddProcessNote adds a process note with an optional type code.
func (b *PaymentReconciliationBuilder) AddProcessNote(text string, noteType *string) *PaymentReconciliationBuilder {
	note := map[string]interface{}{"text": text}
	if noteType != nil && *noteType != "" {
		note["type"] = *noteType
	}
	pushPath(b.Data, "processNote", note)
	return b
}

// SetPaymentDate sets paymentDate.
func (b *PaymentReconciliationBuilder) SetPaymentDate(date string) *PaymentReconciliationBuilder {
	b.Data["paymentDate"] = date
	return b
}

// SetPaymentAmount writes paymentAmount through nested paths.
func (b *PaymentReconciliationBuilder) SetPaymentAmount(value float64, code string) *PaymentReconciliationBuilder {
	setPath(b.Data, "paymentAmount/value", value)
	setPath(b.Data, "paymentAmount/currency", code)
	return b
}

// SetPaymentIdentifier sets paymentIdentifier as a plain identifier object.
func (b *PaymentReconciliationBuilder) SetPaymentIdentifier(system, value string) *PaymentReconciliationBuilder {
	b.Data["paymentIdentifier"] = map[string]interface{}{"system": system, "value": value}
	return b
}

// AddDetail adds a payment detail with an optional amount and request reference.
func (b *PaymentReconciliationBuilder) AddDetail(typeSystem, typeCode string, amount *float64, requestReference *string) *PaymentReconciliationBuilder {
	detail := map[string]interface{}{
		"type": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": typeSystem, "code": typeCode}},
		},
	}
	if amount != nil {
		// detail.amount bertipe Money, jadi nilainya dibungkus; angka telanjang ditolak server.
		detail["amount"] = map[string]interface{}{"value": *amount}
	}
	if requestReference != nil && *requestReference != "" {
		detail["request"] = map[string]interface{}{"reference": *requestReference}
	}
	pushPath(b.Data, "detail", detail)
	return b
}

// SetTotalAmount writes totalAmount through nested paths.
func (b *PaymentReconciliationBuilder) SetTotalAmount(value float64, code string) *PaymentReconciliationBuilder {
	setPath(b.Data, "totalAmount/value", value)
	setPath(b.Data, "totalAmount/currency", code)
	return b
}
