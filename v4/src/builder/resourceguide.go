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

func (b *ResourceGuideBuilder) AddIdentifier(system, value string) *ResourceGuideBuilder {
	ids, _ := b.Data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.Data["identifier"] = ids
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
