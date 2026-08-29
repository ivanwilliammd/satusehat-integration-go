package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// FlagBuilder builds FHIR Flag payload
type FlagBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewFlagBuilder() *FlagBuilder {
	b := &FlagBuilder{ResourceType: "Flag", Data: make(map[string]interface{})}
	return b
}

func (b *FlagBuilder) setId(id string) *FlagBuilder {
	b.Data["id"] = id
	return b
}

func (b *FlagBuilder) addIdentifier(identifier *datatype.Identifier) *FlagBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *FlagBuilder) setStatus(status string) *FlagBuilder {
	b.Data["status"] = status
	return b
}

func (b *FlagBuilder) setCode(code *datatype.CodeableConcept) *FlagBuilder {
	b.Data["code"] = code.ToArray()
	return b
}

func (b *FlagBuilder) setSubject(subject *datatype.Reference) *FlagBuilder {
	b.Data["subject"] = subject.ToArray()
	return b
}

func (b *FlagBuilder) setPeriod(period *datatype.Period) *FlagBuilder {
	b.Data["period"] = period.ToArray()
	return b
}

func (b *FlagBuilder) setEncounter(encounter *datatype.Reference) *FlagBuilder {
	b.Data["encounter"] = encounter.ToArray()
	return b
}

func (b *FlagBuilder) setAuthor(author *datatype.Reference) *FlagBuilder {
	b.Data["author"] = author.ToArray()
	return b
}

func (b *FlagBuilder) addExtension(url string, value interface{}, valueType string) *FlagBuilder {
	ext := map[string]interface{}{"url": url}
	if valueType != "" {
		capitalized := strings.ToUpper(valueType[:1]) + valueType[1:]
		ext["value"+capitalized] = value
	} else {
		ext["valueString"] = value
	}
	if _, ok := b.Data["extension"]; !ok {
		b.Data["extension"] = make([]interface{}, 0)
	}
	b.Data["extension"] = append(b.Data["extension"].([]interface{}), ext)
	return b
}

func (b *FlagBuilder) Build() map[string]interface{} {
	return b.Data
}
