package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MedicinalProductIndicationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMedicinalProductIndicationBuilder() *MedicinalProductIndicationBuilder {
	b := &MedicinalProductIndicationBuilder{ResourceType: "MedicinalProductIndication", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "MedicinalProductIndication"
	return b
}

func (b *MedicinalProductIndicationBuilder) SetId(id string) *MedicinalProductIndicationBuilder {
	b.Data["id"] = id
	return b
}

func (b *MedicinalProductIndicationBuilder) AddIdentifier(system, value string) *MedicinalProductIndicationBuilder {
	ids, _ := b.Data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.Data["identifier"] = ids
	return b
}

func (b *MedicinalProductIndicationBuilder) SetStatus(status string) *MedicinalProductIndicationBuilder {
	b.Data["status"] = status
	return b
}

func (b *MedicinalProductIndicationBuilder) SetSubject(reference string, display ...string) *MedicinalProductIndicationBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *MedicinalProductIndicationBuilder) SetDisease(system, code, display string) *MedicinalProductIndicationBuilder {
	b.Data["disease"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *MedicinalProductIndicationBuilder) SetComorbidity(reference string, display ...string) *MedicinalProductIndicationBuilder {
	ref := map[string]interface{}{"reference": reference}
	if len(display) > 0 && display[0] != "" {
		ref["display"] = display[0]
	}
	b.Data["comorbidity"] = ref
	return b
}

func (b *MedicinalProductIndicationBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *MedicinalProductIndicationBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
