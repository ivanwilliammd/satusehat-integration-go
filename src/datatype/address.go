package datatype

type Period struct {
	Start *string `json:"start,omitempty"`
	End   *string `json:"end,omitempty"`
}

func (p *Period) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if p.Start != nil { m["start"] = *p.Start }
	if p.End != nil { m["end"] = *p.End }
	return m
}

type Address struct {
	Use        *string   `json:"use,omitempty"`
	Type       *string   `json:"type,omitempty"`
	Text       *string   `json:"text,omitempty"`
	Line       []string  `json:"line,omitempty"`
	City       *string   `json:"city,omitempty"`
	District   *string   `json:"district,omitempty"`
	State      *string   `json:"state,omitempty"`
	PostalCode *string   `json:"postalCode,omitempty"`
	Country    *string   `json:"country,omitempty"`
	Period     *Period   `json:"period,omitempty"`
}

func (a *Address) ToArray() map[string]interface{} {
	result := map[string]interface{}{}
	if a.Use != nil { result["use"] = *a.Use }
	if a.Type != nil { result["type"] = *a.Type }
	if a.Text != nil { result["text"] = *a.Text }
	if len(a.Line) > 0 { result["line"] = a.Line }
	if a.City != nil && *a.City != "" { result["city"] = *a.City }
	if a.District != nil && *a.District != "" { result["district"] = *a.District }
	if a.State != nil && *a.State != "" { result["state"] = *a.State }
	if a.PostalCode != nil && *a.PostalCode != "" { result["postalCode"] = *a.PostalCode }
	if a.Country != nil && *a.Country != "" { result["country"] = *a.Country }
	if a.Period != nil { result["period"] = a.Period.ToArray() }
	return result
}
