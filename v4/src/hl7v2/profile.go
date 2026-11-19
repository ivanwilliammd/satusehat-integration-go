package hl7v2

import (
	"fmt"
	"sort"
	"strings"
)

// Severity separates a message that must not be processed from one that is merely unusual.
const (
	SeverityError   = "error"
	SeverityWarning = "warning"
)

// Issue is one finding from a profile.
type Issue struct {
	Severity string
	Path     string
	Message  string
}

func (i Issue) String() string { return fmt.Sprintf("%s: %s (%s)", i.Path, i.Message, i.Severity) }

// Profile is a declarative rule set: which segments must exist, which fields must be filled, and
// which values are allowed. It is data, not a type per message, because every LIS and every hospital
// disagrees about the details.
type Profile struct {
	Name             string
	MessageType      string
	Trigger          string
	RequiredSegments []string
	RequiredFields   map[string]string
	AllowedValues    map[string][]string
}

// Validate returns every finding. Errors block; warnings do not.
func (p Profile) Validate(m *Message) []Issue {
	var issues []Issue
	if p.MessageType != "" && m.MessageType() != p.MessageType {
		issues = append(issues, Issue{SeverityError, "MSH.9.1",
			fmt.Sprintf("message type %s is not %s", m.MessageType(), p.MessageType)})
	}
	if p.Trigger != "" && m.TriggerEvent() != p.Trigger {
		issues = append(issues, Issue{SeverityError, "MSH.9.2",
			fmt.Sprintf("trigger %s is not %s", m.TriggerEvent(), p.Trigger)})
	}
	for _, name := range p.RequiredSegments {
		if m.Segment(name, 1) == nil {
			issues = append(issues, Issue{SeverityError, name, "required segment missing"})
		}
	}
	for _, path := range sortedKeys(p.RequiredFields) {
		if strings.TrimSpace(m.Get(path)) == "" {
			issues = append(issues, Issue{SeverityError, path, "required (" + p.RequiredFields[path] + ")"})
		}
	}
	for _, path := range sortedKeys(p.AllowedValues) {
		value := m.Get(path)
		if value == "" {
			continue
		}
		allowed := p.AllowedValues[path]
		if !contains(allowed, value) {
			issues = append(issues, Issue{SeverityWarning, path,
				fmt.Sprintf("%q is outside the usual set (%s)", value, strings.Join(allowed, ", "))})
		}
	}
	return issues
}

// Passes reports whether the message is free of error-level findings.
func (p Profile) Passes(m *Message) bool {
	for _, i := range p.Validate(m) {
		if i.Severity == SeverityError {
			return false
		}
	}
	return true
}

func sortedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func contains(list []string, v string) bool {
	for _, s := range list {
		if s == v {
			return true
		}
	}
	return false
}
