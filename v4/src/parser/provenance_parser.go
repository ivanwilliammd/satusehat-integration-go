package parser

// ProvenanceParser extracts fields from a FHIR Provenance resource map.
//
// The element list is derived from PayloadBuilderProvenance in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ProvenanceParser struct {
	data map[string]any
}

// NewProvenanceParser creates a parser for the given Provenance payload.
func NewProvenanceParser(data map[string]any) *ProvenanceParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ProvenanceParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ProvenanceParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ProvenanceParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ProvenanceParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ProvenanceParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ProvenanceParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *ProvenanceParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *ProvenanceParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *ProvenanceParser) GetEncounter() any {
	return p.data["encounter"]
}
