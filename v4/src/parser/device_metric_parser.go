package parser

// DeviceMetricParser extracts fields from a FHIR DeviceMetric resource map.
//
// The element list is derived from PayloadBuilderDeviceMetric in the PHP SDK, because this resource has no
// official SATUSEHAT example to take the shape from. Getter names follow the FHIR element names;
// the parsers written from examples keep their descriptive names.
type DeviceMetricParser struct {
	data map[string]any
}

// NewDeviceMetricParser creates a parser for the given DeviceMetric payload.
func NewDeviceMetricParser(data map[string]any) *DeviceMetricParser {
	if data == nil {
		data = map[string]any{}
	}
	return &DeviceMetricParser{data: data}
}

// GetRaw returns the whole payload, so callers are never stuck on a missing accessor.
func (p *DeviceMetricParser) GetRaw() map[string]any { return p.data }

// GetResourceType returns the resourceType element.
func (p *DeviceMetricParser) GetResourceType() string {
	if v, ok := p.data["resourceType"].(string); ok {
		return v
	}
	return ""
}

// GetId returns the id element.
func (p *DeviceMetricParser) GetId() any {
	return p.data["id"]
}

// GetIdentifier returns the identifier element.
func (p *DeviceMetricParser) GetIdentifier() any {
	return p.data["identifier"]
}

// GetStatus returns the status element.
func (p *DeviceMetricParser) GetStatus() any {
	return p.data["status"]
}

// GetType returns the type element.
func (p *DeviceMetricParser) GetType() any {
	return p.data["type"]
}

// GetCategory returns the category element.
func (p *DeviceMetricParser) GetCategory() any {
	return p.data["category"]
}

// GetMeasurementPeriod returns the measurementPeriod element.
func (p *DeviceMetricParser) GetMeasurementPeriod() any {
	return p.data["measurementPeriod"]
}
