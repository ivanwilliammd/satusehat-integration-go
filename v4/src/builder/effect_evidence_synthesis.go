package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type EffectEvidenceSynthesisBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewEffectEvidenceSynthesisBuilder() *EffectEvidenceSynthesisBuilder {
    b := &EffectEvidenceSynthesisBuilder{ResourceType: "EffectEvidenceSynthesis", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "EffectEvidenceSynthesis"
    return b
}

func (b *EffectEvidenceSynthesisBuilder) setId(id string) *EffectEvidenceSynthesisBuilder {
    b.Data["id"] = id
    return b
}

func (b *EffectEvidenceSynthesisBuilder) addIdentifier(identifier *datatype.Identifier) *EffectEvidenceSynthesisBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *EffectEvidenceSynthesisBuilder) setStatus(status string) *EffectEvidenceSynthesisBuilder {
    b.Data["status"] = status
    return b
}

func (b *EffectEvidenceSynthesisBuilder) setSubject(reference string) *EffectEvidenceSynthesisBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *EffectEvidenceSynthesisBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *EffectEvidenceSynthesisBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
