package terminology

import (
	"strings"
)

var systemMap = map[string]string{
	"ICD10":   "http://hl7.org/fhir/sid/icd-10",
	"ICD9":    "http://hl7.org/fhir/sid/icd-9-cm",
	"ICD9CM":  "http://hl7.org/fhir/sid/icd-9-cm",
	"LOINC":   "http://loinc.org",
	"SNOMED":  "http://snomed.info/sct",
	"CVX":     "http://hl7.org/fhir/sid/cvx",
	"UCUM":    "http://unitsofmeasure.org",
	"KFA":     "http://fhir.kemkes.go.id/kfa",
	"KPTL":    "http://fhir.kemkes.go.id/kptl",
	"RXNORM":  "http://www.nlm.nih.gov/research/umls/rxnorm",
	"ICDO":    "http://hl7.org/fhir/sid/icd-o",
	"ICDMM":   "http://example.com/icd-mm",
	"ICDPM":   "http://example.com/icd-pm",
	"MTI":     "http://terminology.kemkes.go.id",
}

// Resolve converts a value to FHIR CodeableConcept format.
// Supports: "ICD10:A00", "A00" (bare string), map (pass-through), []interface{} (array).
func Resolve(value interface{}) interface{} {
	switch v := value.(type) {
	case []interface{}:
		result := make([]map[string]interface{}, 0, len(v))
		for _, item := range v {
			result = append(result, Resolve(item).(map[string]interface{}))
		}
		return result
	case map[string]interface{}:
		return v // pass through
	case string:
		colonIdx := strings.Index(v, ":")
		if colonIdx != -1 {
			prefix := strings.ToUpper(v[:colonIdx])
			code := strings.TrimSpace(v[colonIdx+1:])
			system, ok := systemMap[prefix]
			if !ok {
				system = prefix
			}
			return map[string]interface{}{
				"coding": []map[string]string{
					{"system": system, "code": code, "display": code},
				},
				"text": code,
			}
		}
		return map[string]interface{}{"text": v}
	default:
		return map[string]interface{}{"text": v}
	}
}

// ExpandArray expands a []string shorthand into resolved CodeableConcept array.
func ExpandArray(codes []string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0, len(codes))
	for _, code := range codes {
		resolved := Resolve(code)
		if m, ok := resolved.(map[string]interface{}); ok {
			result = append(result, m)
		}
	}
	return result
}

// IsValid checks if a code is valid against a terminology system.
func IsValid(code, system string) bool {
	if strings.TrimSpace(code) == "" {
		return false
	}
	if system != "" {
		_, ok := systemMap[strings.ToUpper(system)]
		return ok
	}
	return true
}
