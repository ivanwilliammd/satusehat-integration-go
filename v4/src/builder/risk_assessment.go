package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type RiskAssessmentBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewRiskAssessmentBuilder() *RiskAssessmentBuilder {
    b := &RiskAssessmentBuilder{ResourceType: "RiskAssessment", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "RiskAssessment"
    return b
}

func (b *RiskAssessmentBuilder) setId(id string) *RiskAssessmentBuilder {
    b.Data["id"] = id
    return b
}

func (b *RiskAssessmentBuilder) addIdentifier(identifier *datatype.Identifier) *RiskAssessmentBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *RiskAssessmentBuilder) setStatus(status string) *RiskAssessmentBuilder {
    b.Data["status"] = status
    return b
}

func (b *RiskAssessmentBuilder) setSubject(reference string) *RiskAssessmentBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *RiskAssessmentBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *RiskAssessmentBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
