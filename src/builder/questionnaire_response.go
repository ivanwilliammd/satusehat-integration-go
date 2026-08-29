package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type QuestionnaireResponseBuilder struct {
	Data map[string]interface{}
}

func NewQuestionnaireResponseBuilder() *QuestionnaireResponseBuilder {
	return &QuestionnaireResponseBuilder{Data: make(map[string]interface{})}
}

func (b *QuestionnaireResponseBuilder) SetId(id string) *QuestionnaireResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *QuestionnaireResponseBuilder) SetStatus(status string) *QuestionnaireResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *QuestionnaireResponseBuilder) SetQuestionnaire(ref string) *QuestionnaireResponseBuilder {
	b.Data["questionnaire"] = map[string]string{"reference": ref}
	return b
}

func (b *QuestionnaireResponseBuilder) SetSubject(ref *datatype.Reference) *QuestionnaireResponseBuilder {
	b.Data["subject"] = ref.ToArray()
	return b
}

func (b *QuestionnaireResponseBuilder) SetEncounter(ref *datatype.Reference) *QuestionnaireResponseBuilder {
	b.Data["encounter"] = ref.ToArray()
	return b
}

func (b *QuestionnaireResponseBuilder) SetAuthored(dt string) *QuestionnaireResponseBuilder {
	b.Data["authored"] = dt
	return b
}

func (b *QuestionnaireResponseBuilder) SetAuthor(ref *datatype.Reference) *QuestionnaireResponseBuilder {
	b.Data["author"] = ref.ToArray()
	return b
}

func (b *QuestionnaireResponseBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "QuestionnaireResponse"
	return dt
}
