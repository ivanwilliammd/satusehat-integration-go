package builder

import (
	"encoding/json"
	"fmt"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type ImagingStudyBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewImagingStudyBuilder() *ImagingStudyBuilder {
	b := &ImagingStudyBuilder{ResourceType: "ImagingStudy", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ImagingStudy"
	return b
}

func (b *ImagingStudyBuilder) SetId(id string) *ImagingStudyBuilder {
	b.Data["id"] = id
	return b
}

func (b *ImagingStudyBuilder) AddIdentifier(identifier *datatype.Identifier) *ImagingStudyBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *ImagingStudyBuilder) SetStatus(status string) *ImagingStudyBuilder {
	b.Data["status"] = status
	return b
}

func (b *ImagingStudyBuilder) SetSubject(reference string, display ...string) *ImagingStudyBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *ImagingStudyBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *ImagingStudyBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetEncounter sets encounter by reference, as the PHP builder does.
func (b *ImagingStudyBuilder) SetEncounter(reference string) *ImagingStudyBuilder {
	b.Data["encounter"] = map[string]interface{}{"reference": reference}
	return b
}

// SetSeriesStarted sets the start of an existing series.
func (b *ImagingStudyBuilder) SetSeriesStarted(seriesIdx int, dateTime string) *ImagingStudyBuilder {
	setPath(b.Data, fmt.Sprintf("series/%d/started", seriesIdx), dateTime)
	return b
}

// SetSeriesBodySite sets the body site coding of an existing series.
func (b *ImagingStudyBuilder) SetSeriesBodySite(seriesIdx int, system, code, display string) *ImagingStudyBuilder {
	coding := map[string]interface{}{"system": system, "code": code}
	if display != "" {
		coding["display"] = display
	}
	setPath(b.Data, fmt.Sprintf("series/%d/bodySite", seriesIdx), map[string]interface{}{
		"coding": []interface{}{coding},
	})
	return b
}

// AddSeriesInstance adds an instance to an existing series.
func (b *ImagingStudyBuilder) AddSeriesInstance(seriesIdx int, uid, sopSystem, sopCode, sopDisplay string) *ImagingStudyBuilder {
	pushPath(b.Data, fmt.Sprintf("series/%d/instance", seriesIdx), map[string]interface{}{
		"uid": uid,
		"sopClass": map[string]interface{}{
			"system": sopSystem, "code": sopCode, "display": sopDisplay,
		},
	})
	return b
}

// AddSeriesInstanceNumber sets the number of an existing instance.
func (b *ImagingStudyBuilder) AddSeriesInstanceNumber(seriesIdx, instanceIdx, number int) *ImagingStudyBuilder {
	setPath(b.Data, fmt.Sprintf("series/%d/instance/%d/number", seriesIdx, instanceIdx), number)
	return b
}
