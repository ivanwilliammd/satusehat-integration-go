package builder

type CoverageBuilder struct {
	data map[string]interface{}
}

func NewCoverageBuilder() *CoverageBuilder {
	return &CoverageBuilder{
		data: map[string]interface{}{"resourceType": "Coverage"},
	}
}

func (b *CoverageBuilder) SetID(id string) *CoverageBuilder {
	b.data["id"] = id
	return b
}

func (b *CoverageBuilder) AddIdentifier(system, value string) *CoverageBuilder {
	ids, _ := b.data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.data["identifier"] = ids
	return b
}

func (b *CoverageBuilder) SetStatus(status string) *CoverageBuilder {
	b.data["status"] = status
	return b
}

func (b *CoverageBuilder) SetType(system, code, display string) *CoverageBuilder {
	b.data["type"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code, "display": display}},
	}
	return b
}

func (b *CoverageBuilder) SetPolicyHolder(ref string) *CoverageBuilder {
	b.data["policyHolder"] = map[string]interface{}{"reference": ref}
	return b
}

func (b *CoverageBuilder) SetSubscriber(ref string) *CoverageBuilder {
	b.data["subscriber"] = map[string]interface{}{"reference": ref}
	return b
}

func (b *CoverageBuilder) SetSubscriberID(value string) *CoverageBuilder {
	b.data["subscriberId"] = value
	return b
}

func (b *CoverageBuilder) SetBeneficiary(ref string) *CoverageBuilder {
	b.data["beneficiary"] = map[string]interface{}{"reference": ref}
	return b
}

func (b *CoverageBuilder) SetDependent(value string) *CoverageBuilder {
	b.data["dependent"] = value
	return b
}

func (b *CoverageBuilder) SetRelationship(system, code string) *CoverageBuilder {
	b.data["relationship"] = map[string]interface{}{
		"coding": []map[string]string{{"system": system, "code": code}},
	}
	return b
}

func (b *CoverageBuilder) AddPayor(ref string, display ...string) *CoverageBuilder {
	payors, _ := b.data["payor"].([]map[string]interface{})
	payor := map[string]interface{}{"reference": ref}
	if len(display) > 0 && display[0] != "" {
		payor["display"] = display[0]
	}
	payors = append(payors, payor)
	b.data["payor"] = payors
	return b
}

func (b *CoverageBuilder) SetClass(typeSystem, typeCode, value, name string) *CoverageBuilder {
	classes, _ := b.data["class"].([]map[string]interface{})
	class := map[string]interface{}{
		"type":  map[string]interface{}{"coding": []map[string]string{{"system": typeSystem, "code": typeCode}}},
		"value": value,
	}
	if name != "" {
		class["name"] = name
	}
	classes = append(classes, class)
	b.data["class"] = classes
	return b
}

func (b *CoverageBuilder) SetOrder(order int) *CoverageBuilder {
	b.data["order"] = order
	return b
}

func (b *CoverageBuilder) SetNetwork(network string) *CoverageBuilder {
	b.data["network"] = network
	return b
}

func (b *CoverageBuilder) SetCostToBeneficiary(value float64, currency string) *CoverageBuilder {
	costs, _ := b.data["costToBeneficiary"].([]map[string]interface{})
	costs = append(costs, map[string]interface{}{"value": value, "currency": currency})
	b.data["costToBeneficiary"] = costs
	return b
}

func (b *CoverageBuilder) SetPeriod(start, end string) *CoverageBuilder {
	period := map[string]string{"start": start}
	if end != "" {
		period["end"] = end
	}
	b.data["period"] = period
	return b
}

func (b *CoverageBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}
