package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type QuestionnaireBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewQuestionnaireBuilder() *QuestionnaireBuilder {
	b := &QuestionnaireBuilder{ResourceType: "Questionnaire", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Questionnaire"
	return b
}

func (b *QuestionnaireBuilder) SetId(id string) *QuestionnaireBuilder {
	b.Data["id"] = id
	return b
}

func (b *QuestionnaireBuilder) AddIdentifier(identifier *datatype.Identifier) *QuestionnaireBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *QuestionnaireBuilder) SetStatus(status string) *QuestionnaireBuilder {
	b.Data["status"] = status
	return b
}

func (b *QuestionnaireBuilder) SetSubject(reference string, display ...string) *QuestionnaireBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *QuestionnaireBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *QuestionnaireBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *QuestionnaireBuilder) SetCode(system, code, display string) *QuestionnaireBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *QuestionnaireBuilder) SetEncounter(reference string) *QuestionnaireBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
