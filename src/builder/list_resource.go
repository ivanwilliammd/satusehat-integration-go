package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type ListBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewListBuilder() *ListBuilder {
	b := &ListBuilder{ResourceType: "List", Data: make(map[string]interface{})}
	return b
}

func (b *ListBuilder) SetId(id string) *ListBuilder { b.Data["id"] = id; return b }
func (b *ListBuilder) SetStatus(status string) *ListBuilder { b.Data["status"] = status; return b }
func (b *ListBuilder) SetMode(mode string) *ListBuilder { b.Data["mode"] = mode; return b }
func (b *ListBuilder) SetTitle(title string) *ListBuilder { b.Data["title"] = title; return b }
func (b *ListBuilder) SetCode(code *datatype.CodeableConcept) *ListBuilder { b.Data["code"] = code.ToArray(); return b }
func (b *ListBuilder) SetSubject(subjectRef string) *ListBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *ListBuilder) SetEncounter(encRef string) *ListBuilder { b.Data["encounter"] = map[string]interface{}{"reference": encRef}; return b }
func (b *ListBuilder) SetDate(date string) *ListBuilder { b.Data["date"] = date; return b }
func (b *ListBuilder) SetSource(sourceRef string) *ListBuilder { b.Data["source"] = map[string]interface{}{"reference": sourceRef}; return b }
func (b *ListBuilder) SetOrderedBy(code *datatype.CodeableConcept) *ListBuilder { b.Data["orderedBy"] = code.ToArray(); return b }

func (b *ListBuilder) AddEntry(itemRef string, flag *datatype.CodeableConcept, note string) *ListBuilder {
	if _, ok := b.Data["entry"]; !ok { b.Data["entry"] = make([]interface{}, 0) }
	entry := map[string]interface{}{"flag": []interface{}{map[string]interface{}{"text": flag.ToArray()["text"]}}}
	if itemRef != "" { entry["item"] = map[string]interface{}{"reference": itemRef} }
	if note != "" { entry["note"] = []interface{}{map[string]interface{}{"text": note}}} 
	b.Data["entry"] = append(b.Data["entry"].([]interface{}), entry)
	return b
}

func (b *ListBuilder) Build() map[string]interface{} { return b.Data }
