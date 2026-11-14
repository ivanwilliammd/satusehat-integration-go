package builder

import (
	"encoding/json"
)

// CareTeamBuilder builds a FHIR R4 CareTeam resource.
//
// Counterpart of PayloadBuilderCareTeam in the PHP SDK. AddIdentifier works here; the PHP method is
// broken the same way HealthcareService's is — it calls $this->add(), which Builder does not define,
// so the call dies instead of building a payload. The parity golden therefore covers the remaining
// setters, and a separate test covers the identifier.
//
// @link https://www.hl7.org/fhir/careteam.html
type CareTeamBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewCareTeamBuilder creates an empty CareTeam with its resourceType set.
func NewCareTeamBuilder() *CareTeamBuilder {
	b := &CareTeamBuilder{ResourceType: "CareTeam", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *CareTeamBuilder) set(k string, v interface{}) *CareTeamBuilder {
	b.Data[k] = v
	return b
}

func (b *CareTeamBuilder) push(k string, v interface{}) *CareTeamBuilder {
	list, _ := b.Data[k].([]interface{})
	b.Data[k] = append(list, v)
	return b
}

// SetId sets the logical id.
func (b *CareTeamBuilder) SetId(id string) *CareTeamBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier. use and typeCode are optional; a missing typeDisplay falls
// back to the code, as PHP intends.
func (b *CareTeamBuilder) AddIdentifier(system, value, use, typeCode, typeDisplay string) *CareTeamBuilder {
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

// SetStatus sets status (proposed | active | suspended | inactive | entered-in-error).
func (b *CareTeamBuilder) SetStatus(status string) *CareTeamBuilder { return b.set("status", status) }

// SetCode sets code from system, code and display.
func (b *CareTeamBuilder) SetCode(system, code, display string) *CareTeamBuilder {
	return b.set("code", map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{
			"system": system, "code": code, "display": display,
		}},
	})
}

// SetSubject sets subject from a reference. display is optional.
func (b *CareTeamBuilder) SetSubject(reference string, display *string) *CareTeamBuilder {
	m := map[string]interface{}{"reference": reference}
	if display != nil {
		m["display"] = *display
	}
	return b.set("subject", m)
}

// SetEncounter sets encounter from a bare reference.
func (b *CareTeamBuilder) SetEncounter(reference string) *CareTeamBuilder {
	return b.set("encounter", map[string]interface{}{"reference": reference})
}

// Build returns the payload.
func (b *CareTeamBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *CareTeamBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
