package builder

import (
	"encoding/json"
)

type MedicinalProductPharmaceuticalBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMedicinalProductPharmaceuticalBuilder() *MedicinalProductPharmaceuticalBuilder {
	b := &MedicinalProductPharmaceuticalBuilder{ResourceType: "MedicinalProductPharmaceutical", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "MedicinalProductPharmaceutical"
	return b
}

func (b *MedicinalProductPharmaceuticalBuilder) SetId(id string) *MedicinalProductPharmaceuticalBuilder {
	b.Data["id"] = id
	return b
}

func (b *MedicinalProductPharmaceuticalBuilder) AddIdentifier(system, value string) *MedicinalProductPharmaceuticalBuilder {
	ids, _ := b.Data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.Data["identifier"] = ids
	return b
}

func (b *MedicinalProductPharmaceuticalBuilder) SetStatus(status string) *MedicinalProductPharmaceuticalBuilder {
	b.Data["status"] = status
	return b
}

func (b *MedicinalProductPharmaceuticalBuilder) SetRoute(system, code, display string) *MedicinalProductPharmaceuticalBuilder {
	b.Data["route"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *MedicinalProductPharmaceuticalBuilder) SetDoseForm(system, code, display string) *MedicinalProductPharmaceuticalBuilder {
	b.Data["doseForm"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *MedicinalProductPharmaceuticalBuilder) SetIngredient(system, code, display string) *MedicinalProductPharmaceuticalBuilder {
	b.Data["ingredient"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *MedicinalProductPharmaceuticalBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *MedicinalProductPharmaceuticalBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
