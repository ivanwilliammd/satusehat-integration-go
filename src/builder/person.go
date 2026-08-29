package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// PersonBuilder builds FHIR Person payload
type PersonBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewPersonBuilder() *PersonBuilder {
	b := &PersonBuilder{ResourceType: "Person", Data: make(map[string]interface{})}
	return b
}

func (b *PersonBuilder) setId(id string) *PersonBuilder {
	b.Data["id"] = id
	return b
}

func (b *PersonBuilder) addIdentifier(identifier *datatype.Identifier) *PersonBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *PersonBuilder) setActive(active bool) *PersonBuilder {
	b.Data["active"] = active
	return b
}

func (b *PersonBuilder) addName(name *datatype.HumanName) *PersonBuilder {
	if _, ok := b.Data["name"]; !ok {
		b.Data["name"] = make([]interface{}, 0)
	}
	b.Data["name"] = append(b.Data["name"].([]interface{}), name.ToArray())
	return b
}

func (b *PersonBuilder) setGender(gender string) *PersonBuilder {
	b.Data["gender"] = gender
	return b
}

func (b *PersonBuilder) setBirthDate(birthDate string) *PersonBuilder {
	b.Data["birthDate"] = birthDate
	return b
}

func (b *PersonBuilder) addAddress(address *datatype.Address) *PersonBuilder {
	if _, ok := b.Data["address"]; !ok {
		b.Data["address"] = make([]interface{}, 0)
	}
	b.Data["address"] = append(b.Data["address"].([]interface{}), address.ToArray())
	return b
}

func (b *PersonBuilder) addTelecom(telecom *datatype.ContactPoint) *PersonBuilder {
	if _, ok := b.Data["telecom"]; !ok {
		b.Data["telecom"] = make([]interface{}, 0)
	}
	b.Data["telecom"] = append(b.Data["telecom"].([]interface{}), telecom.ToArray())
	return b
}

func (b *PersonBuilder) setLink(target *datatype.Reference, assurance string) *PersonBuilder {
	link := map[string]interface{}{}
	if target != nil {
		link["target"] = target.ToArray()
	}
	if assurance != "" {
		link["assurance"] = assurance
	}
	if _, ok := b.Data["link"]; !ok {
		b.Data["link"] = make([]interface{}, 0)
	}
	b.Data["link"] = append(b.Data["link"].([]interface{}), link)
	return b
}

func (b *PersonBuilder) addExtension(url string, value interface{}, valueType string) *PersonBuilder {
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

func (b *PersonBuilder) Build() map[string]interface{} {
	return b.Data
}
