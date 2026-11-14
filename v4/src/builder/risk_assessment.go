package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type RiskAssessmentBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewRiskAssessmentBuilder() *RiskAssessmentBuilder {
	b := &RiskAssessmentBuilder{ResourceType: "RiskAssessment", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "RiskAssessment"
	return b
}

func (b *RiskAssessmentBuilder) SetId(id string) *RiskAssessmentBuilder {
	b.Data["id"] = id
	return b
}

func (b *RiskAssessmentBuilder) AddIdentifier(identifier *datatype.Identifier) *RiskAssessmentBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *RiskAssessmentBuilder) SetStatus(status string) *RiskAssessmentBuilder {
	b.Data["status"] = status
	return b
}

func (b *RiskAssessmentBuilder) SetSubject(reference string, display ...string) *RiskAssessmentBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *RiskAssessmentBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *RiskAssessmentBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code from a CodeableConcept, as the PHP builder does.
func (b *RiskAssessmentBuilder) SetCode(code *datatype.CodeableConcept) *RiskAssessmentBuilder {
	b.Data["code"] = code.ToArray()
	return b
}

// SetEncounter sets encounter from a Reference, as the PHP builder does.
func (b *RiskAssessmentBuilder) SetEncounter(encounter *datatype.Reference) *RiskAssessmentBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}
