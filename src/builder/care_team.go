package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type CareTeamBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewCareTeamBuilder() *CareTeamBuilder {
	b := &CareTeamBuilder{ResourceType: "CareTeam", Data: make(map[string]interface{})}
	return b
}

func (b *CareTeamBuilder) SetId(id string) *CareTeamBuilder { b.Data["id"] = id; return b }

func (b *CareTeamBuilder) AddIdentifier(id *datatype.Identifier) *CareTeamBuilder {
	if _, ok := b.Data["identifier"]; !ok { b.Data["identifier"] = make([]interface{}, 0) }
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), id.ToArray())
	return b
}

func (b *CareTeamBuilder) SetStatus(status string) *CareTeamBuilder { b.Data["status"] = status; return b }
func (b *CareTeamBuilder) SetName(name string) *CareTeamBuilder { b.Data["name"] = name; return b }
func (b *CareTeamBuilder) SetSubject(subjectRef string) *CareTeamBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *CareTeamBuilder) SetPeriod(start string, end string) *CareTeamBuilder {
	p := map[string]interface{}{"start": start}
	if end != "" { p["end"] = end }
	b.Data["period"] = p
	return b
}
func (b *CareTeamBuilder) AddParticipant(role *datatype.CodeableConcept, memberRef string) *CareTeamBuilder {
	if _, ok := b.Data["participant"]; !ok { b.Data["participant"] = make([]interface{}, 0) }
	b.Data["participant"] = append(b.Data["participant"].([]interface{}), map[string]interface{}{
		"role":     []interface{}{role.ToArray()},
		"member":  map[string]interface{}{"reference": memberRef},
	})
	return b
}
func (b *CareTeamBuilder) SetManagingOrganization(orgRef string) *CareTeamBuilder {
	b.Data["managingOrganization"] = []interface{}{map[string]interface{}{"reference": orgRef}}
	return b
}

func (b *CareTeamBuilder) Build() map[string]interface{} { return b.Data }
