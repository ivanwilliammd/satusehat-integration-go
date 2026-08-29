package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type DocumentReferenceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewDocumentReferenceBuilder() *DocumentReferenceBuilder {
	b := &DocumentReferenceBuilder{ResourceType: "DocumentReference", Data: make(map[string]interface{})}
	return b
}

func (b *DocumentReferenceBuilder) SetId(id string) *DocumentReferenceBuilder { b.Data["id"] = id; return b }

func (b *DocumentReferenceBuilder) AddIdentifier(id *datatype.Identifier) *DocumentReferenceBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}

func (b *DocumentReferenceBuilder) SetStatus(status string) *DocumentReferenceBuilder { b.Data["status"] = status; return b }
func (b *DocumentReferenceBuilder) SetDocStatus(docStatus string) *DocumentReferenceBuilder { b.Data["docStatus"] = docStatus; return b }
func (b *DocumentReferenceBuilder) SetType(type_ *datatype.CodeableConcept) *DocumentReferenceBuilder { b.Data["type"] = type_.ToArray(); return b }
func (b *DocumentReferenceBuilder) SetCategory(category *datatype.CodeableConcept) *DocumentReferenceBuilder { b.Data["category"] = []interface{}{category.ToArray()}; return b }
func (b *DocumentReferenceBuilder) SetSubject(subjectRef string) *DocumentReferenceBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *DocumentReferenceBuilder) SetDate(date string) *DocumentReferenceBuilder { b.Data["date"] = date; return b }
func (b *DocumentReferenceBuilder) SetAuthor(authorRef string) *DocumentReferenceBuilder { b.Data["author"] = []interface{}{map[string]interface{}{"reference": authorRef}}; return b }
func (b *DocumentReferenceBuilder) SetDescription(description string) *DocumentReferenceBuilder { b.Data["description"] = description; return b }
func (b *DocumentReferenceBuilder) SetContent(contentType string, url string) *DocumentReferenceBuilder {
	b.Data["content"] = []interface{}{
		map[string]interface{}{
			"attachment": map[string]interface{}{"contentType": contentType, "url": url},
		},
	}
	return b
}
func (b *DocumentReferenceBuilder) SetContext(encRef string, periodStart string) *DocumentReferenceBuilder {
	b.Data["context"] = map[string]interface{}{
		"encounter": []interface{}{map[string]interface{}{"reference": encRef}},
		"period":    map[string]interface{}{"start": periodStart},
	}
	return b
}

func (b *DocumentReferenceBuilder) Build() map[string]interface{} { return b.Data }
