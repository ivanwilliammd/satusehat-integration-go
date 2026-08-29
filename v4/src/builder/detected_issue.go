package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type DetectedIssueBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewDetectedIssueBuilder() *DetectedIssueBuilder {
    b := &DetectedIssueBuilder{ResourceType: "DetectedIssue", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "DetectedIssue"
    return b
}

func (b *DetectedIssueBuilder) setId(id string) *DetectedIssueBuilder {
    b.Data["id"] = id
    return b
}

func (b *DetectedIssueBuilder) addIdentifier(identifier *datatype.Identifier) *DetectedIssueBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *DetectedIssueBuilder) setStatus(status string) *DetectedIssueBuilder {
    b.Data["status"] = status
    return b
}

func (b *DetectedIssueBuilder) setSubject(reference string) *DetectedIssueBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *DetectedIssueBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *DetectedIssueBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
