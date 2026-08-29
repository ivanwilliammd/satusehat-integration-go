package builder

type GroupBuilder struct {
	Data map[string]interface{}
}

func NewGroupBuilder() *GroupBuilder {
	return &GroupBuilder{Data: make(map[string]interface{})}
}

func (b *GroupBuilder) SetId(id string) *GroupBuilder {
	b.Data["id"] = id
	return b
}

func (b *GroupBuilder) SetType(grpType string) *GroupBuilder {
	b.Data["type"] = grpType
	return b
}

func (b *GroupBuilder) SetActive(active bool) *GroupBuilder {
	b.Data["active"] = active
	return b
}

func (b *GroupBuilder) SetName(name string) *GroupBuilder {
	b.Data["name"] = name
	return b
}

func (b *GroupBuilder) SetQuantity(qty int) *GroupBuilder {
	b.Data["quantity"] = qty
	return b
}

func (b *GroupBuilder) AddMemberEntity(ref string) *GroupBuilder {
	if _, ok := b.Data["member"]; !ok {
		b.Data["member"] = make([]interface{}, 0)
	}
	b.Data["member"] = append(b.Data["member"].([]interface{}), map[string]interface{}{"entity": map[string]string{"reference": ref}})
	return b
}

func (b *GroupBuilder) Build() map[string]interface{} {
	dt := make(map[string]interface{})
	for k, v := range b.Data {
		if v == nil || v == "" {
			continue
		}
		dt[k] = v
	}
	dt["resourceType"] = "Group"
	return dt
}
