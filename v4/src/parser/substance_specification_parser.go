package parser

// SubstanceSpecificationParser extracts fields from a FHIR SubstanceSpecification resource map.
//
// The element list is derived from PayloadBuilderSubstanceSpecification in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type SubstanceSpecificationParser struct {
	data map[string]any
}

// NewSubstanceSpecificationParser creates a parser for the given SubstanceSpecification payload.
func NewSubstanceSpecificationParser(data map[string]any) *SubstanceSpecificationParser {
	if data == nil {
		data = map[string]any{}
	}
	return &SubstanceSpecificationParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *SubstanceSpecificationParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *SubstanceSpecificationParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *SubstanceSpecificationParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *SubstanceSpecificationParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *SubstanceSpecificationParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *SubstanceSpecificationParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *SubstanceSpecificationParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *SubstanceSpecificationParser) GetEncounter() any {
	return p.data["encounter"]
}
