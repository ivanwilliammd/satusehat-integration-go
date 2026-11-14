package parser

// EffectEvidenceSynthesisParser extracts fields from a FHIR EffectEvidenceSynthesis resource map.
//
// The element list is derived from PayloadBuilderEffectEvidenceSynthesis in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type EffectEvidenceSynthesisParser struct {
	data map[string]any
}

// NewEffectEvidenceSynthesisParser creates a parser for the given EffectEvidenceSynthesis payload.
func NewEffectEvidenceSynthesisParser(data map[string]any) *EffectEvidenceSynthesisParser {
	if data == nil {
		data = map[string]any{}
	}
	return &EffectEvidenceSynthesisParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *EffectEvidenceSynthesisParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *EffectEvidenceSynthesisParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *EffectEvidenceSynthesisParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *EffectEvidenceSynthesisParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *EffectEvidenceSynthesisParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *EffectEvidenceSynthesisParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *EffectEvidenceSynthesisParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *EffectEvidenceSynthesisParser) GetEncounter() any {
	return p.data["encounter"]
}
