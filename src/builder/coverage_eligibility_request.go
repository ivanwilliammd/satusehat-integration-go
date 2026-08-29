package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type CoverageEligibilityRequestBuilder struct {
	Data map[string]interface{}
}

func NewCoverageEligibilityRequestBuilder() *CoverageEligibilityRequestBuilder {
	return &CoverageEligibilityRequestBuilder{Data: make(map[string]interface{})}
}

func (b *CoverageEligibilityRequestBuilder) SetId(id string) *CoverageEligibilityRequestBuilder {
	b.Data["id"] = id
	return b
}

func (b *CoverageEligibilityRequestBuilder) AddIdentifier(system, value string) *CoverageEligibilityRequestBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), map[string]string{"system": system, "value": value})
	return b
}

func (b *CoverageEligibilityRequestBuilder) SetStatus(status string) *CoverageEligibilityRequestBuilder {
	b.Data["status"] = status
	return b
}

func (b *CoverageEligibilityRequestBuilder) SetPriority(cc *datatype.CodeableConcept) *CoverageEligibilityRequestBuilder {
	b.Data["priority"] = cc.ToArray()
	return b
}

func (b *CoverageEligibilityRequestBuilder) SetPurpose(purpose []string) *CoverageEligibilityRequestBuilder {
	b.Data["purpose"] = purpose
	return b
}

func (b *CoverageEligibilityRequestBuilder) SetPatient(ref *datatype.Reference) *CoverageEligibilityRequestBuilder {
	b.Data["patient"] = ref.ToArray()
	return b
}

func (b *CoverageEligibilityRequestBuilder) SetServicedDate(dt string) *CoverageEligibilityRequestBuilder {
	b.Data["servicedDate"] = dt
	return b
}

func (b *CoverageEligibilityRequestBuilder) SetCreated(dt string) *CoverageEligibilityRequestBuilder {
	b.Data["created"] = dt
	return b
}

func (b *CoverageEligibilityRequestBuilder) SetRequestor(ref *datatype.Reference) *CoverageEligibilityRequestBuilder {
	b.Data["requestor"] = ref.ToArray()
	return b
}

func (b *CoverageEligibilityRequestBuilder) SetInsurer(ref *datatype.Reference) *CoverageEligibilityRequestBuilder {
	b.Data["insurer"] = ref.ToArray()
	return b
}

func (b *CoverageEligibilityRequestBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "CoverageEligibilityRequest"
	return dt
}
