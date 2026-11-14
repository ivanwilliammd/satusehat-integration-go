package parser

// SubstanceNucleicAcidParser extracts fields from a FHIR SubstanceNucleicAcid resource map.
//
// The element list is derived from PayloadBuilderSubstanceNucleicAcid in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type SubstanceNucleicAcidParser struct {
	data map[string]any
}

// NewSubstanceNucleicAcidParser creates a parser for the given SubstanceNucleicAcid payload.
func NewSubstanceNucleicAcidParser(data map[string]any) *SubstanceNucleicAcidParser {
	if data == nil {
		data = map[string]any{}
	}
	return &SubstanceNucleicAcidParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *SubstanceNucleicAcidParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *SubstanceNucleicAcidParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *SubstanceNucleicAcidParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *SubstanceNucleicAcidParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *SubstanceNucleicAcidParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *SubstanceNucleicAcidParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *SubstanceNucleicAcidParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *SubstanceNucleicAcidParser) GetEncounter() any {
	return p.data["encounter"]
}
