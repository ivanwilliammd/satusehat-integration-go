package datatype

type Dosage struct {
	Sequence   *int                `json:"sequence,omitempty"`
	Text       *string             `json:"text,omitempty"`
	Timing     *Timing             `json:"timing,omitempty"`
	AsNeeded   interface{}         `json:"asNeeded,omitempty"`
	Site       *CodeableConcept    `json:"site,omitempty"`
	Route      *CodeableConcept    `json:"route,omitempty"`
	Method     *CodeableConcept    `json:"method,omitempty"`
	DoseAndRate []DosageDoseAndRate `json:"doseAndRate,omitempty"`
}

func (d *Dosage) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if d.Sequence != nil { m["sequence"] = *d.Sequence }
	if d.Text != nil { m["text"] = *d.Text }
	if d.Timing != nil { m["timing"] = d.Timing.ToArray() }
	if d.AsNeeded != nil { m["asNeeded"] = d.asNeededToArray() }
	if d.Site != nil { m["site"] = d.Site.ToArray() }
	if d.Route != nil { m["route"] = d.Route.ToArray() }
	if d.Method != nil { m["method"] = d.Method.ToArray() }
	if len(d.DoseAndRate) > 0 {
		dr := make([]map[string]interface{}, 0, len(d.DoseAndRate))
		for _, v := range d.DoseAndRate {
			dr = append(dr, v.ToArray())
		}
		m["doseAndRate"] = dr
	}
	return m
}

func (d *Dosage) asNeededToArray() interface{} {
	if d.AsNeeded == nil {
		return nil
	}
	switch v := d.AsNeeded.(type) {
	case bool:
		return v
	case *bool:
		if v != nil {
			return *v
		}
		return nil
	default:
		return nil
	}
}

func (d *Dosage) SetSequence(seq int) *Dosage {
	d.Sequence = &seq
	return d
}

func (d *Dosage) SetText(text string) *Dosage {
	d.Text = &text
	return d
}

func (d *Dosage) SetTiming(timing *Timing) *Dosage {
	d.Timing = timing
	return d
}

func (d *Dosage) SetAsNeeded(v interface{}) *Dosage {
	d.AsNeeded = v
	return d
}

func (d *Dosage) SetSite(site *CodeableConcept) *Dosage {
	d.Site = site
	return d
}

func (d *Dosage) SetRoute(route *CodeableConcept) *Dosage {
	d.Route = route
	return d
}

func (d *Dosage) SetMethod(method *CodeableConcept) *Dosage {
	d.Method = method
	return d
}

func (d *Dosage) AddDoseAndRate(dr DosageDoseAndRate) *Dosage {
	d.DoseAndRate = append(d.DoseAndRate, dr)
	return d
}
