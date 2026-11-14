package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ClaimResponseBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewClaimResponseBuilder() *ClaimResponseBuilder {
	b := &ClaimResponseBuilder{ResourceType: "ClaimResponse", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ClaimResponse"
	return b
}

func (b *ClaimResponseBuilder) SetId(id string) *ClaimResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *ClaimResponseBuilder) AddIdentifier(identifier *datatype.Identifier) *ClaimResponseBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ClaimResponseBuilder) SetStatus(status string) *ClaimResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *ClaimResponseBuilder) SetSubject(reference string, display ...string) *ClaimResponseBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ClaimResponseBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ClaimResponseBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient by reference, as the PHP builder does.
func (b *ClaimResponseBuilder) SetPatient(reference string) *ClaimResponseBuilder {
	b.Data["patient"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}
