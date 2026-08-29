package datatype

type Timing struct {
	Event  []string          `json:"event,omitempty"`
	Repeat *TimingRepeat     `json:"repeat,omitempty"`
	Code   *CodeableConcept  `json:"code,omitempty"`
}

func (t *Timing) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if len(t.Event) > 0 { m["event"] = t.Event }
	if t.Repeat != nil { m["repeat"] = t.Repeat.ToArray() }
	if t.Code != nil { m["code"] = t.Code.ToArray() }
	return m
}

func (t *Timing) AddEvent(event string) *Timing {
	t.Event = append(t.Event, event)
	return t
}

func (t *Timing) SetRepeat(repeat *TimingRepeat) *Timing {
	t.Repeat = repeat
	return t
}

func (t *Timing) SetCode(code *CodeableConcept) *Timing {
	t.Code = code
	return t
}
