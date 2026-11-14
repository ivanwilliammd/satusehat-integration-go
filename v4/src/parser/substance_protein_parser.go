package parser

// SubstanceProteinParser extracts fields from a FHIR SubstanceProtein resource map.
//
// The element list is derived from PayloadBuilderSubstanceProtein in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type SubstanceProteinParser struct {
	data map[string]any
}

// NewSubstanceProteinParser creates a parser for the given SubstanceProtein payload.
func NewSubstanceProteinParser(data map[string]any) *SubstanceProteinParser {
	if data == nil {
		data = map[string]any{}
	}
	return &SubstanceProteinParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *SubstanceProteinParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *SubstanceProteinParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *SubstanceProteinParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *SubstanceProteinParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *SubstanceProteinParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *SubstanceProteinParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *SubstanceProteinParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *SubstanceProteinParser) GetEncounter() any {
	return p.data["encounter"]
}
