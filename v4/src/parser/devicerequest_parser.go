package parser

// DeviceRequestParser extracts fields from a FHIR DeviceRequest resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/DeviceRequest.json, copied unchanged to
// src/parser/testdata/DeviceRequest.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type DeviceRequestParser struct {
	payload map[string]any
}

// NewDeviceRequestParser creates a new DeviceRequestParser instance.
func NewDeviceRequestParser(payload map[string]any) *DeviceRequestParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &DeviceRequestParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *DeviceRequestParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *DeviceRequestParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *DeviceRequestParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *DeviceRequestParser) GetStatus() any {
	return p.payload["status"]
}

// GetIntent returns the intent element unchanged.
func (p *DeviceRequestParser) GetIntent() any {
	return p.payload["intent"]
}

// GetPriority returns the priority element unchanged.
func (p *DeviceRequestParser) GetPriority() any {
	return p.payload["priority"]
}

// GetCodeReference returns the codeReference element unchanged.
func (p *DeviceRequestParser) GetCodeReference() any {
	return p.payload["codeReference"]
}

// GetSubject returns the subject element unchanged.
func (p *DeviceRequestParser) GetSubject() any {
	return p.payload["subject"]
}

// GetAuthoredOn returns the authoredOn element unchanged.
func (p *DeviceRequestParser) GetAuthoredOn() any {
	return p.payload["authoredOn"]
}

// GetRequester returns the requester element unchanged.
func (p *DeviceRequestParser) GetRequester() any {
	return p.payload["requester"]
}

// GetPerformer returns the performer element unchanged.
func (p *DeviceRequestParser) GetPerformer() any {
	return p.payload["performer"]
}

// GetReasonCode returns the reasonCode element unchanged.
func (p *DeviceRequestParser) GetReasonCode() any {
	return p.payload["reasonCode"]
}
