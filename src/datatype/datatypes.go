package datatype

type Identifier struct {
	System string `json:"system"`
	Value  string `json:"value"`
}

func (i *Identifier) ToArray() map[string]interface{} {
	return map[string]interface{}{"system": i.System, "value": i.Value}
}

type HumanNameUse string
const (
	HumanNameUseOfficial HumanNameUse = "official"
	HumanNameUseTemp     HumanNameUse = "temp"
	HumanNameUseNickname HumanNameUse = "nickname"
	HumanNameUseAnon     HumanNameUse = "anonymous"
	HumanNameUseOld      HumanNameUse = "old"
)

type HumanName struct {
	Use    *HumanNameUse `json:"use,omitempty"`
	Text   string        `json:"text,omitempty"`
	Family string        `json:"family,omitempty"`
	Given  []string      `json:"given,omitempty"`
}

func (h *HumanName) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if h.Use != nil { m["use"] = string(*h.Use) }
	if h.Text != "" { m["text"] = h.Text }
	if h.Family != "" { m["family"] = h.Family }
	if len(h.Given) > 0 { m["given"] = h.Given }
	return m
}

type Coding struct {
	System       string `json:"system,omitempty"`
	Version      string `json:"version,omitempty"`
	Code         string `json:"code,omitempty"`
	Display      string `json:"display,omitempty"`
	UserSelected *bool  `json:"userSelected,omitempty"`
}

func (c *Coding) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if c.System != "" { m["system"] = c.System }
	if c.Version != "" { m["version"] = c.Version }
	if c.Code != "" { m["code"] = c.Code }
	if c.Display != "" { m["display"] = c.Display }
	if c.UserSelected != nil { m["userSelected"] = *c.UserSelected }
	return m
}

type CodeableConcept struct {
	Coding []Coding `json:"coding,omitempty"`
	Text   string   `json:"text,omitempty"`
}

func (cc *CodeableConcept) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if len(cc.Coding) > 0 {
		codings := make([]map[string]interface{}, 0, len(cc.Coding))
		for _, c := range cc.Coding {
			codings = append(codings, c.ToArray())
		}
		m["coding"] = codings
	}
	if cc.Text != "" { m["text"] = cc.Text }
	return m
}

func (cc *CodeableConcept) AddCoding(c Coding) *CodeableConcept {
	cc.Coding = append(cc.Coding, c)
	return cc
}

type Reference struct {
	Reference string      `json:"reference,omitempty"`
	Type      string      `json:"type,omitempty"`
	Identifier *Identifier `json:"identifier,omitempty"`
	Display   string      `json:"display,omitempty"`
}

func (r *Reference) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if r.Reference != "" { m["reference"] = r.Reference }
	if r.Type != "" { m["type"] = r.Type }
	if r.Identifier != nil { m["identifier"] = r.Identifier.ToArray() }
	if r.Display != "" { m["display"] = r.Display }
	return m
}

type ParameterComponent struct {
	Name    string             `json:"name,omitempty"`
	Value   interface{}        `json:"value,omitempty"`
	Element map[string]interface{} `json:"_value,omitempty"`
}
