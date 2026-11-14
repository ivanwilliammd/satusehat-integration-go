package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type DeviceDefinitionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewDeviceDefinitionBuilder() *DeviceDefinitionBuilder {
	b := &DeviceDefinitionBuilder{ResourceType: "DeviceDefinition", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "DeviceDefinition"
	return b
}

func (b *DeviceDefinitionBuilder) SetId(id string) *DeviceDefinitionBuilder {
	b.Data["id"] = id
	return b
}

func (b *DeviceDefinitionBuilder) AddIdentifier(identifier *datatype.Identifier) *DeviceDefinitionBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *DeviceDefinitionBuilder) SetStatus(status string) *DeviceDefinitionBuilder {
	b.Data["status"] = status
	return b
}

func (b *DeviceDefinitionBuilder) SetSubject(reference string, display ...string) *DeviceDefinitionBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *DeviceDefinitionBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *DeviceDefinitionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *DeviceDefinitionBuilder) SetCode(system, code, display string) *DeviceDefinitionBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *DeviceDefinitionBuilder) SetEncounter(reference string) *DeviceDefinitionBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
