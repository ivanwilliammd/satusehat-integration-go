package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ImagingManifestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewImagingManifestBuilder() *ImagingManifestBuilder {
	b := &ImagingManifestBuilder{ResourceType: "ImagingManifest", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ImagingManifest"
	return b
}

func (b *ImagingManifestBuilder) SetId(id string) *ImagingManifestBuilder {
	b.Data["id"] = id
	return b
}

func (b *ImagingManifestBuilder) AddIdentifier(identifier *datatype.Identifier) *ImagingManifestBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ImagingManifestBuilder) SetStatus(status string) *ImagingManifestBuilder {
	b.Data["status"] = status
	return b
}

func (b *ImagingManifestBuilder) SetSubject(reference string, display ...string) *ImagingManifestBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ImagingManifestBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ImagingManifestBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *ImagingManifestBuilder) SetCode(system, code, display string) *ImagingManifestBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ImagingManifestBuilder) SetEncounter(reference string) *ImagingManifestBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
