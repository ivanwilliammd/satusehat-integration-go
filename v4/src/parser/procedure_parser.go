package parser

// ProcedureParser extracts fields from a FHIR Procedure resource map.
type ProcedureParser struct {
	procedure map[string]any
}

// NewProcedureParser creates a new ProcedureParser instance.
func NewProcedureParser(procedure map[string]any) *ProcedureParser {
	if procedure == nil {
		procedure = map[string]any{}
	}
	return &ProcedureParser{procedure: procedure}
}

func (p *ProcedureParser) GetIdentifiers() any {
	return p.procedure["identifier"]
}

func (p *ProcedureParser) GetStatus() string {
	return strVal(p.procedure["status"])
}

func (p *ProcedureParser) GetCategory() any {
	return p.procedure["category"]
}

func (p *ProcedureParser) GetCode() any {
	return p.procedure["code"]
}

func (p *ProcedureParser) GetSubject() any {
	return p.procedure["subject"]
}

func (p *ProcedureParser) GetSubjectReference() string {
	return removePrefix(nestedStr(p.procedure, "subject", "reference"), "Patient/")
}

func (p *ProcedureParser) GetSubjectDisplay() string {
	return nestedStr(p.procedure, "subject", "display")
}

func (p *ProcedureParser) GetEncounter() any {
	return p.procedure["encounter"]
}

func (p *ProcedureParser) GetEncounterReference() string {
	return removePrefix(nestedStr(p.procedure, "encounter", "reference"), "Encounter/")
}

func (p *ProcedureParser) GetEncounterDisplay() string {
	return nestedStr(p.procedure, "encounter", "display")
}

func (p *ProcedureParser) GetPerformers() any {
	return p.procedure["performer"]
}

func (p *ProcedureParser) GetPerformerReferences() []string {
	items := sliceMaps(p.procedure["performer"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, removePrefix(nestedStr(item, "actor", "reference"), "Practitioner/"))
	}
	return out
}

func (p *ProcedureParser) GetPerformerDisplays() []string {
	items := sliceMaps(p.procedure["performer"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, nestedStr(item, "actor", "display"))
	}
	return out
}

func (p *ProcedureParser) GetPerformedPeriod() any {
	return p.procedure["performedPeriod"]
}

func (p *ProcedureParser) GetReasonCodes() any {
	return p.procedure["reasonCode"]
}

func (p *ProcedureParser) GetBodySites() any {
	return p.procedure["bodySite"]
}

func (p *ProcedureParser) GetNotes() any {
	return p.procedure["note"]
}

func (p *ProcedureParser) GetUsedCodes() any {
	return p.procedure["usedCode"]
}

func (p *ProcedureParser) GetBasedOn() any {
	return p.procedure["basedOn"]
}

func (p *ProcedureParser) GetLocation() string {
	return removePrefix(nestedStr(p.procedure, "location", "reference"), "Location/")
}
