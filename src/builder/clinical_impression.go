package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type ClinicalImpressionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewClinicalImpressionBuilder() *ClinicalImpressionBuilder {
    b := &ClinicalImpressionBuilder{ResourceType: "ClinicalImpression", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ClinicalImpression"
    return b
}

func (b *ClinicalImpressionBuilder) setId(id string) *ClinicalImpressionBuilder {
    b.Data["id"] = id
    return b
}

func (b *ClinicalImpressionBuilder) addIdentifier(identifier *datatype.Identifier) *ClinicalImpressionBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ClinicalImpressionBuilder) setStatus(status string) *ClinicalImpressionBuilder {
    b.Data["status"] = status
    return b
}

func (b *ClinicalImpressionBuilder) setSubject(reference string) *ClinicalImpressionBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *ClinicalImpressionBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ClinicalImpressionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
