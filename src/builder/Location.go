package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// LocationBuilder builds FHIR Location payload
type LocationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewLocationBuilder() *LocationBuilder {
	b := &LocationBuilder{ResourceType: "Location", Data: make(map[string]interface{})}
	return b
}

func (b *LocationBuilder) setId(id string) *LocationBuilder {
	b.Data["id"] = id
	return b
}

func (b *LocationBuilder) addIdentifier(identifier *datatype.Identifier) *LocationBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *LocationBuilder) setStatus(status string) *LocationBuilder {
	b.Data["status"] = status
	return b
}

func (b *LocationBuilder) setOperationalStatus(operationalStatus *datatype.CodeableConcept) *LocationBuilder {
	b.Data["operationalStatus"] = operationalStatus.ToArray()
	return b
}

func (b *LocationBuilder) setName(name string) *LocationBuilder {
	b.Data["name"] = name
	return b
}

func (b *LocationBuilder) addAlias(alias string) *LocationBuilder {
	if _, ok := b.Data["alias"]; !ok {
		b.Data["alias"] = make([]interface{}, 0)
	}
	b.Data["alias"] = append(b.Data["alias"].([]interface{}), alias)
	return b
}

func (b *LocationBuilder) setDescription(description string) *LocationBuilder {
	b.Data["description"] = description
	return b
}

func (b *LocationBuilder) setType(typeCode *datatype.CodeableConcept) *LocationBuilder {
	b.Data["type"] = typeCode.ToArray()
	return b
}

func (b *LocationBuilder) addTelecom(telecom *datatype.ContactPoint) *LocationBuilder {
	if _, ok := b.Data["telecom"]; !ok {
		b.Data["telecom"] = make([]interface{}, 0)
	}
	b.Data["telecom"] = append(b.Data["telecom"].([]interface{}), telecom.ToArray())
	return b
}

func (b *LocationBuilder) addAddress(address *datatype.Address) *LocationBuilder {
	if _, ok := b.Data["address"]; !ok {
		b.Data["address"] = make([]interface{}, 0)
	}
	b.Data["address"] = append(b.Data["address"].([]interface{}), address.ToArray())
	return b
}

func (b *LocationBuilder) setPhysicalType(physicalType *datatype.CodeableConcept) *LocationBuilder {
	b.Data["physicalType"] = physicalType.ToArray()
	return b
}

func (b *LocationBuilder) setPosition(latitude, longitude, altitude *float64) *LocationBuilder {
	position := make(map[string]interface{})
	if latitude != nil {
		position["latitude"] = *latitude
	}
	if longitude != nil {
		position["longitude"] = *longitude
	}
	if altitude != nil {
		position["altitude"] = *altitude
	}
	b.Data["position"] = position
	return b
}

func (b *LocationBuilder) setManagingOrganization(managingOrganization *datatype.Reference) *LocationBuilder {
	b.Data["managingOrganization"] = managingOrganization.ToArray()
	return b
}

func (b *LocationBuilder) setPartOf(partOf *datatype.Reference) *LocationBuilder {
	b.Data["partOf"] = partOf.ToArray()
	return b
}

func (b *LocationBuilder) addEndpoint(endpoint *datatype.Reference) *LocationBuilder {
	if _, ok := b.Data["endpoint"]; !ok {
		b.Data["endpoint"] = make([]interface{}, 0)
	}
	b.Data["endpoint"] = append(b.Data["endpoint"].([]interface{}), endpoint.ToArray())
	return b
}

func (b *LocationBuilder) addExtension(url string, value interface{}, valueType string) *LocationBuilder {
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

func (b *LocationBuilder) Build() map[string]interface{} {
	return b.Data
}
