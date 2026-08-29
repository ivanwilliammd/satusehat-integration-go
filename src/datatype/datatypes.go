package datatype

type Identifier struct {
	System string `json:"system"`
	Value  string `json:"value"`
}

type HumanName struct {
	Use    string   `json:"use,omitempty"`
	Text   string   `json:"text,omitempty"`
	Family string   `json:"family,omitempty"`
	Given  []string `json:"given,omitempty"`
}

func (h *HumanName) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if h.Use != "" { m["use"] = h.Use }
	if h.Text != "" { m["text"] = h.Text }
	if h.Family != "" { m["family"] = h.Family }
	if len(h.Given) > 0 { m["given"] = h.Given }
	return m
}

type ContactPoint struct {
	System string `json:"system,omitempty"`
	Value  string `json:"value,omitempty"`
	Use    string `json:"use,omitempty"`
}

func (c *ContactPoint) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if c.System != "" { m["system"] = c.System }
	if c.Value != "" { m["value"] = c.Value }
	if c.Use != "" { m["use"] = c.Use }
	return m
}

type Coding struct {
	System string `json:"system,omitempty"`
	Version string `json:"version,omitempty"`
	Code   string `json:"code,omitempty"`
	Display string `json:"display,omitempty"`
	OriginalText string `json:"originalText,omitempty"`
}

func (c *Coding) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if c.System != "" { m["system"] = c.System }
	if c.Version != "" { m["version"] = c.Version }
	if c.Code != "" { m["code"] = c.Code }
	if c.Display != "" { m["display"] = c.Display }
	if c.OriginalText != "" { m["originalText"] = c.OriginalText }
	return m
}

type CodeableConcept struct {
	Coding []Coding `json:"coding,omitempty"`
	Text   string   `json:"text,omitempty"`
}

func (c *CodeableConcept) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if len(c.Coding) > 0 { m["coding"] = c.Coding }
	if c.Text != "" { m["text"] = c.Text }
	return m
}

type Reference struct {
	Reference string `json:"reference,omitempty"`
	Type      string `json:"type,omitempty"`
	Identifier interface{} `json:"identifier,omitempty"`
	Display   string `json:"display,omitempty"`
}

func (r *Reference) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if r.Reference != "" { m["reference"] = r.Reference }
	if r.Type != "" { m["type"] = r.Type }
	if r.Identifier != nil { m["identifier"] = r.Identifier }
	if r.Display != "" { m["display"] = r.Display }
	return m
}

type ParameterComponent struct {
	Name    string             `json:"name,omitempty"`
	Value   interface{}        `json:"value,omitempty"`
	Element map[string]interface{} `json:"_value,omitempty"`
}
