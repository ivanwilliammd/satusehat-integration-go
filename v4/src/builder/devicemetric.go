package builder

import (
    "encoding/json"
)

type DeviceMetricBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewDeviceMetricBuilder() *DeviceMetricBuilder {
    b := &DeviceMetricBuilder{ResourceType: "DeviceMetric", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "DeviceMetric"
    return b
}

func (b *DeviceMetricBuilder) setId(id string) *DeviceMetricBuilder {
    b.Data["id"] = id
    return b
}

func (b *DeviceMetricBuilder) addIdentifier(system, value string) *DeviceMetricBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *DeviceMetricBuilder) setStatus(status string) *DeviceMetricBuilder {
    b.Data["status"] = status
    return b
}

func (b *DeviceMetricBuilder) setType(system, code, display string) *DeviceMetricBuilder {
    b.Data["type"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *DeviceMetricBuilder) setCategory(system, code, display string) *DeviceMetricBuilder {
    b.Data["category"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *DeviceMetricBuilder) setMeasurementPeriod(value string) *DeviceMetricBuilder {
    b.Data["measurementPeriod"] = value
    return b
}

func (b *DeviceMetricBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *DeviceMetricBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
