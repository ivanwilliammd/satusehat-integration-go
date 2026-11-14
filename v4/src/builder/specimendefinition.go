package builder

import (
	"encoding/json"
)

type SpecimenDefinitionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSpecimenDefinitionBuilder() *SpecimenDefinitionBuilder {
	b := &SpecimenDefinitionBuilder{ResourceType: "SpecimenDefinition", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "SpecimenDefinition"
	return b
}

func (b *SpecimenDefinitionBuilder) SetId(id string) *SpecimenDefinitionBuilder {
	b.Data["id"] = id
	return b
}

func (b *SpecimenDefinitionBuilder) AddIdentifier(system, value string) *SpecimenDefinitionBuilder {
	ids, _ := b.Data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.Data["identifier"] = ids
	return b
}

func (b *SpecimenDefinitionBuilder) SetStatus(status string) *SpecimenDefinitionBuilder {
	b.Data["status"] = status
	return b
}

func (b *SpecimenDefinitionBuilder) SetType(system, code, display string) *SpecimenDefinitionBuilder {
	b.Data["type"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *SpecimenDefinitionBuilder) SetCollection(value string) *SpecimenDefinitionBuilder {
	b.Data["collection"] = value
	return b
}

func (b *SpecimenDefinitionBuilder) SetPreparation(system, code, display string) *SpecimenDefinitionBuilder {
	b.Data["preparation"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *SpecimenDefinitionBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *SpecimenDefinitionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
