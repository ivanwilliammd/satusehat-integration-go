package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type CoverageEligibilityResponseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewCoverageEligibilityResponseBuilder() *CoverageEligibilityResponseBuilder {
	b := &CoverageEligibilityResponseBuilder{ResourceType: "CoverageEligibilityResponse", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "CoverageEligibilityResponse"
	return b
}

func (b *CoverageEligibilityResponseBuilder) SetId(id string) *CoverageEligibilityResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *CoverageEligibilityResponseBuilder) AddIdentifier(identifier *datatype.Identifier) *CoverageEligibilityResponseBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *CoverageEligibilityResponseBuilder) SetStatus(status string) *CoverageEligibilityResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *CoverageEligibilityResponseBuilder) SetSubject(reference string, display ...string) *CoverageEligibilityResponseBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *CoverageEligibilityResponseBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *CoverageEligibilityResponseBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient by reference, with an optional display as the PHP builder allows.
func (b *CoverageEligibilityResponseBuilder) SetPatient(reference string, display ...string) *CoverageEligibilityResponseBuilder {
	ref := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		ref.Display = display[0]
	}
	b.Data["patient"] = ref.ToArray()
	return b
}
