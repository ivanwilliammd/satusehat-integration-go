package parser

// TestScriptParser extracts fields from a FHIR TestScript resource map.
//
// The element list is derived from PayloadBuilderTestScript in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type TestScriptParser struct {
	data map[string]any
}

// NewTestScriptParser creates a parser for the given TestScript payload.
func NewTestScriptParser(data map[string]any) *TestScriptParser {
	if data == nil {
		data = map[string]any{}
	}
	return &TestScriptParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *TestScriptParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *TestScriptParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *TestScriptParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *TestScriptParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *TestScriptParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *TestScriptParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *TestScriptParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *TestScriptParser) GetEncounter() any {
	return p.data["encounter"]
}
