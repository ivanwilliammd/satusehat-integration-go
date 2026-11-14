package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type BodyStructureBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewBodyStructureBuilder() *BodyStructureBuilder {
	b := &BodyStructureBuilder{ResourceType: "BodyStructure", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "BodyStructure"
	return b
}

func (b *BodyStructureBuilder) SetId(id string) *BodyStructureBuilder {
	b.Data["id"] = id
	return b
}

func (b *BodyStructureBuilder) AddIdentifier(identifier *datatype.Identifier) *BodyStructureBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *BodyStructureBuilder) SetStatus(status string) *BodyStructureBuilder {
	b.Data["status"] = status
	return b
}

func (b *BodyStructureBuilder) SetSubject(reference string, display ...string) *BodyStructureBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *BodyStructureBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *BodyStructureBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *BodyStructureBuilder) SetCode(system, code, display string) *BodyStructureBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *BodyStructureBuilder) SetEncounter(reference string) *BodyStructureBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
