package builder

import (
	"encoding/json"
	"fmt"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type DeviceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewDeviceBuilder() *DeviceBuilder {
	b := &DeviceBuilder{ResourceType: "Device", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Device"
	return b
}

func (b *DeviceBuilder) SetId(id string) *DeviceBuilder {
	b.Data["id"] = id
	return b
}

func (b *DeviceBuilder) AddIdentifier(identifier *datatype.Identifier) *DeviceBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *DeviceBuilder) SetStatus(status string) *DeviceBuilder {
	b.Data["status"] = status
	return b
}

func (b *DeviceBuilder) SetSubject(reference string, display ...string) *DeviceBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *DeviceBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *DeviceBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient from a Reference, as the PHP builder does.
func (b *DeviceBuilder) SetPatient(patient *datatype.Reference) *DeviceBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}

// SetType sets type from a CodeableConcept, as the PHP builder does.
func (b *DeviceBuilder) SetType(kind *datatype.CodeableConcept) *DeviceBuilder {
	b.Data["type"] = kind.ToArray()
	return b
}

// AddNote accepts an annotation or a plain string, as the PHP builder does; a string becomes {text}.
func (b *DeviceBuilder) AddNote(note interface{}) *DeviceBuilder {
	entri := map[string]interface{}{}
	switch v := note.(type) {
	case string:
		entri["text"] = v
	case *datatype.Annotation:
		entri = v.ToArray()
	case datatype.Annotation:
		entri = v.ToArray()
	default:
		entri["text"] = fmt.Sprint(v)
	}
	daftar, _ := b.Data["note"].([]interface{})
	b.Data["note"] = append(daftar, entri)
	return b
}
