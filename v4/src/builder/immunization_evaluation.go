package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ImmunizationEvaluationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewImmunizationEvaluationBuilder() *ImmunizationEvaluationBuilder {
	b := &ImmunizationEvaluationBuilder{ResourceType: "ImmunizationEvaluation", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ImmunizationEvaluation"
	return b
}

func (b *ImmunizationEvaluationBuilder) SetId(id string) *ImmunizationEvaluationBuilder {
	b.Data["id"] = id
	return b
}

func (b *ImmunizationEvaluationBuilder) AddIdentifier(identifier *datatype.Identifier) *ImmunizationEvaluationBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ImmunizationEvaluationBuilder) SetStatus(status string) *ImmunizationEvaluationBuilder {
	b.Data["status"] = status
	return b
}

func (b *ImmunizationEvaluationBuilder) SetSubject(reference string) *ImmunizationEvaluationBuilder {
	b.Data["subject"] = datatype.Reference{}.ToArray()
	return b
}

func (b *ImmunizationEvaluationBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ImmunizationEvaluationBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
