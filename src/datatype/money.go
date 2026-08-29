package datatype

type Money struct {
	Value    *float64 `json:"value,omitempty"`
	Currency *string  `json:"currency,omitempty"`
}

func (m *Money) ToArray() map[string]interface{} {
	result := map[string]interface{}{}
	if m.Value != nil { result["value"] = *m.Value }
	if m.Currency != nil { result["currency"] = *m.Currency }
	return result
}
