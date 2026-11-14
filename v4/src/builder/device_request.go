package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type DeviceRequestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewDeviceRequestBuilder() *DeviceRequestBuilder {
	b := &DeviceRequestBuilder{ResourceType: "DeviceRequest", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "DeviceRequest"
	return b
}

func (b *DeviceRequestBuilder) SetId(id string) *DeviceRequestBuilder {
	b.Data["id"] = id
	return b
}

func (b *DeviceRequestBuilder) AddIdentifier(identifier *datatype.Identifier) *DeviceRequestBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *DeviceRequestBuilder) SetStatus(status string) *DeviceRequestBuilder {
	b.Data["status"] = status
	return b
}

func (b *DeviceRequestBuilder) SetSubject(reference string, display ...string) *DeviceRequestBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *DeviceRequestBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *DeviceRequestBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *DeviceRequestBuilder) SetCode(system, code, display string) *DeviceRequestBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *DeviceRequestBuilder) SetEncounter(reference string) *DeviceRequestBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
