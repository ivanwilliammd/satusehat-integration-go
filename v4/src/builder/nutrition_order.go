package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type NutritionOrderBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewNutritionOrderBuilder() *NutritionOrderBuilder {
	b := &NutritionOrderBuilder{ResourceType: "NutritionOrder", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "NutritionOrder"
	return b
}

func (b *NutritionOrderBuilder) SetId(id string) *NutritionOrderBuilder {
	b.Data["id"] = id
	return b
}

func (b *NutritionOrderBuilder) AddIdentifier(identifier *datatype.Identifier) *NutritionOrderBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *NutritionOrderBuilder) SetStatus(status string) *NutritionOrderBuilder {
	b.Data["status"] = status
	return b
}

func (b *NutritionOrderBuilder) SetSubject(reference string, display ...string) *NutritionOrderBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *NutritionOrderBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *NutritionOrderBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetEncounter sets encounter from a Reference, as the PHP builder does.
func (b *NutritionOrderBuilder) SetEncounter(encounter *datatype.Reference) *NutritionOrderBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

// SetPatient sets patient from a Reference, as the PHP builder does.
func (b *NutritionOrderBuilder) SetPatient(patient *datatype.Reference) *NutritionOrderBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}
