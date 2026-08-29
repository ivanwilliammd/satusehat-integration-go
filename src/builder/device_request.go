package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type DeviceRequestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewDeviceRequestBuilder() *DeviceRequestBuilder {
	b := &DeviceRequestBuilder{ResourceType: "DeviceRequest", Data: make(map[string]interface{})}
	return b
}

func (b *DeviceRequestBuilder) SetId(id string) *DeviceRequestBuilder { b.Data["id"] = id; return b }

func (b *DeviceRequestBuilder) AddIdentifier(id *datatype.Identifier) *DeviceRequestBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}

func (b *DeviceRequestBuilder) SetStatus(status string) *DeviceRequestBuilder { b.Data["status"] = status; return b }
func (b *DeviceRequestBuilder) SetIntent(intent string) *DeviceRequestBuilder { b.Data["intent"] = intent; return b }
func (b *DeviceRequestBuilder) SetPriority(priority string) *DeviceRequestBuilder { b.Data["priority"] = priority; return b }
func (b *DeviceRequestBuilder) SetCategory(code *datatype.CodeableConcept) *DeviceRequestBuilder { b.Data["category"] = []interface{}{code.ToArray()}; return b }
func (b *DeviceRequestBuilder) SetCode(code *datatype.CodeableConcept) *DeviceRequestBuilder { b.Data["codeCodeableConcept"] = code.ToArray(); return b }
func (b *DeviceRequestBuilder) SetSubject(subjectRef string) *DeviceRequestBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *DeviceRequestBuilder) SetEncounter(encRef string) *DeviceRequestBuilder { b.Data["encounter"] = map[string]interface{}{"reference": encRef}; return b }
func (b *DeviceRequestBuilder) SetAuthoredOn(authoredOn string) *DeviceRequestBuilder { b.Data["authoredOn"] = authoredOn; return b }
func (b *DeviceRequestBuilder) SetRequester(requesterRef string) *DeviceRequestBuilder { b.Data["requester"] = map[string]interface{}{"reference": requesterRef}; return b }
func (b *DeviceRequestBuilder) SetPerformer(performerRef string) *DeviceRequestBuilder { b.Data["performer"] = map[string]interface{}{"reference": performerRef}; return b }
func (b *DeviceRequestBuilder) SetReasonCode(reason *datatype.CodeableConcept) *DeviceRequestBuilder { b.Data["reasonCode"] = []interface{}{reason.ToArray()}; return b }
func (b *DeviceRequestBuilder) SetNote(note string) *DeviceRequestBuilder {
	b.Data["note"] = []interface{}{map[string]interface{}{"text": note}}; return b
}

func (b *DeviceRequestBuilder) Build() map[string]interface{} { return b.Data }
