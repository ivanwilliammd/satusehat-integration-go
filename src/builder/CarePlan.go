package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type CarePlanBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewCarePlanBuilder() *CarePlanBuilder {
	b := &CarePlanBuilder{ResourceType: "CarePlan", Data: make(map[string]interface{})}
	return b
}

func (b *CarePlanBuilder) SetId(id string) *CarePlanBuilder { b.Data["id"] = id; return b }

func (b *CarePlanBuilder) AddIdentifier(id *datatype.Identifier) *CarePlanBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}

func (b *CarePlanBuilder) SetStatus(status string) *CarePlanBuilder { b.Data["status"] = status; return b }
func (b *CarePlanBuilder) SetIntent(intent string) *CarePlanBuilder { b.Data["intent"] = intent; return b }
func (b *CarePlanBuilder) SetCategory(code *datatype.CodeableConcept) *CarePlanBuilder { b.Data["category"] = []interface{}{code.ToArray()}; return b }
func (b *CarePlanBuilder) SetTitle(title string) *CarePlanBuilder { b.Data["title"] = title; return b }
func (b *CarePlanBuilder) SetDescription(description string) *CarePlanBuilder { b.Data["description"] = description; return b }
func (b *CarePlanBuilder) SetSubject(subjectRef string) *CarePlanBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *CarePlanBuilder) SetEncounter(encRef string) *CarePlanBuilder { b.Data["encounter"] = map[string]interface{}{"reference": encRef}; return b }
func (b *CarePlanBuilder) SetPeriod(start string, end string) *CarePlanBuilder {
	p := map[string]interface{}{"start": start}
	if end != "" { p["end"] = end }
	b.Data["period"] = p
	return b
}
func (b *CarePlanBuilder) SetAuthor(authorRef string) *CarePlanBuilder { b.Data["author"] = map[string]interface{}{"reference": authorRef}; return b }
func (b *CarePlanBuilder) AddContributor(contributorRef string) *CarePlanBuilder {
	if _, ok := b.Data["contributor"]; !ok { b.Data["contributor"] = make([]interface{}, 0) }
	b.Data["contributor"] = append(b.Data["contributor"].([]interface{}), map[string]interface{}{"reference": contributorRef})
	return b
}
func (b *CarePlanBuilder) AddCareTeam(careTeamRef string) *CarePlanBuilder {
	if _, ok := b.Data["careTeam"]; !ok { b.Data["careTeam"] = make([]interface{}, 0) }
	b.Data["careTeam"] = append(b.Data["careTeam"].([]interface{}), map[string]interface{}{"reference": careTeamRef})
	return b
}
func (b *CarePlanBuilder) AddAddresses(conditionRef string) *CarePlanBuilder {
	if _, ok := b.Data["addresses"]; !ok { b.Data["addresses"] = make([]interface{}, 0) }
	b.Data["addresses"] = append(b.Data["addresses"].([]interface{}), map[string]interface{}{"reference": conditionRef})
	return b
}
func (b *CarePlanBuilder) AddActivity(detailCode *datatype.CodeableConcept, status string) *CarePlanBuilder {
	if _, ok := b.Data["activity"]; !ok { b.Data["activity"] = make([]interface{}, 0) }
	b.Data["activity"] = append(b.Data["activity"].([]interface{}), map[string]interface{}{
		"detail": map[string]interface{}{"code": detailCode.ToArray(), "status": status},
	})
	return b
}

func (b *CarePlanBuilder) Build() map[string]interface{} { return b.Data }
