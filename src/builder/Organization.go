package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// OrganizationBuilder builds FHIR Organization payload
type OrganizationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewOrganizationBuilder() *OrganizationBuilder {
	b := &OrganizationBuilder{ResourceType: "Organization", Data: make(map[string]interface{})}
	return b
}

func (b *OrganizationBuilder) setId(id string) *OrganizationBuilder {
	b.Data["id"] = id
	return b
}

func (b *OrganizationBuilder) addIdentifier(identifier *datatype.Identifier) *OrganizationBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *OrganizationBuilder) setActive(active bool) *OrganizationBuilder {
	b.Data["active"] = active
	return b
}

func (b *OrganizationBuilder) setName(name string) *OrganizationBuilder {
	b.Data["name"] = name
	return b
}

func (b *OrganizationBuilder) addAlias(alias string) *OrganizationBuilder {
	if _, ok := b.Data["alias"]; !ok {
		b.Data["alias"] = make([]interface{}, 0)
	}
	b.Data["alias"] = append(b.Data["alias"].([]interface{}), alias)
	return b
}

func (b *OrganizationBuilder) setType(typeCode *datatype.CodeableConcept) *OrganizationBuilder {
	b.Data["type"] = typeCode.ToArray()
	return b
}

func (b *OrganizationBuilder) addTelecom(telecom *datatype.ContactPoint) *OrganizationBuilder {
	if _, ok := b.Data["telecom"]; !ok {
		b.Data["telecom"] = make([]interface{}, 0)
	}
	b.Data["telecom"] = append(b.Data["telecom"].([]interface{}), telecom.ToArray())
	return b
}

func (b *OrganizationBuilder) addAddress(address *datatype.Address) *OrganizationBuilder {
	if _, ok := b.Data["address"]; !ok {
		b.Data["address"] = make([]interface{}, 0)
	}
	b.Data["address"] = append(b.Data["address"].([]interface{}), address.ToArray())
	return b
}

func (b *OrganizationBuilder) setPartOf(partOf *datatype.Reference) *OrganizationBuilder {
	b.Data["partOf"] = partOf.ToArray()
	return b
}

func (b *OrganizationBuilder) addContact(
	telecom *datatype.ContactPoint,
	purpose *string,
	name *string,
	address *datatype.Address,
) *OrganizationBuilder {
	contact := map[string]interface{}{
		"telecom": []interface{}{telecom.ToArray()},
	}
	if purpose != nil {
		contact["purpose"] = map[string]interface{}{"text": *purpose}
	}
	if name != nil {
		contact["name"] = map[string]interface{}{"text": *name}
	}
	if address != nil {
		contact["address"] = address.ToArray()
	}
	if _, ok := b.Data["contact"]; !ok {
		b.Data["contact"] = make([]interface{}, 0)
	}
	b.Data["contact"] = append(b.Data["contact"].([]interface{}), contact)
	return b
}

func (b *OrganizationBuilder) addEndpoint(endpoint *datatype.Reference) *OrganizationBuilder {
	if _, ok := b.Data["endpoint"]; !ok {
		b.Data["endpoint"] = make([]interface{}, 0)
	}
	b.Data["endpoint"] = append(b.Data["endpoint"].([]interface{}), endpoint.ToArray())
	return b
}

func (b *OrganizationBuilder) addExtension(url string, value interface{}, valueType string) *OrganizationBuilder {
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

func (b *OrganizationBuilder) Build() map[string]interface{} {
	return b.Data
}
