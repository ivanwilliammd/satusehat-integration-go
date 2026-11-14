package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SubstanceProteinBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSubstanceProteinBuilder() *SubstanceProteinBuilder {
	b := &SubstanceProteinBuilder{ResourceType: "SubstanceProtein", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "SubstanceProtein"
	return b
}

func (b *SubstanceProteinBuilder) SetId(id string) *SubstanceProteinBuilder {
	b.Data["id"] = id
	return b
}

func (b *SubstanceProteinBuilder) AddIdentifier(identifier *datatype.Identifier) *SubstanceProteinBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *SubstanceProteinBuilder) SetStatus(status string) *SubstanceProteinBuilder {
	b.Data["status"] = status
	return b
}

func (b *SubstanceProteinBuilder) SetSubject(reference string, display ...string) *SubstanceProteinBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *SubstanceProteinBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SubstanceProteinBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *SubstanceProteinBuilder) SetCode(system, code, display string) *SubstanceProteinBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *SubstanceProteinBuilder) SetEncounter(reference string) *SubstanceProteinBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
