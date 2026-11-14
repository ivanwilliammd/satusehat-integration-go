package parser

// EndpointParser extracts fields from a FHIR Endpoint resource map.
//
// The element list is derived from PayloadBuilderEndpoint in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type EndpointParser struct {
	data map[string]any
}

// NewEndpointParser creates a parser for the given Endpoint payload.
func NewEndpointParser(data map[string]any) *EndpointParser {
	if data == nil {
		data = map[string]any{}
	}
	return &EndpointParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *EndpointParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *EndpointParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *EndpointParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *EndpointParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *EndpointParser) GetStatus() any {
	return p.data["status"]
}

// GetConnectionType returns the connectionType element.
func (p *EndpointParser) GetConnectionType() any {
	return p.data["connectionType"]
}

// GetName returns the name element.
func (p *EndpointParser) GetName() any {
	return p.data["name"]
}

// GetManagingOrganization returns the managingOrganization element.
func (p *EndpointParser) GetManagingOrganization() any {
	return p.data["managingOrganization"]
}

// GetContact returns the contact element.
func (p *EndpointParser) GetContact() any {
	return p.data["contact"]
}

// GetPeriod returns the period element.
func (p *EndpointParser) GetPeriod() any {
	return p.data["period"]
}

// GetPayloadType returns the payloadType element.
func (p *EndpointParser) GetPayloadType() any {
	return p.data["payloadType"]
}

// GetPayloadMimeType returns the payloadMimeType element.
func (p *EndpointParser) GetPayloadMimeType() any {
	return p.data["payloadMimeType"]
}

// GetAddress returns the address element.
func (p *EndpointParser) GetAddress() any {
	return p.data["address"]
}

// GetHeader returns the header element.
func (p *EndpointParser) GetHeader() any {
	return p.data["header"]
}
