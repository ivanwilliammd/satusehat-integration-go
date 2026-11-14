package parser

// ConsentParser extracts fields from a FHIR Consent resource map.
//
// The element list is derived from PayloadBuilderConsent in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ConsentParser struct {
	data map[string]any
}

// NewConsentParser creates a parser for the given Consent payload.
func NewConsentParser(data map[string]any) *ConsentParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ConsentParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ConsentParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ConsentParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ConsentParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ConsentParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ConsentParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *ConsentParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *ConsentParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *ConsentParser) GetEncounter() any {
	return p.data["encounter"]
}
