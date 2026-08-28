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

type Address struct {
	Use        string   `json:"use,omitempty"`
	Type       string   `json:"type,omitempty"`
	Line       []string `json:"line,omitempty"`
	City       string   `json:"city,omitempty"`
	District   string   `json:"district,omitempty"`
	State      string   `json:"state,omitempty"`
	PostalCode string   `json:"postalCode,omitempty"`
	Country    string   `json:"country,omitempty"`
}

func (a *Address) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if len(a.Line) > 0 { m["line"] = a.Line }
	if a.City != "" { m["city"] = a.City }
	if a.District != "" { m["district"] = a.District }
	if a.State != "" { m["state"] = a.State }
	if a.PostalCode != "" { m["postalCode"] = a.PostalCode }
	if a.Country != "" { m["country"] = a.Country }
	return m
}

type ContactPoint struct {
	System string `json:"system,omitempty"`
	Value  string `json:"value,omitempty"`
	Use    string `json:"use,omitempty"`
}

func (c *ContactPoint) ToArray() map[string]interface{} {
	return map[string]interface{}{"system": c.System, "value": c.Value, "use": c.Use}
}

type CodeableConcept struct {
	Coding []Coding `json:"coding,omitempty"`
	Text   string   `json:"text,omitempty"`
}

func (cc *CodeableConcept) ToArray() map[string]interface{} {
	codings := make([]map[string]interface{}, len(cc.Coding))
	for i, c := range cc.Coding {
		codings[i] = c.ToArray()
	}
	return map[string]interface{}{"coding": codings, "text": cc.Text}
}

type Coding struct {
	System  string `json:"system,omitempty"`
	Code    string `json:"code,omitempty"`
	Display string `json:"display,omitempty"`
}

func (c *Coding) ToArray() map[string]interface{} {
	return map[string]interface{}{"system": c.System, "code": c.Code, "display": c.Display}
}

type Reference struct {
	Reference string `json:"reference,omitempty"`
	Display   string `json:"display,omitempty"`
}

func (r *Reference) ToArray() map[string]interface{} {
	return map[string]interface{}{"reference": r.Reference, "display": r.Display}
}
