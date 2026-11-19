package cda

import (
	"os"
	"strings"
	"testing"
)

// The fixture is byte-identical to the PHP, Node and Python SDKs': parity is a test, not a claim.
func fixture(t *testing.T) string {
	t.Helper()
	b, err := os.ReadFile("testdata/clinical_document.xml")
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func document(t *testing.T) *Document {
	t.Helper()
	d, err := Parse(fixture(t))
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func TestReadsHeader(t *testing.T) {
	d := document(t)
	if got := d.Title(); got != "Resume Medis Rawat Inap" {
		t.Fatalf("title = %q", got)
	}
	if got := d.TypeCode(); got != "34133-9" {
		t.Fatalf("type = %q", got)
	}
	if got := d.EffectiveTime(); got != "20261118103000" {
		t.Fatalf("effectiveTime = %q", got)
	}
	if got := d.Custodian(); got != "RS Umum Contoh" {
		t.Fatalf("custodian = %q", got)
	}
	if got := d.DocumentID(); got != "2.16.840.1.113883.3.3157.1004.1|DOC-20261118-0001" {
		t.Fatalf("id = %q", got)
	}
}

func TestKeepsIdentifierRootsApart(t *testing.T) {
	p := document(t).Patient()
	if p.Name() != "Siti ANGGRAENI" {
		t.Fatalf("name = %q", p.Name())
	}
	if p.AdministrativeGender != "F" || p.BirthTime != "19850312" {
		t.Fatalf("gender/birth = %q/%q", p.AdministrativeGender, p.BirthTime)
	}
	if len(p.Identifiers) != 2 {
		t.Fatalf("identifiers = %d", len(p.Identifiers))
	}
	if got := p.Identifier("2.16.840.1.113883.3.3157.1003.1"); got != "3171234567890001" {
		t.Fatalf("NIK = %q", got)
	}
	if got := p.Identifier("1.2.840.114350.1.13.1.1"); got != "MRN-0098123" {
		t.Fatalf("MRN = %q", got)
	}
	if got := p.Identifier("9.9.9"); got != "" {
		t.Fatalf("unknown root returned %q", got)
	}
}

func TestReadsAuthors(t *testing.T) {
	authors := document(t).Authors()
	if len(authors) != 1 || authors[0] != "Ivan HARSONO" {
		t.Fatalf("authors = %v", authors)
	}
}

func TestFindsNestedSectionsByLOINC(t *testing.T) {
	d := document(t)
	if got := len(d.Sections()); got != 2 {
		t.Fatalf("sections = %d", got)
	}
	if got := d.Sections()[0].Title; got != "Riwayat Penyakit Sekarang" {
		t.Fatalf("first section = %q", got)
	}
	if diagnosis := d.Section("29548-5"); diagnosis == nil || len(diagnosis.Children) != 1 {
		t.Fatal("nested section under diagnosis not found")
	}
	if d.Section("11450-4") == nil {
		t.Fatal("nested problem section not found")
	}
	if d.Section("00000-0") != nil {
		t.Fatal("unknown LOINC returned a section")
	}
}

func TestSectionTextIsPlainTextWithParagraphBreaks(t *testing.T) {
	text := document(t).SectionText("10164-2")
	if !strings.Contains(text, "demam tiga hari") {
		t.Fatalf("text = %q", text)
	}
	if !strings.Contains(text, "demam tiga hari.\nRiwayat perjalanan") {
		t.Fatalf("paragraph break lost: %q", text)
	}
	if strings.Contains(text, "<paragraph>") {
		t.Fatalf("markup leaked into text: %q", text)
	}
}

func TestValidateAcceptsCompleteDocument(t *testing.T) {
	if problems := document(t).Validate(); len(problems) != 0 {
		t.Fatalf("problems = %v", problems)
	}
}

func TestValidateReportsWhatIsMissing(t *testing.T) {
	d, err := Parse(`<ClinicalDocument xmlns="urn:hl7-org:v3"><title>Kosong</title></ClinicalDocument>`)
	if err != nil {
		t.Fatal(err)
	}
	problems := strings.Join(d.Validate(), "|")
	for _, want := range []string{
		"ClinicalDocument/id wajib ada.",
		"recordTarget wajib ada.",
		"custodian wajib ada.",
		"author wajib ada.",
		"structuredBody wajib berisi minimal satu section.",
	} {
		if !strings.Contains(problems, want) {
			t.Fatalf("missing problem %q in %q", want, problems)
		}
	}
}

func TestBuildProducesValidDocument(t *testing.T) {
	d, err := Build(BuildInput{
		ID:          "DOC-TEST-1",
		Title:       "Resume Medis",
		Custodian:   "RS Contoh",
		Author:      "Dokter Contoh",
		PatientID:   "3171234567890001",
		PatientName: "ANGGRAENI",
		Gender:      "F",
		BirthTime:   "19850312",
		Sections:    []BuildSection{{LOINC: "10164-2", Title: "Riwayat", Text: "<paragraph>Demam.</paragraph>"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if problems := d.Validate(); len(problems) != 0 {
		t.Fatalf("problems = %v", problems)
	}
	if d.Title() != "Resume Medis" || d.Patient().Name() != "ANGGRAENI" {
		t.Fatalf("round trip lost data: %q %q", d.Title(), d.Patient().Name())
	}
	if got := d.Patient().Identifier("2.16.840.1.113883.3.3157.1003.1"); got != "3171234567890001" {
		t.Fatalf("NIK = %q", got)
	}
	if got := d.SectionText("10164-2"); got != "Demam." {
		t.Fatalf("section text = %q", got)
	}
}

func TestRoundTripPreservesUnknownStructure(t *testing.T) {
	again, err := Parse(document(t).XML())
	if err != nil {
		t.Fatal(err)
	}
	xml := again.XML()
	if !strings.Contains(xml, "encompassingEncounter") || !strings.Contains(xml, "ENC-556677") {
		t.Fatal("structure the SDK does not model was dropped")
	}
	if again.Title() != "Resume Medis Rawat Inap" || len(again.Sections()) != 2 {
		t.Fatal("round trip lost data")
	}
}

func TestRefusesNonClinicalDocument(t *testing.T) {
	if _, err := Parse(`<Patient xmlns="http://hl7.org/fhir"><id value="1"/></Patient>`); err == nil {
		t.Fatal("accepted a non-CDA document")
	}
}

func TestRefusesMalformedXML(t *testing.T) {
	if _, err := Parse("<ClinicalDocument><unclosed>"); err == nil {
		t.Fatal("accepted malformed XML")
	}
}
