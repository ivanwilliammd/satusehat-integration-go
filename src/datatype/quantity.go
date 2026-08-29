package datatype

type Quantity struct {
	Value      *float64 `json:"value,omitempty"`
	Comparator *string  `json:"comparator,omitempty"`
	Unit       *string  `json:"unit,omitempty"`
	System     *string  `json:"system,omitempty"`
	Code       *string  `json:"code,omitempty"`
}

func (q *Quantity) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if q.Value != nil { m["value"] = *q.Value }
	if q.Comparator != nil { m["comparator"] = *q.Comparator }
	if q.Unit != nil { m["unit"] = *q.Unit }
	if q.System != nil { m["system"] = *q.System }
	if q.Code != nil { m["code"] = *q.Code }
	return m
}
