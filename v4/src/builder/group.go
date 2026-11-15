package builder

import (
	"encoding/json"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
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

func (b *GroupBuilder) SetMetaProfile(profile string) *GroupBuilder {
	// "meta/profile" sebagai kunci literal menghasilkan {"meta/profile": …}, bukan meta.profile —
	// JSON sah, tetapi bukan FHIR. Ditulis sebagai objek meta, dan profile MENUMPUK seperti push di kanon:
	// panggilan kedua tidak boleh membuang profil pertama.
	meta, _ := b.Data["meta"].(map[string]interface{})
	if meta == nil {
		meta = map[string]interface{}{}
		b.Data["meta"] = meta
	}
	list, _ := meta["profile"].([]interface{})
	meta["profile"] = append(list, profile)
	return b
}

func (b *GroupBuilder) SetId(id string) *GroupBuilder {
	b.Data["id"] = id
	return b
}

func (b *GroupBuilder) AddIdentifier(systemOrId string, value ...string) *GroupBuilder {
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

func (b *GroupBuilder) SetActive(active bool) *GroupBuilder {
	b.Data["active"] = active
	return b
}

func (b *GroupBuilder) SetType(type_ string) *GroupBuilder {
	b.Data["type"] = type_
	return b
}

func (b *GroupBuilder) SetActual(actual bool) *GroupBuilder {
	b.Data["actual"] = actual
	return b
}

func (b *GroupBuilder) SetCode(code map[string]interface{}) *GroupBuilder {
	b.Data["code"] = code
	return b
}

func (b *GroupBuilder) SetName(name string) *GroupBuilder {
	b.Data["name"] = name
	return b
}

func (b *GroupBuilder) SetQuantity(quantity int) *GroupBuilder {
	b.Data["quantity"] = quantity
	return b
}

func (b *GroupBuilder) SetManagingEntity(managingEntity map[string]interface{}) *GroupBuilder {
	b.Data["managingEntity"] = managingEntity
	return b
}

// addMember supports member entities with reference auto-prefix (Patient/).
// period/inactive optional via variadic display/period strings and inactive flag.
// AddMember appends a member. The canonical builder accepts four positional arrangements — display alone,
// period alone, display and period, period and inactive — and decides by the type of what it is given, not by
// position. This does the same, so a caller porting from PHP does not have to discover the difference.
func (b *GroupBuilder) AddMember(reference string, displayOrPeriod ...interface{}) *GroupBuilder {
	entity := map[string]interface{}{"reference": groupPrefix(reference)}

	var period interface{}
	var inactive *bool

	if len(displayOrPeriod) > 0 {
		switch v := displayOrPeriod[0].(type) {
		case string:
			if v != "" {
				entity["display"] = v
			}
		case bool:
			inactive = &v
		case map[string]interface{}:
			period = v
		}
	}
	if len(displayOrPeriod) > 1 {
		switch v := displayOrPeriod[1].(type) {
		case bool:
			inactive = &v
		case map[string]interface{}:
			period = v
		}
	}

	member := map[string]interface{}{"entity": entity}
	if period != nil {
		member["period"] = period
	}
	if inactive != nil {
		member["inactive"] = *inactive
	}

	members, _ := b.Data["member"].([]map[string]interface{})
	b.Data["member"] = append(members, member)
	return b
}

// AddMemberReference appends a member from a full Reference, keeping every element it carries. The canonical
// builder accepts a Reference directly and does not prefix or reshape it; the string form cannot express that.
// groupPrefix prefixes a bare id the way the canonical builder does.
func groupPrefix(ref string) string {
	if hasPrefix(ref) || containsSlash(ref) {
		return ref
	}
	return "Patient/" + ref
}

func (b *GroupBuilder) AddMemberReference(entity *datatype.Reference, displayOrPeriod ...interface{}) *GroupBuilder {
	members, _ := b.Data["member"].([]map[string]interface{})
	member := map[string]interface{}{"entity": entity.ToArray()}
	if len(displayOrPeriod) > 0 {
		if p, ok := displayOrPeriod[0].(map[string]interface{}); ok {
			member["period"] = p
		}
	}
	if len(displayOrPeriod) > 1 {
		if in, ok := displayOrPeriod[1].(bool); ok {
			member["inactive"] = in
		}
	}
	b.Data["member"] = append(members, member)
	return b
}

func (b *GroupBuilder) AddExtension(url string, value interface{}) *GroupBuilder {
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
