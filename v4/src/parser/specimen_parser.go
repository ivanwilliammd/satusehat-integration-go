package parser

// SpecimenParser extracts fields from a FHIR Specimen resource map.
type SpecimenParser struct {
	specimen map[string]any
}

// NewSpecimenParser creates a new SpecimenParser instance.
func NewSpecimenParser(specimen map[string]any) *SpecimenParser {
	if specimen == nil {
		specimen = map[string]any{}
	}
	return &SpecimenParser{specimen: specimen}
}

func (p *SpecimenParser) GetIdentifiers() any {
	return p.specimen["identifier"]
}

func (p *SpecimenParser) GetStatus() string {
	return strVal(p.specimen["status"])
}

func (p *SpecimenParser) GetType() any {
	return p.specimen["type"]
}

func (p *SpecimenParser) GetSubject() any {
	return p.specimen["subject"]
}

func (p *SpecimenParser) GetSubjectReference() string {
	return removePrefix(nestedStr(p.specimen, "subject", "reference"), "Patient/")
}

func (p *SpecimenParser) GetSubjectDisplay() string {
	return nestedStr(p.specimen, "subject", "display")
}

func (p *SpecimenParser) GetReceivedTime() string {
	return strVal(p.specimen["receivedTime"])
}

func (p *SpecimenParser) GetRequests() any {
	return p.specimen["request"]
}

func (p *SpecimenParser) GetCollectedDateTime() string {
	return nestedStr(p.specimen, "collection", "collectedDateTime")
}

func (p *SpecimenParser) GetCollector() any {
	col := mapVal(p.specimen["collection"])
	if col == nil {
		return nil
	}
	return col["collector"]
}

func (p *SpecimenParser) GetCollectorReference() string {
	col := mapVal(p.specimen["collection"])
	if col == nil {
		return ""
	}
	return removePrefix(nestedStr(col, "collector", "reference"), "Practitioner/")
}

func (p *SpecimenParser) GetCollectorDisplay() string {
	col := mapVal(p.specimen["collection"])
	if col == nil {
		return ""
	}
	return nestedStr(col, "collector", "display")
}

func (p *SpecimenParser) GetFastingStatus() any {
	col := mapVal(p.specimen["collection"])
	if col == nil {
		return nil
	}
	return col["fastingStatusCodeableConcept"]
}

func (p *SpecimenParser) GetMethod() any {
	col := mapVal(p.specimen["collection"])
	if col == nil {
		return nil
	}
	return col["method"]
}

func (p *SpecimenParser) GetQuantity() any {
	col := mapVal(p.specimen["collection"])
	if col == nil {
		return nil
	}
	return col["quantity"]
}

func (p *SpecimenParser) GetBodySite() any {
	col := mapVal(p.specimen["collection"])
	if col == nil {
		return nil
	}
	return col["bodySite"]
}

func (p *SpecimenParser) GetConditions() any {
	return p.specimen["condition"]
}

func (p *SpecimenParser) GetExtensions() any {
	return p.specimen["extension"]
}

func (p *SpecimenParser) GetProcessing() any {
	return p.specimen["processing"]
}
