package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MedicinalProductUndesirableEffectBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMedicinalProductUndesirableEffectBuilder() *MedicinalProductUndesirableEffectBuilder {
	b := &MedicinalProductUndesirableEffectBuilder{ResourceType: "MedicinalProductUndesirableEffect", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "MedicinalProductUndesirableEffect"
	return b
}

func (b *MedicinalProductUndesirableEffectBuilder) SetId(id string) *MedicinalProductUndesirableEffectBuilder {
	b.Data["id"] = id
	return b
}

func (b *MedicinalProductUndesirableEffectBuilder) AddIdentifier(system, value string) *MedicinalProductUndesirableEffectBuilder {
	ids, _ := b.Data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.Data["identifier"] = ids
	return b
}

func (b *MedicinalProductUndesirableEffectBuilder) SetStatus(status string) *MedicinalProductUndesirableEffectBuilder {
	b.Data["status"] = status
	return b
}

func (b *MedicinalProductUndesirableEffectBuilder) SetSubject(reference string, display ...string) *MedicinalProductUndesirableEffectBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *MedicinalProductUndesirableEffectBuilder) SetSymptomConditionEffect(system, code, display string) *MedicinalProductUndesirableEffectBuilder {
	b.Data["symptomConditionEffect"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *MedicinalProductUndesirableEffectBuilder) SetClassification(system, code, display string) *MedicinalProductUndesirableEffectBuilder {
	b.Data["classification"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *MedicinalProductUndesirableEffectBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *MedicinalProductUndesirableEffectBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
