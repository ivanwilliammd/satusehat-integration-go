package builder

type PurificationDecisionBuilder struct {
	data map[string]interface{}
}

func NewPurificationDecisionBuilder() *PurificationDecisionBuilder {
	return &PurificationDecisionBuilder{
		data: map[string]interface{}{"resourceType": "PurificationDecision"},
	}
}

func (b *PurificationDecisionBuilder) SetID(id string) *PurificationDecisionBuilder {
	b.data["id"] = id
	return b
}

func (b *PurificationDecisionBuilder) AddIdentifier(system, value string) *PurificationDecisionBuilder {
	ids, _ := b.data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.data["identifier"] = ids
	return b
}

func (b *PurificationDecisionBuilder) SetStatus(code string, displayAndSystem ...string) *PurificationDecisionBuilder {
	display := code
	if len(displayAndSystem) > 0 && displayAndSystem[0] != "" {
		display = displayAndSystem[0]
	}
	coding := map[string]string{"code": code, "display": display}
	if len(displayAndSystem) > 1 && displayAndSystem[1] != "" {
		coding["system"] = displayAndSystem[1]
	}
	b.data["status"] = map[string]interface{}{
		"coding": []map[string]string{coding},
	}
	return b
}

func (b *PurificationDecisionBuilder) SetInsurer(ref string, display ...string) *PurificationDecisionBuilder {
	if !hasPrefix(ref) && !containsSlash(ref) {
		ref = "Organization/" + ref
	}
	res := map[string]interface{}{"reference": ref}
	if len(display) > 0 && display[0] != "" {
		res["display"] = display[0]
	}
	b.data["insurer"] = res
	return b
}

func (b *PurificationDecisionBuilder) SetProvider(ref string, display ...string) *PurificationDecisionBuilder {
	if !hasPrefix(ref) && !containsSlash(ref) {
		ref = "Organization/" + ref
	}
	res := map[string]interface{}{"reference": ref}
	if len(display) > 0 && display[0] != "" {
		res["display"] = display[0]
	}
	b.data["provider"] = res
	return b
}

func (b *PurificationDecisionBuilder) SetClaimResponse(ref string, display ...string) *PurificationDecisionBuilder {
	if !hasPrefix(ref) && !containsSlash(ref) {
		ref = "ClaimResponse/" + ref
	}
	res := map[string]interface{}{"reference": ref}
	if len(display) > 0 && display[0] != "" {
		res["display"] = display[0]
	}
	b.data["claimResponse"] = res
	return b
}

func (b *PurificationDecisionBuilder) SetCreated(created string) *PurificationDecisionBuilder {
	b.data["created"] = created
	return b
}

func (b *PurificationDecisionBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}
