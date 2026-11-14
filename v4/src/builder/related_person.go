package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// RelatedPersonBuilder builds a FHIR R4 RelatedPerson resource.
//
// Counterpart of PayloadBuilderRelatedPerson in the PHP SDK. Most setters there are union typed
// (datatype or plain string), so this port takes interface{} and switches, like refArray elsewhere.
//
// PHP quirks reproduced on purpose, because they change the JSON:
//   - AddRelationship with a bare code writes display unconditionally, so an omitted display becomes
//     "display": null;
//   - AddName with a bare string forces use=official;
//   - AddTelecom with a bare system defaults use=home;
//   - AddExtension dispatches on the Go/PHP type of the value (bool, string, int, map).
//
// @link https://www.hl7.org/fhir/relatedperson.html
type RelatedPersonBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewRelatedPersonBuilder creates an empty RelatedPerson with its resourceType set.
func NewRelatedPersonBuilder() *RelatedPersonBuilder {
	b := &RelatedPersonBuilder{ResourceType: "RelatedPerson", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *RelatedPersonBuilder) set(k string, v interface{}) *RelatedPersonBuilder {
	b.Data[k] = v
	return b
}

func (b *RelatedPersonBuilder) push(k string, v interface{}) *RelatedPersonBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

func refPrefixed(typ, v string) string {
	if absolut(v) || strings.Contains(v, "/") {
		return v
	}
	return typ + "/" + v
}

// SetId sets the logical id.
func (b *RelatedPersonBuilder) SetId(id string) *RelatedPersonBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier, from a datatype or from system + value.
func (b *RelatedPersonBuilder) AddIdentifier(identifier interface{}, value *string) *RelatedPersonBuilder {
	if id, ok := identifier.(*datatype.Identifier); ok {
		return b.push("identifier", id.ToArray())
	}
	if system, ok := identifier.(string); ok {
		m := map[string]interface{}{"system": system, "value": nil}
		if value != nil {
			m["value"] = *value
		}
		return b.push("identifier", m)
	}
	return b
}

// SetActive sets active.
func (b *RelatedPersonBuilder) SetActive(active bool) *RelatedPersonBuilder {
	return b.set("active", active)
}

// SetPatient sets patient, from a datatype or from an id/reference. display is optional.
func (b *RelatedPersonBuilder) SetPatient(patient interface{}, display *string) *RelatedPersonBuilder {
	if ref, ok := patient.(*datatype.Reference); ok {
		return b.set("patient", ref.ToArray())
	}
	if v, ok := patient.(string); ok {
		m := map[string]interface{}{"reference": refPrefixed("Patient", v)}
		if display != nil {
			m["display"] = *display
		}
		return b.set("patient", m)
	}
	return b
}

// AddRelationship appends a relationship, from a datatype or a bare code in the v2-0131 system.
// A nil display is written as null, as PHP does.
func (b *RelatedPersonBuilder) AddRelationship(relationship interface{}, display *string) *RelatedPersonBuilder {
	if cc, ok := relationship.(*datatype.CodeableConcept); ok {
		return b.push("relationship", cc.ToArray())
	}
	if code, ok := relationship.(string); ok {
		var tampil interface{}
		if display != nil {
			tampil = *display
		}
		return b.push("relationship", map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{
				"system":  "http://terminology.hl7.org/CodeSystem/v2-0131",
				"code":    code,
				"display": tampil,
			}},
		})
	}
	return b
}

// AddName appends a name. A bare string becomes {use: official, text: …}, as PHP does.
func (b *RelatedPersonBuilder) AddName(name interface{}) *RelatedPersonBuilder {
	if hn, ok := name.(*datatype.HumanName); ok {
		return b.push("name", hn.ToArray())
	}
	if text, ok := name.(string); ok {
		return b.push("name", map[string]interface{}{"use": "official", "text": text})
	}
	return b
}

// AddTelecom appends a telecom. A bare system defaults use to "home"; an empty use means home.
func (b *RelatedPersonBuilder) AddTelecom(system interface{}, value *string, use string) *RelatedPersonBuilder {
	if cp, ok := system.(*datatype.ContactPoint); ok {
		return b.push("telecom", cp.ToArray())
	}
	if sys, ok := system.(string); ok {
		if use == "" {
			use = "home"
		}
		m := map[string]interface{}{"system": sys, "use": use}
		if value != nil {
			m["value"] = *value
		}
		return b.push("telecom", m)
	}
	return b
}

// AddAddress appends an address, from a datatype or a ready-made map.
func (b *RelatedPersonBuilder) AddAddress(address interface{}) *RelatedPersonBuilder {
	switch v := address.(type) {
	case *datatype.Address:
		return b.push("address", v.ToArray())
	case map[string]interface{}:
		return b.push("address", v)
	}
	return b
}

// AddCommunication appends a communication. A bare language code uses urn:ietf:bcp:47, and preferred
// defaults to true, as in PHP.
func (b *RelatedPersonBuilder) AddCommunication(language interface{}, preferred *bool) *RelatedPersonBuilder {
	comm := map[string]interface{}{}
	switch v := language.(type) {
	case *datatype.CodeableConcept:
		comm["language"] = v.ToArray()
	case string:
		comm["language"] = map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": "urn:ietf:bcp:47", "code": v}},
		}
	default:
		return b
	}
	if preferred != nil {
		comm["preferred"] = *preferred
	} else {
		comm["preferred"] = true
	}
	return b.push("communication", comm)
}

// SetGender sets gender.
func (b *RelatedPersonBuilder) SetGender(gender string) *RelatedPersonBuilder {
	return b.set("gender", gender)
}

// SetBirthDate sets birthDate.
func (b *RelatedPersonBuilder) SetBirthDate(birthDate string) *RelatedPersonBuilder {
	return b.set("birthDate", birthDate)
}

// AddExtension appends an extension, choosing the value[x] element from the Go type of value.
func (b *RelatedPersonBuilder) AddExtension(url string, value interface{}) *RelatedPersonBuilder {
	extension := map[string]interface{}{"url": url}
	switch v := value.(type) {
	case bool:
		extension["valueBoolean"] = v
	case string:
		extension["valueString"] = v
	case int:
		extension["valueInteger"] = v
	case map[string]interface{}:
		for k, isi := range v {
			extension[k] = isi
		}
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *RelatedPersonBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *RelatedPersonBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
