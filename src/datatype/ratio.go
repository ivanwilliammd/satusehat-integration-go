package datatype

type Ratio struct {
	Numerator   *Quantity `json:"numerator,omitempty"`
	Denominator *Quantity `json:"denominator,omitempty"`
}

func (r *Ratio) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if r.Numerator != nil { m["numerator"] = r.Numerator.ToArray() }
	if r.Denominator != nil { m["denominator"] = r.Denominator.ToArray() }
	return m
}
