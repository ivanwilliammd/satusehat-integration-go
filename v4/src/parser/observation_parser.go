package parser

// ObservationParser extracts fields from a FHIR Observation resource map.
type ObservationParser struct {
	observation map[string]any
}

// NewObservationParser creates a new ObservationParser instance.
func NewObservationParser(observation map[string]any) *ObservationParser {
	if observation == nil {
		observation = map[string]any{}
	}
	return &ObservationParser{observation: observation}
}

func (p *ObservationParser) GetStatus() string {
	return strVal(p.observation["status"])
}

func (p *ObservationParser) GetCategories() any {
	return p.observation["category"]
}

func (p *ObservationParser) GetCode() any {
	return p.observation["code"]
}

func (p *ObservationParser) GetSubject() any {
	return p.observation["subject"]
}

func (p *ObservationParser) GetSubjectReference() string {
	return removePrefix(nestedStr(p.observation, "subject", "reference"), "Patient/")
}

func (p *ObservationParser) GetEncounter() any {
	return p.observation["encounter"]
}

func (p *ObservationParser) GetEncounterReference() string {
	return removePrefix(nestedStr(p.observation, "encounter", "reference"), "Encounter/")
}

func (p *ObservationParser) GetEffectiveDateTime() string {
	return strVal(p.observation["effectiveDateTime"])
}

func (p *ObservationParser) GetIssued() string {
	return strVal(p.observation["issued"])
}

func (p *ObservationParser) GetPerformers() any {
	return p.observation["performer"]
}

func (p *ObservationParser) GetPerformerReferences() []string {
	items := sliceMaps(p.observation["performer"])
	out := make([]string, 0, len(items))
	for _, item := range items {
		out = append(out, removePrefix(strVal(item["reference"]), "Practitioner/"))
	}
	return out
}

func (p *ObservationParser) GetValueQuantity() any {
	return p.observation["valueQuantity"]
}

func (p *ObservationParser) GetValueString() string {
	return strVal(p.observation["valueString"])
}

func (p *ObservationParser) GetValueCodeableConcept() any {
	return p.observation["valueCodeableConcept"]
}

func (p *ObservationParser) GetBasedOn() any {
	return p.observation["basedOn"]
}

func (p *ObservationParser) GetSpecimen() any {
	return p.observation["specimen"]
}

func (p *ObservationParser) GetSpecimenReference() string {
	return removePrefix(nestedStr(p.observation, "specimen", "reference"), "Specimen/")
}

func (p *ObservationParser) GetInterpretations() any {
	return p.observation["interpretation"]
}

func (p *ObservationParser) GetReferenceRanges() any {
	return p.observation["referenceRange"]
}

func (p *ObservationParser) GetDerivedFrom() any {
	return p.observation["derivedFrom"]
}

func (p *ObservationParser) GetComponents() any {
	return p.observation["component"]
}

func (p *ObservationParser) GetIdentifiers() any {
	return p.observation["identifier"]
}

func (p *ObservationParser) GetPartOf() any {
	return p.observation["partOf"]
}

func (p *ObservationParser) GetMethod() any {
	return p.observation["method"]
}

func (p *ObservationParser) GetNote() any {
	return p.observation["note"]
}
