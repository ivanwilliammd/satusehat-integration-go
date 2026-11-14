package parser

// CommunicationParser extracts fields from a FHIR Communication resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/Communication.json, copied unchanged to
// src/parser/testdata/Communication.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type CommunicationParser struct {
	payload map[string]any
}

// NewCommunicationParser creates a new CommunicationParser instance.
func NewCommunicationParser(payload map[string]any) *CommunicationParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &CommunicationParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *CommunicationParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *CommunicationParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetStatus returns the status element unchanged.
func (p *CommunicationParser) GetStatus() any {
	return p.payload["status"]
}

// GetBasedOn returns the basedOn element unchanged.
func (p *CommunicationParser) GetBasedOn() any {
	return p.payload["basedOn"]
}

// GetPayload returns the payload element unchanged.
func (p *CommunicationParser) GetPayload() any {
	return p.payload["payload"]
}

// GetAbout returns the about element unchanged.
func (p *CommunicationParser) GetAbout() any {
	return p.payload["about"]
}

// GetInResponseTo returns the inResponseTo element unchanged.
func (p *CommunicationParser) GetInResponseTo() any {
	return p.payload["inResponseTo"]
}

// GetSender returns the sender element unchanged.
func (p *CommunicationParser) GetSender() any {
	return p.payload["sender"]
}

// GetRecipient returns the recipient element unchanged.
func (p *CommunicationParser) GetRecipient() any {
	return p.payload["recipient"]
}
