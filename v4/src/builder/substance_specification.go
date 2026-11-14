package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SubstanceSpecificationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSubstanceSpecificationBuilder() *SubstanceSpecificationBuilder {
	b := &SubstanceSpecificationBuilder{ResourceType: "SubstanceSpecification", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "SubstanceSpecification"
	return b
}

func (b *SubstanceSpecificationBuilder) SetId(id string) *SubstanceSpecificationBuilder {
	b.Data["id"] = id
	return b
}

func (b *SubstanceSpecificationBuilder) AddIdentifier(identifier *datatype.Identifier) *SubstanceSpecificationBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *SubstanceSpecificationBuilder) SetStatus(status string) *SubstanceSpecificationBuilder {
	b.Data["status"] = status
	return b
}

func (b *SubstanceSpecificationBuilder) SetSubject(reference string, display ...string) *SubstanceSpecificationBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *SubstanceSpecificationBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SubstanceSpecificationBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *SubstanceSpecificationBuilder) SetCode(system, code, display string) *SubstanceSpecificationBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *SubstanceSpecificationBuilder) SetEncounter(reference string) *SubstanceSpecificationBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
