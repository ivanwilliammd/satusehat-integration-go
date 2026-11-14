package parser

// SubstancePolymerParser extracts fields from a FHIR SubstancePolymer resource map.
//
// The element list is derived from PayloadBuilderSubstancePolymer in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type SubstancePolymerParser struct {
	data map[string]any
}

// NewSubstancePolymerParser creates a parser for the given SubstancePolymer payload.
func NewSubstancePolymerParser(data map[string]any) *SubstancePolymerParser {
	if data == nil {
		data = map[string]any{}
	}
	return &SubstancePolymerParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *SubstancePolymerParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *SubstancePolymerParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *SubstancePolymerParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *SubstancePolymerParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *SubstancePolymerParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *SubstancePolymerParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *SubstancePolymerParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *SubstancePolymerParser) GetEncounter() any {
	return p.data["encounter"]
}
