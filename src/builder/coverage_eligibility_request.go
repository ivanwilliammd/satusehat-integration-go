package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type CoverageEligibilityRequestBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewCoverageEligibilityRequestBuilder() *CoverageEligibilityRequestBuilder {
    b := &CoverageEligibilityRequestBuilder{ResourceType: "CoverageEligibilityRequest", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "CoverageEligibilityRequest"
    return b
}

func (b *CoverageEligibilityRequestBuilder) setId(id string) *CoverageEligibilityRequestBuilder {
    b.Data["id"] = id
    return b
}

func (b *CoverageEligibilityRequestBuilder) addIdentifier(identifier *datatype.Identifier) *CoverageEligibilityRequestBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *CoverageEligibilityRequestBuilder) setStatus(status string) *CoverageEligibilityRequestBuilder {
    b.Data["status"] = status
    return b
}

func (b *CoverageEligibilityRequestBuilder) setSubject(reference string) *CoverageEligibilityRequestBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *CoverageEligibilityRequestBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *CoverageEligibilityRequestBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
