package builder

import (
	"encoding/json"
	"strings"
)

type TaskBuilder struct {
	data map[string]interface{}
}

func NewTaskBuilder() *TaskBuilder {
	return &TaskBuilder{
		data: map[string]interface{}{"resourceType": "Task"},
	}
}

var validTaskStatuses = []string{
	"draft", "requested", "received", "accepted", "rejected", "ready",
	"cancelled", "in-progress", "on-hold", "failed", "completed", "entered-in-error",
}

var validTaskIntents = []string{
	"unknown", "proposal", "plan", "order", "original-order",
	"reflex-order", "filler-order", "instance-order", "option",
}

func (b *TaskBuilder) autoPrefix(ref string, resourceType string) string {
	if !strings.HasPrefix(ref, "urn:") && !strings.HasPrefix(ref, "http://") &&
		!strings.HasPrefix(ref, "https://") && !strings.Contains(ref, "/") {
		return resourceType + "/" + ref
	}
	return ref
}

func (b *TaskBuilder) ref(ref string, resourceType string, display ...string) map[string]interface{} {
	result := map[string]interface{}{"reference": b.autoPrefix(ref, resourceType)}
	if len(display) > 0 && display[0] != "" {
		result["display"] = display[0]
	}
	return result
}

func (b *TaskBuilder) SetID(id string) *TaskBuilder {
	b.data["id"] = id
	return b
}

func (b *TaskBuilder) AddIdentifier(system, value string) *TaskBuilder {
	ids, _ := b.data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.data["identifier"] = ids
	return b
}

func (b *TaskBuilder) SetInstantiatesCanonical(val string) *TaskBuilder {
	b.data["instantiatesCanonical"] = val
	return b
}

func (b *TaskBuilder) SetInstantiatesUri(val string) *TaskBuilder {
	b.data["instantiatesUri"] = val
	return b
}

func (b *TaskBuilder) SetStatus(status string) *TaskBuilder {
	valid := false
	for _, s := range validTaskStatuses {
		if s == status {
			valid = true
			break
		}
	}
	if !valid {
		panic("Invalid Task status: " + status)
	}
	b.data["status"] = status
	return b
}

func (b *TaskBuilder) SetStatusReason(code string, displayAndSystem ...string) *TaskBuilder {
	display := code
	if len(displayAndSystem) > 0 && displayAndSystem[0] != "" {
		display = displayAndSystem[0]
	}
	coding := map[string]string{"code": code, "display": display}
	if len(displayAndSystem) > 1 && displayAndSystem[1] != "" {
		coding["system"] = displayAndSystem[1]
	}
	b.data["statusReason"] = map[string]interface{}{"coding": []map[string]string{coding}}
	return b
}

func (b *TaskBuilder) SetBusinessStatus(code string, displayAndSystem ...string) *TaskBuilder {
	display := code
	if len(displayAndSystem) > 0 && displayAndSystem[0] != "" {
		display = displayAndSystem[0]
	}
	coding := map[string]string{"code": code, "display": display}
	if len(displayAndSystem) > 1 && displayAndSystem[1] != "" {
		coding["system"] = displayAndSystem[1]
	}
	b.data["businessStatus"] = map[string]interface{}{"coding": []map[string]string{coding}}
	return b
}

func (b *TaskBuilder) SetIntent(intent string) *TaskBuilder {
	valid := false
	for _, i := range validTaskIntents {
		if i == intent {
			valid = true
			break
		}
	}
	if !valid {
		panic("Invalid Task intent: " + intent)
	}
	b.data["intent"] = intent
	return b
}

func (b *TaskBuilder) SetPriority(priority string) *TaskBuilder {
	b.data["priority"] = priority
	return b
}

func (b *TaskBuilder) SetCode(code string, displayAndSystem ...string) *TaskBuilder {
	display := code
	if len(displayAndSystem) > 0 && displayAndSystem[0] != "" {
		display = displayAndSystem[0]
	}
	coding := map[string]string{"code": code, "display": display}
	if len(displayAndSystem) > 1 && displayAndSystem[1] != "" {
		coding["system"] = displayAndSystem[1]
	}
	b.data["code"] = map[string]interface{}{"coding": []map[string]string{coding}}
	return b
}

