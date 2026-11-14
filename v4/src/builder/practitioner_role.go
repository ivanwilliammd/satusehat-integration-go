package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
	"strings"
)

type PractitionerRoleBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewPractitionerRoleBuilder() *PractitionerRoleBuilder {
	b := &PractitionerRoleBuilder{ResourceType: "PractitionerRole", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "PractitionerRole"
	return b
}

func (b *PractitionerRoleBuilder) SetId(id string) *PractitionerRoleBuilder {
	b.Data["id"] = id
	return b
}

func (b *PractitionerRoleBuilder) AddIdentifier(identifier *datatype.Identifier) *PractitionerRoleBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *PractitionerRoleBuilder) SetStatus(status string) *PractitionerRoleBuilder {
	b.Data["status"] = status
	return b
}

func (b *PractitionerRoleBuilder) SetSubject(reference string, display ...string) *PractitionerRoleBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *PractitionerRoleBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *PractitionerRoleBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetActive sets active.
func (b *PractitionerRoleBuilder) SetActive(active bool) *PractitionerRoleBuilder {
	b.Data["active"] = active
	return b
}

// SetPractitioner sets practitioner by reference.
func (b *PractitionerRoleBuilder) SetPractitioner(reference string) *PractitionerRoleBuilder {
	b.Data["practitioner"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetOrganization sets organization by reference.
func (b *PractitionerRoleBuilder) SetOrganization(reference string) *PractitionerRoleBuilder {
	b.Data["organization"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// AddCode appends a role coding.
func (b *PractitionerRoleBuilder) AddCode(code *datatype.CodeableConcept) *PractitionerRoleBuilder {
	pushPath(b.Data, "code", code.ToArray())
	return b
}

// AddSpecialty appends a specialty coding.
func (b *PractitionerRoleBuilder) AddSpecialty(specialty *datatype.CodeableConcept) *PractitionerRoleBuilder {
	pushPath(b.Data, "specialty", specialty.ToArray())
	return b
}

// AddLocation appends a location reference.
func (b *PractitionerRoleBuilder) AddLocation(reference string) *PractitionerRoleBuilder {
	pushPath(b.Data, "location", datatype.Reference{Reference: reference}.ToArray())
	return b
}

// AddHealthcareService appends a healthcare service reference.
func (b *PractitionerRoleBuilder) AddHealthcareService(reference string) *PractitionerRoleBuilder {
	pushPath(b.Data, "healthcareService", datatype.Reference{Reference: reference}.ToArray())
	return b
}

// AddTelecom appends a contact point.
func (b *PractitionerRoleBuilder) AddTelecom(telecom *datatype.ContactPoint) *PractitionerRoleBuilder {
	pushPath(b.Data, "telecom", telecom.ToArray())
	return b
}

// AddEndpoint appends an endpoint reference.
func (b *PractitionerRoleBuilder) AddEndpoint(reference string) *PractitionerRoleBuilder {
	pushPath(b.Data, "endpoint", datatype.Reference{Reference: reference}.ToArray())
	return b
}

// SetAvailabilityExceptions sets availabilityExceptions. The canonical builder names this "add" but it
// replaces rather than appends, and that behaviour is kept.
func (b *PractitionerRoleBuilder) SetAvailabilityExceptions(availabilityExceptions string) *PractitionerRoleBuilder {
	b.Data["availabilityExceptions"] = availabilityExceptions
	return b
}

// AddExtension appends an extension, choosing the value key from valueType as the canonical builder does.
func (b *PractitionerRoleBuilder) AddExtension(url string, value interface{}, valueType string) *PractitionerRoleBuilder {
	extension := map[string]interface{}{"url": url}
	if valueType == "" {
		extension["valueString"] = value
	} else {
		extension["value"+strings.ToUpper(valueType[:1])+valueType[1:]] = value
	}
	pushPath(b.Data, "extension", extension)
	return b
}
