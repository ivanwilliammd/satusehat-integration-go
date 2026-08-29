package datatype

type TimingRepeat struct {
	Bounds       interface{} `json:"bounds,omitempty"`
	Count        *int        `json:"count,omitempty"`
	CountMax     *int        `json:"countMax,omitempty"`
	Duration     *float64    `json:"duration,omitempty"`
	DurationMax  *float64    `json:"durationMax,omitempty"`
	DurationUnit *string     `json:"durationUnit,omitempty"`
	Frequency    *int        `json:"frequency,omitempty"`
	FrequencyMax *int        `json:"frequencyMax,omitempty"`
	Period       *float64    `json:"period,omitempty"`
	PeriodMax    *float64    `json:"periodMax,omitempty"`
	PeriodUnit   *string     `json:"periodUnit,omitempty"`
	DayOfWeek    []string    `json:"dayOfWeek,omitempty"`
	TimeOfDay    []string    `json:"timeOfDay,omitempty"`
	When         []string    `json:"when,omitempty"`
	Offset       *int        `json:"offset,omitempty"`
}

func (t *TimingRepeat) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if t.Bounds != nil { m["bounds"] = t.boundsToArray() }
	if t.Count != nil { m["count"] = *t.Count }
	if t.CountMax != nil { m["countMax"] = *t.CountMax }
	if t.Duration != nil { m["duration"] = *t.Duration }
	if t.DurationMax != nil { m["durationMax"] = *t.DurationMax }
	if t.DurationUnit != nil { m["durationUnit"] = *t.DurationUnit }
	if t.Frequency != nil { m["frequency"] = *t.Frequency }
	if t.FrequencyMax != nil { m["frequencyMax"] = *t.FrequencyMax }
	if t.Period != nil { m["period"] = *t.Period }
	if t.PeriodMax != nil { m["periodMax"] = *t.PeriodMax }
	if t.PeriodUnit != nil { m["periodUnit"] = *t.PeriodUnit }
	if len(t.DayOfWeek) > 0 { m["dayOfWeek"] = t.DayOfWeek }
	if len(t.TimeOfDay) > 0 { m["timeOfDay"] = t.TimeOfDay }
	if len(t.When) > 0 { m["when"] = t.When }
	if t.Offset != nil { m["offset"] = *t.Offset }
	return m
}

func (t *TimingRepeat) boundsToArray() interface{} {
	if t.Bounds == nil {
		return nil
	}
	switch v := t.Bounds.(type) {
	case *Range:
		return v.ToArray()
	case *Period:
		return v.ToArray()
	case *Duration:
		return v.ToArray()
	case map[string]interface{}:
		return v
	default:
		return nil
	}
}

func (t *TimingRepeat) SetBounds(bounds interface{}) *TimingRepeat {
	t.Bounds = bounds
	return t
}

func (t *TimingRepeat) AddDayOfWeek(day string) *TimingRepeat {
	t.DayOfWeek = append(t.DayOfWeek, day)
	return t
}

func (t *TimingRepeat) AddTimeOfDay(time string) *TimingRepeat {
	t.TimeOfDay = append(t.TimeOfDay, time)
	return t
}

func (t *TimingRepeat) AddWhen(when string) *TimingRepeat {
	t.When = append(t.When, when)
	return t
}
