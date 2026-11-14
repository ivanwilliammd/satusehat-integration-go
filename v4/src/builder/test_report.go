package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type TestReportBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewTestReportBuilder() *TestReportBuilder {
	b := &TestReportBuilder{ResourceType: "TestReport", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "TestReport"
	return b
}

func (b *TestReportBuilder) SetId(id string) *TestReportBuilder {
	b.Data["id"] = id
	return b
}

func (b *TestReportBuilder) AddIdentifier(identifier *datatype.Identifier) *TestReportBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *TestReportBuilder) SetStatus(status string) *TestReportBuilder {
	b.Data["status"] = status
	return b
}

func (b *TestReportBuilder) SetSubject(reference string, display ...string) *TestReportBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *TestReportBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *TestReportBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetCode sets code with a single coding, as the PHP builder does.
func (b *TestReportBuilder) SetCode(system, code, display string) *TestReportBuilder {
	b.Data["code"] = map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code, "display": display}},
	}
	return b
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *TestReportBuilder) SetEncounter(reference string) *TestReportBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}
