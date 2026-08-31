package builder

import (
    "encoding/json"
)

type ResearchStudyBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewResearchStudyBuilder() *ResearchStudyBuilder {
    b := &ResearchStudyBuilder{ResourceType: "ResearchStudy", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "ResearchStudy"
    return b
}

func (b *ResearchStudyBuilder) setId(id string) *ResearchStudyBuilder {
    b.Data["id"] = id
    return b
}

func (b *ResearchStudyBuilder) addIdentifier(system, value string) *ResearchStudyBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *ResearchStudyBuilder) setStatus(status string) *ResearchStudyBuilder {
    b.Data["status"] = status
    return b
}

func (b *ResearchStudyBuilder) setTitle(value string) *ResearchStudyBuilder {
    b.Data["title"] = value
    return b
}

func (b *ResearchStudyBuilder) setProtocol(reference string, display ...string) *ResearchStudyBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["protocol"] = ref
    return b
}

func (b *ResearchStudyBuilder) setSponsor(reference string, display ...string) *ResearchStudyBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["sponsor"] = ref
    return b
}

func (b *ResearchStudyBuilder) setPhase(system, code, display string) *ResearchStudyBuilder {
    b.Data["phase"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *ResearchStudyBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *ResearchStudyBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
