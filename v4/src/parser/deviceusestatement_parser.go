package parser

// DeviceUseStatementParser extracts fields from a FHIR DeviceUseStatement resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/DeviceUseStatement.json, copied unchanged to
// src/parser/testdata/DeviceUseStatement.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type DeviceUseStatementParser struct {
	payload map[string]any
}

// NewDeviceUseStatementParser creates a new DeviceUseStatementParser instance.
func NewDeviceUseStatementParser(payload map[string]any) *DeviceUseStatementParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &DeviceUseStatementParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *DeviceUseStatementParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *DeviceUseStatementParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *DeviceUseStatementParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *DeviceUseStatementParser) GetStatus() any {
	return p.payload["status"]
}

// GetDevice returns the device element unchanged.
func (p *DeviceUseStatementParser) GetDevice() any {
	return p.payload["device"]
}

// GetSubject returns the subject element unchanged.
func (p *DeviceUseStatementParser) GetSubject() any {
	return p.payload["subject"]
}

// GetTimingPeriod returns the timingPeriod element unchanged.
func (p *DeviceUseStatementParser) GetTimingPeriod() any {
	return p.payload["timingPeriod"]
}

// GetRecordedOn returns the recordedOn element unchanged.
func (p *DeviceUseStatementParser) GetRecordedOn() any {
	return p.payload["recordedOn"]
}

// GetSource returns the source element unchanged.
func (p *DeviceUseStatementParser) GetSource() any {
	return p.payload["source"]
}

// GetReasonCode returns the reasonCode element unchanged.
func (p *DeviceUseStatementParser) GetReasonCode() any {
	return p.payload["reasonCode"]
}

// GetNote returns the note element unchanged.
func (p *DeviceUseStatementParser) GetNote() any {
	return p.payload["note"]
}
