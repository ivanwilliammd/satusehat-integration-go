package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type DeviceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewDeviceBuilder() *DeviceBuilder {
	b := &DeviceBuilder{ResourceType: "Device", Data: make(map[string]interface{})}
	return b
}

func (b *DeviceBuilder) SetId(id string) *DeviceBuilder { b.Data["id"] = id; return b }

func (b *DeviceBuilder) AddIdentifier(id *datatype.Identifier) *DeviceBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}

func (b *DeviceBuilder) SetUdi(udi string) *DeviceBuilder { b.Data["udiCarrier"] = []interface{}{map[string]interface{}{"carrierAIDC": udi}}; return b }
func (b *DeviceBuilder) SetStatus(status string) *DeviceBuilder { b.Data["status"] = status; return b }
func (b *DeviceBuilder) SetType(code *datatype.CodeableConcept) *DeviceBuilder { b.Data["type"] = code.ToArray(); return b }
func (b *DeviceBuilder) SetManufacturer(manufacturer string) *DeviceBuilder { b.Data["manufacturer"] = manufacturer; return b }
func (b *DeviceBuilder) SetDeviceName(name string, type_ string) *DeviceBuilder {
	b.Data["deviceName"] = []interface{}{
		map[string]interface{}{"name": name, "type": type_},
	}
	return b
}
func (b *DeviceBuilder) SetModelNumber(model string) *DeviceBuilder { b.Data["modelNumber"] = model; return b }
func (b *DeviceBuilder) SetSerialNumber(serial string) *DeviceBuilder {
	b.Data["serialNumber"] = serial; return b
}
func (b *DeviceBuilder) SetPatient(patientRef string) *DeviceBuilder { b.Data["patient"] = map[string]interface{}{"reference": patientRef}; return b }
func (b *DeviceBuilder) SetOwner(ownerRef string) *DeviceBuilder { b.Data["owner"] = map[string]interface{}{"reference": ownerRef}; return b }
func (b *DeviceBuilder) SetLocation(locationRef string) *DeviceBuilder { b.Data["location"] = map[string]interface{}{"reference": locationRef}; return b }
func (b *DeviceBuilder) SetNote(note string) *DeviceBuilder {
	b.Data["note"] = []interface{}{map[string]interface{}{"text": note}}; return b
}

func (b *DeviceBuilder) Build() map[string]interface{} { return b.Data }
