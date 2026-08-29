package builder

type RiskAssessmentBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewRiskAssessmentBuilder() *RiskAssessmentBuilder {
	b := &RiskAssessmentBuilder{ResourceType: "RiskAssessment", Data: make(map[string]interface{})}
	return b
}

func (b *RiskAssessmentBuilder) SetId(id string) *RiskAssessmentBuilder { b.Data["id"] = id; return b }
func (b *RiskAssessmentBuilder) SetStatus(status string) *RiskAssessmentBuilder { b.Data["status"] = status; return b }
func (b *RiskAssessmentBuilder) SetSubject(subjectRef string) *RiskAssessmentBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *RiskAssessmentBuilder) SetEncounter(encRef string) *RiskAssessmentBuilder { b.Data["encounter"] = map[string]interface{}{"reference": encRef}; return b }
func (b *RiskAssessmentBuilder) SetOccurrence(occurrence string) *RiskAssessmentBuilder { b.Data["occurrenceDateTime"] = occurrence; return b }
func (b *RiskAssessmentBuilder) SetPrediction(outcome string) *RiskAssessmentBuilder {
	b.Data["prediction"] = []interface{}{map[string]interface{}{"outcomeString": outcome}}
	return b
}
func (b *RiskAssessmentBuilder) SetNote(note string) *RiskAssessmentBuilder {
	b.Data["note"] = []interface{}{map[string]interface{}{"text": note}}; return b
}
func (b *RiskAssessmentBuilder) Build() map[string]interface{} { return b.Data }
