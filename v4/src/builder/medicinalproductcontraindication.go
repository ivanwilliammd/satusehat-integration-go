package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MedicinalProductContraindicationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMedicinalProductContraindicationBuilder() *MedicinalProductContraindicationBuilder {
	b := &MedicinalProductContraindicationBuilder{ResourceType: "MedicinalProductContraindication", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "MedicinalProductContraindication"
	return b
}

func (b *MedicinalProductContraindicationBuilder) SetId(id string) *MedicinalProductContraindicationBuilder {
	b.Data["id"] = id
	return b
}

func (b *MedicinalProductContraindicationBuilder) AddIdentifier(system, value string) *MedicinalProductContraindicationBuilder {
	ids, _ := b.Data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.Data["identifier"] = ids
	return b
}

func (b *MedicinalProductContraindicationBuilder) SetStatus(status string) *MedicinalProductContraindicationBuilder {
	b.Data["status"] = status
	return b
}

func (b *MedicinalProductContraindicationBuilder) SetSubject(reference string, display ...string) *MedicinalProductContraindicationBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *MedicinalProductContraindicationBuilder) SetDisease(system, code, display string) *MedicinalProductContraindicationBuilder {
	b.Data["disease"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *MedicinalProductContraindicationBuilder) SetComorbidity(reference string, display ...string) *MedicinalProductContraindicationBuilder {
	ref := map[string]interface{}{"reference": reference}
	if len(display) > 0 && display[0] != "" {
		ref["display"] = display[0]
	}
	b.Data["comorbidity"] = ref
	return b
}

func (b *MedicinalProductContraindicationBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *MedicinalProductContraindicationBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
