package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// PractitionerRoleBuilder builds FHIR PractitionerRole payload
type PractitionerRoleBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewPractitionerRoleBuilder() *PractitionerRoleBuilder {
	b := &PractitionerRoleBuilder{ResourceType: "PractitionerRole", Data: make(map[string]interface{})}
	return b
}

func (b *PractitionerRoleBuilder) setId(id string) *PractitionerRoleBuilder {
	b.Data["id"] = id
	return b
}

func (b *PractitionerRoleBuilder) addIdentifier(identifier *datatype.Identifier) *PractitionerRoleBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *PractitionerRoleBuilder) setActive(active bool) *PractitionerRoleBuilder {
	b.Data["active"] = active
	return b
}

func (b *PractitionerRoleBuilder) setPeriod(period *datatype.Period) *PractitionerRoleBuilder {
	b.Data["period"] = period.ToArray()
	return b
}

func (b *PractitionerRoleBuilder) setPractitioner(practitioner *datatype.Reference) *PractitionerRoleBuilder {
	b.Data["practitioner"] = practitioner.ToArray()
	return b
}

func (b *PractitionerRoleBuilder) setOrganization(organization *datatype.Reference) *PractitionerRoleBuilder {
	b.Data["organization"] = organization.ToArray()
	return b
}

func (b *PractitionerRoleBuilder) addCode(code *datatype.CodeableConcept) *PractitionerRoleBuilder {
	if _, ok := b.Data["code"]; !ok {
		b.Data["code"] = make([]interface{}, 0)
	}
	b.Data["code"] = append(b.Data["code"].([]interface{}), code.ToArray())
	return b
}

func (b *PractitionerRoleBuilder) addSpecialty(specialty *datatype.CodeableConcept) *PractitionerRoleBuilder {
	if _, ok := b.Data["specialty"]; !ok {
		b.Data["specialty"] = make([]interface{}, 0)
	}
	b.Data["specialty"] = append(b.Data["specialty"].([]interface{}), specialty.ToArray())
	return b
}

func (b *PractitionerRoleBuilder) addLocation(location *datatype.Reference) *PractitionerRoleBuilder {
	if _, ok := b.Data["location"]; !ok {
		b.Data["location"] = make([]interface{}, 0)
	}
	b.Data["location"] = append(b.Data["location"].([]interface{}), location.ToArray())
	return b
}

func (b *PractitionerRoleBuilder) addHealthcareService(service *datatype.Reference) *PractitionerRoleBuilder {
	if _, ok := b.Data["healthcareService"]; !ok {
		b.Data["healthcareService"] = make([]interface{}, 0)
	}
	b.Data["healthcareService"] = append(b.Data["healthcareService"].([]interface{}), service.ToArray())
	return b
}

func (b *PractitionerRoleBuilder) addTelecom(telecom *datatype.ContactPoint) *PractitionerRoleBuilder {
	if _, ok := b.Data["telecom"]; !ok {
		b.Data["telecom"] = make([]interface{}, 0)
	}
	b.Data["telecom"] = append(b.Data["telecom"].([]interface{}), telecom.ToArray())
	return b
}

func (b *PractitionerRoleBuilder) addAddress(address *datatype.Address) *PractitionerRoleBuilder {
	if _, ok := b.Data["address"]; !ok {
		b.Data["address"] = make([]interface{}, 0)
	}
	b.Data["address"] = append(b.Data["address"].([]interface{}), address.ToArray())
	return b
}

func (b *PractitionerRoleBuilder) addEndpoint(endpoint *datatype.Reference) *PractitionerRoleBuilder {
	if _, ok := b.Data["endpoint"]; !ok {
		b.Data["endpoint"] = make([]interface{}, 0)
	}
	b.Data["endpoint"] = append(b.Data["endpoint"].([]interface{}), endpoint.ToArray())
	return b
}

func (b *PractitionerRoleBuilder) addExtension(url string, value interface{}, valueType string) *PractitionerRoleBuilder {
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

func (b *PractitionerRoleBuilder) Build() map[string]interface{} {
	return b.Data
}
