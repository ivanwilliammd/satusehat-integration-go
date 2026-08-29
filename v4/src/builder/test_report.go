package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type TestReportBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewTestReportBuilder() *TestReportBuilder {
    b := &TestReportBuilder{ResourceType: "TestReport", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "TestReport"
    return b
}

func (b *TestReportBuilder) setId(id string) *TestReportBuilder {
    b.Data["id"] = id
    return b
}

func (b *TestReportBuilder) addIdentifier(identifier *datatype.Identifier) *TestReportBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *TestReportBuilder) setStatus(status string) *TestReportBuilder {
    b.Data["status"] = status
    return b
}

func (b *TestReportBuilder) setSubject(reference string) *TestReportBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *TestReportBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *TestReportBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
