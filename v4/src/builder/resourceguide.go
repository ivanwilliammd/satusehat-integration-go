package builder

import (
	"encoding/json"
)

type ResourceGuideBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewResourceGuideBuilder() *ResourceGuideBuilder {
	b := &ResourceGuideBuilder{ResourceType: "ResourceGuide", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "ResourceGuide"
	return b
}

func (b *ResourceGuideBuilder) SetId(id string) *ResourceGuideBuilder {
	b.Data["id"] = id
	return b
}

// AddIdentifier appends an identifier. use, typeCode and typeDisplay are optional and, when given, are
// written the way the canonical builder writes them — the type coding uses the v2-0203 system.
func (b *ResourceGuideBuilder) AddIdentifier(system, value string, useAndType ...string) *ResourceGuideBuilder {
	ident := map[string]interface{}{"system": system, "value": value}
	if len(useAndType) > 0 && useAndType[0] != "" {
		ident["use"] = useAndType[0]
	}
	if len(useAndType) > 1 && useAndType[1] != "" {
		display := useAndType[1]
		if len(useAndType) > 2 && useAndType[2] != "" {
			display = useAndType[2]
		}
		ident["type"] = map[string]interface{}{
			"coding": []map[string]string{{
				"system":  "http://terminology.hl7.org/CodeSystem/v2-0203",
				"code":    useAndType[1],
				"display": display,
			}},
		}
	}
	ids, _ := b.Data["identifier"].([]map[string]interface{})
	b.Data["identifier"] = append(ids, ident)
	return b
}

func (b *ResourceGuideBuilder) SetStatus(status string) *ResourceGuideBuilder {
	b.Data["status"] = status
	return b
}

func (b *ResourceGuideBuilder) SetName(value string) *ResourceGuideBuilder {
	b.Data["name"] = value
	return b
}

func (b *ResourceGuideBuilder) SetDescription(value string) *ResourceGuideBuilder {
	b.Data["description"] = value
	return b
}

func (b *ResourceGuideBuilder) SetVersion(value string) *ResourceGuideBuilder {
	b.Data["version"] = value
	return b
}

func (b *ResourceGuideBuilder) SetPublisher(value string) *ResourceGuideBuilder {
	b.Data["publisher"] = value
	return b
}

func (b *ResourceGuideBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *ResourceGuideBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
