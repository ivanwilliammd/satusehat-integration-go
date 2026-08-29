package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type GoalBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewGoalBuilder() *GoalBuilder {
	b := &GoalBuilder{ResourceType: "Goal", Data: make(map[string]interface{})}
	return b
}

func (b *GoalBuilder) SetId(id string) *GoalBuilder { b.Data["id"] = id; return b }

func (b *GoalBuilder) AddIdentifier(id *datatype.Identifier) *GoalBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}

func (b *GoalBuilder) SetLifecycleStatus(status string) *GoalBuilder { b.Data["lifecycleStatus"] = status; return b }
func (b *GoalBuilder) SetAchievementStatus(status string) *GoalBuilder { b.Data["achievementStatus"] = map[string]interface{}{"coding": []interface{}{map[string]interface{}{"code": status}}}; return b }
func (b *GoalBuilder) SetCategory(code *datatype.CodeableConcept) *GoalBuilder { b.Data["category"] = []interface{}{code.ToArray()}; return b }
func (b *GoalBuilder) SetDescription(description *datatype.CodeableConcept) *GoalBuilder { b.Data["description"] = description.ToArray(); return b }
func (b *GoalBuilder) SetSubject(subjectRef string) *GoalBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *GoalBuilder) SetStart(start *datatype.CodeableConcept) *GoalBuilder { b.Data["start"] = start.ToArray(); return b }
func (b *GoalBuilder) SetTargetMetric(code *datatype.CodeableConcept) *GoalBuilder { b.Data["target"] = []interface{}{map[string]interface{}{"metric": code.ToArray()}}; return b }
func (b *GoalBuilder) SetTargetDetail(minValue, maxValue string) *GoalBuilder {
	tgt := b.Data["target"].([]interface{})[0].(map[string]interface{})
	tgt["detailString"] = minValue
	if maxValue != "" { tgt["max"] = map[string]interface{}{"value": maxValue} }
	return b
}
func (b *GoalBuilder) SetOnsetDateTime(onset string) *GoalBuilder { b.Data["onsetDateTime"] = onset; return b }
func (b *GoalBuilder) SetExpressions(expr string, text string) *GoalBuilder {
	exprArr := b.Data["expressions"].([]interface{})
	last := exprArr[len(exprArr)-1].(map[string]interface{})
	last["text"] = text
	last["dynamicValue"] = map[string]interface{}{"expression": expr}
	return b
}

func (b *GoalBuilder) Build() map[string]interface{} { return b.Data }
