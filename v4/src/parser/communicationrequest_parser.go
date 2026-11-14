package parser

// CommunicationRequestParser extracts fields from a FHIR CommunicationRequest resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/CommunicationRequest.json, copied unchanged to
// src/parser/testdata/CommunicationRequest.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type CommunicationRequestParser struct {
	payload map[string]any
}

// NewCommunicationRequestParser creates a new CommunicationRequestParser instance.
func NewCommunicationRequestParser(payload map[string]any) *CommunicationRequestParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &CommunicationRequestParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *CommunicationRequestParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *CommunicationRequestParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *CommunicationRequestParser) GetStatus() any {
	return p.payload["status"]
}

// GetCategory returns the category element unchanged.
func (p *CommunicationRequestParser) GetCategory() any {
	return p.payload["category"]
}

// GetPayload returns the payload element unchanged.
func (p *CommunicationRequestParser) GetPayload() any {
	return p.payload["payload"]
}

// GetAbout returns the about element unchanged.
func (p *CommunicationRequestParser) GetAbout() any {
	return p.payload["about"]
}

// GetSender returns the sender element unchanged.
func (p *CommunicationRequestParser) GetSender() any {
	return p.payload["sender"]
}

// GetRecipient returns the recipient element unchanged.
func (p *CommunicationRequestParser) GetRecipient() any {
	return p.payload["recipient"]
}
