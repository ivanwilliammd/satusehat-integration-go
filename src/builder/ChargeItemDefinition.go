package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type ChargeItemDefinitionBuilder struct {
	Data map[string]interface{}
}

func NewChargeItemDefinitionBuilder() *ChargeItemDefinitionBuilder {
	return &ChargeItemDefinitionBuilder{Data: make(map[string]interface{})}
}

func (b *ChargeItemDefinitionBuilder) SetId(id string) *ChargeItemDefinitionBuilder {
	b.Data["id"] = id
	return b
}

func (b *ChargeItemDefinitionBuilder) SetUrl(url string) *ChargeItemDefinitionBuilder {
	b.Data["url"] = url
	return b
}

func (b *ChargeItemDefinitionBuilder) SetVersion(version string) *ChargeItemDefinitionBuilder {
	b.Data["version"] = version
	return b
}

func (b *ChargeItemDefinitionBuilder) SetName(name string) *ChargeItemDefinitionBuilder {
	b.Data["name"] = name
	return b
}

func (b *ChargeItemDefinitionBuilder) SetTitle(title string) *ChargeItemDefinitionBuilder {
	b.Data["title"] = title
	return b
}

func (b *ChargeItemDefinitionBuilder) SetStatus(status string) *ChargeItemDefinitionBuilder {
	b.Data["status"] = status
	return b
}

func (b *ChargeItemDefinitionBuilder) SetDate(date string) *ChargeItemDefinitionBuilder {
	b.Data["date"] = date
	return b
}

func (b *ChargeItemDefinitionBuilder) SetPublisher(publisher string) *ChargeItemDefinitionBuilder {
	b.Data["publisher"] = publisher
	return b
}

func (b *ChargeItemDefinitionBuilder) SetDescription(description string) *ChargeItemDefinitionBuilder {
	b.Data["description"] = description
	return b
}

func (b *ChargeItemDefinitionBuilder) SetCode(cc *datatype.CodeableConcept) *ChargeItemDefinitionBuilder {
	b.Data["code"] = cc.ToArray()
	return b
}

func (b *ChargeItemDefinitionBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "ChargeItemDefinition"
	return dt
}
