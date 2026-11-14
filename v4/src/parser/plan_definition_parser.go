package parser

// PlanDefinitionParser extracts fields from a FHIR PlanDefinition resource map.
//
// The element list is derived from PayloadBuilderPlanDefinition in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type PlanDefinitionParser struct {
	data map[string]any
}

// NewPlanDefinitionParser creates a parser for the given PlanDefinition payload.
func NewPlanDefinitionParser(data map[string]any) *PlanDefinitionParser {
	if data == nil {
		data = map[string]any{}
	}
	return &PlanDefinitionParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *PlanDefinitionParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *PlanDefinitionParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *PlanDefinitionParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *PlanDefinitionParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *PlanDefinitionParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *PlanDefinitionParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *PlanDefinitionParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *PlanDefinitionParser) GetEncounter() any {
	return p.data["encounter"]
}
