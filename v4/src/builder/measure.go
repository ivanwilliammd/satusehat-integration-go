package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MeasureBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMeasureBuilder() *MeasureBuilder {
	b := &MeasureBuilder{ResourceType: "Measure", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Measure"
	return b
}

func (b *MeasureBuilder) SetId(id string) *MeasureBuilder {
	b.Data["id"] = id
	return b
}

func (b *MeasureBuilder) AddIdentifier(identifier *datatype.Identifier) *MeasureBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *MeasureBuilder) SetStatus(status string) *MeasureBuilder {
	b.Data["status"] = status
	return b
}

func (b *MeasureBuilder) SetSubject(reference string, display ...string) *MeasureBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *MeasureBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *MeasureBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *MeasureBuilder) SetCode(system, code, display string) *MeasureBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *MeasureBuilder) SetEncounter(reference string) *MeasureBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
