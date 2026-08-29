package builder

type DeviceUseStatementBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewDeviceUseStatementBuilder() *DeviceUseStatementBuilder {
	b := &DeviceUseStatementBuilder{ResourceType: "DeviceUseStatement", Data: make(map[string]interface{})}
	return b
}

func (b *DeviceUseStatementBuilder) SetId(id string) *DeviceUseStatementBuilder { b.Data["id"] = id; return b }

func (b *DeviceUseStatementBuilder) AddIdentifier(id string) *DeviceUseStatementBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), map[string]interface{}{"value": id})
	return b
}

func (b *DeviceUseStatementBuilder) SetStatus(status string) *DeviceUseStatementBuilder { b.Data["status"] = status; return b }
func (b *DeviceUseStatementBuilder) SetPatient(patientRef string) *DeviceUseStatementBuilder { b.Data["patient"] = map[string]interface{}{"reference": patientRef}; return b }
func (b *DeviceUseStatementBuilder) SetDevice(deviceRef string) *DeviceUseStatementBuilder { b.Data["device"] = map[string]interface{}{"reference": deviceRef}; return b }
func (b *DeviceUseStatementBuilder) SetUsedDuring(start string, end string) *DeviceUseStatementBuilder {
	b.Data["usedDuring"] = map[string]interface{}{"start": start}
	if end != "" { b.Data["usedDuring"].(map[string]interface{})["end"] = end }
	return b
}
func (b *DeviceUseStatementBuilder) SetIndication(code string) *DeviceUseStatementBuilder {
	b.Data["indication"] = []interface{}{map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": code}}}}
	return b
}
func (b *DeviceUseStatementBuilder) SetNote(note string) *DeviceUseStatementBuilder {
	b.Data["note"] = []interface{}{map[string]interface{}{"text": note}}; return b
}

func (b *DeviceUseStatementBuilder) Build() map[string]interface{} { return b.Data }
