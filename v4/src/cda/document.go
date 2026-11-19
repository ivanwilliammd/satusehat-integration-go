// Package cda reads and writes the part of HL7 CDA R2 that a record exchange actually uses: the
// header and the sections with narrative text. Everything else is preserved in the parse tree and
// written back, so a document with structures this package does not model survives a round trip.
//
// The HL7 CDA R2 schemas are licensed; this package implements the structural rules from the
// specification and does not redistribute them.
package cda

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

// Namespace of every CDA R2 document.
const Namespace = "urn:hl7-org:v3"

var (
	errNoRoot = errors.New("dokumen kosong")
	errNotCDA = errors.New("akar dokumen bukan ClinicalDocument")
)

// Document is a parsed ClinicalDocument.
type Document struct {
	root *node
}

// Parse reads a ClinicalDocument. Anything whose root element is not ClinicalDocument is refused.
func Parse(data string) (*Document, error) {
	tree, err := parseTree(data)
	if err != nil {
		return nil, fmt.Errorf("XML CDA tidak sah: %w", err)
	}
	doc := tree.child("ClinicalDocument")
	if doc == nil {
		return nil, errNotCDA
	}
	return &Document{root: doc}, nil
}

// Title is ClinicalDocument/title.
func (d *Document) Title() string {
	if n := d.root.child("title"); n != nil {
		return n.trimmed()
	}
	return ""
}

// DocumentID is the II of the document as "root|extension".
func (d *Document) DocumentID() string {
	n := d.root.child("id")
	if n == nil {
		return ""
	}
	if ext := n.attr("extension"); ext != "" {
		return n.attr("root") + "|" + ext
	}
	return n.attr("root")
}

// TypeCode is the document type (LOINC), e.g. 34133-9.
func (d *Document) TypeCode() string {
	if n := d.root.child("code"); n != nil {
		return n.attr("code")
	}
	return ""
}

// EffectiveTime is the clinical event time as YYYYMMDDHHMMSS.
func (d *Document) EffectiveTime() string {
	if n := d.root.child("effectiveTime"); n != nil {
		return n.attr("value")
	}
	return ""
}

// Custodian is the name of the organisation responsible for the document.
func (d *Document) Custodian() string {
	n := d.root.child("custodian")
	if n == nil {
		return ""
	}
	if name := n.descendants("name"); len(name) > 0 {
		return name[0].trimmed()
	}
	return ""
}

// Authors returns author names in document order; softwareName is used when a device signed it.
func (d *Document) Authors() []string {
	var out []string
	for _, author := range d.root.childAll("author") {
		given := textOf(author.descendants("given"))
		family := textOf(author.descendants("family"))
		name := strings.TrimSpace(given + " " + family)
		if name == "" {
			name = textOf(author.descendants("softwareName"))
		}
		if name != "" {
			out = append(out, name)
		}
	}
	return out
}

func textOf(nodes []*node) string {
	if len(nodes) == 0 {
		return ""
	}
	return nodes[0].trimmed()
}

// Patient is the recordTarget/patientRole block.
func (d *Document) Patient() PatientRole {
	target := d.root.child("recordTarget")
	if target == nil {
		return PatientRole{}
	}
	role := target.child("patientRole")
	if role == nil {
		return PatientRole{}
	}
	var out PatientRole
	for _, id := range role.childAll("id") {
		out.Identifiers = append(out.Identifiers, Identifier{Root: id.attr("root"), Extension: id.attr("extension")})
	}
	for _, name := range role.descendants("name") {
		value := strings.TrimSpace(textOf(name.childAll("given")) + " " + textOf(name.childAll("family")))
		if value != "" {
			out.Names = append(out.Names, value)
		}
	}
	if g := role.descendants("administrativeGenderCode"); len(g) > 0 {
		out.AdministrativeGender = g[0].attr("code")
	}
	if b := role.descendants("birthTime"); len(b) > 0 {
		out.BirthTime = b[0].attr("value")
	}
	return out
}

