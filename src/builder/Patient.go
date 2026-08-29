package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// PatientBuilder builds FHIR Patient payload
type PatientBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewPatientBuilder() *PatientBuilder {
	b := &PatientBuilder{ResourceType: "Patient", Data: make(map[string]interface{})}
	return b
}

func (b *PatientBuilder) setMetaProfile(profile string) *PatientBuilder {
	if _, ok := b.Data["meta"]; !ok {
		b.Data["meta"] = make(map[string]interface{})
	}
	b.Data["meta"].(map[string]interface{})["profile"] = profile
	return b
}

func (b *PatientBuilder) setId(id string) *PatientBuilder {
	b.Data["id"] = id
	return b
}

func (b *PatientBuilder) addIdentifier(identifier *datatype.Identifier) *PatientBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *PatientBuilder) setActive(active bool) *PatientBuilder {
	b.Data["active"] = active
	return b
}

func (b *PatientBuilder) addName(name *datatype.HumanName) *PatientBuilder {
	if _, ok := b.Data["name"]; !ok {
		b.Data["name"] = make([]interface{}, 0)
	}
	b.Data["name"] = append(b.Data["name"].([]interface{}), name.ToArray())
	return b
}

func (b *PatientBuilder) setGender(gender string) *PatientBuilder {
	b.Data["gender"] = gender
	return b
}

func (b *PatientBuilder) setBirthDate(birthDate string) *PatientBuilder {
	b.Data["birthDate"] = birthDate
	return b
}

func (b *PatientBuilder) setDeceasedBoolean(deceased bool) *PatientBuilder {
	b.Data["deceasedBoolean"] = deceased
	return b
}

func (b *PatientBuilder) setDeceasedDateTime(dateTime string) *PatientBuilder {
	b.Data["deceasedDateTime"] = dateTime
	return b
}

func (b *PatientBuilder) setMultipleBirthBoolean(multipleBirth bool) *PatientBuilder {
	b.Data["multipleBirthBoolean"] = multipleBirth
	return b
}

func (b *PatientBuilder) setMultipleBirthInteger(multipleBirth int) *PatientBuilder {
	b.Data["multipleBirthInteger"] = multipleBirth
	return b
}

func (b *PatientBuilder) addAddress(address *datatype.Address) *PatientBuilder {
	if _, ok := b.Data["address"]; !ok {
		b.Data["address"] = make([]interface{}, 0)
	}
	b.Data["address"] = append(b.Data["address"].([]interface{}), address.ToArray())
	return b
}

func (b *PatientBuilder) addTelecom(telecom *datatype.ContactPoint) *PatientBuilder {
	if _, ok := b.Data["telecom"]; !ok {
		b.Data["telecom"] = make([]interface{}, 0)
	}
	b.Data["telecom"] = append(b.Data["telecom"].([]interface{}), telecom.ToArray())
	return b
}

func (b *PatientBuilder) setMaritalStatus(maritalStatus *datatype.CodeableConcept) *PatientBuilder {
	b.Data["maritalStatus"] = maritalStatus.ToArray()
	return b
}

func (b *PatientBuilder) addCommunication(language *datatype.CodeableConcept, preferred bool) *PatientBuilder {
	if _, ok := b.Data["communication"]; !ok {
		b.Data["communication"] = make([]interface{}, 0)
	}
	b.Data["communication"] = append(b.Data["communication"].([]interface{}), map[string]interface{}{
		"language":  language.ToArray(),
		"preferred": preferred,
	})
	return b
}

func (b *PatientBuilder) addContact(relationship *datatype.CodeableConcept, name *datatype.HumanName, telecom *datatype.ContactPoint, address *datatype.Address, organization *datatype.Reference) *PatientBuilder {
 contact := map[string]interface{}{
  "relationship": []interface{}{relationship.ToArray()},
  "name":         name.ToArray(),
  "telecom":      []interface{}{telecom.ToArray()},
 }
 if address != nil {
  contact["address"] = address.ToArray()
 }
 if organization != nil {
  contact["organization"] = organization.ToArray()
 }
 if _, ok := b.Data["contact"]; !ok {
  b.Data["contact"] = make([]interface{}, 0)
 }
 b.Data["contact"] = append(b.Data["contact"].([]interface{}), contact)
 return b
}

func (b *PatientBuilder) addExtension(url string, value interface{}, valueType string) *PatientBuilder {
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
