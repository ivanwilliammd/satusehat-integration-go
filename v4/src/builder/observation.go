package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/terminology"
)

type ObservationBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewObservationBuilder() *ObservationBuilder {
    b := &ObservationBuilder{ResourceType: "Observation", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Observation"
    return b
}

func (b *ObservationBuilder) setId(id string) *ObservationBuilder {
    b.Data["id"] = id
    return b
}

func (b *ObservationBuilder) addIdentifier(identifier *datatype.Identifier) *ObservationBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *ObservationBuilder) setStatus(status string) *ObservationBuilder {
    b.Data["status"] = status
    return b
}

func (b *ObservationBuilder) setSubject(reference string) *ObservationBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

// setCode supports "System:Code" castable notation (e.g. "LOINC:2951-2", "ICD10:A00").
func (b *ObservationBuilder) setCode(codeOrConcept interface{}) *ObservationBuilder {
    if s, ok := codeOrConcept.(string); ok {
        b.Data["code"] = terminology.Resolve(s)
    } else {
        b.Data["code"] = codeOrConcept
    }
    return b
}

// addCategory supports castable strings too.
func (b *ObservationBuilder) addCategory(category interface{}) *ObservationBuilder {
    if b.Data["category"] == nil {
        b.Data["category"] = []interface{}{}
    }
    if s, ok := category.(string); ok {
        b.Data["category"] = append(b.Data["category"].([]interface{}), terminology.Resolve(s))
    } else {
        b.Data["category"] = append(b.Data["category"].([]interface{}), category)
    }
    return b
}

func (b *ObservationBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *ObservationBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
