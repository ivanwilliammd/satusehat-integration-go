package parser

// EncounterParser extracts fields from a FHIR Encounter resource map.
type EncounterParser struct {
	encounter map[string]any
}

// NewEncounterParser creates a new EncounterParser instance.
func NewEncounterParser(encounter map[string]any) *EncounterParser {
	if encounter == nil {
		encounter = map[string]any{}
	}
	return &EncounterParser{encounter: encounter}
}

func (p *EncounterParser) GetStatus() string {
	return strVal(p.encounter["status"])
}

func (p *EncounterParser) GetLocations() any {
	return p.encounter["location"]
}

func (p *EncounterParser) GetLocationReferences() []string {
	items := sliceMaps(p.encounter["location"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, removePrefix(nestedStr(item, "location", "reference"), "Location/"))
	}
	return out
}

func (p *EncounterParser) GetLocationDisplays() []string {
	items := sliceMaps(p.encounter["location"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, nestedStr(item, "location", "display"))
	}
	return out
}

func (p *EncounterParser) GetSubject() any {
	return p.encounter["subject"]
}

func (p *EncounterParser) GetSubjectReference() string {
	return removePrefix(nestedStr(p.encounter, "subject", "reference"), "Patient/")
}

func (p *EncounterParser) GetSubjectDisplay() string {
	return nestedStr(p.encounter, "subject", "display")
}

func (p *EncounterParser) GetParticipants() any {
	return p.encounter["participant"]
}

func (p *EncounterParser) GetParticipantReferences() []string {
	items := sliceMaps(p.encounter["participant"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, removePrefix(nestedStr(item, "individual", "reference"), "Practitioner/"))
	}
	return out
}

func (p *EncounterParser) GetParticipantDisplays() []string {
	items := sliceMaps(p.encounter["participant"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, nestedStr(item, "individual", "display"))
	}
	return out
}

func (p *EncounterParser) GetDiagnosis() any {
	return p.encounter["diagnosis"]
}

func (p *EncounterParser) GetDiagnosisReferences() []string {
	items := sliceMaps(p.encounter["diagnosis"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, removePrefix(nestedStr(item, "condition", "reference"), "Condition/"))
	}
	return out
}

func (p *EncounterParser) GetDiagnosisDisplays() []string {
	items := sliceMaps(p.encounter["diagnosis"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, nestedStr(item, "condition", "display"))
	}
	return out
}

func (p *EncounterParser) GetRegistrationIds() []string {
	items := sliceMaps(p.encounter["identifier"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, strVal(item["value"]))
	}
	return out
}

func (p *EncounterParser) GetStatusHistory() any {
	return p.encounter["statusHistory"]
}

func (p *EncounterParser) GetConsultationMethod() any {
	return p.encounter["class"]
}

func (p *EncounterParser) GetServiceProvider() any {
	return p.encounter["serviceProvider"]
}

func (p *EncounterParser) GetPeriod() any {
	return p.encounter["period"]
}

func (p *EncounterParser) GetPeriodStart() string {
	return nestedStr(p.encounter, "period", "start")
}

func (p *EncounterParser) GetPeriodEnd() string {
	return nestedStr(p.encounter, "period", "end")
}
