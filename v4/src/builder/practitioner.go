package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// PractitionerBuilder builds a FHIR R4 Practitioner resource.
//
// Method-for-method counterpart of PayloadBuilderPractitioner in the PHP SDK: name, telecom and
// address are lists, qualification carries its identifier as a list, and photo takes a bare url.
//
// @link https://www.hl7.org/fhir/practitioner.html
type PractitionerBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewPractitionerBuilder creates an empty Practitioner with its resourceType set.
func NewPractitionerBuilder() *PractitionerBuilder {
	b := &PractitionerBuilder{ResourceType: "Practitioner", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *PractitionerBuilder) set(k string, v interface{}) *PractitionerBuilder {
	b.Data[k] = v
	return b
}

func (b *PractitionerBuilder) push(k string, v interface{}) *PractitionerBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *PractitionerBuilder) SetId(id string) *PractitionerBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *PractitionerBuilder) AddIdentifier(identifier *datatype.Identifier) *PractitionerBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetActive sets active.
func (b *PractitionerBuilder) SetActive(active bool) *PractitionerBuilder {
	return b.set("active", active)
}

// AddName appends a name.
func (b *PractitionerBuilder) AddName(name *datatype.HumanName) *PractitionerBuilder {
	return b.push("name", name.ToArray())
}

// AddTelecom appends a telecom.
func (b *PractitionerBuilder) AddTelecom(telecom *datatype.ContactPoint) *PractitionerBuilder {
	return b.push("telecom", telecom.ToArray())
}

// AddAddress appends an address.
func (b *PractitionerBuilder) AddAddress(address *datatype.Address) *PractitionerBuilder {
	return b.push("address", address.ToArray())
}

// SetGender sets gender (male | female | other | unknown).
func (b *PractitionerBuilder) SetGender(gender string) *PractitionerBuilder {
	return b.set("gender", gender)
}

// SetBirthDate sets birthDate.
func (b *PractitionerBuilder) SetBirthDate(birthDate string) *PractitionerBuilder {
	return b.set("birthDate", birthDate)
}

// AddPhoto appends a photo. contentType is optional.
func (b *PractitionerBuilder) AddPhoto(url, contentType string) *PractitionerBuilder {
	photo := map[string]interface{}{"url": url}
	if contentType != "" {
		photo["contentType"] = contentType
	}
	return b.push("photo", photo)
}

// AddQualification appends a qualification. periodStart and issuer are optional; the identifier is
// written as a list, as PHP does.
func (b *PractitionerBuilder) AddQualification(identifier *datatype.Identifier, code *datatype.CodeableConcept, periodStart string, issuer *datatype.Reference) *PractitionerBuilder {
	qualification := map[string]interface{}{
		"identifier": []interface{}{identifier.ToArray()},
		"code":       code.ToArray(),
	}
	if periodStart != "" {
		qualification["period"] = map[string]interface{}{"start": periodStart}
	}
	if issuer != nil {
		qualification["issuer"] = issuer.ToArray()
	}
	return b.push("qualification", qualification)
}

// AddCommunication appends a communication. preferred is optional.
func (b *PractitionerBuilder) AddCommunication(language *datatype.CodeableConcept, preferred *bool) *PractitionerBuilder {
	communication := map[string]interface{}{"language": language.ToArray()}
	if preferred != nil {
		communication["preferred"] = *preferred
	}
	return b.push("communication", communication)
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *PractitionerBuilder) AddExtension(url string, value interface{}, valueType string) *PractitionerBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *PractitionerBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *PractitionerBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
