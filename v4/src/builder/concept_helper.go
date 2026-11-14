package builder

// codedConcept membentuk CodeableConcept satu coding tanpa display, bentuk yang dipakai builder PHP
// untuk elemen status seperti clinicalStatus pada AllergyIntolerance.
func codedConcept(system, code string) map[string]interface{} {
	return map[string]interface{}{
		"coding": []interface{}{map[string]interface{}{"system": system, "code": code}},
	}
}
