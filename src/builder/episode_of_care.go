package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

// EpisodeOfCareBuilder builds FHIR EpisodeOfCare payload
type EpisodeOfCareBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewEpisodeOfCareBuilder() *EpisodeOfCareBuilder {
	b := &EpisodeOfCareBuilder{ResourceType: "EpisodeOfCare", Data: make(map[string]interface{})}
	return b
}

func (b *EpisodeOfCareBuilder) setId(id string) *EpisodeOfCareBuilder {
	b.Data["id"] = id
	return b
}

func (b *EpisodeOfCareBuilder) addIdentifier(identifier *datatype.Identifier) *EpisodeOfCareBuilder {
	if _, ok := b.Data["identifier"]; !ok {
		b.Data["identifier"] = make([]interface{}, 0)
	}
	b.Data["identifier"] = append(b.Data["identifier"].([]interface{}), identifier.ToArray())
	return b
}

func (b *EpisodeOfCareBuilder) setStatus(status string) *EpisodeOfCareBuilder {
	b.Data["status"] = status
	return b
}

func (b *EpisodeOfCareBuilder) setPatient(patient *datatype.Reference) *EpisodeOfCareBuilder {
	b.Data["patient"] = patient.ToArray()
	return b
}

func (b *EpisodeOfCareBuilder) addManagingOrganization(organization *datatype.Reference) *EpisodeOfCareBuilder {
	b.Data["managingOrganization"] = organization.ToArray()
	return b
}

func (b *EpisodeOfCareBuilder) setPeriod(period *datatype.Period) *EpisodeOfCareBuilder {
	b.Data["period"] = period.ToArray()
	return b
}

func (b *EpisodeOfCareBuilder) addDiagnosis(
	condition *datatype.Reference,
	role *datatype.CodeableConcept,
	rank int,
) *EpisodeOfCareBuilder {
	diagnosis := map[string]interface{}{}
	if condition != nil {
		diagnosis["condition"] = condition.ToArray()
	}
	if role != nil {
		diagnosis["role"] = role.ToArray()
	}
	if rank != 0 {
		diagnosis["rank"] = rank
	}
	if _, ok := b.Data["diagnosis"]; !ok {
		b.Data["diagnosis"] = make([]interface{}, 0)
	}
	b.Data["diagnosis"] = append(b.Data["diagnosis"].([]interface{}), diagnosis)
	return b
}

func (b *EpisodeOfCareBuilder) addReferralRequest(referralRequest *datatype.Reference) *EpisodeOfCareBuilder {
	if _, ok := b.Data["referralRequest"]; !ok {
		b.Data["referralRequest"] = make([]interface{}, 0)
	}
	b.Data["referralRequest"] = append(b.Data["referralRequest"].([]interface{}), referralRequest.ToArray())
	return b
}

func (b *EpisodeOfCareBuilder) setCareManager(careManager *datatype.Reference) *EpisodeOfCareBuilder {
	b.Data["careManager"] = careManager.ToArray()
	return b
}

func (b *EpisodeOfCareBuilder) addTeamMember(member *datatype.Reference) *EpisodeOfCareBuilder {
	if _, ok := b.Data["careTeam"]; !ok {
		b.Data["careTeam"] = make([]interface{}, 0)
	}
	b.Data["careTeam"] = append(b.Data["careTeam"].([]interface{}), member.ToArray())
	return b
}

func (b *EpisodeOfCareBuilder) addExtension(url string, value interface{}, valueType string) *EpisodeOfCareBuilder {
	ext := map[string]interface{}{"url": url}
	if valueType != "" {
		capitalized := strings.ToUpper(valueType[:1]) + valueType[1:]
		ext["value"+capitalized] = value
	} else {
		ext["valueString"] = value
	}
	if _, ok := b.Data["extension"]; !ok {
		b.Data["extension"] = make([]interface{}, 0)
	}
	b.Data["extension"] = append(b.Data["extension"].([]interface{}), ext)
	return b
}

func (b *EpisodeOfCareBuilder) Build() map[string]interface{} {
	return b.Data
}
