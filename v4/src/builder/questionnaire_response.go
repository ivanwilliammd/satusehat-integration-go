package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type QuestionnaireResponseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewQuestionnaireResponseBuilder() *QuestionnaireResponseBuilder {
	b := &QuestionnaireResponseBuilder{ResourceType: "QuestionnaireResponse", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "QuestionnaireResponse"
	return b
}

func (b *QuestionnaireResponseBuilder) SetId(id string) *QuestionnaireResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *QuestionnaireResponseBuilder) AddIdentifier(identifier *datatype.Identifier) *QuestionnaireResponseBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *QuestionnaireResponseBuilder) SetStatus(status string) *QuestionnaireResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *QuestionnaireResponseBuilder) SetSubject(reference string, display ...string) *QuestionnaireResponseBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *QuestionnaireResponseBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *QuestionnaireResponseBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetEncounter sets encounter from a Reference, as the PHP builder does.
func (b *QuestionnaireResponseBuilder) SetEncounter(encounter *datatype.Reference) *QuestionnaireResponseBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

// AddItem adds an answered question, as the PHP builder does.
func (b *QuestionnaireResponseBuilder) AddItem(linkId, text string, answer *datatype.CodeableConcept) *QuestionnaireResponseBuilder {
	item := map[string]interface{}{"linkId": linkId}
	if text != "" {
		item["text"] = text
	}
	item["answer"] = []interface{}{answer.ToArray()}
	daftar, _ := b.Data["item"].([]interface{})
	b.Data["item"] = append(daftar, item)
	return b
}

// SetQuestionnaire sets questionnaire.
func (b *QuestionnaireResponseBuilder) SetQuestionnaire(questionnaire string) *QuestionnaireResponseBuilder {
	b.Data["questionnaire"] = questionnaire
	return b
}

// SetAuthored sets authored.
func (b *QuestionnaireResponseBuilder) SetAuthored(dateTime string) *QuestionnaireResponseBuilder {
	b.Data["authored"] = dateTime
	return b
}

// SetAuthor sets author from a reference.
func (b *QuestionnaireResponseBuilder) SetAuthor(author *datatype.Reference) *QuestionnaireResponseBuilder {
	b.Data["author"] = author.ToArray()
	return b
}

// SetSource sets source from a reference.
func (b *QuestionnaireResponseBuilder) SetSource(source *datatype.Reference) *QuestionnaireResponseBuilder {
	b.Data["source"] = source.ToArray()
	return b
}
