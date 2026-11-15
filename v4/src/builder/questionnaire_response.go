package builder

import (
	"encoding/json"
	"fmt"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type QuestionnaireResponseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
	err          error
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

// SetStatus sets status. The canonical builder accepts only the five values FHIR allows for a response and
// throws otherwise; the port records the same refusal on Err() and writes nothing.
func (b *QuestionnaireResponseBuilder) SetStatus(status string) *QuestionnaireResponseBuilder {
	sah := map[string]bool{"in-progress": true, "completed": true, "amended": true, "entered-in-error": true, "stopped": true}
	if !sah[status] {
		b.err = fmt.Errorf("status %q bukan nilai QuestionnaireResponse yang sah", status)
		return b
	}
	b.Data["status"] = status
	return b
}

// Err returns the first validation error recorded by a setter, if any.
func (b *QuestionnaireResponseBuilder) Err() error { return b.err }

func (b *QuestionnaireResponseBuilder) SetSubject(reference string, display ...string) *QuestionnaireResponseBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

// Build returns the payload. The canonical builder fills status with "completed" when the caller set none,
// because FHIR requires the element; the port does the same.
func (b *QuestionnaireResponseBuilder) Build() map[string]interface{} {
	if b.err != nil {
		return b.Data
	}
	if _, ada := b.Data["status"]; !ada {
		b.Data["status"] = "completed"
	}
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
