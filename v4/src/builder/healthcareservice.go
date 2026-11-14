package builder

import (
	"encoding/json"
)

type HealthcareServiceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewHealthcareServiceBuilder() *HealthcareServiceBuilder {
	b := &HealthcareServiceBuilder{ResourceType: "HealthcareService", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "HealthcareService"
	return b
}

func (b *HealthcareServiceBuilder) SetId(id string) *HealthcareServiceBuilder {
	b.Data["id"] = id
	return b
}

func (b *HealthcareServiceBuilder) AddIdentifier(system, value string) *HealthcareServiceBuilder {
	ids, _ := b.Data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.Data["identifier"] = ids
	return b
}

func (b *HealthcareServiceBuilder) SetStatus(status string) *HealthcareServiceBuilder {
	b.Data["status"] = status
	return b
}

func (b *HealthcareServiceBuilder) SetName(value string) *HealthcareServiceBuilder {
	b.Data["name"] = value
	return b
}

func (b *HealthcareServiceBuilder) SetType(system, code, display string) *HealthcareServiceBuilder {
	b.Data["type"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *HealthcareServiceBuilder) SetLocation(reference string, display ...string) *HealthcareServiceBuilder {
	ref := map[string]interface{}{"reference": reference}
	if len(display) > 0 && display[0] != "" {
		ref["display"] = display[0]
	}
	b.Data["location"] = ref
	return b
}

func (b *HealthcareServiceBuilder) SetProvidedBy(reference string, display ...string) *HealthcareServiceBuilder {
	ref := map[string]interface{}{"reference": reference}
	if len(display) > 0 && display[0] != "" {
		ref["display"] = display[0]
	}
	b.Data["providedBy"] = ref
	return b
}

func (b *HealthcareServiceBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *HealthcareServiceBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
