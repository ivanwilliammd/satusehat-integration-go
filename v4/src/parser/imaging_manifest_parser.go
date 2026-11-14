package parser

// ImagingManifestParser extracts fields from a FHIR ImagingManifest resource map.
//
// The element list is derived from PayloadBuilderImagingManifest in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ImagingManifestParser struct {
	data map[string]any
}

// NewImagingManifestParser creates a parser for the given ImagingManifest payload.
func NewImagingManifestParser(data map[string]any) *ImagingManifestParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ImagingManifestParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ImagingManifestParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ImagingManifestParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ImagingManifestParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ImagingManifestParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ImagingManifestParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *ImagingManifestParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *ImagingManifestParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *ImagingManifestParser) GetEncounter() any {
	return p.data["encounter"]
}
