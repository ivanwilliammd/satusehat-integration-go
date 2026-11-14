package parser

// OperationOutcomeParser extracts fields from a FHIR OperationOutcome resource map.
//
// The element list is derived from PayloadBuilderOperationOutcome in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type OperationOutcomeParser struct {
	data map[string]any
}

// NewOperationOutcomeParser creates a parser for the given OperationOutcome payload.
func NewOperationOutcomeParser(data map[string]any) *OperationOutcomeParser {
	if data == nil {
		data = map[string]any{}
	}
	return &OperationOutcomeParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *OperationOutcomeParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *OperationOutcomeParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *OperationOutcomeParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *OperationOutcomeParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *OperationOutcomeParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *OperationOutcomeParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *OperationOutcomeParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *OperationOutcomeParser) GetEncounter() any {
	return p.data["encounter"]
}
