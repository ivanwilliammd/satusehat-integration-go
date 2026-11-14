package parser

// PaymentNoticeParser extracts fields from a FHIR PaymentNotice resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/PaymentNotice.json, copied unchanged to
// src/parser/testdata/PaymentNotice.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type PaymentNoticeParser struct {
	payload map[string]any
}

// NewPaymentNoticeParser creates a new PaymentNoticeParser instance.
func NewPaymentNoticeParser(payload map[string]any) *PaymentNoticeParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &PaymentNoticeParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *PaymentNoticeParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *PaymentNoticeParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *PaymentNoticeParser) GetStatus() any {
	return p.payload["status"]
}

// GetRequest returns the request element unchanged.
func (p *PaymentNoticeParser) GetRequest() any {
	return p.payload["request"]
}

// GetResponse returns the response element unchanged.
func (p *PaymentNoticeParser) GetResponse() any {
	return p.payload["response"]
}

// GetCreated returns the created element unchanged.
func (p *PaymentNoticeParser) GetCreated() any {
	return p.payload["created"]
}

// GetPaymentDate returns the paymentDate element unchanged.
func (p *PaymentNoticeParser) GetPaymentDate() any {
	return p.payload["paymentDate"]
}

// GetProvider returns the provider element unchanged.
func (p *PaymentNoticeParser) GetProvider() any {
	return p.payload["provider"]
}

// GetPayee returns the payee element unchanged.
func (p *PaymentNoticeParser) GetPayee() any {
	return p.payload["payee"]
}

// GetRecipient returns the recipient element unchanged.
func (p *PaymentNoticeParser) GetRecipient() any {
	return p.payload["recipient"]
}

// GetAmount returns the amount element unchanged.
func (p *PaymentNoticeParser) GetAmount() any {
	return p.payload["amount"]
}

// GetPayment returns the payment element unchanged.
func (p *PaymentNoticeParser) GetPayment() any {
	return p.payload["payment"]
}

// GetPaymentStatus returns the paymentStatus element unchanged.
func (p *PaymentNoticeParser) GetPaymentStatus() any {
	return p.payload["paymentStatus"]
}
