package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type DeviceBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewDeviceBuilder() *DeviceBuilder {
    b := &DeviceBuilder{ResourceType: "Device", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "Device"
    return b
}

func (b *DeviceBuilder) setId(id string) *DeviceBuilder {
    b.Data["id"] = id
    return b
}

func (b *DeviceBuilder) addIdentifier(identifier *datatype.Identifier) *DeviceBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *DeviceBuilder) setStatus(status string) *DeviceBuilder {
    b.Data["status"] = status
    return b
}

func (b *DeviceBuilder) setSubject(reference string) *DeviceBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *DeviceBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *DeviceBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
