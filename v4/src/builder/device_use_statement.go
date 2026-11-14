package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type DeviceUseStatementBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewDeviceUseStatementBuilder() *DeviceUseStatementBuilder {
	b := &DeviceUseStatementBuilder{ResourceType: "DeviceUseStatement", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "DeviceUseStatement"
	return b
}

func (b *DeviceUseStatementBuilder) SetId(id string) *DeviceUseStatementBuilder {
	b.Data["id"] = id
	return b
}

func (b *DeviceUseStatementBuilder) AddIdentifier(identifier *datatype.Identifier) *DeviceUseStatementBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *DeviceUseStatementBuilder) SetStatus(status string) *DeviceUseStatementBuilder {
	b.Data["status"] = status
	return b
}

func (b *DeviceUseStatementBuilder) SetSubject(reference string, display ...string) *DeviceUseStatementBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *DeviceUseStatementBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *DeviceUseStatementBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *DeviceUseStatementBuilder) SetCode(system, code, display string) *DeviceUseStatementBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *DeviceUseStatementBuilder) SetEncounter(reference string) *DeviceUseStatementBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
