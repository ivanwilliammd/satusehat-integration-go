package parser

// PaymentReconciliationParser extracts fields from a FHIR PaymentReconciliation resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/PaymentReconciliation.json, copied unchanged to
// src/parser/testdata/PaymentReconciliation.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type PaymentReconciliationParser struct {
	payload map[string]any
}

// NewPaymentReconciliationParser creates a new PaymentReconciliationParser instance.
func NewPaymentReconciliationParser(payload map[string]any) *PaymentReconciliationParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &PaymentReconciliationParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *PaymentReconciliationParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *PaymentReconciliationParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *PaymentReconciliationParser) GetStatus() any {
	return p.payload["status"]
}

// GetPeriod returns the period element unchanged.
func (p *PaymentReconciliationParser) GetPeriod() any {
	return p.payload["period"]
}

// GetCreated returns the created element unchanged.
func (p *PaymentReconciliationParser) GetCreated() any {
	return p.payload["created"]
}

// GetPaymentIssuer returns the paymentIssuer element unchanged.
func (p *PaymentReconciliationParser) GetPaymentIssuer() any {
	return p.payload["paymentIssuer"]
}

// GetRequestor returns the requestor element unchanged.
func (p *PaymentReconciliationParser) GetRequestor() any {
	return p.payload["requestor"]
}

// GetOutcome returns the outcome element unchanged.
func (p *PaymentReconciliationParser) GetOutcome() any {
	return p.payload["outcome"]
}

// GetDisposition returns the disposition element unchanged.
func (p *PaymentReconciliationParser) GetDisposition() any {
	return p.payload["disposition"]
}

// GetPaymentDate returns the paymentDate element unchanged.
func (p *PaymentReconciliationParser) GetPaymentDate() any {
	return p.payload["paymentDate"]
}

// GetPaymentAmount returns the paymentAmount element unchanged.
func (p *PaymentReconciliationParser) GetPaymentAmount() any {
	return p.payload["paymentAmount"]
}

// GetDetail returns the detail element unchanged.
func (p *PaymentReconciliationParser) GetDetail() any {
	return p.payload["detail"]
}

// GetProcessNote returns the processNote element unchanged.
func (p *PaymentReconciliationParser) GetProcessNote() any {
	return p.payload["processNote"]
}
