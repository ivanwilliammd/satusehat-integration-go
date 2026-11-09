package parser

// RiskAssessmentParser extracts fields from a FHIR RiskAssessment resource map.
type RiskAssessmentParser struct {
	riskAssessment map[string]any
}

// NewRiskAssessmentParser creates a new RiskAssessmentParser instance.
func NewRiskAssessmentParser(riskAssessment map[string]any) *RiskAssessmentParser {
	if riskAssessment == nil {
		riskAssessment = map[string]any{}
	}
	return &RiskAssessmentParser{riskAssessment: riskAssessment}
}

func (p *RiskAssessmentParser) GetStatus() string {
	return strVal(p.riskAssessment["status"])
}

func (p *RiskAssessmentParser) GetIdentifiers() any {
	return p.riskAssessment["identifier"]
}

func (p *RiskAssessmentParser) GetCode() any {
	return p.riskAssessment["code"]
}

func (p *RiskAssessmentParser) GetSubject() any {
	return p.riskAssessment["subject"]
}

func (p *RiskAssessmentParser) GetSubjectReference() string {
	return removePrefix(nestedStr(p.riskAssessment, "subject", "reference"), "Patient/")
}

func (p *RiskAssessmentParser) GetEncounter() any {
	return p.riskAssessment["encounter"]
}

func (p *RiskAssessmentParser) GetEncounterReference() string {
	return removePrefix(nestedStr(p.riskAssessment, "encounter", "reference"), "Encounter/")
}

func (p *RiskAssessmentParser) GetOccurrenceDateTime() string {
	return strVal(p.riskAssessment["occurrenceDateTime"])
}

func (p *RiskAssessmentParser) GetCondition() any {
	return p.riskAssessment["condition"]
}

func (p *RiskAssessmentParser) GetConditionReference() string {
	return removePrefix(nestedStr(p.riskAssessment, "condition", "reference"), "Condition/")
}

func (p *RiskAssessmentParser) GetPerformer() any {
	return p.riskAssessment["performer"]
}

func (p *RiskAssessmentParser) GetPerformerReference() string {
	return removePrefix(nestedStr(p.riskAssessment, "performer", "reference"), "Practitioner/")
}

func (p *RiskAssessmentParser) GetReasonReferences() any {
	return p.riskAssessment["reasonReference"]
}

func (p *RiskAssessmentParser) GetBasis() any {
	return p.riskAssessment["basis"]
}

func (p *RiskAssessmentParser) GetPredictions() any {
	return p.riskAssessment["prediction"]
}

func (p *RiskAssessmentParser) GetMitigation() string {
	return strVal(p.riskAssessment["mitigation"])
}

func (p *RiskAssessmentParser) GetNotes() any {
	return p.riskAssessment["note"]
}
