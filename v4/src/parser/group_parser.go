package parser

// GroupParser extracts fields from a FHIR Group resource map.
type GroupParser struct {
	group map[string]any
}

// NewGroupParser creates a new GroupParser instance.
func NewGroupParser(group map[string]any) *GroupParser {
	if group == nil {
		group = map[string]any{}
	}
	return &GroupParser{group: group}
}

func (p *GroupParser) GetIdentifiers() any {
	return p.group["identifier"]
}

func (p *GroupParser) GetActive() any {
	return p.group["active"]
}

func (p *GroupParser) GetType() string {
	return strVal(p.group["type"])
}

func (p *GroupParser) GetActual() any {
	return p.group["actual"]
}

func (p *GroupParser) GetCode() any {
	return p.group["code"]
}

func (p *GroupParser) GetName() string {
	return strVal(p.group["name"])
}

func (p *GroupParser) GetQuantity() any {
	return p.group["quantity"]
}

func (p *GroupParser) GetManagingEntity() any {
	return p.group["managingEntity"]
}

func (p *GroupParser) GetMembers() any {
	return p.group["member"]
}
