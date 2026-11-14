package builder

import (
	"encoding/json"
)

type MedicinalProductInteractionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMedicinalProductInteractionBuilder() *MedicinalProductInteractionBuilder {
	b := &MedicinalProductInteractionBuilder{ResourceType: "MedicinalProductInteraction", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "MedicinalProductInteraction"
	return b
}

func (b *MedicinalProductInteractionBuilder) SetId(id string) *MedicinalProductInteractionBuilder {
	b.Data["id"] = id
	return b
}

func (b *MedicinalProductInteractionBuilder) AddIdentifier(system, value string) *MedicinalProductInteractionBuilder {
	ids, _ := b.Data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.Data["identifier"] = ids
	return b
}

func (b *MedicinalProductInteractionBuilder) SetStatus(status string) *MedicinalProductInteractionBuilder {
	b.Data["status"] = status
	return b
}

func (b *MedicinalProductInteractionBuilder) SetSubject(reference string, display ...string) *MedicinalProductInteractionBuilder {
	ref := map[string]interface{}{"reference": reference}
	if len(display) > 0 && display[0] != "" {
		ref["display"] = display[0]
	}
	b.Data["subject"] = ref
	return b
}

func (b *MedicinalProductInteractionBuilder) SetDescription(value string) *MedicinalProductInteractionBuilder {
	b.Data["description"] = value
	return b
}

func (b *MedicinalProductInteractionBuilder) SetInteractant(system, code, display string) *MedicinalProductInteractionBuilder {
	b.Data["interactant"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *MedicinalProductInteractionBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *MedicinalProductInteractionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
