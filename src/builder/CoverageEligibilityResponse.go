package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type CoverageEligibilityResponseBuilder struct {
	Data map[string]interface{}
}

func NewCoverageEligibilityResponseBuilder() *CoverageEligibilityResponseBuilder {
	return &CoverageEligibilityResponseBuilder{Data: make(map[string]interface{})}
}

func (b *CoverageEligibilityResponseBuilder) SetId(id string) *CoverageEligibilityResponseBuilder {
	b.Data["id"] = id
	return b
}

func (b *CoverageEligibilityResponseBuilder) AddIdentifier(system, value string) *CoverageEligibilityResponseBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), map[string]string{"system": system, "value": value})
	return b
}

func (b *CoverageEligibilityResponseBuilder) SetStatus(status string) *CoverageEligibilityResponseBuilder {
	b.Data["status"] = status
	return b
}

func (b *CoverageEligibilityResponseBuilder) SetPurpose(purpose []string) *CoverageEligibilityResponseBuilder {
	b.Data["purpose"] = purpose
	return b
}

func (b *CoverageEligibilityResponseBuilder) SetPatient(ref *datatype.Reference) *CoverageEligibilityResponseBuilder {
	b.Data["patient"] = ref.ToArray()
	return b
}

func (b *CoverageEligibilityResponseBuilder) SetServicedDate(dt string) *CoverageEligibilityResponseBuilder {
	b.Data["servicedDate"] = dt
	return b
}

func (b *CoverageEligibilityResponseBuilder) SetCreated(dt string) *CoverageEligibilityResponseBuilder {
	b.Data["created"] = dt
	return b
}

func (b *CoverageEligibilityResponseBuilder) SetRequest(ref string) *CoverageEligibilityResponseBuilder {
	b.Data["request"] = map[string]string{"reference": ref}
	return b
}

func (b *CoverageEligibilityResponseBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "CoverageEligibilityResponse"
	return dt
}
