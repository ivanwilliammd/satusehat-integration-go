package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SubstanceSourceMaterialBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSubstanceSourceMaterialBuilder() *SubstanceSourceMaterialBuilder {
	b := &SubstanceSourceMaterialBuilder{ResourceType: "SubstanceSourceMaterial", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "SubstanceSourceMaterial"
	return b
}

func (b *SubstanceSourceMaterialBuilder) SetId(id string) *SubstanceSourceMaterialBuilder {
	b.Data["id"] = id
	return b
}

func (b *SubstanceSourceMaterialBuilder) AddIdentifier(identifier *datatype.Identifier) *SubstanceSourceMaterialBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *SubstanceSourceMaterialBuilder) SetStatus(status string) *SubstanceSourceMaterialBuilder {
	b.Data["status"] = status
	return b
}

func (b *SubstanceSourceMaterialBuilder) SetSubject(reference string, display ...string) *SubstanceSourceMaterialBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *SubstanceSourceMaterialBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SubstanceSourceMaterialBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *SubstanceSourceMaterialBuilder) SetCode(system, code, display string) *SubstanceSourceMaterialBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *SubstanceSourceMaterialBuilder) SetEncounter(reference string) *SubstanceSourceMaterialBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
