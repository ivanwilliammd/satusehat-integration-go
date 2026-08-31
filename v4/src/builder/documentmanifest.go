package builder

import (
    "encoding/json"
)

type DocumentManifestBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewDocumentManifestBuilder() *DocumentManifestBuilder {
    b := &DocumentManifestBuilder{ResourceType: "DocumentManifest", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "DocumentManifest"
    return b
}

func (b *DocumentManifestBuilder) setId(id string) *DocumentManifestBuilder {
    b.Data["id"] = id
    return b
}

func (b *DocumentManifestBuilder) addIdentifier(system, value string) *DocumentManifestBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *DocumentManifestBuilder) setStatus(status string) *DocumentManifestBuilder {
    b.Data["status"] = status
    return b
}

func (b *DocumentManifestBuilder) setType(system, code, display string) *DocumentManifestBuilder {
    b.Data["type"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *DocumentManifestBuilder) setSubject(reference string, display ...string) *DocumentManifestBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["subject"] = ref
    return b
}

func (b *DocumentManifestBuilder) setCreated(value string) *DocumentManifestBuilder {
    b.Data["created"] = value
    return b
}

func (b *DocumentManifestBuilder) setSource(reference string, display ...string) *DocumentManifestBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["source"] = ref
    return b
}

func (b *DocumentManifestBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *DocumentManifestBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
