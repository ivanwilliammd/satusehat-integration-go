package parser

// GoalParser extracts fields from a FHIR Goal resource map.
//
// Shape source: satusehat-conformance/examples/satusehat/Goal.json, copied unchanged to
// src/parser/testdata/Goal.json. One accessor per element the official example contains;
// GetRaw returns the payload so a caller is never blocked by a missing accessor.
type GoalParser struct {
	payload map[string]any
}

// NewGoalParser creates a new GoalParser instance.
func NewGoalParser(payload map[string]any) *GoalParser {
	if payload == nil {
		payload = map[string]any{}
	}
	return &GoalParser{payload: payload}
}

// GetRaw returns the whole payload.
func (p *GoalParser) GetRaw() map[string]any {
	return p.payload
}

// GetResourceType returns the resourceType element unchanged.
func (p *GoalParser) GetResourceType() any {
	return p.payload["resourceType"]
}

// GetId returns the id element unchanged.
func (p *GoalParser) GetId() any {
	return p.payload["id"]
}

// GetLifecycleStatus returns the lifecycleStatus element unchanged.
func (p *GoalParser) GetLifecycleStatus() any {
	return p.payload["lifecycleStatus"]
}

// GetAchievementStatus returns the achievementStatus element unchanged.
func (p *GoalParser) GetAchievementStatus() any {
	return p.payload["achievementStatus"]
}

// GetCategory returns the category element unchanged.
func (p *GoalParser) GetCategory() any {
	return p.payload["category"]
}

// GetDescription returns the description element unchanged.
func (p *GoalParser) GetDescription() any {
	return p.payload["description"]
}

// GetSubject returns the subject element unchanged.
func (p *GoalParser) GetSubject() any {
	return p.payload["subject"]
}

// GetTarget returns the target element unchanged.
func (p *GoalParser) GetTarget() any {
	return p.payload["target"]
}

// GetStatusDate returns the statusDate element unchanged.
func (p *GoalParser) GetStatusDate() any {
	return p.payload["statusDate"]
}

// GetExpressedBy returns the expressedBy element unchanged.
func (p *GoalParser) GetExpressedBy() any {
	return p.payload["expressedBy"]
}

// GetAddresses returns the addresses element unchanged.
func (p *GoalParser) GetAddresses() any {
	return p.payload["addresses"]
}

// GetOutcomeCode returns the outcomeCode element unchanged.
func (p *GoalParser) GetOutcomeCode() any {
	return p.payload["outcomeCode"]
}
