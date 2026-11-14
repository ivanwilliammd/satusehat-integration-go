package parser

// LibraryParser extracts fields from a FHIR Library resource map.
//
// The element list is derived from PayloadBuilderLibrary in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type LibraryParser struct {
	data map[string]any
}

// NewLibraryParser creates a parser for the given Library payload.
func NewLibraryParser(data map[string]any) *LibraryParser {
	if data == nil {
		data = map[string]any{}
	}
	return &LibraryParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *LibraryParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *LibraryParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *LibraryParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *LibraryParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *LibraryParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *LibraryParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *LibraryParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *LibraryParser) GetEncounter() any {
	return p.data["encounter"]
}
