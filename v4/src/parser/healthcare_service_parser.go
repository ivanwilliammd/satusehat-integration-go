package parser

// HealthcareServiceParser extracts fields from a FHIR HealthcareService resource map.
//
// The element list is derived from PayloadBuilderHealthcareService in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type HealthcareServiceParser struct {
	data map[string]any
}

// NewHealthcareServiceParser creates a parser for the given HealthcareService payload.
func NewHealthcareServiceParser(data map[string]any) *HealthcareServiceParser {
	if data == nil {
		data = map[string]any{}
	}
	return &HealthcareServiceParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *HealthcareServiceParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *HealthcareServiceParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *HealthcareServiceParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *HealthcareServiceParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *HealthcareServiceParser) GetStatus() any {
	return p.data["status"]
}

// GetName returns the name element.
func (p *HealthcareServiceParser) GetName() any {
	return p.data["name"]
}

// GetType returns the type element.
func (p *HealthcareServiceParser) GetType() any {
	return p.data["type"]
}

// GetLocation returns the location element.
func (p *HealthcareServiceParser) GetLocation() any {
	return p.data["location"]
}

// GetProvidedBy returns the providedBy element.
func (p *HealthcareServiceParser) GetProvidedBy() any {
	return p.data["providedBy"]
}
