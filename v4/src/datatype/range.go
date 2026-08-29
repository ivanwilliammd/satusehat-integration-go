package datatype

type Range struct {
	Low  *Quantity `json:"low,omitempty"`
	High *Quantity `json:"high,omitempty"`
}

func (r *Range) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if r.Low != nil { m["low"] = r.Low.ToArray() }
	if r.High != nil { m["high"] = r.High.ToArray() }
	return m
}
