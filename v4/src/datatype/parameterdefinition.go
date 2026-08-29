package datatype

type ParameterDefinition struct {
	Name         *string `json:"name,omitempty"`
	Use          *string `json:"use,omitempty"`
	Min          *int    `json:"min,omitempty"`
	Max          *string `json:"max,omitempty"`
	Documentation *string `json:"documentation,omitempty"`
	Type         *string `json:"type,omitempty"`
	Profile      *string `json:"profile,omitempty"`
}

func (p *ParameterDefinition) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if p.Name != nil { m["name"] = *p.Name }
	if p.Use != nil { m["use"] = *p.Use }
	if p.Min != nil { m["min"] = *p.Min }
	if p.Max != nil { m["max"] = *p.Max }
	if p.Documentation != nil { m["documentation"] = *p.Documentation }
	if p.Type != nil { m["type"] = *p.Type }
	if p.Profile != nil { m["profile"] = *p.Profile }
	return m
}
