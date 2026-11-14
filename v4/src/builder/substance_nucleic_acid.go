package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type SubstanceNucleicAcidBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSubstanceNucleicAcidBuilder() *SubstanceNucleicAcidBuilder {
	b := &SubstanceNucleicAcidBuilder{ResourceType: "SubstanceNucleicAcid", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "SubstanceNucleicAcid"
	return b
}

func (b *SubstanceNucleicAcidBuilder) SetId(id string) *SubstanceNucleicAcidBuilder {
	b.Data["id"] = id
	return b
}

func (b *SubstanceNucleicAcidBuilder) AddIdentifier(identifier *datatype.Identifier) *SubstanceNucleicAcidBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *SubstanceNucleicAcidBuilder) SetStatus(status string) *SubstanceNucleicAcidBuilder {
	b.Data["status"] = status
	return b
}

func (b *SubstanceNucleicAcidBuilder) SetSubject(reference string, display ...string) *SubstanceNucleicAcidBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *SubstanceNucleicAcidBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *SubstanceNucleicAcidBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *SubstanceNucleicAcidBuilder) SetCode(system, code, display string) *SubstanceNucleicAcidBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *SubstanceNucleicAcidBuilder) SetEncounter(reference string) *SubstanceNucleicAcidBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
