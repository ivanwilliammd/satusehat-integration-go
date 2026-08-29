package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type DiagnosticReportBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewDiagnosticReportBuilder() *DiagnosticReportBuilder {
    b := &DiagnosticReportBuilder{ResourceType: "DiagnosticReport", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "DiagnosticReport"
    return b
}

func (b *DiagnosticReportBuilder) setId(id string) *DiagnosticReportBuilder {
    b.Data["id"] = id
    return b
}

func (b *DiagnosticReportBuilder) addIdentifier(identifier *datatype.Identifier) *DiagnosticReportBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *DiagnosticReportBuilder) setStatus(status string) *DiagnosticReportBuilder {
    b.Data["status"] = status
    return b
}

func (b *DiagnosticReportBuilder) setSubject(reference string) *DiagnosticReportBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *DiagnosticReportBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *DiagnosticReportBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
