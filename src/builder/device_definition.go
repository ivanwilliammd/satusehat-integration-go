package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type DeviceDefinitionBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewDeviceDefinitionBuilder() *DeviceDefinitionBuilder {
    b := &DeviceDefinitionBuilder{ResourceType: "DeviceDefinition", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "DeviceDefinition"
    return b
}

func (b *DeviceDefinitionBuilder) setId(id string) *DeviceDefinitionBuilder {
    b.Data["id"] = id
    return b
}

func (b *DeviceDefinitionBuilder) addIdentifier(identifier *datatype.Identifier) *DeviceDefinitionBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *DeviceDefinitionBuilder) setStatus(status string) *DeviceDefinitionBuilder {
    b.Data["status"] = status
    return b
}

func (b *DeviceDefinitionBuilder) setSubject(reference string) *DeviceDefinitionBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *DeviceDefinitionBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *DeviceDefinitionBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
