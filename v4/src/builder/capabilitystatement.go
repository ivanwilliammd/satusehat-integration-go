package builder

import (
	"encoding/json"
)

type CapabilityStatementBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewCapabilityStatementBuilder() *CapabilityStatementBuilder {
	b := &CapabilityStatementBuilder{ResourceType: "CapabilityStatement", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "CapabilityStatement"
	return b
}

func (b *CapabilityStatementBuilder) SetId(id string) *CapabilityStatementBuilder {
	b.Data["id"] = id
	return b
}

func (b *CapabilityStatementBuilder) AddIdentifier(system, value string) *CapabilityStatementBuilder {
	ids, _ := b.Data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.Data["identifier"] = ids
	return b
}

func (b *CapabilityStatementBuilder) SetStatus(status string) *CapabilityStatementBuilder {
	b.Data["status"] = status
	return b
}

func (b *CapabilityStatementBuilder) SetDate(value string) *CapabilityStatementBuilder {
	b.Data["date"] = value
	return b
}

func (b *CapabilityStatementBuilder) SetKind(system, code, display string) *CapabilityStatementBuilder {
	b.Data["kind"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *CapabilityStatementBuilder) SetFhirVersion(system, code, display string) *CapabilityStatementBuilder {
	b.Data["fhirVersion"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *CapabilityStatementBuilder) SetFormat(system, code, display string) *CapabilityStatementBuilder {
	b.Data["format"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *CapabilityStatementBuilder) SetRest(system, code, display string) *CapabilityStatementBuilder {
	b.Data["rest"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *CapabilityStatementBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *CapabilityStatementBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
