package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type QuestionnaireResponseBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewQuestionnaireResponseBuilder() *QuestionnaireResponseBuilder {
    b := &QuestionnaireResponseBuilder{ResourceType: "QuestionnaireResponse", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "QuestionnaireResponse"
    return b
}

func (b *QuestionnaireResponseBuilder) setId(id string) *QuestionnaireResponseBuilder {
    b.Data["id"] = id
    return b
}

func (b *QuestionnaireResponseBuilder) addIdentifier(identifier *datatype.Identifier) *QuestionnaireResponseBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *QuestionnaireResponseBuilder) setStatus(status string) *QuestionnaireResponseBuilder {
    b.Data["status"] = status
    return b
}

func (b *QuestionnaireResponseBuilder) setSubject(reference string) *QuestionnaireResponseBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *QuestionnaireResponseBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *QuestionnaireResponseBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