// Sections returns the top level sections of the structured body.
func (d *Document) Sections() []*Section {
	body := d.root.child("component")
	if body == nil {
		return nil
	}
	sb := body.child("structuredBody")
	if sb == nil {
		return nil
	}
	out := make([]*Section, 0, len(sb.children))
	for _, s := range sectionsOf(sb) {
		out = append(out, readSection(s))
	}
	return out
}

// sectionsOf unwraps the <component><section> level used by structuredBody and by nested sections.
func sectionsOf(parent *node) []*node {
	var out []*node
	for _, comp := range parent.childAll("component") {
		out = append(out, comp.childAll("section")...)
	}
	return out
}

func readSection(n *node) *Section {
	s := &Section{Title: sTitle(n), Codes: map[string]string{}}
	for _, code := range n.childAll("code") {
		system, value := code.attr("codeSystem"), code.attr("code")
		if system != "" && value != "" {
			s.Codes[system] = value
		}
	}
	if text := n.child("text"); text != nil {
		var b strings.Builder
		text.write(&b)
		s.Text = b.String()
	}
	for _, child := range sectionsOf(n) {
		s.Children = append(s.Children, readSection(child))
	}
	return s
}

func sTitle(n *node) string {
	if t := n.child("title"); t != nil {
		return t.trimmed()
	}
	return ""
}

// Section returns the first section coded with this LOINC code, depth first.
func (d *Document) Section(loinc string) *Section {
	for _, s := range d.Sections() {
		if hit := s.Find(loinc); hit != nil {
			return hit
		}
	}
	return nil
}

// SectionText returns a section's narrative flattened to text, with paragraph breaks kept.
func (d *Document) SectionText(loinc string) string {
	s := d.Section(loinc)
	if s == nil {
		return ""
	}
	return PlainText(s.Text)
}

// PlainText flattens CDA narrative markup: paragraph-level elements become newlines, the rest of
// the text is kept as written.
func PlainText(xmlText string) string {
	if strings.TrimSpace(xmlText) == "" {
		return ""
	}
	tree, err := parseTree(xmlText)
	if err != nil {
		return strings.TrimSpace(xmlText)
	}
	var b strings.Builder
	var walk func(n *node)
	walk = func(n *node) {
		switch n.name {
		case "paragraph", "item", "tr", "caption":
			b.WriteByte('\n')
		case "#comment":
			return
		}
		if t := strings.TrimSpace(n.text); t != "" {
			b.WriteString(t)
		}
		for _, c := range n.children {
			walk(c)
		}
	}
	walk(tree)
	out := b.String()
	for strings.Contains(out, "\n\n\n") {
		out = strings.ReplaceAll(out, "\n\n\n", "\n\n")
	}
	return strings.TrimSpace(out)
}

// Validate checks the header rules that matter for exchange. An empty result means the document is
// exchangeable, not that it is clinically correct.
func (d *Document) Validate() []string {
	var problems []string
	if d.DocumentID() == "" {
		problems = append(problems, "ClinicalDocument/id wajib ada.")
	}
	if d.Title() == "" {
		problems = append(problems, "ClinicalDocument/title wajib ada.")
	}
	if d.TypeCode() == "" {
		problems = append(problems, "ClinicalDocument/code wajib ada.")
	}
	if d.EffectiveTime() == "" {
		problems = append(problems, "ClinicalDocument/effectiveTime wajib ada.")
	}
	if d.root.child("recordTarget") == nil {
		problems = append(problems, "recordTarget wajib ada.")
	}
	if len(d.Patient().Identifiers) == 0 {
		problems = append(problems, "recordTarget/patientRole/id wajib ada.")
	}
	if d.Custodian() == "" {
		problems = append(problems, "custodian wajib ada.")
	}
	if len(d.Authors()) == 0 {
		problems = append(problems, "author wajib ada.")
	}
	if len(d.Sections()) == 0 {
		problems = append(problems, "structuredBody wajib berisi minimal satu section.")
	}
	return problems
}

// XML renders the document. Structure parsed earlier but not modelled by this package is written
// back unchanged.
func (d *Document) XML() string {
	var b strings.Builder
	b.WriteString(`<?xml version="1.0" encoding="UTF-8"?>`)
	d.root.write(&b)
	return b.String()
}