func (b *TaskBuilder) SetDescription(desc string) *TaskBuilder {
	b.data["description"] = desc
	return b
}

func (b *TaskBuilder) SetFocus(ref string, display ...string) *TaskBuilder {
	b.data["focus"] = b.ref(ref, "QuestionnaireResponse", display...)
	return b
}

func (b *TaskBuilder) SetFor(ref string, display ...string) *TaskBuilder {
	b.data["for"] = b.ref(ref, "Patient", display...)
	return b
}

func (b *TaskBuilder) SetEncounter(ref string, display ...string) *TaskBuilder {
	b.data["encounter"] = b.ref(ref, "Encounter", display...)
	return b
}

func (b *TaskBuilder) SetExecutionPeriod(start string, end ...string) *TaskBuilder {
	period := map[string]string{"start": start}
	if len(end) > 0 && end[0] != "" {
		period["end"] = end[0]
	}
	b.data["executionPeriod"] = period
	return b
}

func (b *TaskBuilder) SetAuthoredOn(dt string) *TaskBuilder {
	b.data["authoredOn"] = dt
	return b
}

func (b *TaskBuilder) SetLastModified(dt string) *TaskBuilder {
	b.data["lastModified"] = dt
	return b
}

func (b *TaskBuilder) SetRequester(ref string, display ...string) *TaskBuilder {
	b.data["requester"] = b.ref(ref, "Practitioner", display...)
	return b
}

func (b *TaskBuilder) SetOwner(ref string, display ...string) *TaskBuilder {
	b.data["owner"] = b.ref(ref, "Practitioner", display...)
	return b
}

func (b *TaskBuilder) SetLocation(ref string, display ...string) *TaskBuilder {
	b.data["location"] = b.ref(ref, "Location", display...)
	return b
}

func (b *TaskBuilder) AddReasonCode(cc map[string]interface{}) *TaskBuilder {
	items, _ := b.data["reasonCode"].([]map[string]interface{})
	items = append(items, cc)
	b.data["reasonCode"] = items
	return b
}

func (b *TaskBuilder) AddReasonReference(ref string) *TaskBuilder {
	items, _ := b.data["reasonReference"].([]map[string]interface{})
	items = append(items, map[string]interface{}{"reference": ref})
	b.data["reasonReference"] = items
	return b
}

func (b *TaskBuilder) AddInput(typeText string, value string) *TaskBuilder {
	items, _ := b.data["input"].([]map[string]interface{})
	items = append(items, map[string]interface{}{
		"type":        map[string]string{"text": typeText},
		"valueString": value,
	})
	b.data["input"] = items
	return b
}

func (b *TaskBuilder) AddOutput(typeText string, value string) *TaskBuilder {
	items, _ := b.data["output"].([]map[string]interface{})
	items = append(items, map[string]interface{}{
		"type":        map[string]string{"text": typeText},
		"valueString": value,
	})
	b.data["output"] = items
	return b
}

func (b *TaskBuilder) AddRestriction(ref string, repetitions ...int) *TaskBuilder {
	items, _ := b.data["restriction"].([]map[string]interface{})
	restriction := map[string]interface{}{
		"requester": map[string]interface{}{
			"reference": b.autoPrefix(ref, "Patient"),
		},
	}
	if len(repetitions) > 0 && repetitions[0] > 0 {
		restriction["repetitions"] = repetitions[0]
	}
	items = append(items, restriction)
	b.data["restriction"] = items
	return b
}

func (b *TaskBuilder) AddNote(text string) *TaskBuilder {
	items, _ := b.data["note"].([]map[string]interface{})
	items = append(items, map[string]interface{}{"text": text})
	b.data["note"] = items
	return b
}

func (b *TaskBuilder) AddExtension(url string, value string) *TaskBuilder {
	items, _ := b.data["extension"].([]map[string]interface{})
	items = append(items, map[string]interface{}{
		"url":         url,
		"valueString": value,
	})
	b.data["extension"] = items
	return b
}

func (b *TaskBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *TaskBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.data)
}
