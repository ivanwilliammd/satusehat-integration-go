package datatype

type Signature struct {
	Type       []Coding    `json:"type,omitempty"`
	When       *string     `json:"when,omitempty"`
	Who        *Reference  `json:"who,omitempty"`
	OnBehalfOf *Reference  `json:"onBehalfOf,omitempty"`
	TargetFormat *string   `json:"targetFormat,omitempty"`
	SigFormat    *string   `json:"sigFormat,omitempty"`
	Data         *string   `json:"data,omitempty"`
}

func (s *Signature) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if len(s.Type) > 0 {
		types := make([]map[string]interface{}, 0, len(s.Type))
		for _, t := range s.Type {
			types = append(types, t.ToArray())
		}
		m["type"] = types
	}
	if s.When != nil { m["when"] = *s.When }
	if s.Who != nil { m["who"] = s.Who.ToArray() }
	if s.OnBehalfOf != nil { m["onBehalfOf"] = s.OnBehalfOf.ToArray() }
	if s.TargetFormat != nil { m["targetFormat"] = *s.TargetFormat }
	if s.SigFormat != nil { m["sigFormat"] = *s.SigFormat }
	if s.Data != nil { m["data"] = *s.Data }
	return m
}

func (s *Signature) AddType(c Coding) *Signature {
	s.Type = append(s.Type, c)
	return s
}
