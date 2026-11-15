package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// PatientBuilder builds a FHIR R4 Patient resource.
//
// Field-for-field counterpart of PayloadBuilderPatient in the PHP SDK: the same setters, the same
// element names, so a payload built here matches one built in PHP. Elements are only written when
// a setter is called; the builder never invents values.
//
// @link https://www.hl7.org/fhir/patient.html
type PatientBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewPatientBuilder creates an empty Patient with its resourceType set.
func NewPatientBuilder() *PatientBuilder {
	b := &PatientBuilder{ResourceType: "Patient", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *PatientBuilder) set(k string, v interface{}) *PatientBuilder {
	b.Data[k] = v
	return b
}

func (b *PatientBuilder) push(k string, v interface{}) *PatientBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetMetaProfile sets meta.profile.
func (b *PatientBuilder) SetMetaProfile(profile string) *PatientBuilder {
	meta, _ := b.Data["meta"].(map[string]interface{})
	if meta == nil {
		meta = map[string]interface{}{}
		b.Data["meta"] = meta
	}
	list, _ := meta["profile"].([]interface{})
	meta["profile"] = append(list, profile)
	return b
}

// SetId sets the logical id.
func (b *PatientBuilder) SetId(id string) *PatientBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *PatientBuilder) AddIdentifier(identifier *datatype.Identifier) *PatientBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetActive sets active.
func (b *PatientBuilder) SetActive(active bool) *PatientBuilder { return b.set("active", active) }

// AddName appends a human name.
func (b *PatientBuilder) AddName(name *datatype.HumanName) *PatientBuilder {
	return b.push("name", name.ToArray())
}

// SetGender sets gender (male | female | other | unknown).
func (b *PatientBuilder) SetGender(gender string) *PatientBuilder { return b.set("gender", gender) }

// SetBirthDate sets birthDate (YYYY-MM-DD).
func (b *PatientBuilder) SetBirthDate(birthDate string) *PatientBuilder {
	return b.set("birthDate", birthDate)
}

// SetDeceasedBoolean sets deceasedBoolean.
func (b *PatientBuilder) SetDeceasedBoolean(deceased bool) *PatientBuilder {
	return b.set("deceasedBoolean", deceased)
}

// SetDeceasedDateTime sets deceasedDateTime.
func (b *PatientBuilder) SetDeceasedDateTime(dateTime string) *PatientBuilder {
	return b.set("deceasedDateTime", dateTime)
}

// SetMultipleBirthBoolean sets multipleBirthBoolean.
func (b *PatientBuilder) SetMultipleBirthBoolean(multiple bool) *PatientBuilder {
	return b.set("multipleBirthBoolean", multiple)
}

// SetMultipleBirthInteger sets multipleBirthInteger.
func (b *PatientBuilder) SetMultipleBirthInteger(multiple int) *PatientBuilder {
	return b.set("multipleBirthInteger", multiple)
}

// AddAddress appends an address.
func (b *PatientBuilder) AddAddress(address *datatype.Address) *PatientBuilder {
	return b.push("address", address.ToArray())
}

// AddTelecom appends a contact point.
func (b *PatientBuilder) AddTelecom(telecom *datatype.ContactPoint) *PatientBuilder {
	return b.push("telecom", telecom.ToArray())
}

// SetMaritalStatus sets maritalStatus.
func (b *PatientBuilder) SetMaritalStatus(status *datatype.CodeableConcept) *PatientBuilder {
	return b.set("maritalStatus", status.ToArray())
}

// AddCommunication appends a communication entry.
func (b *PatientBuilder) AddCommunication(language *datatype.CodeableConcept, preferred bool) *PatientBuilder {
	return b.push("communication", map[string]interface{}{
		"language":  language.ToArray(),
		"preferred": preferred,
	})
}

// AddContact appends an emergency contact. Address and organization are optional; pass nil to omit
// them, mirroring the PHP signature.
func (b *PatientBuilder) AddContact(
	relationship *datatype.CodeableConcept,
	name *datatype.HumanName,
	telecom *datatype.ContactPoint,
	address *datatype.Address,
	organization *datatype.Reference,
) *PatientBuilder {
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
	return b.push("contact", contact)
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *PatientBuilder) AddExtension(url string, value interface{}, valueType string) *PatientBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *PatientBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *PatientBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
