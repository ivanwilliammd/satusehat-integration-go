package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// OrganizationBuilder builds a FHIR R4 Organization resource.
//
// Method-for-method counterpart of PayloadBuilderOrganization in the PHP SDK. setType pushes into a
// list because FHIR declares organization.type as 0..*; the PHP comment says as much, so the name is
// misleading but the shape is right. contact wraps its telecom in a list and its purpose/name in
// {text: …}.
//
// @link https://www.hl7.org/fhir/organization.html
type OrganizationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewOrganizationBuilder creates an empty Organization with its resourceType set.
func NewOrganizationBuilder() *OrganizationBuilder {
	b := &OrganizationBuilder{ResourceType: "Organization", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *OrganizationBuilder) set(k string, v interface{}) *OrganizationBuilder {
	b.Data[k] = v
	return b
}

func (b *OrganizationBuilder) push(k string, v interface{}) *OrganizationBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *OrganizationBuilder) SetId(id string) *OrganizationBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *OrganizationBuilder) AddIdentifier(identifier *datatype.Identifier) *OrganizationBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetActive sets active.
func (b *OrganizationBuilder) SetActive(active bool) *OrganizationBuilder {
	return b.set("active", active)
}

// SetName sets name.
func (b *OrganizationBuilder) SetName(name string) *OrganizationBuilder { return b.set("name", name) }

// AddAlias appends an alias.
func (b *OrganizationBuilder) AddAlias(alias string) *OrganizationBuilder {
	return b.push("alias", alias)
}

// SetType appends a type; the element is 0..* in FHIR, hence the list.
func (b *OrganizationBuilder) SetType(typ *datatype.CodeableConcept) *OrganizationBuilder {
	return b.push("type", typ.ToArray())
}

// AddTelecom appends a telecom.
func (b *OrganizationBuilder) AddTelecom(telecom *datatype.ContactPoint) *OrganizationBuilder {
	return b.push("telecom", telecom.ToArray())
}

// AddAddress appends an address.
func (b *OrganizationBuilder) AddAddress(address *datatype.Address) *OrganizationBuilder {
	return b.push("address", address.ToArray())
}

// SetPartOf sets partOf.
func (b *OrganizationBuilder) SetPartOf(partOf *datatype.Reference) *OrganizationBuilder {
	return b.set("partOf", partOf.ToArray())
}

// AddContact appends a contact. purpose, name and address are optional.
func (b *OrganizationBuilder) AddContact(telecom *datatype.ContactPoint, purpose, name string, address *datatype.Address) *OrganizationBuilder {
	contact := map[string]interface{}{
		"telecom": []interface{}{telecom.ToArray()},
	}
	if purpose != "" {
		contact["purpose"] = map[string]interface{}{"text": purpose}
	}
	if name != "" {
		contact["name"] = map[string]interface{}{"text": name}
	}
	if address != nil {
		contact["address"] = address.ToArray()
	}
	return b.push("contact", contact)
}

// AddEndpoint appends an endpoint reference.
func (b *OrganizationBuilder) AddEndpoint(endpoint *datatype.Reference) *OrganizationBuilder {
	return b.push("endpoint", endpoint.ToArray())
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *OrganizationBuilder) AddExtension(url string, value interface{}, valueType string) *OrganizationBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *OrganizationBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *OrganizationBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
