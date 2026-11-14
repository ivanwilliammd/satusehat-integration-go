package parser

// DeviceDispenseParser extracts fields from a FHIR DeviceDispense resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/DeviceDispense.json, copied unchanged to
// src/parser/testdata/DeviceDispense.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type DeviceDispenseParser struct {
	payload map[string]any
}

// NewDeviceDispenseParser creates a new DeviceDispenseParser instance.
func NewDeviceDispenseParser(payload map[string]any) *DeviceDispenseParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &DeviceDispenseParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *DeviceDispenseParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *DeviceDispenseParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *DeviceDispenseParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *DeviceDispenseParser) GetStatus() any {
	return p.payload["status"]
}

// GetBasedOn returns the basedOn element unchanged.
func (p *DeviceDispenseParser) GetBasedOn() any {
	return p.payload["basedOn"]
}

// GetDevice returns the device element unchanged.
func (p *DeviceDispenseParser) GetDevice() any {
	return p.payload["device"]
}

// GetSubject returns the subject element unchanged.
func (p *DeviceDispenseParser) GetSubject() any {
	return p.payload["subject"]
}

// GetReceiver returns the receiver element unchanged.
func (p *DeviceDispenseParser) GetReceiver() any {
	return p.payload["receiver"]
}

// GetPerformer returns the performer element unchanged.
func (p *DeviceDispenseParser) GetPerformer() any {
	return p.payload["performer"]
}

// GetWhenPrepared returns the whenPrepared element unchanged.
func (p *DeviceDispenseParser) GetWhenPrepared() any {
	return p.payload["whenPrepared"]
}

// GetWhenHandedOver returns the whenHandedOver element unchanged.
func (p *DeviceDispenseParser) GetWhenHandedOver() any {
	return p.payload["whenHandedOver"]
}

// GetDestination returns the destination element unchanged.
func (p *DeviceDispenseParser) GetDestination() any {
	return p.payload["destination"]
}

// GetNote returns the note element unchanged.
func (p *DeviceDispenseParser) GetNote() any {
	return p.payload["note"]
}
