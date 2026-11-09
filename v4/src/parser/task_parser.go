package parser

// TaskParser extracts fields from a FHIR Task resource map.
type TaskParser struct {
	task map[string]any
}

// NewTaskParser creates a new TaskParser instance.
func NewTaskParser(task map[string]any) *TaskParser {
	if task == nil {
		task = map[string]any{}
	}
	return &TaskParser{task: task}
}

func (p *TaskParser) GetIdentifiers() any {
	return p.task["identifier"]
}

func (p *TaskParser) GetInstantiatesUri() string {
	return strVal(p.task["instantiatesUri"])
}

func (p *TaskParser) GetStatus() string {
	return strVal(p.task["status"])
}

func (p *TaskParser) GetIntent() string {
	return strVal(p.task["intent"])
}

func (p *TaskParser) GetPriority() string {
	return strVal(p.task["priority"])
}

func (p *TaskParser) GetDescription() string {
	return strVal(p.task["description"])
}

func (p *TaskParser) GetFor() any {
	return p.task["for"]
}

func (p *TaskParser) GetForReference() string {
	return removePrefix(nestedStr(p.task, "for", "reference"), "Patient/")
}

func (p *TaskParser) GetEncounter() any {
	return p.task["encounter"]
}

func (p *TaskParser) GetEncounterReference() string {
	return removePrefix(nestedStr(p.task, "encounter", "reference"), "Encounter/")
}

func (p *TaskParser) GetAuthoredOn() string {
	return strVal(p.task["authoredOn"])
}

func (p *TaskParser) GetRequester() any {
	return p.task["requester"]
}

func (p *TaskParser) GetRequesterReference() string {
	return removePrefix(nestedStr(p.task, "requester", "reference"), "Practitioner/")
}

func (p *TaskParser) GetInputs() any {
	return p.task["input"]
}
