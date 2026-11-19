package hl7v2

import (
	"fmt"
	"time"
)

// Verdict codes for MSA-1.
const (
	Accept = "AA"
	Error  = "AE"
	Reject = "AR"
)

// BuildAck answers an inbound message.
//
// MSA-1 is the verdict and MSA-2 echoes MSH-10, which is how the sender correlates the ACK with what
// it sent. An ACK without that echo is indistinguishable from a stale one.
func BuildAck(inbound *Message, code, text string) *Message {
	msh := inbound.Segment("MSH", 1)
	field := func(n int) string {
		if msh == nil {
			return ""
		}
		return msh.Field(n)
	}
	version := ""
	if msh != nil {
		version = msh.Get("12.1")
	}
	if version == "" {
		version = "2.5.1"
	}
	controlID := fmt.Sprintf("ACK%d", time.Now().UnixNano())
	parts := []string{
		"MSH",
		"^~\\&",
		field(5), field(6), field(3), field(4),
		time.Now().Format("20060102150405"),
		"", "ACK", controlID, "P", version,
	}
	msa := Segment{Name: "MSA", Sep: "|", Fields: map[int]string{1: code, 2: inbound.ControlID(), 3: text}}
	return NewMessage([]Segment{NewSegmentFromParts(parts, "|"), msa}, "|")
}

func BuildNak(inbound *Message, reason string) *Message {
	if len(reason) > 200 {
		reason = reason[:200]
	}
	return BuildAck(inbound, Error, reason)
}
