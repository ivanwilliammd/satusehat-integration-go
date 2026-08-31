package builder

type BillingStatusBuilder struct {
	data map[string]interface{}
}

func NewBillingStatusBuilder() *BillingStatusBuilder {
	return &BillingStatusBuilder{
		data: map[string]interface{}{"resourceType": "BillingStatus"},
	}
}

func (b *BillingStatusBuilder) SetID(id string) *BillingStatusBuilder {
	b.data["id"] = id
	return b
}

func (b *BillingStatusBuilder) AddIdentifier(system, value string) *BillingStatusBuilder {
	ids, _ := b.data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.data["identifier"] = ids
	return b
}

func (b *BillingStatusBuilder) SetStatus(status string) *BillingStatusBuilder {
	b.data["status"] = status
	return b
}

func (b *BillingStatusBuilder) SetInsurer(ref string, display ...string) *BillingStatusBuilder {
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

func (b *BillingStatusBuilder) SetRecipient(ref string, display ...string) *BillingStatusBuilder {
	if !hasPrefix(ref) && !containsSlash(ref) {
		ref = "Organization/" + ref
	}
	res := map[string]interface{}{"reference": ref}
	if len(display) > 0 && display[0] != "" {
		res["display"] = display[0]
	}
	b.data["recipient"] = res
	return b
}

func (b *BillingStatusBuilder) SetSubject(ref string, display ...string) *BillingStatusBuilder {
	if !hasPrefix(ref) && !containsSlash(ref) {
		ref = "Patient/" + ref
	}
	res := map[string]interface{}{"reference": ref}
	if len(display) > 0 && display[0] != "" {
		res["display"] = display[0]
	}
	b.data["subject"] = res
	return b
}

func (b *BillingStatusBuilder) SetRequest(ref string, display ...string) *BillingStatusBuilder {
	if !hasPrefix(ref) && !containsSlash(ref) {
		ref = "CoverageEligibilityRequest/" + ref
	}
	res := map[string]interface{}{"reference": ref}
	if len(display) > 0 && display[0] != "" {
		res["display"] = display[0]
	}
	b.data["request"] = res
	return b
}

func (b *BillingStatusBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}
