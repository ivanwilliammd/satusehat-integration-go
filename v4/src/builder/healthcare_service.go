package builder

import (
	"encoding/json"
)

// HealthcareServiceBuilder builds a FHIR R4 HealthcareService resource.
//
// Counterpart of PayloadBuilderHealthcareService in the PHP SDK, with one deliberate difference:
// AddIdentifier works. The PHP method calls $this->add(...), which Builder does not define, so any
// call to it dies with "Call to undefined method". There is therefore no PHP payload to compare that
// element against; the parity golden covers the remaining setters, and a separate test asserts that
// the identifier is written here.
//
// @link https://www.hl7.org/fhir/healthcareservice.html
type HealthcareServiceBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewHealthcareServiceBuilder creates an empty HealthcareService with its resourceType set.
func NewHealthcareServiceBuilder() *HealthcareServiceBuilder {
	b := &HealthcareServiceBuilder{ResourceType: "HealthcareService", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *HealthcareServiceBuilder) set(k string, v interface{}) *HealthcareServiceBuilder {
	b.Data[k] = v
	return b
}

func (b *HealthcareServiceBuilder) push(k string, v interface{}) *HealthcareServiceBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *HealthcareServiceBuilder) SetId(id string) *HealthcareServiceBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier. use and typeCode are optional; when typeCode is given the
// coding goes into the v2-0203 system and display defaults to the code, as PHP intends.
func (b *HealthcareServiceBuilder) AddIdentifier(system, value, use, typeCode, typeDisplay string) *HealthcareServiceBuilder {
	ident := map[string]interface{}{"system": system, "value": value}
	if use != "" {
		ident["use"] = use
	}
	if typeCode != "" {
		display := typeDisplay
		if display == "" {
			display = typeCode
		}
		ident["type"] = map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{
				"system":  "http://terminology.hl7.org/CodeSystem/v2-0203",
				"code":    typeCode,
				"display": display,
			}},
		}
	}
	return b.push("identifier", ident)
}

// SetStatus sets status.
func (b *HealthcareServiceBuilder) SetStatus(status string) *HealthcareServiceBuilder {
	return b.set("status", status)
}

// SetName sets name.
func (b *HealthcareServiceBuilder) SetName(name string) *HealthcareServiceBuilder {
	return b.set("name", name)
}

// SetType appends a type. Repeated calls add rather than overwrite, because the element is 0..*.
func (b *HealthcareServiceBuilder) SetType(system, code, display string) *HealthcareServiceBuilder {
	return b.push("type", map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{
			"system": system, "code": code, "display": display,
		}},
	})
}

// SetLocation sets location. display is optional.
func (b *HealthcareServiceBuilder) SetLocation(reference string, display *string) *HealthcareServiceBuilder {
	ref := map[string]interface{}{"reference": reference}
	if display != nil {
		ref["display"] = *display
	}
	return b.set("location", ref)
}

// SetProvidedBy sets providedBy. display is optional.
func (b *HealthcareServiceBuilder) SetProvidedBy(reference string, display *string) *HealthcareServiceBuilder {
	ref := map[string]interface{}{"reference": reference}
	if display != nil {
		ref["display"] = *display
	}
	return b.set("providedBy", ref)
}

// Build returns the payload.
func (b *HealthcareServiceBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *HealthcareServiceBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
