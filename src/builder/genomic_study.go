package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type GenomicStudyBuilder struct {
	Data map[string]interface{}
}

func NewGenomicStudyBuilder() *GenomicStudyBuilder {
	return &GenomicStudyBuilder{Data: make(map[string]interface{})}
}

func (b *GenomicStudyBuilder) SetId(id string) *GenomicStudyBuilder {
	b.Data["id"] = id
	return b
}

func (b *GenomicStudyBuilder) SetStatus(status string) *GenomicStudyBuilder {
	b.Data["status"] = status
	return b
}

func (b *GenomicStudyBuilder) SetType(cc *datatype.CodeableConcept) *GenomicStudyBuilder {
	b.Data["type"] = cc.ToArray()
	return b
}

func (b *GenomicStudyBuilder) SetSubject(ref *datatype.Reference) *GenomicStudyBuilder {
	b.Data["subject"] = ref.ToArray()
	return b
}

func (b *GenomicStudyBuilder) SetEncounter(ref *datatype.Reference) *GenomicStudyBuilder {
	b.Data["encounter"] = ref.ToArray()
	return b
}

func (b *GenomicStudyBuilder) SetStarted(dt string) *GenomicStudyBuilder {
	b.Data["started"] = dt
	return b
}

func (b *GenomicStudyBuilder) SetBasedOn(ref *datatype.Reference) *GenomicStudyBuilder {
	if _, ok := b.Data["basedOn"]; !ok {
		b.Data["basedOn"] = make([]interface{}, 0)
	}
	b.Data["basedOn"] = append(b.Data["basedOn"].([]interface{}), ref.ToArray())
	return b
}

func (b *GenomicStudyBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "GenomicStudy"
	return dt
}
