package hl7v2

import (
	"regexp"
	"strings"
)

// HL7 v2 escape sequences (chapter 2).
//
// A naive strings.Split(line, "|") corrupts real messages: a pipe inside a name arrives as \F\, a
// caret inside an address as \S\. Decoding happens on read and encoding on write, so what is stored
// is exactly what went over the wire.
var decodePairs = [][2]string{
	{`\F\`, "|"},
	{`\S\`, "^"},
	{`\R\`, "~"},
	{`\E\`, `\`},
	{`\T\`, "&"},
}

var hexEscape = regexp.MustCompile(`\\X([0-9A-Fa-f]+)\\`)

// Decode resolves escape sequences. \Z..\ (locally defined) and \.br\ (formatting) are left
// untouched, because inventing a meaning for them would change clinical text.
func Decode(v string) string {
	if !strings.Contains(v, `\`) {
		return v
	}
	v = hexEscape.ReplaceAllStringFunc(v, func(m string) string {
		hex := hexEscape.FindStringSubmatch(m)[1]
		b := make([]byte, 0, len(hex)/2)
		for i := 0; i+1 < len(hex); i += 2 {
			b = append(b, byte(hexVal(hex[i])<<4|hexVal(hex[i+1])))
		}
		return string(b)
	})
	for _, p := range decodePairs {
		v = strings.ReplaceAll(v, p[0], p[1])
	}
	return v
}

// Encode is the inverse of Decode.
func Encode(v string) string {
	replacer := strings.NewReplacer(`\`, `\E\`, "|", `\F\`, "^", `\S\`, "~", `\R\`, "&", `\T\`)
	return replacer.Replace(v)
}

// CharsetOf reports MSH-18. It is never used to re-encode bytes: the declared charset in the field
// is often wrong in the field, and converting on a wrong declaration corrupts patient names.
func CharsetOf(msh18 string) string {
	switch strings.ToUpper(strings.TrimSpace(msh18)) {
	case "", "ASCII":
		return "ASCII"
	case "8859/1", "ISO IR6", "ISO IR100":
		return "ISO-8859-1"
	case "UNICODE UTF-8", "UNICODE UTF8", "UTF-8":
		return "UTF-8"
	default:
		return strings.ToUpper(strings.TrimSpace(msh18))
	}
}

func hexVal(c byte) byte {
	switch {
	case c >= '0' && c <= '9':
		return c - '0'
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10
	default:
		return c - 'A' + 10
	}
}
