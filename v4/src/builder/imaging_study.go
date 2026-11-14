package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ImagingStudyBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewImagingStudyBuilder() *ImagingStudyBuilder {
	b := &ImagingStudyBuilder{ResourceType: "ImagingStudy", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ImagingStudy"
	return b
}

func (b *ImagingStudyBuilder) SetId(id string) *ImagingStudyBuilder {
	b.Data["id"] = id
	return b
}

func (b *ImagingStudyBuilder) AddIdentifier(identifier *datatype.Identifier) *ImagingStudyBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ImagingStudyBuilder) SetStatus(status string) *ImagingStudyBuilder {
	b.Data["status"] = status
	return b
}

func (b *ImagingStudyBuilder) SetSubject(reference string, display ...string) *ImagingStudyBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ImagingStudyBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ImagingStudyBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ImagingStudyBuilder) SetEncounter(reference string) *ImagingStudyBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
