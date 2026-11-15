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

// SetInsurer sets insurer by reference, as the PHP builder does.
func (b *CoverageEligibilityResponseBuilder) SetInsurer(reference string) *CoverageEligibilityResponseBuilder {
	b.Data["insurer"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetCreated sets created, as the PHP builder does.
func (b *CoverageEligibilityResponseBuilder) SetCreated(dateTime string) *CoverageEligibilityResponseBuilder {
	b.Data["created"] = dateTime
	return b
}

// AddCoverage adds a coverage entry with an optional pre-authorisation reference.
func (b *CoverageEligibilityResponseBuilder) AddCoverage(reference string, preAuthRef string) *CoverageEligibilityResponseBuilder {
	entri := map[string]interface{}{"reference": reference}
	if preAuthRef != "" {
		entri["preAuthRef"] = []interface{}{preAuthRef}
	}
	daftar, _ := b.Data["coverage"].([]interface{})
	b.Data["coverage"] = append(daftar, entri)
	return b
}

// SetOutcome sets the outcome code, as the PHP builder does.
func (b *CoverageEligibilityResponseBuilder) SetOutcome(outcome string) *CoverageEligibilityResponseBuilder {
	b.Data["outcome"] = outcome
	return b
}

// SetRequest sets request by reference, as the PHP builder does.
func (b *CoverageEligibilityResponseBuilder) SetRequest(reference string) *CoverageEligibilityResponseBuilder {
	b.Data["request"] = datatype.Reference{Reference: reference}.ToArray()
	return b
}

// SetDisposition sets disposition, as the PHP builder does.
func (b *CoverageEligibilityResponseBuilder) SetDisposition(disposition string) *CoverageEligibilityResponseBuilder {
	b.Data["disposition"] = disposition
	return b
}

// SetPurpose sets the purpose list.
func (b *CoverageEligibilityResponseBuilder) SetPurpose(purpose []string) *CoverageEligibilityResponseBuilder {
	b.Data["purpose"] = purpose
	return b
}

// SetServicedDate sets servicedDate.
func (b *CoverageEligibilityResponseBuilder) SetServicedDate(date string) *CoverageEligibilityResponseBuilder {
	b.Data["servicedDate"] = date
	return b
}

// SetServicedPeriod writes servicedPeriod through nested paths.
func (b *CoverageEligibilityResponseBuilder) SetServicedPeriod(start, end string) *CoverageEligibilityResponseBuilder {
	setPath(b.Data, "servicedPeriod/start", start)
	if end != "" {
		setPath(b.Data, "servicedPeriod/end", end)
	}
	return b
}

// SetRequestProvider sets requestProvider by reference.
func (b *CoverageEligibilityResponseBuilder) SetRequestProvider(reference string) *CoverageEligibilityResponseBuilder {
	b.Data["requestProvider"] = map[string]interface{}{"reference": reference}
	return b
}

// AddInsurance appends an insurance block with an optional benefit balance.
func (b *CoverageEligibilityResponseBuilder) AddInsurance(coverageReference string, balanceName, typeSystem, typeCode, typeDisplay *string) *CoverageEligibilityResponseBuilder {
	insurance := map[string]interface{}{"coverage": map[string]interface{}{"reference": coverageReference}}
	if balanceName != nil && *balanceName != "" {
		item := map[string]interface{}{"name": *balanceName}
		if typeSystem != nil && typeCode != nil {
			display := ""
			if typeDisplay != nil {
				display = *typeDisplay
			}
			item["type"] = map[string]interface{}{
				"coding": []interface{}{map[string]interface{}{"system": *typeSystem, "code": *typeCode, "display": display}},
			}
		}
		insurance["benefitBalance"] = []interface{}{item}
	}
	pushPath(b.Data, "insurance", insurance)
	return b
}
