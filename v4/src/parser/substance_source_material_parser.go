package parser

// SubstanceSourceMaterialParser extracts fields from a FHIR SubstanceSourceMaterial resource map.
//
// The element list is derived from PayloadBuilderSubstanceSourceMaterial in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type SubstanceSourceMaterialParser struct {
	data map[string]any
}

// NewSubstanceSourceMaterialParser creates a parser for the given SubstanceSourceMaterial payload.
func NewSubstanceSourceMaterialParser(data map[string]any) *SubstanceSourceMaterialParser {
	if data == nil {
		data = map[string]any{}
	}
	return &SubstanceSourceMaterialParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *SubstanceSourceMaterialParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *SubstanceSourceMaterialParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *SubstanceSourceMaterialParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *SubstanceSourceMaterialParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *SubstanceSourceMaterialParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *SubstanceSourceMaterialParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *SubstanceSourceMaterialParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *SubstanceSourceMaterialParser) GetEncounter() any {
	return p.data["encounter"]
}
