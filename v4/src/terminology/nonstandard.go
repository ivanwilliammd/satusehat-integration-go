package terminology

// Objek non-standar SATUSEHAT — identifier dan extension di luar FHIR R4 inti.
//
// URL di bawah berasal dari dokumentasi resmi SATUSEHAT dan Implementation Guide, bukan karangan.
// Payload yang "sah secara R4" masih bisa ditolak karena NIK wajib memakai sistem identifier Kemkes.
const (
	NIKSystem          = "https://fhir.kemkes.go.id/id/nik"
	PasporSystem       = "https://fhir.kemkes.go.id/id/paspor"
	KKSystem           = "https://fhir.kemkes.go.id/id/kk"
	IHSSystem          = "https://fhir.kemkes.go.id/id/patient-ihs-number"
	Administrative     = "https://fhir.kemkes.go.id/r4/StructureDefinition/administrativeCode"
	BirthPlace         = "https://fhir.kemkes.go.id/r4/StructureDefinition/birthPlace"
	Citizenship        = "https://fhir.kemkes.go.id/r4/StructureDefinition/citizenshipStatus"
	PatientReligion    = "https://fhir.kemkes.go.id/r4/StructureDefinition/patient-religion"
	PatientDisability  = "https://fhir.kemkes.go.id/r4/StructureDefinition/patient-disability"
	PatientCitizenship = "https://fhir.kemkes.go.id/r4/StructureDefinition/patient-citizenship"
	PatientBirthTime   = "https://fhir.kemkes.go.id/r4/StructureDefinition/patient-birthTime"
	ProfilePatient     = "https://fhir.kemkes.go.id/r4/StructureDefinition/Patient"
)

// NonStandardAll memuat seluruh konstanta di atas, untuk uji dan dokumentasi.
var NonStandardAll = [][2]string{
	{"nik", NIKSystem},
	{"paspor", PasporSystem},
	{"kk", KKSystem},
	{"ihs", IHSSystem},
	{"administrative_code", Administrative},
	{"birth_place", BirthPlace},
	{"citizenship_status", Citizenship},
	{"patient_religion", PatientReligion},
	{"patient_disability", PatientDisability},
	{"patient_citizenship", PatientCitizenship},
	{"patient_birth_time", PatientBirthTime},
}

// OfficialIdentifier membuat identifier resmi SATUSEHAT. Use selalu "official".
func OfficialIdentifier(system, value string) map[string]any {
	return map[string]any{"use": "official", "system": system, "value": value}
}

// AdministrativeCode membuat extension administrativeCode dengan irisan bertingkat.
// Satu valueCode datar sah secara R4 tetapi tidak terbaca sebagai wilayah.
func AdministrativeCode(slices [][2]string) map[string]any {
	isi := make([]any, 0, len(slices))
	for _, s := range slices {
		isi = append(isi, map[string]any{"url": s[0], "valueCode": s[1]})
	}
	return map[string]any{"url": Administrative, "extension": isi}
}
