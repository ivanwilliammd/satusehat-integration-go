package parser

// GenomicStudyParser extracts fields from a FHIR GenomicStudy resource map.
type GenomicStudyParser struct {
	genomicStudy map[string]any
}

// NewGenomicStudyParser creates a new GenomicStudyParser instance.
func NewGenomicStudyParser(genomicStudy map[string]any) *GenomicStudyParser {
	if genomicStudy == nil {
		genomicStudy = map[string]any{}
	}
	return &GenomicStudyParser{genomicStudy: genomicStudy}
}

func (p *GenomicStudyParser) GetIdentifiers() any {
	return p.genomicStudy["identifier"]
}

func (p *GenomicStudyParser) GetStatus() string {
	return strVal(p.genomicStudy["status"])
}

func (p *GenomicStudyParser) GetTypes() any {
	return p.genomicStudy["type"]
}

func (p *GenomicStudyParser) GetSubjectReference() string {
	return nestedStr(p.genomicStudy, "subject", "reference")
}

func (p *GenomicStudyParser) GetEncounterReference() string {
	return nestedStr(p.genomicStudy, "encounter", "reference")
}

func (p *GenomicStudyParser) GetStartDate() string {
	return strVal(p.genomicStudy["startDate"])
}

func (p *GenomicStudyParser) GetBasedOnReferences() any {
	return p.genomicStudy["basedOn"]
}

func (p *GenomicStudyParser) GetReferrerReference() string {
	return nestedStr(p.genomicStudy, "referrer", "reference")
}

func (p *GenomicStudyParser) GetInterpreterReferences() any {
	return p.genomicStudy["interpreter"]
}

func (p *GenomicStudyParser) GetReasons() any {
	return p.genomicStudy["reason"]
}

func (p *GenomicStudyParser) GetNotes() any {
	return p.genomicStudy["note"]
}

func (p *GenomicStudyParser) GetDescription() string {
	return strVal(p.genomicStudy["description"])
}

func (p *GenomicStudyParser) GetAnalyses() any {
	return p.genomicStudy["analysis"]
}

func (p *GenomicStudyParser) GetPerformers() any {
	return p.genomicStudy["performer"]
}
