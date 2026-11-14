package parser

// TestReportParser extracts fields from a FHIR TestReport resource map.
//
// The element list is derived from PayloadBuilderTestReport in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type TestReportParser struct {
	data map[string]any
}

// NewTestReportParser creates a parser for the given TestReport payload.
func NewTestReportParser(data map[string]any) *TestReportParser {
	if data == nil {
		data = map[string]any{}
	}
	return &TestReportParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *TestReportParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *TestReportParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *TestReportParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *TestReportParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *TestReportParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *TestReportParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *TestReportParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *TestReportParser) GetEncounter() any {
	return p.data["encounter"]
}
