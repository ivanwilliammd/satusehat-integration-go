package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type RiskEvidenceSynthesisBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewRiskEvidenceSynthesisBuilder() *RiskEvidenceSynthesisBuilder {
    b := &RiskEvidenceSynthesisBuilder{ResourceType: "RiskEvidenceSynthesis", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "RiskEvidenceSynthesis"
    return b
}

func (b *RiskEvidenceSynthesisBuilder) setId(id string) *RiskEvidenceSynthesisBuilder {
    b.Data["id"] = id
    return b
}

func (b *RiskEvidenceSynthesisBuilder) addIdentifier(identifier *datatype.Identifier) *RiskEvidenceSynthesisBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *RiskEvidenceSynthesisBuilder) setStatus(status string) *RiskEvidenceSynthesisBuilder {
    b.Data["status"] = status
    return b
}

func (b *RiskEvidenceSynthesisBuilder) setSubject(reference string) *RiskEvidenceSynthesisBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *RiskEvidenceSynthesisBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *RiskEvidenceSynthesisBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
