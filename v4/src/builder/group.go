package builder

import (
	"encoding/json"
)

type GroupBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewGroupBuilder() *GroupBuilder {
	b := &GroupBuilder{ResourceType: "Group", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "Group"
	return b
}

func (b *GroupBuilder) setMetaProfile(profile string) *GroupBuilder {
	b.Data["meta/profile"] = []string{profile}
	return b
}

func (b *GroupBuilder) setId(id string) *GroupBuilder {
	b.Data["id"] = id
	return b
}

func (b *GroupBuilder) addIdentifier(systemOrId string, value ...string) *GroupBuilder {
	// If a bare identifier-like value passed, treat first arg as system, second as value.
	ids, _ := b.Data["identifier"].([]map[string]interface{})
	entry := map[string]interface{}{"system": systemOrId}
	if len(value) > 0 {
		entry["value"] = value[0]
	}
	ids = append(ids, entry)
	b.Data["identifier"] = ids
	return b
}

func (b *GroupBuilder) setActive(active bool) *GroupBuilder {
	b.Data["active"] = active
	return b
}

func (b *GroupBuilder) setType(type_ string) *GroupBuilder {
	b.Data["type"] = type_
	return b
}

func (b *GroupBuilder) setActual(actual bool) *GroupBuilder {
	b.Data["actual"] = actual
	return b
}

func (b *GroupBuilder) setCode(code map[string]interface{}) *GroupBuilder {
	b.Data["code"] = code
	return b
}

func (b *GroupBuilder) setName(name string) *GroupBuilder {
	b.Data["name"] = name
	return b
}

func (b *GroupBuilder) setQuantity(quantity int) *GroupBuilder {
	b.Data["quantity"] = quantity
	return b
}

func (b *GroupBuilder) setManagingEntity(managingEntity map[string]interface{}) *GroupBuilder {
	b.Data["managingEntity"] = managingEntity
	return b
}

// addMember supports member entities with reference auto-prefix (Patient/).
// period/inactive optional via variadic display/period strings and inactive flag.
func (b *GroupBuilder) addMember(reference string, displayOrPeriod ...interface{}) *GroupBuilder {
	members, _ := b.Data["member"].([]map[string]interface{})

	entity := map[string]interface{}{"reference": reference}
	if !hasPrefix(reference) && !containsSlash(reference) {
		entity["reference"] = "Patient/" + reference
	}
	if len(displayOrPeriod) > 0 {
		if d, ok := displayOrPeriod[0].(string); ok && d != "" {
			entity["display"] = d
		}
	}

	member := map[string]interface{}{"entity": entity}
	if len(displayOrPeriod) > 1 {
		if p, ok := displayOrPeriod[1].(map[string]interface{}); ok {
			member["period"] = p
		}
	}
	if len(displayOrPeriod) > 2 {
		if in, ok := displayOrPeriod[2].(bool); ok {
			member["inactive"] = in
		}
	}

	members = append(members, member)
	b.Data["member"] = members
	return b
}

func (b *GroupBuilder) addExtension(url string, value interface{}) *GroupBuilder {
	extensions, _ := b.Data["extension"].([]map[string]interface{})
	ext := map[string]interface{}{"url": url}
	switch v := value.(type) {
	case bool:
		ext["valueBoolean"] = v
	case string:
		ext["valueString"] = v
	case int:
		ext["valueInteger"] = v
	case map[string]interface{}:
		for k, val := range v {
			ext[k] = val
		}
	}
	extensions = append(extensions, ext)
	b.Data["extension"] = extensions
	return b
}

func (b *GroupBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *GroupBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
