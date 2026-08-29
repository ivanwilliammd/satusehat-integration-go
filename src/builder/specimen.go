package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type SpecimenBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSpecimenBuilder() *SpecimenBuilder {
	b := &SpecimenBuilder{ResourceType: "Specimen", Data: make(map[string]interface{})}
	return b
}

func (b *SpecimenBuilder) SetId(id string) *SpecimenBuilder { b.Data["id"] = id; return b }

func (b *SpecimenBuilder) AddIdentifier(id *datatype.Identifier) *SpecimenBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}

func (b *SpecimenBuilder) SetAccessionIdentifier(system string, value string) *SpecimenBuilder {
	b.Data["accessionIdentifier"] = map[string]interface{}{"system": system, "value": value}
	return b
}
func (b *SpecimenBuilder) SetStatus(status string) *SpecimenBuilder { b.Data["status"] = status; return b }
func (b *SpecimenBuilder) SetType(code *datatype.CodeableConcept) *SpecimenBuilder { b.Data["type"] = code.ToArray(); return b }
func (b *SpecimenBuilder) SetSubject(subjectRef string) *SpecimenBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *SpecimenBuilder) SetReceivedTime(receivedTime string) *SpecimenBuilder { b.Data["receivedTime"] = receivedTime; return b }
func (b *SpecimenBuilder) SetCollectionCollected(collected string) *SpecimenBuilder {
	b.Data["collection"] = map[string]interface{}{"collectedDateTime": collected}
	return b
}
func (b *SpecimenBuilder) SetCollectionBodySite(bodySite *datatype.CodeableConcept) *SpecimenBuilder {
	if _, ok := b.Data["collection"]; !ok { b.Data["collection"] = make(map[string]interface{}) }
	b.Data["collection"].(map[string]interface{})["bodySite"] = bodySite.ToArray()
	return b
}
func (b *SpecimenBuilder) AddProcessing(description string, procedure *datatype.CodeableConcept) *SpecimenBuilder {
	proc := map[string]interface{}{"description": description}
	if procedure != nil { proc["procedure"] = procedure.ToArray() }
	if _, ok := b.Data["processing"]; !ok { b.Data["processing"] = make([]interface{}, 0) }
	b.Data["processing"] = append(b.Data["processing"].([]interface{}), proc)
	return b
}
func (b *SpecimenBuilder) SetContainer(identifier string, type_ *datatype.CodeableConcept) *SpecimenBuilder {
	container := map[string]interface{}{"identifier": []interface{}{map[string]interface{}{"value": identifier}}}
	if type_ != nil { container["type"] = type_.ToArray() }
	if _, ok := b.Data["container"]; !ok { b.Data["container"] = make([]interface{}, 0) }
	b.Data["container"] = append(b.Data["container"].([]interface{}), container)
	return b
}
func (b *SpecimenBuilder) Build() map[string]interface{} { return b.Data }
