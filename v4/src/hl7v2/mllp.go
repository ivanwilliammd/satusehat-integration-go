package hl7v2

import (
	"errors"
	"strings"
)

// Minimal Lower Layer Protocol framing: VT ... FS CR.
const (
	VT = "\x0b"
	FS = "\x1c"
	CR = "\x0d"
	// MaxFrame guards against a peer that never sends FS. Four megabytes is far above any real
	// message and far below anything that would exhaust memory.
	MaxFrame = 4 << 20
)

var errFrameTooLarge = errors.New("MLLP frame exceeds the 4 MB limit")

func Frame(message string) string { return VT + message + FS + CR }

func Unframe(framed string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(framed, VT, ""), FS+CR, ""))
}

// Shift extracts one frame: (message, remaining). An incomplete frame returns ("", buffer).
func Shift(buffer string) (string, string) {
	start := strings.Index(buffer, VT)
	if start < 0 {
		return "", buffer
	}
	end := strings.Index(buffer[start:], FS)
	if end < 0 {
		return "", buffer
	}
	end += start
	// The frame is complete at FS; the trailing CR may not have arrived yet, and slicing past the
	// end of the buffer would panic instead of waiting for the rest.
	rest := ""
	if end+2 <= len(buffer) {
		rest = buffer[end+2:]
	}
	return buffer[start+1 : end], rest
}

// Split decodes every complete frame in a stream and discards the incomplete tail.
func Split(stream string) []string {
	var out []string
	buf := stream
	for {
		msg, rest := Shift(buf)
		if msg == "" {
			return out
		}
		out = append(out, msg)
		buf = rest
	}
}

// Reader keeps the bytes that arrived after the last complete frame.
//
// This is not a convenience: TCP does not respect message boundaries, and one Write of three ACKs is
// routinely delivered as a single read. A stateless reader would return the first ACK and drop the
// rest, leaving the client waiting for a reply it had already received.
type Reader struct {
	buf string
	max int
}

func NewReader() *Reader { return &Reader{max: MaxFrame} }

func (r *Reader) Push(chunk string) error {
	r.buf += chunk
	if len(r.buf) > r.max {
		return errFrameTooLarge
	}
	return nil
}

func (r *Reader) Take() (string, bool) {
	msg, rest := Shift(r.buf)
	if msg == "" {
		return "", false
	}
	r.buf = rest
	return msg, true
}
