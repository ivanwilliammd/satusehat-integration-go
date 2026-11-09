package parser

// DiagnosticReportParser extracts fields from a FHIR DiagnosticReport resource map.
type DiagnosticReportParser struct {
	diagnosticReport map[string]any
}

// NewDiagnosticReportParser creates a new DiagnosticReportParser instance.
func NewDiagnosticReportParser(diagnosticReport map[string]any) *DiagnosticReportParser {
	if diagnosticReport == nil {
		diagnosticReport = map[string]any{}
	}
	return &DiagnosticReportParser{diagnosticReport: diagnosticReport}
}

func (p *DiagnosticReportParser) GetIdentifiers() any {
	return p.diagnosticReport["identifier"]
}

func (p *DiagnosticReportParser) GetStatus() string {
	return strVal(p.diagnosticReport["status"])
}

func (p *DiagnosticReportParser) GetCategories() any {
	return p.diagnosticReport["category"]
}

func (p *DiagnosticReportParser) GetCode() any {
	return p.diagnosticReport["code"]
}

func (p *DiagnosticReportParser) GetSubject() any {
	return p.diagnosticReport["subject"]
}

func (p *DiagnosticReportParser) GetSubjectReference() string {
	return removePrefix(nestedStr(p.diagnosticReport, "subject", "reference"), "Patient/")
}

func (p *DiagnosticReportParser) GetEncounter() any {
	return p.diagnosticReport["encounter"]
}

func (p *DiagnosticReportParser) GetEncounterReference() string {
	return removePrefix(nestedStr(p.diagnosticReport, "encounter", "reference"), "Encounter/")
}

func (p *DiagnosticReportParser) GetEffectiveDateTime() string {
	return strVal(p.diagnosticReport["effectiveDateTime"])
}

func (p *DiagnosticReportParser) GetIssued() string {
	return strVal(p.diagnosticReport["issued"])
}

func (p *DiagnosticReportParser) GetPerformers() any {
	return p.diagnosticReport["performer"]
}

func (p *DiagnosticReportParser) GetPerformerReferences() []string {
	items := sliceMaps(p.diagnosticReport["performer"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, removePrefix(strVal(item["reference"]), "Practitioner/"))
	}
	return out
}

func (p *DiagnosticReportParser) GetResults() any {
	return p.diagnosticReport["result"]
}

func (p *DiagnosticReportParser) GetResultReferences() []string {
	items := sliceMaps(p.diagnosticReport["result"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, removePrefix(strVal(item["reference"]), "Observation/"))
	}
	return out
}

func (p *DiagnosticReportParser) GetSpecimens() any {
	return p.diagnosticReport["specimen"]
}

func (p *DiagnosticReportParser) GetSpecimenReferences() []string {
	items := sliceMaps(p.diagnosticReport["specimen"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, removePrefix(strVal(item["reference"]), "Specimen/"))
	}
	return out
}

func (p *DiagnosticReportParser) GetConclusionCodes() any {
	return p.diagnosticReport["conclusionCode"]
}

func (p *DiagnosticReportParser) GetBasedOn() any {
	return p.diagnosticReport["basedOn"]
}

func (p *DiagnosticReportParser) GetConclusion() string {
	return strVal(p.diagnosticReport["conclusion"])
}
