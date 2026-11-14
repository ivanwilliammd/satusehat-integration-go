package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ClaimBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewClaimBuilder() *ClaimBuilder {
	b := &ClaimBuilder{ResourceType: "Claim", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Claim"
	return b
}

func (b *ClaimBuilder) SetId(id string) *ClaimBuilder {
	b.Data["id"] = id
	return b
}

func (b *ClaimBuilder) AddIdentifier(identifier *datatype.Identifier) *ClaimBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ClaimBuilder) SetStatus(status string) *ClaimBuilder {
	b.Data["status"] = status
	return b
}

func (b *ClaimBuilder) SetSubject(reference string, display ...string) *ClaimBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ClaimBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ClaimBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient by reference, with an optional display as the PHP builder allows.
func (b *ClaimBuilder) SetPatient(reference string, display ...string) *ClaimBuilder {
	ref := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		ref.Display = display[0]
	}
	b.Data["patient"] = ref.ToArray()
	return b
}
