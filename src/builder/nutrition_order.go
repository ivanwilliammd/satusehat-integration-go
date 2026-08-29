package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type NutritionOrderBuilder struct {
	Data map[string]interface{}
}

func NewNutritionOrderBuilder() *NutritionOrderBuilder {
	return &NutritionOrderBuilder{Data: make(map[string]interface{})}
}

func (b *NutritionOrderBuilder) SetId(id string) *NutritionOrderBuilder {
	b.Data["id"] = id
	return b
}

func (b *NutritionOrderBuilder) SetStatus(status string) *NutritionOrderBuilder {
	b.Data["status"] = status
	return b
}

func (b *NutritionOrderBuilder) SetIntent(intent string) *NutritionOrderBuilder {
	b.Data["intent"] = intent
	return b
}

func (b *NutritionOrderBuilder) SetPatient(ref *datatype.Reference) *NutritionOrderBuilder {
	b.Data["patient"] = ref.ToArray()
	return b
}

func (b *NutritionOrderBuilder) SetDateTime(dt string) *NutritionOrderBuilder {
	b.Data["dateTime"] = dt
	return b
}

func (b *NutritionOrderBuilder) SetOrderer(ref *datatype.Reference) *NutritionOrderBuilder {
	b.Data["orderer"] = ref.ToArray()
	return b
}

func (b *NutritionOrderBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "NutritionOrder"
	return dt
}
