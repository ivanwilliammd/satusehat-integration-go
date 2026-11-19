package hl7v2

import "strings"

// Segment is one HL7 v2 segment: fields kept as raw text and split on demand.
type Segment struct {
	Name   string
	Fields map[int]string
	Sep    string
}

// NewSegmentFromParts builds a segment from already-split parts. MSH is special: MSH-1 is the field
// separator itself and MSH-2 holds the encoding characters, so its split starts one position later
// than every other segment.
func NewSegmentFromParts(parts []string, sep string) Segment {
	if len(parts) == 0 {
		return Segment{Fields: map[int]string{}, Sep: sep}
	}
	name := parts[0]
	offset := 0
	if name == "MSH" {
		offset = 1
	}
	fields := make(map[int]string, len(parts))
	for i, v := range parts[1:] {
		fields[i+1+offset] = v
	}
	if name == "MSH" {
		fields[1] = sep
	}
	return Segment{Name: name, Fields: fields, Sep: sep}
}

func (s Segment) Field(n int) string { return s.Fields[n] }

// Get reads "5", "5.1" (component) or "5.1.2" (subcomponent); the first repetition only.
func (s Segment) Get(path string) string {
	parts := strings.Split(path, ".")
	n := atoi(parts[0])
	raw := s.Field(n)
	if i := strings.Index(raw, "~"); i >= 0 {
		raw = raw[:i]
	}
	if len(parts) == 1 {
		return Decode(raw)
	}
	comps := strings.Split(raw, "^")
	idx := atoi(parts[1]) - 1
	value := ""
	if idx >= 0 && idx < len(comps) {
		value = comps[idx]
	}
	if len(parts) > 2 {
		subs := strings.Split(value, "&")
		idx = atoi(parts[2]) - 1
		if idx >= 0 && idx < len(subs) {
			value = subs[idx]
		} else {
			value = ""
		}
	}
	return Decode(value)
}

// FieldAll returns every repetition of a field, decoded. Identifiers and phone numbers live here.
func (s Segment) FieldAll(n int) []string {
	raw := s.Field(n)
	if raw == "" {
		return nil
	}
	reps := strings.Split(raw, "~")
	out := make([]string, len(reps))
	for i, r := range reps {
		out[i] = Decode(r)
	}
	return out
}

func (s Segment) Set(path, value string) {
	parts := strings.Split(path, ".")
	n := atoi(parts[0])
	if len(parts) == 1 {
		s.Fields[n] = Encode(value)
		return
	}
	raw := s.Field(n)
	var reps []string
	if raw == "" {
		reps = []string{""}
	} else {
		reps = strings.Split(raw, "~")
	}
	comps := strings.Split(reps[0], "^")
	idx := atoi(parts[1]) - 1
	for len(comps) <= idx {
		comps = append(comps, "")
	}
	comps[idx] = Encode(value)
	reps[0] = strings.Join(comps, "^")
	s.Fields[n] = strings.Join(reps, "~")
}

func (s Segment) String() string {
	highest := 0
	for k := range s.Fields {
		if k > highest {
			highest = k
		}
	}
	// MSH keeps its separator in field 1; rendering from field 2 avoids "MSH|||".
	start := 1
	if s.Name == "MSH" {
		start = 2
	}
	out := []string{s.Name}
	for i := start; i <= highest; i++ {
		out = append(out, s.Fields[i])
	}
	for len(out) > 2 && out[len(out)-1] == "" {
		out = out[:len(out)-1]
	}
	return strings.Join(out, s.Sep)
}

func atoi(s string) int {
	n := 0
	for _, c := range s {
		if c < '0' || c > '9' {
			return n
		}
		n = n*10 + int(c-'0')
	}
	return n
}
