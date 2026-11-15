package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type CoverageEligibilityRequestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewCoverageEligibilityRequestBuilder() *CoverageEligibilityRequestBuilder {
	b := &CoverageEligibilityRequestBuilder{ResourceType: "CoverageEligibilityRequest", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "CoverageEligibilityRequest"
	return b
}

func (b *CoverageEligibilityRequestBuilder) SetId(id string) *CoverageEligibilityRequestBuilder {
	b.Data["id"] = id
	return b
}

func (b *CoverageEligibilityRequestBuilder) AddIdentifier(identifier *datatype.Identifier) *CoverageEligibilityRequestBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *CoverageEligibilityRequestBuilder) SetStatus(status string) *CoverageEligibilityRequestBuilder {
	b.Data["status"] = status
	return b
}

func (b *CoverageEligibilityRequestBuilder) SetSubject(reference string, display ...string) *CoverageEligibilityRequestBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *CoverageEligibilityRequestBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *CoverageEligibilityRequestBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient by reference, with an optional display as the PHP builder allows.
func (b *CoverageEligibilityRequestBuilder) SetPatient(reference string, display ...string) *CoverageEligibilityRequestBuilder {
	ref := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		ref.Display = display[0]
	}
	b.Data["patient"] = ref.ToArray()
	return b
}

// SetInsurer sets insurer by reference, as the PHP builder does.
func (b *CoverageEligibilityRequestBuilder) SetInsurer(reference string) *CoverageEligibilityRequestBuilder {
	b.Data["insurer"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetCreated sets created, as the PHP builder does.
func (b *CoverageEligibilityRequestBuilder) SetCreated(dateTime string) *CoverageEligibilityRequestBuilder {
	b.Data["created"] = dateTime
	return b
}

// AddCoverage adds a coverage entry with an optional pre-authorisation reference.
func (b *CoverageEligibilityRequestBuilder) AddCoverage(reference string, preAuthRef string) *CoverageEligibilityRequestBuilder {
	entri := map[string]interface{}{"reference": reference}
	if preAuthRef != "" {
		entri["preAuthRef"] = []interface{}{preAuthRef}
	}
	daftar, _ := b.Data["coverage"].([]interface{})
	b.Data["coverage"] = append(daftar, entri)
	return b
}

// SetEnterer sets enterer by reference, as the PHP builder does.
func (b *CoverageEligibilityRequestBuilder) SetEnterer(reference string) *CoverageEligibilityRequestBuilder {
	b.Data["enterer"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetProvider sets provider by reference, as the PHP builder does.
func (b *CoverageEligibilityRequestBuilder) SetProvider(reference string) *CoverageEligibilityRequestBuilder {
	b.Data["provider"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// AddItem adds an item with its service coding and optional category, as the PHP builder does.
func (b *CoverageEligibilityRequestBuilder) AddItem(system, code, display string, categorySystem, categoryCode string) *CoverageEligibilityRequestBuilder {
	item := map[string]interface{}{
		"productOrService": map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
		},
	}
	if categorySystem != "" && categoryCode != "" {
		item["category"] = map[string]interface{}{
			"coding": []interface{}{map[string]interface{}{"system": categorySystem, "code": categoryCode}},
		}
	}
	daftar, _ := b.Data["item"].([]interface{})
	b.Data["item"] = append(daftar, item)
	return b
}

// SetPriority sets priority with a single coding.
func (b *CoverageEligibilityRequestBuilder) SetPriority(system, code, display string) *CoverageEligibilityRequestBuilder {
	b.Data["priority"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetPurpose sets the purpose list.
func (b *CoverageEligibilityRequestBuilder) SetPurpose(purpose []string) *CoverageEligibilityRequestBuilder {
	b.Data["purpose"] = purpose
	return b
}

// SetServicedDate sets servicedDate.
func (b *CoverageEligibilityRequestBuilder) SetServicedDate(date string) *CoverageEligibilityRequestBuilder {
	b.Data["servicedDate"] = date
	return b
}

// SetServicedPeriod writes servicedPeriod through nested paths.
func (b *CoverageEligibilityRequestBuilder) SetServicedPeriod(start, end string) *CoverageEligibilityRequestBuilder {
	setPath(b.Data, "servicedPeriod/start", start)
	if end != "" {
		setPath(b.Data, "servicedPeriod/end", end)
	}
	return b
}
