package parser

// QuestionnaireParser extracts fields from a FHIR Questionnaire resource map.
//
// The element list is derived from PayloadBuilderQuestionnaire in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type QuestionnaireParser struct {
	data map[string]any
}

// NewQuestionnaireParser creates a parser for the given Questionnaire payload.
func NewQuestionnaireParser(data map[string]any) *QuestionnaireParser {
	if data == nil {
		data = map[string]any{}
	}
	return &QuestionnaireParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *QuestionnaireParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *QuestionnaireParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *QuestionnaireParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *QuestionnaireParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *QuestionnaireParser) GetStatus() any {
	return p.data["status"]
}

// GetCode returns the code element.
func (p *QuestionnaireParser) GetCode() any {
	return p.data["code"]
}

// GetSubject returns the subject element.
func (p *QuestionnaireParser) GetSubject() any {
	return p.data["subject"]
}

// GetEncounter returns the encounter element.
func (p *QuestionnaireParser) GetEncounter() any {
	return p.data["encounter"]
}
