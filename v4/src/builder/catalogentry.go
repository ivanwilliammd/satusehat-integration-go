package builder

import (
    "encoding/json"
)

type CatalogEntryBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewCatalogEntryBuilder() *CatalogEntryBuilder {
    b := &CatalogEntryBuilder{ResourceType: "CatalogEntry", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "CatalogEntry"
    return b
}

func (b *CatalogEntryBuilder) setId(id string) *CatalogEntryBuilder {
    b.Data["id"] = id
    return b
}

func (b *CatalogEntryBuilder) addIdentifier(system, value string) *CatalogEntryBuilder {
    ids, _ := b.Data["identifier"].([]map[string]string)
    ids = append(ids, map[string]string{"system": system, "value": value})
    b.Data["identifier"] = ids
    return b
}

func (b *CatalogEntryBuilder) setStatus(status string) *CatalogEntryBuilder {
    b.Data["status"] = status
    return b
}

func (b *CatalogEntryBuilder) setType(system, code, display string) *CatalogEntryBuilder {
    b.Data["type"] = map[string]interface{}{
        "coding": []map[string]string{{"system": system, "code": code, "display": display}},
    }
    return b
}

func (b *CatalogEntryBuilder) setReference(reference string, display ...string) *CatalogEntryBuilder {
    ref := map[string]interface{}{"reference": reference}
    if len(display) > 0 && display[0] != "" {
        ref["display"] = display[0]
    }
    b.Data["reference"] = ref
    return b
}

func (b *CatalogEntryBuilder) setValidityPeriod(value string) *CatalogEntryBuilder {
    b.Data["validityPeriod"] = value
    return b
}

func (b *CatalogEntryBuilder) Build() map[string]interface{} {
    clean := make(map[string]interface{})
    for k, v := range b.Data {
        if v != nil {
            clean[k] = v
        }
    }
    return clean
}

func (b *CatalogEntryBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Build())
}
