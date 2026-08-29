package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type CompositionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewCompositionBuilder() *CompositionBuilder {
	b := &CompositionBuilder{ResourceType: "Composition", Data: make(map[string]interface{})}
	return b
}

func (b *CompositionBuilder) SetId(id string) *CompositionBuilder { b.Data["id"] = id; return b }

func (b *CompositionBuilder) AddIdentifier(id *datatype.Identifier) *CompositionBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}

func (b *CompositionBuilder) SetStatus(status string) *CompositionBuilder { b.Data["status"] = status; return b }
func (b *CompositionBuilder) SetType(type_ *datatype.CodeableConcept) *CompositionBuilder { b.Data["type"] = type_.ToArray(); return b }
func (b *CompositionBuilder) SetCategory(category *datatype.CodeableConcept) *CompositionBuilder { b.Data["category"] = []interface{}{category.ToArray()}; return b }
func (b *CompositionBuilder) SetSubject(subjectRef string) *CompositionBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *CompositionBuilder) SetEncounter(encRef string) *CompositionBuilder { b.Data["encounter"] = map[string]interface{}{"reference": encRef}; return b }
func (b *CompositionBuilder) SetDate(date string) *CompositionBuilder { b.Data["date"] = date; return b }
func (b *CompositionBuilder) SetAuthor(authorRef string) *CompositionBuilder { b.Data["author"] = []interface{}{map[string]interface{}{"reference": authorRef}}; return b }
func (b *CompositionBuilder) SetTitle(title string) *CompositionBuilder { b.Data["title"] = title; return b }
func (b *CompositionBuilder) SetConfidentiality(conf string) *CompositionBuilder { b.Data["confidentiality"] = conf; return b }
func (b *CompositionBuilder) AddSection(sectionCode *datatype.CodeableConcept, title string, entryRef string) *CompositionBuilder {
	if _, ok := b.Data["section"]; !ok { b.Data["section"] = make([]interface{}, 0) }
	sec := map[string]interface{}{"title": title}
	if sectionCode != nil { sec["code"] = sectionCode.ToArray() }
	if entryRef != "" { sec["entry"] = []interface{}{map[string]interface{}{"reference": entryRef}} }
	b.Data["section"] = append(b.Data["section"].([]interface{}), sec)
	return b
}

func (b *CompositionBuilder) Build() map[string]interface{} { return b.Data }
