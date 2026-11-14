package parser

// EvidenceVariableParser extracts fields from a FHIR EvidenceVariable resource map.
//
// The element list is derived from PayloadBuilderEvidenceVariable in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type EvidenceVariableParser struct {
	data map[string]any
}

// NewEvidenceVariableParser creates a parser for the given EvidenceVariable payload.
func NewEvidenceVariableParser(data map[string]any) *EvidenceVariableParser {
	if data == nil {
		data = map[string]any{}
	}
	return &EvidenceVariableParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *EvidenceVariableParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *EvidenceVariableParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *EvidenceVariableParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *EvidenceVariableParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *EvidenceVariableParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *EvidenceVariableParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *EvidenceVariableParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *EvidenceVariableParser) GetEncounter() any {
	return p.data["encounter"]
}
