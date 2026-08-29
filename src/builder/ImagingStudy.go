package builder

type ImagingStudyBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewImagingStudyBuilder() *ImagingStudyBuilder {
	b := &ImagingStudyBuilder{ResourceType: "ImagingStudy", Data: make(map[string]interface{})}
	return b
}

func (b *ImagingStudyBuilder) SetId(id string) *ImagingStudyBuilder { b.Data["id"] = id; return b }
func (b *ImagingStudyBuilder) SetStatus(status string) *ImagingStudyBuilder { b.Data["status"] = status; return b }
func (b *ImagingStudyBuilder) SetSubject(subjectRef string) *ImagingStudyBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *ImagingStudyBuilder) SetEncounter(encRef string) *ImagingStudyBuilder { b.Data["encounter"] = map[string]interface{}{"reference": encRef}; return b }
func (b *ImagingStudyBuilder) SetStarted(started string) *ImagingStudyBuilder { b.Data["started"] = started; return b }
func (b *ImagingStudyBuilder) SetBasedOn(basedOnRef string) *ImagingStudyBuilder { b.Data["basedOn"] = []interface{}{map[string]interface{}{"reference": basedOnRef}}; return b }
func (b *ImagingStudyBuilder) SetReferrer(referrerRef string) *ImagingStudyBuilder { b.Data["referrer"] = map[string]interface{}{"reference": referrerRef}; return b }
func (b *ImagingStudyBuilder) SetInterpreter(interpreterRef string) *ImagingStudyBuilder { b.Data["interpreter"] = []interface{}{map[string]interface{}{"reference": interpreterRef}}; return b }
func (b *ImagingStudyBuilder) SetModalityList(system string, code string) *ImagingStudyBuilder {
	b.Data["modality"] = []interface{}{map[string]interface{}{"system": system, "code": code}}
	return b
}
func (b *ImagingStudyBuilder) Build() map[string]interface{} { return b.Data }
