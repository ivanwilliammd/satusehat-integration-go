package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type MeasureReportBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewMeasureReportBuilder() *MeasureReportBuilder {
	b := &MeasureReportBuilder{ResourceType: "MeasureReport", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "MeasureReport"
	return b
}

func (b *MeasureReportBuilder) SetId(id string) *MeasureReportBuilder {
	b.Data["id"] = id
	return b
}

func (b *MeasureReportBuilder) AddIdentifier(identifier *datatype.Identifier) *MeasureReportBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *MeasureReportBuilder) SetStatus(status string) *MeasureReportBuilder {
	b.Data["status"] = status
	return b
}

func (b *MeasureReportBuilder) SetSubject(reference string, display ...string) *MeasureReportBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *MeasureReportBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *MeasureReportBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *MeasureReportBuilder) SetCode(system, code, display string) *MeasureReportBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *MeasureReportBuilder) SetEncounter(reference string) *MeasureReportBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
