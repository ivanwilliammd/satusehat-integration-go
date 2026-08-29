package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type TaskBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewTaskBuilder() *TaskBuilder {
	b := &TaskBuilder{ResourceType: "Task", Data: make(map[string]interface{})}
	return b
}

func (b *TaskBuilder) SetId(id string) *TaskBuilder { b.Data["id"] = id; return b }

func (b *TaskBuilder) AddIdentifier(id *datatype.Identifier) *TaskBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}

func (b *TaskBuilder) SetStatus(status string) *TaskBuilder { b.Data["status"] = status; return b }
func (b *TaskBuilder) SetIntent(intent string) *TaskBuilder { b.Data["intent"] = intent; return b }
func (b *TaskBuilder) SetPriority(priority string) *TaskBuilder { b.Data["priority"] = priority; return b }
func (b *TaskBuilder) SetCode(code *datatype.CodeableConcept) *TaskBuilder { b.Data["code"] = code.ToArray(); return b }
func (b *TaskBuilder) SetDescription(description string) *TaskBuilder { b.Data["description"] = description; return b }
func (b *TaskBuilder) SetFocus(focusRef string) *TaskBuilder { b.Data["focus"] = map[string]interface{}{"reference": focusRef}; return b }
func (b *TaskBuilder) SetFor(forRef string) *TaskBuilder { b.Data["for"] = map[string]interface{}{"reference": forRef}; return b }
func (b *TaskBuilder) SetExecutionPeriod(start string, end string) *TaskBuilder {
	p := map[string]interface{}{"start": start}
	if end != "" { p["end"] = end }
	b.Data["executionPeriod"] = p
	return b
}
func (b *TaskBuilder) SetAuthoredOn(authoredOn string) *TaskBuilder { b.Data["authoredOn"] = authoredOn; return b }
func (b *TaskBuilder) SetLastModified(lastModified string) *TaskBuilder { b.Data["lastModified"] = lastModified; return b }
func (b *TaskBuilder) SetRequester(requesterRef string) *TaskBuilder { b.Data["requester"] = map[string]interface{}{"reference": requesterRef}; return b }
func (b *TaskBuilder) SetOwner(ownerRef string) *TaskBuilder { b.Data["owner"] = map[string]interface{}{"reference": ownerRef}; return b }
func (b *TaskBuilder) SetReason(reason *datatype.CodeableConcept) *TaskBuilder { b.Data["reasonCode"] = reason.ToArray(); return b }
func (b *TaskBuilder) SetNote(note string) *TaskBuilder {
	b.Data["note"] = []interface{}{map[string]interface{}{"text": note}}; return b
}
func (b *TaskBuilder) AddInput(name string, value string) *TaskBuilder {
	if _, ok := b.Data["input"]; !ok { b.Data["input"] = make([]interface{}, 0) }
	b.Data["input"] = append(b.Data["input"].([]interface{}), map[string]interface{}{"type": map[string]interface{}{"text": name}, "valueString": value})
	return b
}
func (b *TaskBuilder) AddOutput(name string, value string) *TaskBuilder {
	if _, ok := b.Data["output"]; !ok { b.Data["output"] = make([]interface{}, 0) }
	b.Data["output"] = append(b.Data["output"].([]interface{}), map[string]interface{}{"type": map[string]interface{}{"text": name}, "valueString": value})
	return b
}

func (b *TaskBuilder) Build() map[string]interface{} { return b.Data }
