package parser

// SupplyDeliveryParser extracts fields from a FHIR SupplyDelivery resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/SupplyDelivery.json, copied unchanged to
// src/parser/testdata/SupplyDelivery.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type SupplyDeliveryParser struct {
	payload map[string]any
}

// NewSupplyDeliveryParser creates a new SupplyDeliveryParser instance.
func NewSupplyDeliveryParser(payload map[string]any) *SupplyDeliveryParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &SupplyDeliveryParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *SupplyDeliveryParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *SupplyDeliveryParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *SupplyDeliveryParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetBasedOn returns the basedOn element unchanged.
func (p *SupplyDeliveryParser) GetBasedOn() any {
	return p.payload["basedOn"]
}

// GetStatus returns the status element unchanged.
func (p *SupplyDeliveryParser) GetStatus() any {
	return p.payload["status"]
}

// GetPatient returns the patient element unchanged.
func (p *SupplyDeliveryParser) GetPatient() any {
	return p.payload["patient"]
}

// GetSuppliedItem returns the suppliedItem element unchanged.
func (p *SupplyDeliveryParser) GetSuppliedItem() any {
	return p.payload["suppliedItem"]
}

// GetOccurrenceDateTime returns the occurrenceDateTime element unchanged.
func (p *SupplyDeliveryParser) GetOccurrenceDateTime() any {
	return p.payload["occurrenceDateTime"]
}

// GetSupplier returns the supplier element unchanged.
func (p *SupplyDeliveryParser) GetSupplier() any {
	return p.payload["supplier"]
}

// GetDestination returns the destination element unchanged.
func (p *SupplyDeliveryParser) GetDestination() any {
	return p.payload["destination"]
}

// GetReceiver returns the receiver element unchanged.
func (p *SupplyDeliveryParser) GetReceiver() any {
	return p.payload["receiver"]
}
