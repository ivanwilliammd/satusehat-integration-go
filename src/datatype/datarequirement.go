package datatype

type DataRequirement struct {
	Type        *string    `json:"type,omitempty"`
	Profile     []string   `json:"profile,omitempty"`
	Subject     interface{} `json:"subject,omitempty"`
	CodeFilter  []map[string]interface{} `json:"codeFilter,omitempty"`
	DateFilter  []map[string]interface{} `json:"dateFilter,omitempty"`
	Sort        []map[string]interface{} `json:"sort,omitempty"`
}

func (d *DataRequirement) ToArray() map[string]interface{} {
	m := make(map[string]interface{})
	if d.Type != nil { m["type"] = *d.Type }
	if len(d.Profile) > 0 { m["profile"] = d.Profile }
	if d.Subject != nil { m["subject"] = d.subjectToArray() }
	if len(d.CodeFilter) > 0 { m["codeFilter"] = d.CodeFilter }
	if len(d.DateFilter) > 0 { m["dateFilter"] = d.DateFilter }
	if len(d.Sort) > 0 { m["sort"] = d.Sort }
	return m
}

func (d *DataRequirement) subjectToArray() interface{} {
	if d.Subject == nil {
		return nil
	}
	switch v := d.Subject.(type) {
	case *CodeableConcept:
		return v.ToArray()
	case *Reference:
		return v.ToArray()
	default:
		return nil
	}
}

func (d *DataRequirement) AddProfile(profile string) *DataRequirement {
	d.Profile = append(d.Profile, profile)
	return d
}

func (d *DataRequirement) SetSubject(subject interface{}) *DataRequirement {
	d.Subject = subject
	return d
}

func (d *DataRequirement) AddCodeFilter(filter map[string]interface{}) *DataRequirement {
	d.CodeFilter = append(d.CodeFilter, filter)
	return d
}

func (d *DataRequirement) AddDateFilter(filter map[string]interface{}) *DataRequirement {
	d.DateFilter = append(d.DateFilter, filter)
	return d
}

func (d *DataRequirement) AddSort(sort map[string]interface{}) *DataRequirement {
	d.Sort = append(d.Sort, sort)
	return d
}
