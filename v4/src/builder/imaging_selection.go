package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ImagingSelectionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewImagingSelectionBuilder() *ImagingSelectionBuilder {
	b := &ImagingSelectionBuilder{ResourceType: "ImagingSelection", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ImagingSelection"
	return b
}

func (b *ImagingSelectionBuilder) SetId(id string) *ImagingSelectionBuilder {
	b.Data["id"] = id
	return b
}

func (b *ImagingSelectionBuilder) AddIdentifier(identifier *datatype.Identifier) *ImagingSelectionBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ImagingSelectionBuilder) SetStatus(status string) *ImagingSelectionBuilder {
	b.Data["status"] = status
	return b
}

func (b *ImagingSelectionBuilder) SetSubject(reference string, display ...string) *ImagingSelectionBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ImagingSelectionBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ImagingSelectionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ImagingSelectionBuilder) SetCode(system, code, display string) *ImagingSelectionBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ImagingSelectionBuilder) SetEncounter(reference string) *ImagingSelectionBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
