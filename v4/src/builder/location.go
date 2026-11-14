package builder

import (
	"encoding/json"
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// LocationBuilder builds a FHIR R4 Location resource.
//
// Method-for-method counterpart of PayloadBuilderLocation in the PHP SDK. setType pushes into a list
// (location.type is 0..*), and setPosition always writes the position object, even when all three
// arguments are nil — PHP does the same, so an empty {} is expected rather than an omitted element.
//
// @link https://www.hl7.org/fhir/location.html
type LocationBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewLocationBuilder creates an empty Location with its resourceType set.
func NewLocationBuilder() *LocationBuilder {
	b := &LocationBuilder{ResourceType: "Location", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *LocationBuilder) set(k string, v interface{}) *LocationBuilder {
	b.Data[k] = v
	return b
}

func (b *LocationBuilder) push(k string, v interface{}) *LocationBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *LocationBuilder) SetId(id string) *LocationBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *LocationBuilder) AddIdentifier(identifier *datatype.Identifier) *LocationBuilder {
	return b.push("identifier", identifier.ToArray())
}

// SetStatus sets status (active | suspended | inactive).
func (b *LocationBuilder) SetStatus(status string) *LocationBuilder { return b.set("status", status) }

// SetOperationalStatus sets operationalStatus.
func (b *LocationBuilder) SetOperationalStatus(status *datatype.CodeableConcept) *LocationBuilder {
	return b.set("operationalStatus", status.ToArray())
}

// SetName sets name.
func (b *LocationBuilder) SetName(name string) *LocationBuilder { return b.set("name", name) }

// AddAlias appends an alias.
func (b *LocationBuilder) AddAlias(alias string) *LocationBuilder { return b.push("alias", alias) }

// SetDescription sets description.
func (b *LocationBuilder) SetDescription(description string) *LocationBuilder {
	return b.set("description", description)
}

// SetType appends a type; the element is 0..* in FHIR.
func (b *LocationBuilder) SetType(typ *datatype.CodeableConcept) *LocationBuilder {
	return b.push("type", typ.ToArray())
}

// AddTelecom appends a telecom.
func (b *LocationBuilder) AddTelecom(telecom *datatype.ContactPoint) *LocationBuilder {
	return b.push("telecom", telecom.ToArray())
}

// AddAddress appends an address.
func (b *LocationBuilder) AddAddress(address *datatype.Address) *LocationBuilder {
	return b.push("address", address.ToArray())
}

// SetPhysicalType sets physicalType.
func (b *LocationBuilder) SetPhysicalType(physicalType *datatype.CodeableConcept) *LocationBuilder {
	return b.set("physicalType", physicalType.ToArray())
}

// SetPosition sets position. Nil arguments are omitted, but the object itself is always written.
func (b *LocationBuilder) SetPosition(latitude, longitude, altitude *float64) *LocationBuilder {
	position := map[string]interface{}{}
	if latitude != nil {
		position["latitude"] = *latitude
	}
	if longitude != nil {
		position["longitude"] = *longitude
	}
	if altitude != nil {
		position["altitude"] = *altitude
	}
	return b.set("position", position)
}

// SetManagingOrganization sets managingOrganization.
func (b *LocationBuilder) SetManagingOrganization(organization *datatype.Reference) *LocationBuilder {
	return b.set("managingOrganization", organization.ToArray())
}

// SetPartOf sets partOf.
func (b *LocationBuilder) SetPartOf(partOf *datatype.Reference) *LocationBuilder {
	return b.set("partOf", partOf.ToArray())
}

// AddEndpoint appends an endpoint reference.
func (b *LocationBuilder) AddEndpoint(endpoint *datatype.Reference) *LocationBuilder {
	return b.push("endpoint", endpoint.ToArray())
}

// AddExtension appends an extension. When valueType is empty the value is written as valueString.
func (b *LocationBuilder) AddExtension(url string, value interface{}, valueType string) *LocationBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	return b.push("extension", extension)
}

// Build returns the payload.
func (b *LocationBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *LocationBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
