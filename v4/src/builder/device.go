package builder

import (
	"encoding/json"
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
