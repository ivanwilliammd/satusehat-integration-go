package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type QuestionnaireBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewQuestionnaireBuilder() *QuestionnaireBuilder {
    b := &QuestionnaireBuilder{ResourceType: "Questionnaire", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Questionnaire"
    return b
}

func (b *QuestionnaireBuilder) setId(id string) *QuestionnaireBuilder {
    b.Data["id"] = id
    return b
}

func (b *QuestionnaireBuilder) addIdentifier(identifier *datatype.Identifier) *QuestionnaireBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *QuestionnaireBuilder) setStatus(status string) *QuestionnaireBuilder {
    b.Data["status"] = status
    return b
}

func (b *QuestionnaireBuilder) setSubject(reference string) *QuestionnaireBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *QuestionnaireBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *QuestionnaireBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
