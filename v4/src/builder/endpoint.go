package builder

import (
	"encoding/json"
	"fmt"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// EndpointBuilder builds a FHIR R4 Endpoint resource.
//
// Method-for-method counterpart of PayloadBuilderEndpoint in the PHP SDK. connectionType and
// payloadType are built from code + display with the HL7 systems baked in, and the status setter
// validates against the six values PHP allows.
//
// @link https://www.hl7.org/fhir/endpoint.html
type EndpointBuilder struct {
	ResourceType string
	Data         map[string]interface{}
	err          error
}

// EndpointStatuses lists the values accepted by SetStatus, mirroring VALID_STATUSES in PHP.
var EndpointStatuses = []string{"active", "suspended", "error", "off", "entered-in-error", "test"}

// NewEndpointBuilder creates an empty Endpoint with its resourceType set.
func NewEndpointBuilder() *EndpointBuilder {
	b := &EndpointBuilder{ResourceType: "Endpoint", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

// Err returns the first validation error recorded by a setter, if any.
func (b *EndpointBuilder) Err() error { return b.err }

func (b *EndpointBuilder) set(k string, v interface{}) *EndpointBuilder {
	b.Data[k] = v
	return b
}

func (b *EndpointBuilder) push(k string, v interface{}) *EndpointBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *EndpointBuilder) SetId(id string) *EndpointBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier, from a datatype or from system + value.
func (b *EndpointBuilder) AddIdentifier(identifier interface{}, value *string) *EndpointBuilder {
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

// SetStatus sets status. Only the six PHP values are accepted; anything else is recorded in Err().
func (b *EndpointBuilder) SetStatus(status string) *EndpointBuilder {
	for _, sah := range EndpointStatuses {
		if status == sah {
			return b.set("status", status)
		}
	}
	if b.err == nil {
		b.err = fmt.Errorf("Invalid status: %s", status)
	}
	return b
}

// SetConnectionType sets connectionType. The endpoint-connection-type system is the default.
func (b *EndpointBuilder) SetConnectionType(code, display, system string) *EndpointBuilder {
	if system == "" {
		system = "http://terminology.hl7.org/CodeSystem/endpoint-connection-type"
	}
	return b.set("connectionType", map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{
			"code": code, "display": display, "system": system,
		}},
	})
}

// SetName sets name.
func (b *EndpointBuilder) SetName(name string) *EndpointBuilder { return b.set("name", name) }

// SetManagingOrganization sets managingOrganization. A bare id is prefixed with Organization/.
func (b *EndpointBuilder) SetManagingOrganization(reference string, display *string) *EndpointBuilder {
	m := map[string]interface{}{"reference": refPrefixed("Organization", reference)}
	if display != nil {
		m["display"] = *display
	}
	return b.set("managingOrganization", m)
}

// AddContact appends a contact. use is optional.
func (b *EndpointBuilder) AddContact(system, value, use string) *EndpointBuilder {
	contact := map[string]interface{}{"system": system, "value": value}
	if use != "" {
		contact["use"] = use
	}
	return b.push("contact", contact)
}

// SetPeriod sets period. end is optional.
func (b *EndpointBuilder) SetPeriod(start, end string) *EndpointBuilder {
	period := map[string]interface{}{"start": start}
	if end != "" {
		period["end"] = end
	}
	return b.set("period", period)
}

// AddPayloadType appends a payloadType. The endpoint-payload-type system is the default.
func (b *EndpointBuilder) AddPayloadType(code, display, system string) *EndpointBuilder {
	if system == "" {
		system = "http://terminology.hl7.org/CodeSystem/endpoint-payload-type"
	}
	return b.push("payloadType", map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{
			"code": code, "display": display, "system": system,
		}},
	})
}

// AddPayloadMimeType appends a payloadMimeType.
func (b *EndpointBuilder) AddPayloadMimeType(mime string) *EndpointBuilder {
	return b.push("payloadMimeType", mime)
}

// SetAddress sets address.
func (b *EndpointBuilder) SetAddress(address string) *EndpointBuilder {
	return b.set("address", address)
}

// AddHeader appends a header.
func (b *EndpointBuilder) AddHeader(header string) *EndpointBuilder { return b.push("header", header) }

// Build returns the payload.
func (b *EndpointBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *EndpointBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
