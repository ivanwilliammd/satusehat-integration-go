package builder

import (
	"encoding/json"
	"fmt"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ClinicalImpressionBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewClinicalImpressionBuilder() *ClinicalImpressionBuilder {
	b := &ClinicalImpressionBuilder{ResourceType: "ClinicalImpression", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ClinicalImpression"
	return b
}

func (b *ClinicalImpressionBuilder) SetId(id string) *ClinicalImpressionBuilder {
	b.Data["id"] = id
	return b
}

func (b *ClinicalImpressionBuilder) AddIdentifier(identifier *datatype.Identifier) *ClinicalImpressionBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ClinicalImpressionBuilder) SetStatus(status string) *ClinicalImpressionBuilder {
	b.Data["status"] = status
	return b
}

func (b *ClinicalImpressionBuilder) SetSubject(reference string, display ...string) *ClinicalImpressionBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ClinicalImpressionBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ClinicalImpressionBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code from a CodeableConcept, as the PHP builder does.
func (b *ClinicalImpressionBuilder) SetCode(code *datatype.CodeableConcept) *ClinicalImpressionBuilder {
	b.Data["code"] = code.ToArray()
	return b
}

// SetEncounter sets encounter from a Reference, as the PHP builder does.
func (b *ClinicalImpressionBuilder) SetEncounter(encounter *datatype.Reference) *ClinicalImpressionBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

// SetDate sets date, as the PHP builder does.
func (b *ClinicalImpressionBuilder) SetDate(date string) *ClinicalImpressionBuilder {
	b.Data["date"] = date
	return b
}

// AddNote accepts an annotation or a plain string, as the PHP builder does; a string becomes {text}.
func (b *ClinicalImpressionBuilder) AddNote(note interface{}) *ClinicalImpressionBuilder {
	entri := map[string]interface{}{}
	switch v := note.(type) {
	case string:
		entri["text"] = v
	case *datatype.Annotation:
		entri = v.ToArray()
	case datatype.Annotation:
		entri = v.ToArray()
	default:
		entri["text"] = fmt.Sprint(v)
	}
	daftar, _ := b.Data["note"].([]interface{})
	b.Data["note"] = append(daftar, entri)
	return b
}
