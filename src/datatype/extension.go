package datatype

type Extension struct {
	URL  string       `json:"url"`
	Value interface{} `json:"value,omitempty"`
}

func (e *Extension) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	m["url"] = e.URL
	if e.Value != nil {
		m["value"] = e.Value
	}
	return m
}
