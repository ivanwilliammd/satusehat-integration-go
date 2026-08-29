package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type DeviceUseStatementBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewDeviceUseStatementBuilder() *DeviceUseStatementBuilder {
    b := &DeviceUseStatementBuilder{ResourceType: "DeviceUseStatement", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "DeviceUseStatement"
    return b
}

func (b *DeviceUseStatementBuilder) setId(id string) *DeviceUseStatementBuilder {
    b.Data["id"] = id
    return b
}

func (b *DeviceUseStatementBuilder) addIdentifier(identifier *datatype.Identifier) *DeviceUseStatementBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *DeviceUseStatementBuilder) setStatus(status string) *DeviceUseStatementBuilder {
    b.Data["status"] = status
    return b
}

func (b *DeviceUseStatementBuilder) setSubject(reference string) *DeviceUseStatementBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *DeviceUseStatementBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *DeviceUseStatementBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
