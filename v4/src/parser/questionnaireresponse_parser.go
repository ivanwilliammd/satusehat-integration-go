package parser

// QuestionnaireResponseParser extracts fields from a FHIR QuestionnaireResponse resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/QuestionnaireResponse.json, copied unchanged to
// src/parser/testdata/QuestionnaireResponse.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type QuestionnaireResponseParser struct {
	payload map[string]any
}

// NewQuestionnaireResponseParser creates a new QuestionnaireResponseParser instance.
func NewQuestionnaireResponseParser(payload map[string]any) *QuestionnaireResponseParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &QuestionnaireResponseParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *QuestionnaireResponseParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *QuestionnaireResponseParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetQuestionnaire returns the questionnaire element unchanged.
func (p *QuestionnaireResponseParser) GetQuestionnaire() any {
	return p.payload["questionnaire"]
}

// GetStatus returns the status element unchanged.
func (p *QuestionnaireResponseParser) GetStatus() any {
	return p.payload["status"]
}

// GetSubject returns the subject element unchanged.
func (p *QuestionnaireResponseParser) GetSubject() any {
	return p.payload["subject"]
}

// GetEncounter returns the encounter element unchanged.
func (p *QuestionnaireResponseParser) GetEncounter() any {
	return p.payload["encounter"]
}

// GetAuthored returns the authored element unchanged.
func (p *QuestionnaireResponseParser) GetAuthored() any {
	return p.payload["authored"]
}

// GetAuthor returns the author element unchanged.
func (p *QuestionnaireResponseParser) GetAuthor() any {
	return p.payload["author"]
}

// GetSource returns the source element unchanged.
func (p *QuestionnaireResponseParser) GetSource() any {
	return p.payload["source"]
}

// GetItem returns the item element unchanged.
func (p *QuestionnaireResponseParser) GetItem() any {
	return p.payload["item"]
}
