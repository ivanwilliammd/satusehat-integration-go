package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type DeviceRequestBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewDeviceRequestBuilder() *DeviceRequestBuilder {
    b := &DeviceRequestBuilder{ResourceType: "DeviceRequest", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "DeviceRequest"
    return b
}

func (b *DeviceRequestBuilder) setId(id string) *DeviceRequestBuilder {
    b.Data["id"] = id
    return b
}

func (b *DeviceRequestBuilder) addIdentifier(identifier *datatype.Identifier) *DeviceRequestBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *DeviceRequestBuilder) setStatus(status string) *DeviceRequestBuilder {
    b.Data["status"] = status
    return b
}

func (b *DeviceRequestBuilder) setSubject(reference string) *DeviceRequestBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *DeviceRequestBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *DeviceRequestBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
