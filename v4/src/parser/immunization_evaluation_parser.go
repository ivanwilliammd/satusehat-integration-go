package parser

// ImmunizationEvaluationParser extracts fields from a FHIR ImmunizationEvaluation resource map.
//
// The element list is derived from PayloadBuilderImmunizationEvaluation in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type ImmunizationEvaluationParser struct {
	data map[string]any
}

// NewImmunizationEvaluationParser creates a parser for the given ImmunizationEvaluation payload.
func NewImmunizationEvaluationParser(data map[string]any) *ImmunizationEvaluationParser {
	if data == nil {
		data = map[string]any{}
	}
	return &ImmunizationEvaluationParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *ImmunizationEvaluationParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *ImmunizationEvaluationParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *ImmunizationEvaluationParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *ImmunizationEvaluationParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *ImmunizationEvaluationParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *ImmunizationEvaluationParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *ImmunizationEvaluationParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *ImmunizationEvaluationParser) GetEncounter() any {
	return p.data["encounter"]
}
