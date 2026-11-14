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

// SetInsurer sets insurer by reference, as the PHP builder does.
func (b *ClaimResponseBuilder) SetInsurer(reference string) *ClaimResponseBuilder {
	b.Data["insurer"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetCreated sets created, as the PHP builder does.
func (b *ClaimResponseBuilder) SetCreated(dateTime string) *ClaimResponseBuilder {
	b.Data["created"] = dateTime
	return b
}

// SetType sets type with a single coding, as the PHP builder does.
func (b *ClaimResponseBuilder) SetType(system, code, display string) *ClaimResponseBuilder {
	b.Data["type"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// AddCoverage adds a coverage entry with its sequence, as the PHP builder does.
func (b *ClaimResponseBuilder) AddCoverage(sequence int, reference string, focal *bool, prioritySystem, priorityCode *string) *ClaimResponseBuilder {
	entri := map[string]interface{}{"sequence": sequence, "reference": reference}
	if focal != nil {
		entri["focal"] = *focal
	}
	if prioritySystem != nil && priorityCode != nil {
		entri["priority"] = map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": *prioritySystem, "code": *priorityCode}},
		}
	}
	daftar, _ := b.Data["coverage"].([]interface{})
	b.Data["coverage"] = append(daftar, entri)
	return b
}

// SetOutcome sets the outcome code, as the PHP builder does.
func (b *ClaimResponseBuilder) SetOutcome(outcome string) *ClaimResponseBuilder {
	b.Data["outcome"] = outcome
	return b
}

// SetRequest sets request by reference, as the PHP builder does.
func (b *ClaimResponseBuilder) SetRequest(reference string) *ClaimResponseBuilder {
	b.Data["request"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}
