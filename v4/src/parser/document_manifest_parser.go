package parser

// DocumentManifestParser extracts fields from a FHIR DocumentManifest resource map.
//
// The element list is derived from PayloadBuilderDocumentManifest in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type DocumentManifestParser struct {
	data map[string]any
}

// NewDocumentManifestParser creates a parser for the given DocumentManifest payload.
func NewDocumentManifestParser(data map[string]any) *DocumentManifestParser {
	if data == nil {
		data = map[string]any{}
	}
	return &DocumentManifestParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *DocumentManifestParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *DocumentManifestParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *DocumentManifestParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *DocumentManifestParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *DocumentManifestParser) GetStatus() any {
	return p.data["status"]
}

// GetType returns the type element.
func (p *DocumentManifestParser) GetType() any {
	return p.data["type"]
}

// GetSubject returns the subject element.
func (p *DocumentManifestParser) GetSubject() any {
	return p.data["subject"]
}

// GetCreated returns the created element.
func (p *DocumentManifestParser) GetCreated() any {
	return p.data["created"]
}

// GetSource returns the source element.
func (p *DocumentManifestParser) GetSource() any {
	return p.data["source"]
}
