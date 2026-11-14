package parser

// SubstanceParser extracts fields from a FHIR Substance resource map.
//
// The element list is derived from PayloadBuilderSubstance in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type SubstanceParser struct {
	data map[string]any
}

// NewSubstanceParser creates a parser for the given Substance payload.
func NewSubstanceParser(data map[string]any) *SubstanceParser {
	if data == nil {
		data = map[string]any{}
	}
	return &SubstanceParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *SubstanceParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *SubstanceParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *SubstanceParser) GetId() any {
	return p.data["id"]
}

// GetStatus returns the status element.
func (p *SubstanceParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *SubstanceParser) GetCode() any {
	return p.data["code"]
}

// GetText returns the text element.
func (p *SubstanceParser) GetText() any {
	return p.data["text"]
}
