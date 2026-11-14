package parser

// NutritionOrderParser extracts fields from a FHIR NutritionOrder resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/NutritionOrder.json, copied unchanged to
// src/parser/testdata/NutritionOrder.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type NutritionOrderParser struct {
	payload map[string]any
}

// NewNutritionOrderParser creates a new NutritionOrderParser instance.
func NewNutritionOrderParser(payload map[string]any) *NutritionOrderParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &NutritionOrderParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *NutritionOrderParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *NutritionOrderParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetIdentifier returns the identifier element unchanged.
func (p *NutritionOrderParser) GetIdentifier() any {
	return p.payload["identifier"]
}

// GetStatus returns the status element unchanged.
func (p *NutritionOrderParser) GetStatus() any {
	return p.payload["status"]
}

// GetIntent returns the intent element unchanged.
func (p *NutritionOrderParser) GetIntent() any {
	return p.payload["intent"]
}

// GetPatient returns the patient element unchanged.
func (p *NutritionOrderParser) GetPatient() any {
	return p.payload["patient"]
}

// GetEncounter returns the encounter element unchanged.
func (p *NutritionOrderParser) GetEncounter() any {
	return p.payload["encounter"]
}

// GetDateTime returns the dateTime element unchanged.
func (p *NutritionOrderParser) GetDateTime() any {
	return p.payload["dateTime"]
}

// GetOralDiet returns the oralDiet element unchanged.
func (p *NutritionOrderParser) GetOralDiet() any {
	return p.payload["oralDiet"]
}
