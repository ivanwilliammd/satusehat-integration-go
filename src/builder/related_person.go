package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// RelatedPersonBuilder builds FHIR RelatedPerson payload
type RelatedPersonBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewRelatedPersonBuilder() *RelatedPersonBuilder {
	b := &RelatedPersonBuilder{ResourceType: "RelatedPerson", Data: make(map[string]interface{})}
	return b
}

func (b *RelatedPersonBuilder) setId(id string) *RelatedPersonBuilder {
	b.Data["id"] = id
	return b
}

func (b *RelatedPersonBuilder) addIdentifier(identifier *datatype.Identifier) *RelatedPersonBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *RelatedPersonBuilder) setActive(active bool) *RelatedPersonBuilder {
	b.Data["active"] = active
	return b
}

func (b *RelatedPersonBuilder) addName(name *datatype.HumanName) *RelatedPersonBuilder {
	if _, ok := b.Data["name"]; !ok {
		b.Data["name"] = make([]interface{}, 0)
	}
	b.Data["name"] = append(b.Data["name"].([]interface{}), name.ToArray())
	return b
}

func (b *RelatedPersonBuilder) setGender(gender string) *RelatedPersonBuilder {
	b.Data["gender"] = gender
	return b
}

func (b *RelatedPersonBuilder) setBirthDate(birthDate string) *RelatedPersonBuilder {
	b.Data["birthDate"] = birthDate
	return b
}

func (b *RelatedPersonBuilder) addAddress(address *datatype.Address) *RelatedPersonBuilder {
	if _, ok := b.Data["address"]; !ok {
		b.Data["address"] = make([]interface{}, 0)
	}
	b.Data["address"] = append(b.Data["address"].([]interface{}), address.ToArray())
	return b
}

func (b *RelatedPersonBuilder) addTelecom(telecom *datatype.ContactPoint) *RelatedPersonBuilder {
	if _, ok := b.Data["telecom"]; !ok {
		b.Data["telecom"] = make([]interface{}, 0)
	}
	b.Data["telecom"] = append(b.Data["telecom"].([]interface{}), telecom.ToArray())
	return b
}

func (b *RelatedPersonBuilder) setPatient(patient *datatype.Reference) *RelatedPersonBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}

func (b *RelatedPersonBuilder) addRelationship(relationship *datatype.CodeableConcept) *RelatedPersonBuilder {
	if _, ok := b.Data["relationship"]; !ok {
		b.Data["relationship"] = make([]interface{}, 0)
	}
	b.Data["relationship"] = append(b.Data["relationship"].([]interface{}), relationship.ToArray())
	return b
}

func (b *RelatedPersonBuilder) setPeriod(period *datatype.Period) *RelatedPersonBuilder {
	b.Data["period"] = period.ToArray()
	return b
}

func (b *RelatedPersonBuilder) addExtension(url string, value interface{}, valueType string) *RelatedPersonBuilder {
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

func (b *RelatedPersonBuilder) Build() map[string]interface{} {
	return b.Data
}
