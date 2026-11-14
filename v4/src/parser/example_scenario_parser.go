package parser

// ExampleScenarioParser extracts fields from a FHIR ExampleScenario resource map.
//
// The element list is derived from PayloadBuilderExampleScenario in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ExampleScenarioParser struct {
	data map[string]any
}

// NewExampleScenarioParser creates a parser for the given ExampleScenario payload.
func NewExampleScenarioParser(data map[string]any) *ExampleScenarioParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ExampleScenarioParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ExampleScenarioParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ExampleScenarioParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ExampleScenarioParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ExampleScenarioParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ExampleScenarioParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *ExampleScenarioParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *ExampleScenarioParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *ExampleScenarioParser) GetEncounter() any {
	return p.data["encounter"]
}
