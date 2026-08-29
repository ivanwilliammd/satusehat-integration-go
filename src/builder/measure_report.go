package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

type MeasureReportBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewMeasureReportBuilder() *MeasureReportBuilder {
    b := &MeasureReportBuilder{ResourceType: "MeasureReport", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "MeasureReport"
    return b
}

func (b *MeasureReportBuilder) setId(id string) *MeasureReportBuilder {
    b.Data["id"] = id
    return b
}

func (b *MeasureReportBuilder) addIdentifier(identifier *datatype.Identifier) *MeasureReportBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *MeasureReportBuilder) setStatus(status string) *MeasureReportBuilder {
    b.Data["status"] = status
    return b
}

func (b *MeasureReportBuilder) setSubject(reference string) *MeasureReportBuilder {
    b.Data["subject"] = (&datatype.Reference{Reference: reference}).ToArray()
    return b
}

func (b *MeasureReportBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *MeasureReportBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
