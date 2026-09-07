package terminology

import (
	"regexp"
	"strings"
)

// TerminologyQueryBuilder — fluent query builder untuk SmarTerm.
//
//	term.Query("pemasangan kateter").On("kptl").Limit(10).Get()
//	term.Query("demam").Category("diagnosis").First()
//	term.Query("paracetamol").On("kfa").All()
//	term.Query("93004944").AsCode().Get()
type TerminologyQueryBuilder struct {
	client   *TerminologyClient
	text     string
	system   string // "" = unified search (/search)
	category string
	limit    int
	asCode   bool
}

var (
	reKFA      = regexp.MustCompile(`^93\d+$`)
	reKFAAlkes = regexp.MustCompile(`^83\d+$`)
	reKFABza   = regexp.MustCompile(`^91\d+$`)
	reSnomed   = regexp.MustCompile(`^\d{5,7}$`)
	reICD10    = regexp.MustCompile(`^[A-Z]\d{2}(\.\d+)?$`)
	reNumeric  = regexp.MustCompile(`^\d+(\.\d+)?$`)
)

// On — scope ke satu sistem: icd10|icd9cm|snomed|loinc|kfa|kfa_alkes|kfa_bza|kptl.
func (b *TerminologyQueryBuilder) On(system string) *TerminologyQueryBuilder {
	b.system = system
	return b
}

// Category — unified search category: all|diagnosis|procedure|medication|lab|imaging|...
func (b *TerminologyQueryBuilder) Category(c string) *TerminologyQueryBuilder {
	b.category = c
	return b
}

// Limit — clamp 1..100.
func (b *TerminologyQueryBuilder) Limit(n int) *TerminologyQueryBuilder {
	b.limit = clampLimit(n)
	return b
}

// First — alias Limit(1).
func (b *TerminologyQueryBuilder) First() *TerminologyQueryBuilder {
	return b.Limit(1)
}

// All — maksimal 100 hasil.
func (b *TerminologyQueryBuilder) All() *TerminologyQueryBuilder {
	return b.Limit(100)
}

// AsCode — treat text sebagai kode eksak → Validate (POST /validate).
// Sistem ditebak dari prefix kode kalau belum di-scope via On().
func (b *TerminologyQueryBuilder) AsCode() *TerminologyQueryBuilder {
	b.asCode = true
	return b
}

// Get — eksekusi: search (GET) atau validate (POST). Return map envelope.
func (b *TerminologyQueryBuilder) Get() (map[string]any, error) {
	if b.asCode {
		system := b.system
		if system == "" {
			system = guessSystemFromCode(b.text)
		}
		return b.client.Validate(system, b.text)
	}
	if b.system != "" {
		return b.client.SearchSystem(b.system, b.text, b.limit)
	}
	return b.client.Search(b.text, b.category, b.limit)
}

// guessSystemFromCode — tebak sistem dari bentuk kode.
func guessSystemFromCode(code string) string {
	switch {
	case reKFA.MatchString(code):
		return "kfa"
	case reKFAAlkes.MatchString(code):
		return "kfa_alkes"
	case reKFABza.MatchString(code):
		return "kfa_bza"
	case reSnomed.MatchString(code):
		return "snomed"
	case reICD10.MatchString(code):
		return "icd10"
	case strings.Contains(code, "-"):
		return "loinc"
	case reNumeric.MatchString(code):
		return "icd9cm"
	default:
		return "snomed"
	}
}
