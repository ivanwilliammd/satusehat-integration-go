package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// PractitionerBuilder builds FHIR Practitioner payload
type PractitionerBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewPractitionerBuilder() *PractitionerBuilder {
	b := &PractitionerBuilder{ResourceType: "Practitioner", Data: make(map[string]interface{})}
	return b
}

func (b *PractitionerBuilder) setId(id string) *PractitionerBuilder {
	b.Data["id"] = id
	return b
}

func (b *PractitionerBuilder) addIdentifier(identifier *datatype.Identifier) *PractitionerBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *PractitionerBuilder) setActive(active bool) *PractitionerBuilder {
	b.Data["active"] = active
	return b
}

func (b *PractitionerBuilder) addName(name *datatype.HumanName) *PractitionerBuilder {
	if _, ok := b.Data["name"]; !ok {
		b.Data["name"] = make([]interface{}, 0)
	}
	b.Data["name"] = append(b.Data["name"].([]interface{}), name.ToArray())
	return b
}

func (b *PractitionerBuilder) addTelecom(telecom *datatype.ContactPoint) *PractitionerBuilder {
	if _, ok := b.Data["telecom"]; !ok {
		b.Data["telecom"] = make([]interface{}, 0)
	}
	b.Data["telecom"] = append(b.Data["telecom"].([]interface{}), telecom.ToArray())
	return b
}

func (b *PractitionerBuilder) addAddress(address *datatype.Address) *PractitionerBuilder {
	if _, ok := b.Data["address"]; !ok {
		b.Data["address"] = make([]interface{}, 0)
	}
	b.Data["address"] = append(b.Data["address"].([]interface{}), address.ToArray())
	return b
}

func (b *PractitionerBuilder) setGender(gender string) *PractitionerBuilder {
	b.Data["gender"] = gender
	return b
}

func (b *PractitionerBuilder) setBirthDate(birthDate string) *PractitionerBuilder {
	b.Data["birthDate"] = birthDate
	return b
}

func (b *PractitionerBuilder) addPhoto(url string, contentType *string) *PractitionerBuilder {
	photo := map[string]interface{}{"url": url}
	if contentType != nil {
		photo["contentType"] = *contentType
	}
	if _, ok := b.Data["photo"]; !ok {
		b.Data["photo"] = make([]interface{}, 0)
	}
	b.Data["photo"] = append(b.Data["photo"].([]interface{}), photo)
	return b
}

func (b *PractitionerBuilder) addQualification(
	identifier *datatype.Identifier,
	code *datatype.CodeableConcept,
	periodStart *string,
	issuer *datatype.Reference,
) *PractitionerBuilder {
	qualification := map[string]interface{}{
		"identifier": []interface{}{identifier.ToArray()},
		"code":       code.ToArray(),
	}
	if periodStart != nil {
		qualification["period"] = map[string]interface{}{"start": *periodStart}
	}
	if issuer != nil {
		qualification["issuer"] = issuer.ToArray()
	}
	if _, ok := b.Data["qualification"]; !ok {
		b.Data["qualification"] = make([]interface{}, 0)
	}
	b.Data["qualification"] = append(b.Data["qualification"].([]interface{}), qualification)
	return b
}

func (b *PractitionerBuilder) addCommunication(language *datatype.CodeableConcept, preferred *bool) *PractitionerBuilder {
	communication := map[string]interface{}{"language": language.ToArray()}
	if preferred != nil {
		communication["preferred"] = *preferred
	}
	if _, ok := b.Data["communication"]; !ok {
		b.Data["communication"] = make([]interface{}, 0)
	}
	b.Data["communication"] = append(b.Data["communication"].([]interface{}), communication)
	return b
}

func (b *PractitionerBuilder) addExtension(url string, value interface{}, valueType string) *PractitionerBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType != "" {
		capitalized := strings.ToUpper(valueType[:1]) + valueType[1:]
		extension["value"+capitalized] = value
	} else {
		extension["valueString"] = value
	}
	if _, ok := b.Data["extension"]; !ok {
		b.Data["extension"] = make([]interface{}, 0)
	}
	b.Data["extension"] = append(b.Data["extension"].([]interface{}), extension)
	return b
}

func (b *PractitionerBuilder) Build() map[string]interface{} {
	return b.Data
}
