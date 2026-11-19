package hl7v2

import (
	"errors"
	"strings"
)

// ErrEmpty is returned for a message with no content at all.
var ErrEmpty = errors.New("empty HL7 message")

// ErrNoMSH is returned when the first segment is not MSH: the bytes may be a batch (FHS/BHS) or a
// fragment, and guessing would put unvalidated data into a database.
var ErrNoMSH = errors.New("HL7 message must start with an MSH segment")

// Message is an ordered list of segments addressed by name.
type Message struct {
	Segments []Segment
	Sep      string
}

// Parse reads ER7 text. Line endings may be CR, LF or CRLF, and the encoding characters are taken
// from MSH-2 rather than assumed.
func Parse(er7 string) (*Message, error) {
	cleaned := strings.NewReplacer("\x0b", "", "\x1c", "").Replace(er7)
	cleaned = strings.TrimSpace(cleaned)
	if cleaned == "" {
		return nil, ErrEmpty
	}
	cleaned = strings.ReplaceAll(cleaned, "\r\n", "\n")
	cleaned = strings.ReplaceAll(cleaned, "\r", "\n")

	msg := &Message{Sep: "|"}
	for _, line := range strings.Split(cleaned, "\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		sep := "|"
		if len(line) > 3 {
			sep = string(line[3])
		}
		parts := strings.Split(line, sep)
		if parts[0] == "" {
			continue
		}
		if parts[0] == "MSH" {
			msg.Sep = sep
		}
		msg.Segments = append(msg.Segments, NewSegmentFromParts(parts, sep))
	}
	if len(msg.Segments) == 0 || msg.Segments[0].Name != "MSH" {
		return nil, ErrNoMSH
	}
	return msg, nil
}

// NewMessage builds a message from segments.
func NewMessage(segments []Segment, sep string) *Message {
	if sep == "" {
		sep = "|"
	}
	return &Message{Segments: segments, Sep: sep}
}

// Segment returns the nth occurrence of a segment, or nil.
func (m *Message) Segment(name string, occurrence int) *Segment {
	if occurrence < 1 {
		occurrence = 1
	}
	seen := 0
	for i := range m.Segments {
		if m.Segments[i].Name == name {
			seen++
			if seen == occurrence {
				return &m.Segments[i]
			}
		}
	}
	return nil
}

// AllOf returns every occurrence of a segment.
func (m *Message) AllOf(name string) []Segment {
	var out []Segment
	for _, s := range m.Segments {
		if s.Name == name {
			out = append(out, s)
		}
	}
	return out
}

func (m *Message) Get(path string) string {
	name, rest, _ := strings.Cut(path, ".")
	seg := m.Segment(name, 1)
	if seg == nil || rest == "" {
		return ""
	}
	return seg.Get(rest)
}

func (m *Message) Set(path, value string) *Message {
	name, rest, _ := strings.Cut(path, ".")
	if seg := m.Segment(name, 1); seg != nil && rest != "" {
		seg.Set(rest, value)
	}
	return m
}

func (m *Message) MessageType() string   { return m.Get("MSH.9.1") }
func (m *Message) TriggerEvent() string  { return m.Get("MSH.9.2") }
func (m *Message) ControlID() string     { return m.Get("MSH.10") }
func (m *Message) Version() string       { return m.Get("MSH.12.1") }
func (m *Message) Charset() string {
	if seg := m.Segment("MSH", 1); seg != nil {
		return CharsetOf(seg.Field(18))
	}
	return "ASCII"
}

// String renders ER7 with CR line endings, byte-for-byte with what was parsed.
func (m *Message) String() string {
	parts := make([]string, 0, len(m.Segments))
	for _, s := range m.Segments {
		parts = append(parts, s.String())
	}
	body := strings.Join(parts, "\r")
	if body == "" {
		return ""
	}
	return body + "\r"
}
