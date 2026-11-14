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

func (b *CoverageEligibilityResponseBuilder) SetSubject(reference string) *CoverageEligibilityResponseBuilder {
	b.Data["subject"] = datatype.Reference{}.ToArray()
	return b
}

func (b *CoverageEligibilityResponseBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *CoverageEligibilityResponseBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}
