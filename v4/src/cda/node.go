package cda

import (
	"encoding/xml"
	"io"
	"strings"
)

// node is a minimal XML element tree. It exists because encoding/xml unmarshals into typed structs,
// which would drop every element the SDK does not know about; keeping the tree means an unknown
// extension survives a read/write round trip.
type node struct {
	name     string
	attrs    []xml.Attr
	text     string
	children []*node
}

func parseTree(data string) (*node, error) {
	dec := xml.NewDecoder(strings.NewReader(data))
	dec.Strict = true
	root := &node{name: "#document"}
	stack := []*node{root}

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		switch t := tok.(type) {
		case xml.StartElement:
			child := &node{name: t.Name.Local, attrs: append([]xml.Attr(nil), t.Attr...)}
			parent := stack[len(stack)-1]
			parent.children = append(parent.children, child)
			stack = append(stack, child)
		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			if len(stack) == 0 {
				continue
			}
			stack[len(stack)-1].text += string(t)
		}
	}
	if len(root.children) == 0 {
		return nil, errNoRoot
	}
	return root, nil
}

func (n *node) attr(name string) string {
	for _, a := range n.attrs {
		if a.Name.Local == name {
			return a.Value
		}
	}
	return ""
}

func (n *node) child(name string) *node {
	for _, c := range n.children {
		if c.name == name {
			return c
		}
	}
	return nil
}

func (n *node) childAll(name string) []*node {
	var out []*node
	for _, c := range n.children {
		if c.name == name {
			out = append(out, c)
		}
	}
	return out
}

// descendants walks the tree depth first, in document order.
func (n *node) descendants(name string) []*node {
	var out []*node
	for _, c := range n.children {
		if c.name == name {
			out = append(out, c)
		}
		out = append(out, c.descendants(name)...)
	}
	return out
}

func (n *node) trimmed() string { return strings.TrimSpace(n.text) }

// write serialises the tree. Attributes keep their original order and text keeps its original
// spacing, so the output is stable enough to diff against the input.
func (n *node) write(b *strings.Builder) {
	if n.name != "" && n.name != "#document" {
		b.WriteByte('<')
		b.WriteString(n.name)
		for _, a := range n.attrs {
			b.WriteByte(' ')
			b.WriteString(a.Name.Local)
			b.WriteString(`="`)
			b.WriteString(escapeAttr(a.Value))
			b.WriteString(`"`)
		}
		if len(n.children) == 0 && strings.TrimSpace(n.text) == "" {
			b.WriteString("/>")
			return
		}
		b.WriteByte('>')
	}
	if t := strings.TrimSpace(n.text); t != "" || len(n.children) > 0 {
		b.WriteString(escapeText(n.text))
	}
	for _, c := range n.children {
		c.write(b)
	}
	if n.name != "" && n.name != "#document" {
		b.WriteString("</")
		b.WriteString(n.name)
		b.WriteByte('>')
	}
}

func escapeText(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;")
	return r.Replace(s)
}

func escapeAttr(s string) string {
	r := strings.NewReplacer("&", "&amp;", "<", "&lt;", ">", "&gt;", `"`, "&quot;")
	return r.Replace(s)
}
