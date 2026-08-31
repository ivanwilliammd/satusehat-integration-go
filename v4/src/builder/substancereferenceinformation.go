package builder

import (
    "encoding/json"
)

type SubstanceReferenceInformationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewSubstanceReferenceInformationBuilder() *SubstanceReferenceInformationBuilder {
    b := &SubstanceReferenceInformationBuilder{ResourceType: "SubstanceReferenceInformation", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "SubstanceReferenceInformation"
    return b
}

func (b *SubstanceReferenceInformationBuilder) setId(id string) *SubstanceReferenceInformationBuilder {
    b.Data["id"] = id
    return b
}

func (b *SubstanceReferenceInformationBuilder) addIdentifier(system, value string) *SubstanceReferenceInformationBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *SubstanceReferenceInformationBuilder) setStatus(status string) *SubstanceReferenceInformationBuilder {
    b.Data["status"] = status
    return b
}

func (b *SubstanceReferenceInformationBuilder) setComment(value string) *SubstanceReferenceInformationBuilder {
    b.Data["comment"] = value
    return b
}

func (b *SubstanceReferenceInformationBuilder) setGene(system, code, display string) *SubstanceReferenceInformationBuilder {
    b.Data["gene"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *SubstanceReferenceInformationBuilder) setGeneElement(system, code, display string) *SubstanceReferenceInformationBuilder {
    b.Data["geneElement"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *SubstanceReferenceInformationBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *SubstanceReferenceInformationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
