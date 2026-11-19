package cda

// LOINC is the code system every CDA section is usually coded in.
const LOINC = "2.16.840.1.113883.6.1"

// Section is one narrative section: a coded heading, its text, and any nested sections.
type Section struct {
	Title    string
	Text     string // narrative XML, not plain text: CDA narrative carries markup
	Codes    map[string]string
	Children []*Section
}

// Code returns the LOINC code of this section, or the first code when it is not LOINC-coded.
func (s *Section) Code() string {
	if v, ok := s.Codes[LOINC]; ok {
		return v
	}
	for _, v := range s.Codes {
		return v
	}
	return ""
}

// Find returns the first section in the tree with this LOINC code, depth first.
func (s *Section) Find(loinc string) *Section {
	if s.Codes[LOINC] == loinc {
		return s
	}
	for _, c := range s.Children {
		if hit := c.Find(loinc); hit != nil {
			return hit
		}
	}
	return nil
}

// PatientRole is the patient identity block: identifiers, name, sex and birth date.
//
// Identifiers keep their full (root, extension) pair. A NIK without its root cannot be matched
// against anything, and dropping the root is how a local MRN silently becomes a national NIK.
type PatientRole struct {
	Identifiers          []Identifier
	Names                []string
	AdministrativeGender string
	BirthTime            string
}

// Identifier is a CDA II: root names the assigning authority, extension the value.
type Identifier struct {
	Root      string
	Extension string
}

// Name returns the first name, or "" when the document has none.
func (p PatientRole) Name() string {
	if len(p.Names) == 0 {
		return ""
	}
	return p.Names[0]
}

// Identifier returns the extension for this assigning authority, or "".
func (p PatientRole) Identifier(root string) string {
	for _, id := range p.Identifiers {
		if id.Root == root {
			return id.Extension
		}
	}
	return ""
}
