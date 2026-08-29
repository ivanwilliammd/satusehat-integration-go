package datatype

// SampledData represents a series of measurements taken by a device
type SampledData struct {
	Origin    *Quantity `json:"origin,omitempty"`
	Period    *float64  `json:"period,omitempty"`
	Factor    *float64  `json:"factor,omitempty"`
	LowerLimit *float64 `json:"lowerLimit,omitempty"`
	UpperLimit *float64 `json:"upperLimit,omitempty"`
	Dimensions *int     `json:"dimensions,omitempty"`
	Data      *string   `json:"data,omitempty"`
}

func (s *SampledData) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if s.Origin != nil { m["origin"] = s.Origin.ToArray() }
	if s.Period != nil { m["period"] = *s.Period }
	if s.Factor != nil { m["factor"] = *s.Factor }
	if s.LowerLimit != nil { m["lowerLimit"] = *s.LowerLimit }
	if s.UpperLimit != nil { m["upperLimit"] = *s.UpperLimit }
	if s.Dimensions != nil { m["dimensions"] = *s.Dimensions }
	if s.Data != nil { m["data"] = *s.Data }
	return m
}
