package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type MediaBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMediaBuilder() *MediaBuilder {
	b := &MediaBuilder{ResourceType: "Media", Data: make(map[string]interface{})}
	return b
}

func (b *MediaBuilder) SetId(id string) *MediaBuilder { b.Data["id"] = id; return b }
func (b *MediaBuilder) AddIdentifier(id *datatype.Identifier) *MediaBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}
func (b *MediaBuilder) SetStatus(status string) *MediaBuilder { b.Data["status"] = status; return b }
func (b *MediaBuilder) SetType(code *datatype.CodeableConcept) *MediaBuilder { b.Data["type"] = code.ToArray(); return b }
func (b *MediaBuilder) SetSubject(subjectRef string) *MediaBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *MediaBuilder) SetEncounter(encRef string) *MediaBuilder { b.Data["encounter"] = map[string]interface{}{"reference": encRef}; return b }
func (b *MediaBuilder) SetDateCreated(date string) *MediaBuilder { b.Data["createdDateTime"] = date; return b }
func (b *MediaBuilder) SetOperator(operatorRef string) *MediaBuilder { b.Data["operator"] = map[string]interface{}{"reference": operatorRef}; return b }
func (b *MediaBuilder) SetReason(reason *datatype.CodeableConcept) *MediaBuilder { b.Data["reasonCode"] = []interface{}{reason.ToArray()}; return b }
func (b *MediaBuilder) SetBodySite(bodySite *datatype.CodeableConcept) *MediaBuilder { b.Data["bodySite"] = bodySite.ToArray(); return b }
func (b *MediaBuilder) SetContent(contentType string, url string) *MediaBuilder { b.Data["content"] = map[string]interface{}{"contentType": contentType, "url": url}; return b }
func (b *MediaBuilder) SetNote(note string) *MediaBuilder { b.Data["note"] = []interface{}{map[string]interface{}{"text": note}}; return b }
func (b *MediaBuilder) Build() map[string]interface{} { return b.Data }
