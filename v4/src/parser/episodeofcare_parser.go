package parser

// EpisodeOfCareParser extracts fields from a FHIR EpisodeOfCare resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/EpisodeOfCare.json, copied unchanged to
// src/parser/testdata/EpisodeOfCare.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type EpisodeOfCareParser struct {
	payload map[string]any
}

// NewEpisodeOfCareParser creates a new EpisodeOfCareParser instance.
func NewEpisodeOfCareParser(payload map[string]any) *EpisodeOfCareParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &EpisodeOfCareParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *EpisodeOfCareParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *EpisodeOfCareParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetId returns the id element unchanged.
func (p *EpisodeOfCareParser) GetId() any {
	return p.payload["id"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *EpisodeOfCareParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *EpisodeOfCareParser) GetStatus() any {
	return p.payload["status"]
}

// GetStatusHistory returns the statusHistory element unchanged.
func (p *EpisodeOfCareParser) GetStatusHistory() any {
	return p.payload["statusHistory"]
}

// GetType returns the type element unchanged.
func (p *EpisodeOfCareParser) GetType() any {
	return p.payload["type"]
}

// GetDiagnosis returns the diagnosis element unchanged.
func (p *EpisodeOfCareParser) GetDiagnosis() any {
	return p.payload["diagnosis"]
}

// GetPatient returns the patient element unchanged.
func (p *EpisodeOfCareParser) GetPatient() any {
	return p.payload["patient"]
}

// GetManagingOrganization returns the managingOrganization element unchanged.
func (p *EpisodeOfCareParser) GetManagingOrganization() any {
	return p.payload["managingOrganization"]
}

// GetPeriod returns the period element unchanged.
func (p *EpisodeOfCareParser) GetPeriod() any {
	return p.payload["period"]
}

// GetCareManager returns the careManager element unchanged.
func (p *EpisodeOfCareParser) GetCareManager() any {
	return p.payload["careManager"]
}
