package datatype

type TriggerDefinition struct {
	Type       *string           `json:"type,omitempty"`
	EventName  *string           `json:"eventName,omitempty"`
	EventTiming interface{}       `json:"eventTiming,omitempty"`
	EventData  *DataRequirement  `json:"eventData,omitempty"`
}

func (t *TriggerDefinition) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if t.Type != nil { m["type"] = *t.Type }
	if t.EventName != nil { m["eventName"] = *t.EventName }
	if t.EventTiming != nil { m["eventTiming"] = t.eventTimingToArray() }
	if t.EventData != nil { m["eventData"] = t.EventData.ToArray() }
	return m
}

func (t *TriggerDefinition) eventTimingToArray() interface{} {
	if t.EventTiming == nil {
		return nil
	}
	switch v := t.EventTiming.(type) {
	case *Timing:
		return v.ToArray()
	case *Period:
		return v.ToArray()
	case string:
		return v
	default:
		return nil
	}
}

func (t *TriggerDefinition) SetEventTiming(eventTiming interface{}) *TriggerDefinition {
	t.EventTiming = eventTiming
	return t
}

func (t *TriggerDefinition) SetEventData(eventData *DataRequirement) *TriggerDefinition {
	t.EventData = eventData
	return t
}
