package datatype

type DosageDoseAndRate struct {
	Type *CodeableConcept `json:"type,omitempty"`
	Dose interface{}      `json:"dose,omitempty"`
	Rate interface{}      `json:"rate,omitempty"`
}

func (d *DosageDoseAndRate) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if d.Type != nil { m["type"] = d.Type.ToArray() }
	if d.Dose != nil { m["dose"] = d.doseToArray() }
	if d.Rate != nil { m["rate"] = d.rateToArray() }
	return m
}

func (d *DosageDoseAndRate) doseToArray() interface{} {
	if d.Dose == nil {
		return nil
	}
	switch v := d.Dose.(type) {
	case *Range:
		return v.ToArray()
	case *SimpleQuantity:
		return v.ToArray()
	default:
		return nil
	}
}

func (d *DosageDoseAndRate) rateToArray() interface{} {
	if d.Rate == nil {
		return nil
	}
	switch v := d.Rate.(type) {
	case *Range:
		return v.ToArray()
	case *Ratio:
		return v.ToArray()
	case *SimpleQuantity:
		return v.ToArray()
	default:
		return nil
	}
}

func (d *DosageDoseAndRate) SetType(t *CodeableConcept) *DosageDoseAndRate {
	d.Type = t
	return d
}

func (d *DosageDoseAndRate) SetDose(dose interface{}) *DosageDoseAndRate {
	d.Dose = dose
	return d
}

func (d *DosageDoseAndRate) SetRate(rate interface{}) *DosageDoseAndRate {
	d.Rate = rate
	return d
}
