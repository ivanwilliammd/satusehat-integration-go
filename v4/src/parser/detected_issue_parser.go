package parser

// DetectedIssueParser extracts fields from a FHIR DetectedIssue resource map.
//
// The element list is derived from PayloadBuilderDetectedIssue in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type DetectedIssueParser struct {
	data map[string]any
}

// NewDetectedIssueParser creates a parser for the given DetectedIssue payload.
func NewDetectedIssueParser(data map[string]any) *DetectedIssueParser {
	if data == nil {
		data = map[string]any{}
	}
	return &DetectedIssueParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *DetectedIssueParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *DetectedIssueParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *DetectedIssueParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *DetectedIssueParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *DetectedIssueParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *DetectedIssueParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *DetectedIssueParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *DetectedIssueParser) GetEncounter() any {
	return p.data["encounter"]
}
