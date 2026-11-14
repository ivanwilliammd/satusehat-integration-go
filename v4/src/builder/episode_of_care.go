package builder

import (
	"encoding/json"
	"fmt"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type EpisodeOfCareBuilder struct {
	ResourceType string
	Data         map[string]interface{}
	err          error
}

func NewEpisodeOfCareBuilder() *EpisodeOfCareBuilder {
	b := &EpisodeOfCareBuilder{ResourceType: "EpisodeOfCare", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "EpisodeOfCare"
	return b
}

func (b *EpisodeOfCareBuilder) SetId(id string) *EpisodeOfCareBuilder {
	b.Data["id"] = id
	return b
}

func (b *EpisodeOfCareBuilder) AddIdentifier(identifier *datatype.Identifier) *EpisodeOfCareBuilder {
	if b.Data["identifier"] == nil {
		b.Data["identifier"] = []*datatype.Identifier{}
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
	return b
}

func (b *EpisodeOfCareBuilder) SetStatus(status string) *EpisodeOfCareBuilder {
	b.Data["status"] = status
	return b
}

func (b *EpisodeOfCareBuilder) SetSubject(reference string, display ...string) *EpisodeOfCareBuilder {
	sub := datatype.Reference{Reference: reference}
	if len(display) > 0 {
		sub.Display = display[0]
	}
	b.Data["subject"] = sub.ToArray()
	return b
}

func (b *EpisodeOfCareBuilder) Build() map[string]interface{} {
	return b.Data
}

func (b *EpisodeOfCareBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Data)
}

// SetPatient sets patient from a Reference, as the PHP builder does.
func (b *EpisodeOfCareBuilder) SetPatient(patient *datatype.Reference) *EpisodeOfCareBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}

// Err returns the first validation error recorded by a setter, if any.
func (b *EpisodeOfCareBuilder) Err() error { return b.err }

// EpisodeOfCareStatuses lists the values accepted by AddStatusHistory, mirroring the canonical builder.
var EpisodeOfCareStatuses = []string{"planned", "waitlist", "active", "onhold", "finished", "cancelled", "entered-in-error"}

// AddStatusHistory appends a status entry. An unknown status is recorded as an error and nothing is written,
// matching the canonical builder that throws — writing nothing silently would lose the entry.
func (b *EpisodeOfCareBuilder) AddStatusHistory(status, start string, end *string) *EpisodeOfCareBuilder {
	sah := false
	for _, s := range EpisodeOfCareStatuses {
		if s == status {
			sah = true
			break
		}
	}
	if !sah {
		b.err = fmt.Errorf("invalid status: %s", status)
		return b
	}
	periode := map[string]interface{}{"start": start}
	if end != nil && *end != "" {
		periode["end"] = *end
	}
	pushPath(b.Data, "statusHistory", map[string]interface{}{"status": status, "period": periode})
	return b
}

// SetManagingOrganization sets managingOrganization from a reference.
func (b *EpisodeOfCareBuilder) SetManagingOrganization(organization *datatype.Reference) *EpisodeOfCareBuilder {
	b.Data["managingOrganization"] = organization.ToArray()
	return b
}

// AddType appends a type coding.
func (b *EpisodeOfCareBuilder) AddType(episodeType *datatype.CodeableConcept) *EpisodeOfCareBuilder {
	pushPath(b.Data, "type", episodeType.ToArray())
	return b
}

// SetPeriod sets period from a period.
func (b *EpisodeOfCareBuilder) SetPeriod(period *datatype.Period) *EpisodeOfCareBuilder {
	b.Data["period"] = period.ToArray()
	return b
}

// AddDiagnosis appends a diagnosis with its role and optional rank.
func (b *EpisodeOfCareBuilder) AddDiagnosis(condition *datatype.Reference, role *datatype.CodeableConcept, rank *int) *EpisodeOfCareBuilder {
	diagnosis := map[string]interface{}{"condition": condition.ToArray(), "role": role.ToArray()}
	if rank != nil {
		diagnosis["rank"] = *rank
	}
	pushPath(b.Data, "diagnosis", diagnosis)
	return b
}

// AddReferralRequest appends a referral request reference.
func (b *EpisodeOfCareBuilder) AddReferralRequest(referralRequest *datatype.Reference) *EpisodeOfCareBuilder {
	pushPath(b.Data, "referralRequest", referralRequest.ToArray())
	return b
}

// SetCareManager sets careManager from a reference.
func (b *EpisodeOfCareBuilder) SetCareManager(careManager *datatype.Reference) *EpisodeOfCareBuilder {
	b.Data["careManager"] = careManager.ToArray()
	return b
}

// AddTeam appends a team reference.
func (b *EpisodeOfCareBuilder) AddTeam(team *datatype.Reference) *EpisodeOfCareBuilder {
	pushPath(b.Data, "team", team.ToArray())
	return b
}

// AddAccount appends an account reference.
func (b *EpisodeOfCareBuilder) AddAccount(account *datatype.Reference) *EpisodeOfCareBuilder {
	pushPath(b.Data, "account", account.ToArray())
	return b
}
