package parser

// InvoiceParser extracts fields from a FHIR Invoice resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/Invoice.json, copied unchanged to
// src/parser/testdata/Invoice.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type InvoiceParser struct {
	payload map[string]any
}

// NewInvoiceParser creates a new InvoiceParser instance.
func NewInvoiceParser(payload map[string]any) *InvoiceParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &InvoiceParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *InvoiceParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *InvoiceParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *InvoiceParser) GetStatus() any {
	return p.payload["status"]
}

// GetSubject returns the subject element unchanged.
func (p *InvoiceParser) GetSubject() any {
	return p.payload["subject"]
}

// GetDate returns the date element unchanged.
func (p *InvoiceParser) GetDate() any {
	return p.payload["date"]
}

// GetIssuer returns the issuer element unchanged.
func (p *InvoiceParser) GetIssuer() any {
	return p.payload["issuer"]
}

// GetAccount returns the account element unchanged.
func (p *InvoiceParser) GetAccount() any {
	return p.payload["account"]
}

// GetLineItem returns the lineItem element unchanged.
func (p *InvoiceParser) GetLineItem() any {
	return p.payload["lineItem"]
}

// GetTotalNet returns the totalNet element unchanged.
func (p *InvoiceParser) GetTotalNet() any {
	return p.payload["totalNet"]
}