// BuildInput is the data accepted by Build.
type BuildInput struct {
	ID            string
	Title         string
	Code          string
	CodeDisplay   string
	EffectiveTime string
	Custodian     string
	Author        string
	PatientID     string
	PatientIDRoot string
	PatientName   string
	Gender        string
	BirthTime     string
	Sections      []BuildSection
}

// BuildSection is one section passed to Build; Text is CDA narrative markup, not escaped text.
type BuildSection struct {
	LOINC string
	Title string
	Text  string
}

// Build assembles a minimal but structurally valid ClinicalDocument.
func Build(in BuildInput) (*Document, error) {
	if in.Code == "" {
		in.Code = "34133-9"
	}
	if in.CodeDisplay == "" {
		in.CodeDisplay = "Summarization of episode note"
	}
	if in.Title == "" {
		in.Title = "Catatan Medis"
	}
	if in.PatientIDRoot == "" {
		in.PatientIDRoot = "2.16.840.1.113883.3.3157.1003.1"
	}
	now := in.EffectiveTime
	if now == "" {
		now = time.Now().Format("20060102150405")
	}
	id := in.ID
	if id == "" {
		id = "DOC" + now
	}

	var sections strings.Builder
	for _, s := range in.Sections {
		sections.WriteString("<component><section>")
		if s.LOINC != "" {
			sections.WriteString(`<code code="` + escapeAttr(s.LOINC) + `" codeSystem="` + LOINC + `"/>`)
		}
		sections.WriteString("<title>" + escapeText(s.Title) + "</title>")
		sections.WriteString("<text>" + s.Text + "</text>")
		sections.WriteString("</section></component>")
	}

	var b strings.Builder
	b.WriteString(`<ClinicalDocument xmlns="` + Namespace + `">`)
	b.WriteString(`<realmCode code="ID"/>`)
	b.WriteString(`<typeId root="2.16.840.1.113883.1.3" extension="POCD_MT000040"/>`)
	b.WriteString(`<templateId root="2.16.840.1.113883.10.20.22.1.1"/>`)
	b.WriteString(`<id root="` + escapeAttr(id) + `"/>`)
	b.WriteString(`<code code="` + escapeAttr(in.Code) + `" codeSystem="` + LOINC + `" displayName="` + escapeAttr(in.CodeDisplay) + `"/>`)
	b.WriteString("<title>" + escapeText(in.Title) + "</title>")
	b.WriteString(`<effectiveTime value="` + escapeAttr(now) + `"/>`)
	b.WriteString(`<confidentialityCode code="N" codeSystem="2.16.840.1.113883.5.25"/>`)
	b.WriteString(`<recordTarget><patientRole>`)
	if in.PatientID != "" {
		b.WriteString(`<id root="` + escapeAttr(in.PatientIDRoot) + `" extension="` + escapeAttr(in.PatientID) + `"/>`)
	}
	b.WriteString(`<patient>`)
	if in.PatientName != "" {
		b.WriteString(`<name><family>` + escapeText(in.PatientName) + `</family></name>`)
	}
	if in.Gender != "" {
		b.WriteString(`<administrativeGenderCode code="` + escapeAttr(in.Gender) + `" codeSystem="2.16.840.1.113883.5.1"/>`)
	}
	if in.BirthTime != "" {
		b.WriteString(`<birthTime value="` + escapeAttr(in.BirthTime) + `"/>`)
	}
	b.WriteString(`</patient></patientRole></recordTarget>`)
	b.WriteString(`<author><time value="` + escapeAttr(now) + `"/><assignedAuthor><id root="` + escapeAttr(id) + `.AUTH"/>`)
	b.WriteString(`<assignedPerson><name><given>` + escapeText(in.Author) + `</given></name></assignedPerson>`)
	b.WriteString(`</assignedAuthor></author>`)
	b.WriteString(`<custodian><assignedCustodian><representedCustodianOrganization>`)
	b.WriteString(`<name>` + escapeText(in.Custodian) + `</name>`)
	b.WriteString(`</representedCustodianOrganization></assignedCustodian></custodian>`)
	b.WriteString(`<component><structuredBody>` + sections.String() + `</structuredBody></component>`)
	b.WriteString(`</ClinicalDocument>`)

	return Parse(b.String())
}
