package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type SubstanceBuilder struct {
	Data map[string]interface{}
}

func NewSubstanceBuilder() *SubstanceBuilder {
	return &SubstanceBuilder{Data: make(map[string]interface{})}
}

func (b *SubstanceBuilder) SetId(id string) *SubstanceBuilder {
	b.Data["id"] = id
	return b
}

func (b *SubstanceBuilder) AddIdentifier(system, value string) *SubstanceBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), map[string]string{"system": system, "value": value})
	return b
}

func (b *SubstanceBuilder) SetStatus(status string) *SubstanceBuilder {
	b.Data["status"] = status
	return b
}

func (b *SubstanceBuilder) SetCategory(cc *datatype.CodeableConcept) *SubstanceBuilder {
	if _, ok := b.Data["category"]; !ok {
		b.Data["category"] = make([]interface{}, 0)
	}
	b.Data["category"] = append(b.Data["category"].([]interface{}), cc.ToArray())
	return b
}

func (b *SubstanceBuilder) SetCode(cc *datatype.CodeableConcept) *SubstanceBuilder {
	b.Data["code"] = cc.ToArray()
	return b
}

func (b *SubstanceBuilder) SetDescription(desc string) *SubstanceBuilder {
	b.Data["description"] = desc
	return b
}

func (b *SubstanceBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "Substance"
	return dt
}
