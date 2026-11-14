package ssresponse

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// RuleInfo holds metadata for a SATUSEHAT validation rule.
type RuleInfo struct {
	Path        string
	Terminology string
	Description string
	Fix         string
	Version     string
	LastUpdate  string
}

// ToMap converts RuleInfo to a map[string]string matching the PHP associative array.
func (r RuleInfo) ToMap() map[string]string {
	return map[string]string{
		"path":        r.Path,
		"terminology": r.Terminology,
		"description": r.Description,
		"fix":         r.Fix,
		"version":     r.Version,
		"last_update": r.LastUpdate,
	}
}

// ResponseTypeInfo holds metadata for a SATUSEHAT validation error response type.
type ResponseTypeInfo struct {
	Title string
	Cause string
	Fix   string
}

// ToMap converts ResponseTypeInfo to a map[string]string.
func (rt ResponseTypeInfo) ToMap() map[string]string {
	return map[string]string{
		"title": rt.Title,
		"cause": rt.Cause,
		"fix":   rt.Fix,
	}
}

// RuleMap contains all 558 SATUSEHAT validation rules.
var RuleMap = map[int]RuleInfo{
	10003: {Path: "Condition.code", Terminology: "http://hl7.org/fhir/sid/icd-10\nhttp://terminology.kemkes.go.id/CodeSystem/clinical-term\nhttp://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Condition.code", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10006: {Path: "Condition.clinicalStatus", Terminology: "http://terminology.hl7.org/CodeSystem/condition-clinical", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT sub-bab Condition.clinicalStatus", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10007: {Path: "Condition.category", Terminology: "http://terminology.hl7.org/CodeSystem/condition-category\nhttp://terminology.kemkes.go.id", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Condition.category", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10008: {Path: "Observation.status", Terminology: "http://hl7.org/fhir/observation-status", Description: "Code", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10009: {Path: "Observation.category", Terminology: "http://terminology.hl7.org/CodeSystem/observation-category", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.category", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10010: {Path: "Observation.code", Terminology: "ANC\nhttp://loinc.org\nhttp://snomed.info/sct\nhttp://terminology.kemkes.go.id/CodeSystem/examination", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.code\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.code dan Akses playbook secara lengkap disini.\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.code.coding", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10011: {Path: "Observation.bodySite", Terminology: "http://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.bodySite", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10012: {Path: "Observation.valueQuantity.code", Terminology: "http://unitsofmeasure.org", Description: "Data Observation.valueQuantity.code harus menggunakan satuan yang tersedia pada codesystem dari http://unitsofmeasure.org dapat dilihat disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10013: {Path: "Procedure.category", Terminology: "http://snomed.info/sct\nhttp://terminology.kemkes.go.id", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Procedure.category", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10015: {Path: "Procedure.code", Terminology: "http://hl7.org/fhir/sid/icd-9-cm\nhttp://snomed.info/sct\nhttp://terminology.kemkes.go.id/CodeSystem/clinical-term\nhttp://terminology.kemkes.go.id/CodeSystem/kptl", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Procedure.code", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10016: {Path: "Procedure.bodySite", Terminology: "http://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Procedure.bodySite", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10018: {Path: "Procedure.reasonCode", Terminology: "http://hl7.org/fhir/sid/icd-10", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Procedure.reasonCode", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10019: {Path: "Procedure.status", Terminology: "http://hl7.org/fhir/event-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Procedure.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10020: {Path: "Composition.status", Terminology: "http://hl7.org/fhir/composition-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Composition.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10021: {Path: "Composition.type.coding", Terminology: "http://loinc.org", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Composition.type.coding", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10022: {Path: "Composition.category", Terminology: "http://loinc.org", Description: "Data harus menggunakan terminologi LOINC yang sesuai dengan use-case terkait, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Composition.category", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10023: {Path: "Composition.section.code", Terminology: "http://loinc.org", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Composition.section[i].code", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10024: {Path: "Medication.code", Terminology: "http://sys-ids.kemkes.go.id/kfa\nhttp://snomed.info/sct", Description: "Data harus menggunakan kode obat yang tersedia pada KFA (kamus farmasi dan alat kesehatan). Pengisian kode KFA tergantung jenis obat yang diresepkan atau dikeluarkan tergantung apakah dalam bentuk produk obat template (92xxxxxx) atau produk obat aktual (93xxxxxx). Daftar kode obat KFA dapat dilihat dalam browser kamus KFA pada link berikut : https://dto.kemkes.go.id/kfa-browser.\nNilai pada elemen Medication.code harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10025: {Path: "Medication.form", Terminology: "http://terminology.kemkes.go.id/CodeSystem/medication-form", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Medication.form", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10026: {Path: "Medication.status", Terminology: "http://hl7.org/fhir/CodeSystem/medication-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Medication.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10027: {Path: "Medication.ingredient.strength.numerator", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for\nUnits of Measure (UCUM). Daftar kode satuan dapat diakses disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10028: {Path: "Medication.ingredient.strength.denominator", Terminology: "http://snomed.info/sct\nhttp://terminology.hl7.org/CodeSystem/v3-orderableDrugForm\nhttp://unitsofmeasure.org", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Medication.ingredient.strength.denominator\nData harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for\nUnits of Measure (UCUM). Daftar kode satuan dapat diakses disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10030: {Path: "Medication.ingredient.itemCodableConcept.coding", Terminology: "http://sys-ids.kemkes.go.id/kfa", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Medication.ingredient.itemCodableConcept.coding", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10031: {Path: "Medication.extension.valueCodableConcept.coding", Terminology: "http://terminology.kemkes.go.id/CodeSystem/medication-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Medication.extension.valueCodableConcept.coding", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10032: {Path: "MedicationRequest.status", Terminology: "http://hl7.org/fhir/CodeSystem/medicationrequest-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10033: {Path: "MedicationRequest.intent", Terminology: "http://hl7.org/fhir/CodeSystem/medicationrequest-intent", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.intent", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10034: {Path: "MedicationRequest.category", Terminology: "http://terminology.hl7.org/CodeSystem/medicationrequest-category", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.category", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10035: {Path: "MedicationRequest.reasonCode", Terminology: "http://hl7.org/fhir/sid/icd-10", Description: "Data harus menggunakan Kode dapat diisi dengan kode ICD-10 versi 2010, cara penulisan lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.reasonCode", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10036: {Path: "MedicationRequest.dosageInstruction.timing.periodUnit", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.dosageInstruction.timing.periodUnit", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10037: {Path: "MedicationRequest.dosageInstruction.site", Terminology: "http://snomed.info/sct", Description: "Data harus menggunakan Kode SNOMED CT yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.dosageInstruction.site", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10038: {Path: "MedicationRequest.dosageInstruction.route", Terminology: "http://www.whocc.no/atc", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.dosageInstruction.route", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10039: {Path: "MedicationRequest.dosageInstruction.doseAndRate.type", Terminology: "http://terminology.hl7.org/CodeSystem/dose-rate-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.dosageInstruction.doseAndRate.type", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10040: {Path: "MedicationRequest.dosageInstruction.method", Terminology: "http://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.dosageInstruction.method", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10041: {Path: "MedicationRequest.dispenseRequest.initialFill.duration", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for Units of Measure (UCUM) untuk satuan waktu.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10042: {Path: "MedicationRequest.dispenseRequest.dispenseInterval", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.dispenseRequest.dispenseInterval", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10043: {Path: "MedicationRequest.dispenseRequest.expectedSupplyDuration", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.dispenseRequest.expectedSupplyDuration", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10044: {Path: "MedicationRequest.priority", Terminology: "http://hl7.org/fhir/request-priority", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.priority", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10045: {Path: "MedicationRequest.courseOfTherapyType", Terminology: "http://terminology.hl7.org/CodeSystem/medicationrequest-course-of-therapy", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.courseOfTherapyType", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10046: {Path: "MedicationRequest.dosageInstruction.timing.code", Terminology: "http://terminology.hl7.org/CodeSystem/v3-GTSAbbreviation", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.dosageInstruction[i].timing.co\nde", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10047: {Path: "MedicationRequest.dosageInstruction.additionalInstruction.coding.code", Terminology: "http://snomed.info/sct", Description: "Data harus menggunakan Kode dapat diisi dengan kode SNOMED-CT. Ruang lingkup kode SNOMED CT versi\nDesember 2022 yang dapat dipakai : Expression Constraint Language (ECL) Query : < 419492006 |Additional dosage instructions (qualifier value)| Contoh daftar kode SNOMED CT yang dapat digunakan dapat diakses disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10048: {Path: "MedicationDispense.category", Terminology: "http://terminology.hl7.org/fhir/CodeSystem/medicationdispense-category", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationDispense.category", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10049: {Path: "MedicationDispense.status", Terminology: "http://terminology.hl7.org/CodeSystem/medicationdispense-status\nnot null", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationDispense.status\nData harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10050: {Path: "MedicationDispense.quantity.code", Terminology: "http://terminology.hl7.org/CodeSystem/v3-orderableDrugForm", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationDispense.quantity.code", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10051: {Path: "MedicationDispense.daysSupply.code", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for\nUnits of Measure (UCUM). Daftar kode satuan dapat diakses disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10052: {Path: "MedicationDispense.dosageInstruction.timing.periodUnit", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for\nUnits of Measure (UCUM). Daftar kode satuan dapat diakses disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10053: {Path: "MedicationDispense.dosageInstruction.doseAndRate.type", Terminology: "http://terminology.hl7.org/CodeSystem/dose-rate-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationDispense.dosageInstruction.doseAndRate.type", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10054: {Path: "MedicationDispense.dosageInstruction.additionalInstruction", Terminology: "http://snomed.info/sct", Description: "Data harus menggunakan Kode dapat diisi dengan kode SNOMED-CT. Ruang lingkup kode SNOMED CT versi Desember 2022 yang dapat dipakai : Expression Constraint Language (ECL) Query : < 419492006 |Additional dosage instructions (qualifier value)| Contoh daftar kode SNOMED CT yang dapat digunakan dapat diakses disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10055: {Path: "MedicationDispense.dosageInstruction.timing.code", Terminology: "http://terminology.hl7.org/CodeSystem/v3-GTSAbbreviation", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationDispense.dosageInstruction[i].timing.c\node", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10056: {Path: "MedicationDispense.dosageInstruction.route", Terminology: "http://www.whocc.no/atc", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationDispense.dosageInstruction.route", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10057: {Path: "MedicationDispense.dosageInstruction.doseAndRate.type", Terminology: "http://terminology.hl7.org/CodeSystem/dose-rate-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationDispense.dosageInstruction.doseAndRate.type", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10058: {Path: "MedicationDispense.dosageInstruction.timing.repeat.when", Terminology: "http://hl7.org/fhir/ValueSet/event-timing", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationDispense.dosageInstruction[i].timing.repeat.when", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10059: {Path: "ServiceRequest.bodySite", Terminology: "http://snomed.info/sct", Description: "Data harus menggunakan Kode dapat diisi dengan kode SNOMED-CT. Ruang lingkup kode SNOMED CT versi Desember 2022 yang dapat dipakai : Expression Constraint Language (ECL) Query : < 123037004 |Body structure (body structure)|", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10060: {Path: "ServiceRequest.code", Terminology: "http://hl7.org/fhir/sid/icd-9-cm\nhttp://loinc.org\nhttp://snomed.info/sct\nhttp://terminology.kemkes.go.id/CodeSystem/examination", Description: "Data harus menggunakan Kode ICD-9 CM untuk tindakan atau prosedur medis untuk keperluan klaim, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT sub-bab ServiceRequest.code\nData harus menggunakan terminologi LOINC yang sesuai dengan use-case terkait. Akses playbook secara lengkap disini.\nData harus menggunakan Kode dapat diisi dengan kode SNOMED-CT. Ruang lingkup kode SNOMED CT yang dapat dipakai :\nExpression Constraint Language (ECL) Query : < 71388002 |Procedure (procedure)|\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ServiceRequest.code", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10063: {Path: "ServiceRequest.intent", Terminology: "http://hl7.org/fhir/request-intent", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ServiceRequest.intent", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10064: {Path: "ServiceRequest.reasonCode", Terminology: "http://hl7.org/fhir/sid/icd-10\nhttp://hl7.org/fhir/sid/icd-9-cm\nhttp://snomed.info/sct", Description: "Data harus menggunakan Kode dapat diisi dengan kode ICD-10 versi 2010, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT sub-bab ServiceRequest.reasonCode\nData harus menggunakan Kode ICD-9 CM untuk tindakan atau prosedur medis untuk keperluan klaim, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT sub-bab ServiceRequest.reasonCode\nNilai pada elemen ServiceRequest.reasonCode.coding harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10065: {Path: "ServiceRequest.status", Terminology: "http://hl7.org/fhir/request-status\nnot null", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ServiceRequest.status\nData harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10066: {Path: "ServiceRequest.priority", Terminology: "http://hl7.org/fhir/request-priority", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ServiceRequest.priority", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10067: {Path: "ServiceRequest.locationCode", Terminology: "http://terminology.hl7.org/CodeSystem/v3-RoleCode", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ServiceRequest.locationCode", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10068: {Path: "Specimen.collection.method", Terminology: "http://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Specimen.collection.method", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10069: {Path: "Specimen.status", Terminology: "http://hl7.org/fhir/specimen-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Specimen.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10070: {Path: "Specimen.type", Terminology: "http://terminology.kemkes.go.id/CodeSystem/clinical-term\nhttp://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Specimen.type", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10071: {Path: "DiagnosticReport.code", Terminology: "http://terminology.kemkes.go.id/CodeSystem/examination\nhttp://loinc.org", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab DiagnosticReport.code", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10072: {Path: "DiagnosticReport.status", Terminology: "http://hl7.org/fhir/diagnostic-report-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab DiagnosticReport.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10073: {Path: "DiagnosticReport.category", Terminology: "http://terminology.hl7.org/CodeSystem/v2-0074", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab DiagnosticReport.category", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10074: {Path: "DiagnosticReport.conclusionCode", Terminology: "http://loinc.org\nhttp://snomed.info/sct\nhttp://terminology.hl7.org/CodeSystem/v3-ObservationInterpretation\nhttp://terminology.kemkes.go.id/CodeSystem/examination", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab DiagnosticReport.conclusionCode", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10075: {Path: "AllergyIntolerance.category", Terminology: "http://hl7.org/fhir/allergy-intolerance-category", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab AllergyIntolerance.category", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10076: {Path: "AllergyIntolerance.clinicalStatus", Terminology: "http://terminology.hl7.org/CodeSystem/allergyintolerance-clinical", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab AllergyIntolerance.clinicalStatus", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10077: {Path: "AllergyIntolerance.criticality", Terminology: "http://hl7.org/fhir/allergy-intolerance-criticality", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab AllergyIntolerance.criticality", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10078: {Path: "AllergyIntolerance.code", Terminology: "http://snomed.info/sct\nhttp://sys-ids.kemkes.go.id/kfa\nhttp://terminology.kemkes.go.id/CodeSystem/clinical-term\nnot null", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab AllergyIntolerance.code\nData harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10079: {Path: "AllergyIntolerance.verificationStatus", Terminology: "http://terminology.hl7.org/CodeSystem/allergyintolerance-verification", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab AllergyIntolerance.verificationStatus", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10080: {Path: "ClinicalImpression.status", Terminology: "http://hl7.org/fhir/event-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ClinicalImpression.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10081: {Path: "ClinicalImpression.investigation.code", Terminology: "http://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ClinicalImpression.investigation.code", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10082: {Path: "ClinicalImpression.finding.itemCodeableConcept", Terminology: "http://hl7.org/fhir/sid/icd-10", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ClinicalImpression.finding.itemCodeableConcept", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10083: {Path: "ClinicalImpression.prognosisCodeableConcept", Terminology: "http://terminology.kemkes.go.id/CodeSystem/clinical-term\nhttp://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Clinicallimpression.prognosisCodeableConcept", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10089: {Path: "Organization.type", Terminology: "http://terminology.hl7.org/CodeSystem/organization-type\nhttp://terminology.kemkes.go.id/CodeSystem/organization-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Organization.type", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10090: {Path: "Organization.telecom.system", Terminology: "http://hl7.org/fhir/contact-point-system", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Organization.telecom[i].system", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10091: {Path: "Organization.telecom.use", Terminology: "http://hl7.org/fhir/contact-point-use", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Organization.telecom[i].use", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10092: {Path: "Organization.address.use", Terminology: "http://hl7.org/fhir/address-use", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Organization.address[i].use", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10093: {Path: "Organization.address.type", Terminology: "http://hl7.org/fhir/address-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Organization.address[i].type", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10094: {Path: "Organization.contact.purpose", Terminology: "http://terminology.hl7.org/CodeSystem/contactentity-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Organization.contact[i].purpose", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10095: {Path: "Location.status", Terminology: "http://hl7.org/fhir/location-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Location.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10096: {Path: "Location.operationalStatus", Terminology: "http://terminology.hl7.org/CodeSystem/v2-0116", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Location.operationalStatus", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10097: {Path: "Location.mode", Terminology: "http://hl7.org/fhir/location-mode", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Location.mode", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10098: {Path: "Location.telecom.system", Terminology: "http://hl7.org/fhir/contact-point-system", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Location.telecom[i].system", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10099: {Path: "Location.telecom.use", Terminology: "http://hl7.org/fhir/contact-point-use", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Location.telecom[i].use", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10100: {Path: "Location.address.use", Terminology: "http://hl7.org/fhir/address-use", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Location.address.use", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10101: {Path: "Location.address.type", Terminology: "http://hl7.org/fhir/address-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Location.address.type", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10102: {Path: "Location.physicalType", Terminology: "http://terminology.hl7.org/CodeSystem/location-physical-type\nhttp://terminology.kemkes.go.id/CodeSystem/location-physical-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Location.physicalType", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10103: {Path: "Immunization.vaccineCode", Terminology: "http://hl7.org/fhir/sid/cvx\nhttp://sys-ids.kemkes.go.id/kfa", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Immunization.vaccineCode", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10104: {Path: "Immunization.performer.function", Terminology: "http://terminology.hl7.org/3.1.0/CodeSystem-v2-0443.html", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Immunization.performer.function", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10105: {Path: "Immunization.reasonCode", Terminology: "http://terminology.kemkes.go.id/CodeSystem/immunization-reason\nhttp://terminology.kemkes.go.id/CodeSystem/immunization-routine-timing\nhttp://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Immunization.reasonCode", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10106: {Path: "Immunization.reportOrigin", Terminology: "http://terminology.hl7.org/CodeSystem/immunization-origin", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Immunization.reportOrigin", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10107: {Path: "Immunization.status", Terminology: "http://hl7.org/fhir/event-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Immunization.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10108: {Path: "Immunization.statusReason", Terminology: "http://snomed.info/sct\nhttp://terminology.hl7.org/CodeSystem/v3-ActReason\nhttp://terminology.kemkes.go.id/CodeSystem/immunization-statusReason", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Immunization.statusReason", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10109: {Path: "EpisodeOfCare.status", Terminology: "http://hl7.org/fhir/episode-of-care-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab EpisodeOfCare.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10110: {Path: "EpisodeOfCare.type", Terminology: "http://terminology.hl7.org/CodeSystem/episodeofcare-type\nhttp://terminology.kemkes.go.id/CodeSystem/episodeofcare-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab EpisodeOfCare.type", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10111: {Path: "EpisodeOfCare.diagnosis.role", Terminology: "http://terminology.hl7.org/CodeSystem/diagnosis-role", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab EpisodeOfCare.diagnosis.role", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10112: {Path: "QuestionnaireResponse.status", Terminology: "http://hl7.org/fhir/ValueSet/questionnaire-answers-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab QuestionnaireResponse.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10113: {Path: "FamilyMemberHistory.status", Terminology: "http://hl7.org/fhir/codesystem-history-status.html", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab FamilyMemberHistory.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10114: {Path: "FamilyMemberHistory.relationship", Terminology: "http://snomed.info/sct\nhttp://terminology.hl7.org/CodeSystem/v3-RoleCode", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab FamilyMemberHistory.relationship", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10115: {Path: "CarePlan.status", Terminology: "http://hl7.org/fhir/ValueSet/request-status\nhttp://hl7.org/fhir/request-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab CarePlan.status\nNilai pada elemen CarePlan.status harus sesuai dengan sistem http://hl7.org/fhir/request-status.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10116: {Path: "CarePlan.intent", Terminology: "http://hl7.org/fhir/ValueSet/care-plan-intent\nhttp://hl7.org/fhir/request-intent", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab CarePlan.intent\nNilai pada elemen CarePlan.intent harus sesuai dengan sistem http://hl7.org/fhir/request-intent.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10117: {Path: "Encounter.identifier", Terminology: "http://terminology.kemkes.go.id/CodeSystem/episodeofcare/ANC\nhttp://terminology.kemkes.go.id/CodeSystem/episodeofcare/neonate\nhttp://terminology.kemkes.go.id/CodeSystem/episodeofcare/puerperium\nnot null\nhttp://sys-ids.kemkes.go.id/asik/{{org-id}}\nhttp://sys-ids.kemkes.go.id/sitb/{{org-id}}", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.identifier\nData harus terisi (tidak boleh kosong)\nData harus menggunakan codesystem http://sys-ids.kemkes.go.id/asik/{{org-id}} dengan {orgID} adalah Organization ID faskes\nData harus menggunakan codesystem http://sys-ids.kemkes.go.id/sitb/{{org-id}} dengan {{org_id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10118: {Path: "Encounter.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10119: {Path: "Encounter.participant.individual", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10120: {Path: "Encounter.location.location", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10121: {Path: "Encounter.period", Terminology: "\nnot null", Description: "Jika elemen Encounter.status bernilai 'finished', maka Encounter.period.start dan Encounter.period.end wajib diisi, serta nilai Encounter.period.end harus lebih besar dari Encounter.period.start.\"\nData harus terisi (tidak boleh kosong)", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10122: {Path: "Encounter.statusHistory", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10123: {Path: "Encounter.period.start", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10124: {Path: "Encounter.serviceProvider", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10125: {Path: "Condition.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10126: {Path: "Condition.encounter", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10127: {Path: "Procedure.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10128: {Path: "Procedure.encounter", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10129: {Path: "Observation.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10130: {Path: "Observation.encounter", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10131: {Path: "Composition.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10132: {Path: "Composition.date", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10133: {Path: "Composition.author", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10134: {Path: "Composition.title", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10135: {Path: "MedicationRequest.medication[x]", Terminology: "not null\nReference ke Medication", Description: "Data harus terisi (tidak boleh kosong)\nPath MedicationRequest.medicationReference wajib terisi dan merujuk pada nomor ID Medication yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10136: {Path: "MedicationRequest.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10138: {Path: "MedicationDispense.medication", Terminology: "not null\nMedication/{{medication-id}}", Description: "Data harus terisi (tidak boleh kosong)\nData MedicationDispense.medicationReference wajib terisi dengan Reference yang merujuk pada data Medication dan {medication-id}} adalah ID data medication yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10139: {Path: "MedicationDispense.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10141: {Path: "ServiceRequest.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10142: {Path: "ServiceRequest.encounter", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10143: {Path: "Specimen.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10144: {Path: "DiagnosticReport.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10145: {Path: "DiagnosticReport.encounter", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10147: {Path: "AllergyIntolerance.patient", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10148: {Path: "ClinicalImpression.subject", Terminology: "not null\nReference ke Patient", Description: "Data harus terisi (tidak boleh kosong)\nReferensi pada elemen ClinicalImpression.subject harus mengarah ke resource Patient.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10149: {Path: "ClinicalImpression.encounter", Terminology: "not null\nReference ke Encounter", Description: "Data harus terisi (tidak boleh kosong)\nReferensi pada elemen ClinicalImpression.encounter harus mengarah ke resource Encounter.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10150: {Path: "ImagingStudy.identifier", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10151: {Path: "ImagingStudy.status", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10152: {Path: "ImagingStudy.modality", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10153: {Path: "ImagingStudy.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10154: {Path: "ImagingStudy.basedOn", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10155: {Path: "Immunization.patient", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10156: {Path: "Immunization.occurrence[x]", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10157: {Path: "EpisodeOfCare.patient", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10158: {Path: "FamilyMemberHistory.patient", Terminology: "not null\nReference ke Patient", Description: "Data harus terisi (tidak boleh kosong)\nReferensi pada elemen FamilyMemberHistory.patient harus mengarah ke resource Patient.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10159: {Path: "MedicationStatement.status", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10160: {Path: "MedicationStatement.medication[x]", Terminology: "MedicationStatement.medication[x]", Description: "Data medication[x] wajib dikirimkan untuk resource MedicationStatement, lihat ketentuan MedicationStatement disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10161: {Path: "MedicationStatement.subject", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10162: {Path: "RelatedPerson.patient", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10163: {Path: "Location.extension.serviceClass", Terminology: "http://terminology.kemkes.go.id/CodeSystem/locationServiceClass", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Location.extension.serviceClass", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10164: {Path: "Encounter.location.extension:serviceClass.value.valueCodeableConcept.coding.system", Terminology: "http://terminology.kemkes.go.id/CodeSystem/locationServiceClass", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.location.extension:serviceClass.value.valueCodeableConcept.coding.system", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10165: {Path: "Encounter.location.extension:serviceClass.upgradeClassIndicator", Terminology: "http://terminology.kemkes.go.id/CodeSystem/locationUpgradeClass", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.location.extension:serviceClass.upgradeClassIndicator", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10166: {Path: "Patient.identifier", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10167: {Path: "Patient.multipleBirthInteger", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10168: {Path: "Patient.name.text", Terminology: "not null and \"\"", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10169: {Path: "QuestionnaireResponse.questionnaire", Terminology: "https://fhir.kemkes.go.id/Questionnaire/Q0006\nhttps://fhir.kemkes.go.id/Questionnaire/Q0007\nhttps://fhir.kemkes.go.id/Questionnaire/Q0008\nhttps://fhir.kemkes.go.id/Questionnaire/Q0011\nhttps://fhir.kemkes.go.id/Questionnaire/Q0023", Description: "\nNilai pada elemen QuestionnaireResponse.questionnaire harus sesuai dengan sistem https://fhir.kemkes.go.id/Questionnaire/Q0023.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10170: {Path: "Condition.verificationStatus", Terminology: "http://terminology.hl7.org/CodeSystem/condition-ver-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Condition.verificationStatus", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10171: {Path: "Location.operationalStatus", Terminology: "http://terminology.hl7.org/CodeSystem/v2-0116", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Location.operationalStatus", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10172: {Path: "Location.type", Terminology: "http://terminology.hl7.org/CodeSystem/v3-RoleCode\nhttp://terminology.kemkes.go.id/CodeSystem/location-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Location.type", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10173: {Path: "Encounter.period.end", Terminology: "future date", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00 dan value lebih dari value pada path Encounter.period.start", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10208: {Path: "CoverageEligibilityResponse.request", Terminology: "Reference ke CoverageEligibilityRequest", Description: "Referensi pada elemen CoverageEligibilityResponse.request harus mengarah ke resource CoverageEligibilityRequest.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10212: {Path: "CoverageEligibilityResponse.insurance.coverage", Terminology: "Reference ke Coverage", Description: "Referensi pada elemen CoverageEligibilityResponse.insurance.coverage harus mengarah ke resource Coverage.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10217: {Path: "ChargeItem.subject", Terminology: "Reference ke Patient", Description: "Referensi pada elemen ChargeItem.subject harus mengarah ke resource Patient.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10218: {Path: "ChargeItem.context", Terminology: "Reference ke Encounter", Description: "Referensi pada elemen ChargeItem.context harus mengarah ke resource Encounter.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10222: {Path: "ChargeItem.account", Terminology: "Reference ke Account", Description: "Referensi pada elemen ChargeItem.account harus mengarah ke resource Account.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10225: {Path: "Account.name", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10226: {Path: "Account.subject", Terminology: "Reference ke Patient", Description: "Referensi pada elemen Account.subject harus mengarah ke resource Patient.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10227: {Path: "Account.servicePeriod.start", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10228: {Path: "Account.coverage", Terminology: "not null\nReference ke Coverage", Description: "Data harus terisi (tidak boleh kosong)\nReferensi pada elemen Account.coverage.coverage harus mengarah ke resource Coverage.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10229: {Path: "Account.owner", Terminology: "not null\nReference ke Organization", Description: "Data harus terisi (tidak boleh kosong)\nReferensi pada elemen Account.owner harus mengarah ke resource Organization.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10231: {Path: "Invoice.subject", Terminology: "Reference ke Patient", Description: "Referensi pada elemen Invoice.subject harus mengarah ke resource Patient.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10233: {Path: "Invoice.issuer", Terminology: "Reference ke Organization", Description: "Referensi pada elemen Invoice.issuer harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10234: {Path: "Invoice.account", Terminology: "Reference ke Account", Description: "Referensi pada elemen Invoice.account harus mengarah ke resource Account.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10236: {Path: "Invoice.chargeItemRefrence", Terminology: "Reference ke ChargeItem", Description: "Referensi pada elemen Invoice.chargeItemRefrence harus mengarah ke resource ChargeItem.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10243: {Path: "Claim.patient", Terminology: "Reference ke Patient", Description: "Referensi pada elemen Claim.patient harus mengarah ke resource Patient.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10246: {Path: "Claim.enterer", Terminology: "Reference ke Practitioner", Description: "Referensi pada elemen Claim.enterer harus mengarah ke resource Practitioner.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10247: {Path: "Claim.insurer", Terminology: "Reference ke Organization", Description: "Referensi pada elemen Claim.insurer harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10248: {Path: "Claim.provider", Terminology: "Reference ke Organization", Description: "Referensi pada elemen Claim.provider harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10252: {Path: "Claim.payee.party", Terminology: "Reference ke Organization", Description: "Referensi pada elemen Claim.payee.party harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10253: {Path: "Claim.facility", Terminology: "Reference ke Location", Description: "Referensi pada elemen Claim.facility harus mengarah ke resource Location.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10264: {Path: "Claim.insurance.coverage", Terminology: "Reference ke Coverage", Description: "Referensi pada elemen Claim.insurance.coverage harus mengarah ke resource Coverage.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10266: {Path: "Claim.total.currency", Terminology: "urn:iso:std:iso:4217\ndecimal", Description: "Nilai pada elemen Claim.total.currency dapat diisi dengan kode 'IDR' sesuai standar ISO 4217.\nNilai pada elemen Claim.total.value harus valid sesuai aturan: Tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10272: {Path: "ClaimResponse.subType", Terminology: "http://terminology.kemkes.go.id", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ClaimResponse.subType", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10274: {Path: "ClaimResponse.patient", Terminology: "Reference ke Patient", Description: "Referensi pada elemen ClaimResponse.patient harus mengarah ke resource Patient.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10276: {Path: "ClaimResponse.insurer", Terminology: "Reference ke Organization", Description: "Referensi pada elemen ClaimResponse.insurer harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10277: {Path: "ClaimResponse.requestor", Terminology: "Reference ke Organization", Description: "Referensi pada elemen ClaimResponse.requestor harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10278: {Path: "ClaimResponse.request", Terminology: "Reference ke Claim", Description: "Referensi pada elemen ClaimResponse.request harus mengarah ke resource Claim.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10283: {Path: "ClaimResponse.adjudication.category", Terminology: "http://terminology.kemkes.go.id\nhttp://snomed.info/sct\nhttp://loinc.org\nhttp://sys-ids.kemkes.go.id/administrative-area\nhttp://terminology.kemkes.go.id/CodeSystem/clinical-speciality\nhttp://terminology.kemkes.go.id/CodeSystem/clinical-term\nhttp://terminology.kemkes.go.id/Codesystem/kpsp\nhttp://terminology.kemkes.go.id/CodeSystem/prev-tb-treatment\nhttp://terminology.kemkes.go.id/CodeSystem/tb-anatomical\nhttp://terminology.kemkes.go.id/CodeSystem/tb-case-definition\nhttp://terminology.kemkes.go.id/CodeSystem/tb-outcome-class", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ClaimResponse.adjudication.category\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab QuestionnaireResponse.item.item.answer.valueCoding", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10284: {Path: "ClaimResponse.adjudication.amount.currency", Terminology: "urn:iso:std:iso:4217\ndecimal\nhttp://terminology.kemkes.go.id", Description: "Nilai pada elemen ClaimResponse.adjudication.amount.currency dapat diisi dengan kode 'IDR' sesuai standar ISO 4217.\nNilai pada elemen ClaimResponse.adjudication.amount.value harus valid sesuai aturan: Tipe data decimal.\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ClaimResponse.adjudication.reason", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10286: {Path: "ClaimResponse.total.currency", Terminology: "urn:iso:std:iso:4217\ndecimal", Description: "Nilai pada elemen ClaimResponse.total.currency dapat diisi dengan kode 'IDR' sesuai standar ISO 4217.\nNilai pada elemen ClaimResponse.total.value harus valid sesuai aturan: Tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10289: {Path: "Account.servicePeriod.end", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10290: {Path: "Encounter.location.period.start", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10291: {Path: "Encounter.location.period.end", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10292: {Path: "Immunization.primarySource", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10293: {Path: "Immunization.encounter", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10294: {Path: "Immunization.performer.actor", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10295: {Path: "Observation.effectiveDateTime", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10296: {Path: "Observation.issued", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10297: {Path: "Observation.valueCodeableConcept", Terminology: "http://hl7.org/fhir/sid/icd-10\nhttp://snomed.info/sct\nhttp://terminology.kemkes.go.id/CodeSystem/clinical-term\nhttp://terminology.kemkes.go.id/CodeSystem/examination\nhttp://terminology.kemkes.go.id/CodeSystem/organization-type\nhttp://terminology.kemkes.go.id/CodeSystem/tb-case-definition", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.code dan Akses playbook secara lengkap disini.\nKode Observation.valueCodeableConcept dapat diisi dengan kode LOINC maupun kode SNOMED CT sesuai dengan use case masing-masing. Untuk keterangan lebih lanjut dapat melihat dalam playbook masing-masing usecase.\nKode Observation.valueCodeableConcept dapat diisi dengan kode LOINC maupun kode SNOMED CT sesuai dengan use case  masing-masing. Untuk keterangan lebih lanjut dapat melihat dalam playbook masing-masing usecase disini\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.valueCodeableConcept\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.valueCodeableConcept.coding", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10299: {Path: "MedicationStatement.dateAsserted", Terminology: "dateTime", Description: "Data tidak boleh kosong (harus diisi) dan harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10301: {Path: "Encounter.participant.period.end", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10302: {Path: "Encounter.classHistory.period.start", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10303: {Path: "Encounter.classHistory.period.end", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10304: {Path: "Encounter.classHistory.class", Terminology: "http://terminology.hl7.org/CodeSystem/v3-ActCode", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.classHistory[i].class", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10305: {Path: "Immunization.location", Terminology: "not null if", Description: "Data immunization.location wajib diisi pada pengiriman resource Immunization jika Immunization.status = completed", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10306: {Path: "Immunization.lotNumber", Terminology: "not null if", Description: "Data immunization.lotNumber harus terisi pada pengiriman resource Immunization dengan use case imunisasi dilakukan oleh nakes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10307: {Path: "Immunization.expirationDate", Terminology: "\nImmunization.expiratonDate > Immunization.occurenceDateTime", Description: "Value pada path Immunization.expirationDate wajib diisi pada pengiriman resource Immunization dengan use case imunisasi dilakukan oleh nakes\nNilai Immunization.expiratonDate harus lebih besar dari Immunization.occurenceDateTime.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10309: {Path: "Medication.batch.lotnumber", Terminology: "data string, gabungan alphabet dan numerik. Contoh : 1625042A", Description: "Input data tipe string", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10310: {Path: "Medication.batch.expirationDate", Terminology: "dateTime", Description: "Data harus menggunakan format datetime", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10317: {Path: "Encounter.diagnosis.use", Terminology: "http://terminology.hl7.org/CodeSystem/diagnosis-role", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.diagnosis.use", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10318: {Path: "Encounter.hospitalization.dischargeDisposition", Terminology: "http://hl7.org/fhir/ValueSet/encounter-discharge-disposition\nhttp://terminology.hl7.org/CodeSystem/discharge-disposition\nhttp://terminology.kemkes.go.id/CodeSystem/encounter-discharge-disposition", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.hospitalization.dischargeDisposition", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10319: {Path: "Encounter.hospitalization.admitSource", Terminology: "http://terminology.hl7.org/CodeSystem/admit-source\nhttp://terminology.kemkes.go.id/CodeSystem/clinical-term \nhttp://terminology.kemkes.go.id/CodeSystem/organization-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.hospitalization.admitSource\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.hospitalization.admitSource.coding", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10320: {Path: "Encounter.hospitalization.dietPreference", Terminology: "http://terminology.hl7.org/CodeSystem/diet", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.hospitalization.dietPreference", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10321: {Path: "Encounter.hospitalization.reAdmission", Terminology: "http://terminology.hl7.org/CodeSystem/v2-0092", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.hospitalization.reAdmission", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10322: {Path: "Encounter.hospitalization.specialArrangement", Terminology: "http://terminology.hl7.org/CodeSystem/encounter-special-arrangements", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.hospitalization.specialArrangement", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10324: {Path: "CarePlan.created", Terminology: "Validasi format nilai timestamp dengan menggunakan format (YYYY-MM-DD HH:MM:SS UTC)", Description: "Data harus menggunakan format datetime", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10325: {Path: "Encounter.location.status", Terminology: "http://hl7.org/fhir/encounter-location-status", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.location.status", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10326: {Path: "Encounter.reasonCode", Terminology: "http://hl7.org/fhir/sid/icd-10", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Immunization.vaccineCode", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10327: {Path: "Encounter.serviceType", Terminology: "http://terminology.hl7.org/CodeSystem/service-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.serviceType", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10328: {Path: "CarePlan.description", Terminology: "\nhttp://terminology.hl7.org/CodeSystem/v3-ActPriority", Description: "Path CarePlan.description wajib dikirimkan saat mengirim resouce CarePlan\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.priority", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10329: {Path: "CarePlan.title", Terminology: "", Description: "Path CarePlan.title wajib terisi saat mengirim resource CarePlan dengan nama rencana perawatan dalam bentuk string.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10330: {Path: "CarePlan.category", Terminology: "http://terminology.kemkes.go.id/CodeSystem/care-plan-category \nhttp://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab CarePlan.category", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10331: {Path: "Observation.valueQuantity.value", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for Units of Measure (UCUM). Daftar kode satuan dapat diakses disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10332: {Path: "Observation.component.valueQuantity.value", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab 5.4 Observation.valueQuantity", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10333: {Path: "Patient.address", Terminology: "", Description: "Value pada path Patient.address wajib diisi pada pengiriman resource Patient", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10334: {Path: "Medication.ingredient.strength.numerator.value", Terminology: "http://unitsofmeasure.org", Description: "Data type pada path Medication.ingredient.strength.numerator.value adalah decimal", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10335: {Path: "Medication.ingredient.strength.denominator.value", Terminology: "http://terminology.hl7.org/CodeSystem/v3-orderableDrugForm", Description: "Data type pada path Medication.ingredient.strength.denominator.value adalah decimal", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10336: {Path: "Encounter.participant", Terminology: "non mandatory if mpdn identifier", Description: "Path Encounter.participant tidak diwajibkan apabila pengirim berasal dari system MPDN", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10337: {Path: "MedicationRequest.authoredOn", Terminology: "datetime\nnot null", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00\nData harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10338: {Path: "MedicationRequest.dispenseRequest.validityPeriod.start", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10339: {Path: "MedicationRequest.dispenseRequest.validityPeriod.end", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10340: {Path: "MedicationRequest.dosageInstruction.sequence", Terminology: "data numerik, mulai dari 0", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10341: {Path: "MedicationRequest.dosageInstruction.timing.repeat.frequency", Terminology: "data numerik, mulai dari 0", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10342: {Path: "MedicationRequest.dosageInstruction.timing.repeat.period", Terminology: "data numerik, mulai dari 0", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10343: {Path: "MedicationRequest.dosageInstruction.doseAndRate.doseQuantity.value", Terminology: "data numerik, mulai dari 0", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10344: {Path: "MedicationRequest.dispenseRequest.dispenseInterval.value", Terminology: "data numerik, mulai dari 0", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10345: {Path: "MedicationRequest.dispenseRequest.numberOfRepeatsAllowed", Terminology: "data numerik, mulai dari 0", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10346: {Path: "MedicationRequest.dispenseRequest.quantity.value", Terminology: "data numerik, mulai dari 0", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10347: {Path: "MedicationRequest.dispenseRequest.expectedSupplyDuration.value", Terminology: "data numerik, mulai dari 0", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10348: {Path: "MedicationRequest.dispenseRequest.quantity.code", Terminology: "http://snomed.info/sct\nhttp://terminology.hl7.org/CodeSystem/v3-orderableDrugForm\nhttp://unitsofmeasure.org", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.dispenseRequest.quantity", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10349: {Path: "MedicationRequest.dosageInstruction.doseAndRate.doseQuantity.code", Terminology: "http://terminology.hl7.org/CodeSystem/v3-orderableDrugForm dan http://unitsofmeasure.org", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationRequest.dosageInstruction[i].doseAndRate.doseQuantity", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10350: {Path: "MedicationDispense.daysSupply.value", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for\nUnits of Measure (UCUM). Daftar kode satuan dapat diakses disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10351: {Path: "MedicationDispense.whenPrepared", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10352: {Path: "MedicationDispense.whenHandedOver", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10353: {Path: "MedicationDispense.dosageInstruction.sequence", Terminology: "data numerik, contoh: 1, 60, tidak bisa (-)", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10354: {Path: "MedicationDispense.dosageInstruction.timing.repeat.frequency", Terminology: "data numerik, contoh: 1, 60, tidak bisa (-)", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10355: {Path: "MedicationDispense.dosageInstruction.timing.repeat.period", Terminology: "data numerik, contoh: 1, 60, tidak bisa (-)", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10356: {Path: "MedicationDispense.dosageInstruction.doseAndRate.doseQuantity.value", Terminology: "data numerik, contoh: 1, 60, tidak bisa (-)", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10357: {Path: "MedicationDispense.dosageInstruction.doseAndRate.doseQuantity.code", Terminology: "http://terminology.hl7.org/CodeSystem/v3-orderableDrugForm\nhttp://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationDispense.dosageInstruction[i].doseAndRate.doseQuantity\nNilai pada elemen MedicationDispense.dosageInstruction.doseAndRate.doseQuantity.code harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10358: {Path: "AllergyIntolerance.recordedDate", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10359: {Path: "DiagnosticReport.issued", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10360: {Path: "ServiceRequest.occurrenceDateTime", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10361: {Path: "ServiceRequest.authoredOn", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10362: {Path: "Specimen.collection.collectedDateTime", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10363: {Path: "Specimen.collection.fastingStatusCodeableConcept.coding", Terminology: "http://terminology.hl7.org/CodeSystem/v2-0916", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Specimen.collection.fastingStatusCodeableConcept", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10364: {Path: "Specimen.collection.quantity.value", Terminology: "data numerik, contoh: 1, 60,dst. tidak bisa (-)", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10365: {Path: "Specimen.processing.timeDateTime", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10366: {Path: "Specimen.receivedTime", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10367: {Path: "Specimen.extension.valueDateTime", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10368: {Path: "Observation.referenceRange.high.value", Terminology: "data numerik, contoh: 1, 60,dst. tidak bisa (-)", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10369: {Path: "Observation.referenceRange.low.value", Terminology: "data numerik, contoh: 1, 60,dst. tidak bisa (-)", Description: "Data input harus numerik dan tidak bisa diisi \"-\"", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10370: {Path: "QuestionnaireResponse.subject", Terminology: "\"Patient/{patient-ihs-number}\"", Description: "Path QuestionnaireResponse.subject wajib terisi dengan Reference yang merujuk pada data Patient dan {patient-ihs-number} adalah nomor ID Patient yang didapatkan dari Master Patient Index", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10371: {Path: "QuestionnaireResponse.author", Terminology: "(Device | Practitioner | PractitionerRole | Patient | RelatedPerson | Organization)", Description: "Path QuestionnaireResponse.author wajib terisi dan merujuk pada resource (Device | Practitioner | PractitionerRole | Patient | RelatedPerson | Organization) yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10372: {Path: "QuestionnaireResponse.source", Terminology: "(Patient | Practitioner | PractitionerRole | RelatedPerson)", Description: "Path QuestionnaireResponse.source wajib diisi dan merujuk pada resource (Patient | Practitioner | PractitionerRole | RelatedPerson) yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10373: {Path: "QuestionnaireResponse.authored", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10374: {Path: "AllergyIntolerance.encounter", Terminology: "Encounter/{id-resource-Encounter}\"", Description: "Path AllergyIntolerance.encounter wajib diisi dan merujuk pada data Encounter dan {id-resource-Encounter} adalah nomor ID Encounter yang didapatkan dari server ketika melakukan POST data encounter", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10375: {Path: "AllergyIntolerance.recorder", Terminology: "Practitioner/{practitioner-ihs-number}", Description: "Path AllergyIntolerance.recorder wajib terisi dengan Reference yang merujuk pada data Practitioner dan {practitioner-ihs-number} adalah nomor ID tenaga kesehatan yang didapatkan dari Master Nakes Indeks.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10376: {Path: "ServiceRequest.requester", Terminology: "Practitioner/{practitioner-ihs-number}", Description: "Path ServiceRequest.performer  diisi dengan Reference merujuk pada data Practitioner dengan {practitioner-ihs-number} adalah ID Nakes organisasi induk yang didapatkan dari master Nakes indeks.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10377: {Path: "ServiceRequest.performer", Terminology: "Practitioner/{practitioner-ihs-number}", Description: "Path ServiceRequest.performer  diisi dengan Reference merujuk pada data Practitioner dengan {practitioner-ihs-number} adalah ID Nakes organisasi induk yang didapatkan dari master Nakes indeks.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10378: {Path: "QuestionnaireResponse.encounter", Terminology: "Encounter", Description: "Path QuestionnaireResponse.encounter wajib terisi dengan Reference yang merujuk pada data Encounter dan {id-resource-Encounter} adalah nomor ID Encounter yang didapatkan dari server ketika melakukan POST data encounter", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10379: {Path: "DiagnosticReport.effectiveDateTime", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10380: {Path: "CarePlan.subject", Terminology: "\"Patient/{patient-ihs-number}\"\nnot null", Description: "Path CarePlan.subject wajib terisi dengan Reference yang merujuk pada data Patient dan {patient-ihs-number} adalah nomor ID Patient yang didapatkan dari Master Patient Index\nData harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10381: {Path: "CarePlan.encounter", Terminology: "\"Encounter/{id-resource-Encounter}\"\ncoding http://unitsofmeasure.org", Description: "Path CarePlan.encounter wajib terisi dengan Reference yang merujuk pada data Encounter dan {id-resource-Encounter} adalah nomor ID Encounter yang didapatkan dari server ketika melakukan POST data encounter\nData Observation.referenceRange.high.code harus menggunakan satuan yang tersedia pada codesystem dari http://unitsofmeasure.org", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10382: {Path: "CarePlan.author", Terminology: "\"Practitioner/{practitioner-ihs-number}\"\ncoding http://unitsofmeasure.org", Description: "Path CarePlan.author wajib terisi dengan Reference yang merujuk pada data Practitioner dan {practitioner-ihs-number} adalah nomor ID tenaga kesehatan yang didapatkan dari Master Nakes Indeks.\nData Observation.referenceRange.low.code harus menggunakan satuan yang tersedia pada codesystem dari http://unitsofmeasure.org", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10383: {Path: "Observation.performer", Terminology: "Practitioner", Description: "Path Observation.performer wajib terisi dengan Reference yang merujuk pada data Practitioner dan {practitioner-ihs-number} adalah nomor ID tenaga kesehatan yang didapatkan dari Master Nakes Indeks.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10384: {Path: "DiagnosticReport.performer", Terminology: "Practitioner", Description: "Path DiagnosticReport.performer wajib dikirim dan merujuk pada resource Practitioner yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10385: {Path: "DiagnosticReport.result", Terminology: "Observation/{{observation-id}}", Description: "Data DiagnosticReport.result wajib terisi dengan Reference merujuk pada data Observation dan {{observation-id}} adalah nomor ID Observation yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10386: {Path: "DiagnosticReport.specimen", Terminology: "Specimen/{{specimen-id}}", Description: "Data DiagnosticReport.specimen wajib terisi dengan Reference merujuk pada data Specimen dan {{specimen-id}} adalah ID specimen yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10387: {Path: "DiagnosticReport.basedOn", Terminology: "ServiceRequest/{id-resource-ServiceRequest}", Description: "Data DiagnosticReport.basedOn wajib dikirim dan Reference merujuk pada data ServiceRequest dengan {id-resource-ServiceRequest} adalah ID data yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10388: {Path: "MedicationRequest.identifier", Terminology: "not null\nhttp://sys-ids.kemkes.go.id/prescription-item/{{org_id}}", Description: "Data harus terisi (tidak boleh kosong)\nData harus menggunakan codesystem http://sys-ids.kemkes.go.id/prescription-item/{{org_id}} \ndengan {{org_id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10389: {Path: "MedicationDispense.identifier", Terminology: "not null\nhttp://terminology.kemkes.go.id/CodeSystem/coverage-type\nhttp://sys-ids.kemkes.go.id/claim-number/{{org_id}}\nhttp://sys-ids.kemkes.go.id/prescription/{{org_id}}\nhttp://sys-ids.kemkes.go.id/prescription-item/{{org_id}}", Description: "Data harus terisi (tidak boleh kosong)\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationDispense.identifier\nData harus menggunakan codesystem http://sys-ids.kemkes.go.id/claim-number/{{org_id}} dengan {{org_id}} adalah Organization ID faskes\nData harus menggunakan codesystem http://sys-ids.kemkes.go.id/prescription/{{org_id}} adalah Organization ID faskes\nData harus menggunakan codesystem http://sys-ids.kemkes.go.id/prescription-item/{{org_id}} dengan {{org_id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10390: {Path: "MedicationDispense.context", Terminology: "Encounter", Description: "Path MedicationDispense.context wajib diisi dan merujuk pada resource Encounter yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10391: {Path: "Medication.manufacturer", Terminology: "Organization/{{org-id}}", Description: "Data Medication.manufacturer diisi dengan Reference merujuk pada data yang tersimpan di resource Organization, yang menyimpan data pabrik obat.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10392: {Path: "MedicationDispense.performer", Terminology: "Practitioner/{practitioner-ihs-number}", Description: "Path MedicationDispense.performer diisi dengan Reference yang merujuk pada data Practitioner dan {practitioner-ihs-number} adalah nomor ID tenaga kesehatan yang didapatkan dari Master Nakes Indeks.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10393: {Path: "MedicationDispense.location", Terminology: "Location/{{location-id}}", Description: "Data MedicationDispense.location diisi dengan Reference merujuk pada data Location dengan Location/{{location-id}} adalah ID lokasi yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10394: {Path: "MedicationDispense.authorizingPrescription", Terminology: "MedicationRequest", Description: "Path MedicationDispense.authorizingPrescription diisi dengan reference yang merujuk pada data MedicationRequest dengan nomor ID data medicationrequest yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10395: {Path: "Specimen.identifier.assigner", Terminology: "Organization", Description: "Path Specimen.identifier.assigner diisi dengan reference yang merujuk pada data Organization dan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10396: {Path: "Specimen.collection.collector", Terminology: "(Practitioner | PractitionerRole)", Description: "Path Specimen.collection.collector diisi dengan reference yang merujuk pada resource (Practitioner | PractitionerRole) yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10397: {Path: "Patient.birthDate", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10398: {Path: "Observation.interpretation", Terminology: "http://terminology.kemkes.go.id/CodeSystem/examination\nhttp://snomed.info/sct\nhttp://terminology.hl7.org/CodeSystem/v3-ObservationInterpretation\nhttp://loinc.org", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.interpretation\nNilai pada elemen Observation.interpretation.coding.code harus diisi menggunakan kode yang valid dari sistem terminologi http://loinc.org.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10399: {Path: "Procedure.performer", Terminology: "Practitioner", Description: "Path Procedure.performer diisi dengan reference merujuk pada data Practitioner dan {practitioner-ihs-number} adalah nomor ID tenaga kesehatan yang didapatkan dari Master Nakes Indeks.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10400: {Path: "Procedure.performedPeriod.start", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10401: {Path: "Procedure.performedPeriod.end", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10402: {Path: "Condition.onsetDateTime", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10403: {Path: "Condition.recordedDate", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10404: {Path: "EpisodeOfCare.statusHistory.period.start", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10405: {Path: "EpisodeOfCare.statusHistory.period.end", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10406: {Path: "EpisodeOfCare.period.start", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10407: {Path: "EpisodeOfCare.period.end", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10408: {Path: "EpisodeOfCare.managingOrganization", Terminology: "Organization/{{org-id}}", Description: "Data EpisodeOfCare.managingOrganization diisi dengan Reference merujuk pada data Organization dan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10409: {Path: "Condition.bodySite", Terminology: "http://terminology.kemkes.go.id/CodeSystem/clinical-term", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Condition.bodySite atau akses Lampiran Terminology bodysite disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10410: {Path: "Condition.evidence.code", Terminology: "http://hl7.org/fhir/sid/icd-10\nhttp://terminology.kemkes.go.id/CodeSystem/clinical-term\nhttp://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT sub-bab Condition.evidence.code", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10411: {Path: "Condition.stage.summary", Terminology: "http://terminology.kemkes.go.id/CodeSystem/clinical-term\nhttp://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Condition.stage.summary", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10412: {Path: "Condition.stage.type", Terminology: "http://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Condition.stage.summary", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10413: {Path: "Patient.gender", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10414: {Path: "Specimen.identifier.system", Terminology: "http://sys-ids.kemkes.go.id/sitb/{{org-id}}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/sitb/{{org-id}}  dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10415: {Path: "ServiceRequest.identifier", Terminology: "http://sys-ids.kemkes.go.id/servicerequest/{{organization-ihs-number}}\nhttp://sys-ids.kemkes.go.id/sitb/{{org-id}}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/servicerequest/{{organization-ihs-number}} dengan {{organization-ihs-number}} adalah Organization ID faskes\nData harus menggunakan codesystem http://sys-ids.kemkes.go.id/sitb/{{org-id}}  dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10416: {Path: "Observation.method", Terminology: "http://terminology.kemkes.go.id/CodeSystem/clinical-term", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.method", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10417: {Path: "ClinicalImpression.effectiveDateTime", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10418: {Path: "ClinicalImpression.identifier", Terminology: "http://sys-ids.kemkes.go.id/clinicalimpression/{{org_id}}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/clinicalimpression/{{org_id}}  dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10419: {Path: "ClinicalImpression.date", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10420: {Path: "ClinicalImpression.assessor", Terminology: "Reference to Practitioner | PractitionerRole", Description: "Path ClinicalImpression.assessor diisi dengan Reference yang merujuk pada data Practitioner dan {practitioner-ihs-number} adalah nomor ID tenaga kesehatan yang didapatkan dari Master Nakes Indeks atau merujuk pada PractitionerRole", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10421: {Path: "ClinicalImpression.problem", Terminology: "Reference to Condition | Allergy Intolerance", Description: "Data ClinicalImpression.problem dengan reference yang merujuk pada data resource (Condition | Allergy Intolerance) dengan nomor ID yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10422: {Path: "ClinicalImpression.investigation.item", Terminology: "Reference to Observation | QuestionnaireResponse | FamilyMemberHistory | DiagnosticReport | RiskAssessment | ImagingStudy | Media", Description: "Path ClinicalImpression.investigation.item diisi dengan reference yang merujuk pada resource (Observation | QuestionnaireResponse | FamilyMemberHistory | DiagnosticReport | RiskAssessment | ImagingStudy | Media) yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10423: {Path: "ClinicalImpression.finding.itemReference", Terminology: "Reference to Condition | Observation | Media", Description: "Path ClinicalImpression.finding.itemReference diisi dengan reference yang merujuk pada resource (Condition | Observation | Media) yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10424: {Path: "ImagingStudy.started", Terminology: "dateTime", Description: "Data harus menggunakan format datetime", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10425: {Path: "Procedure.basedOn", Terminology: "(CarePlan | ServiceRequest)", Description: "Path Procedure.basedOn diisi dengan reference merujuk pada data (CarePlan | ServiceRequest) dengan nomor ID data terkait yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10426: {Path: "Immunization.fundingSource", Terminology: "http://terminology.hl7.org/CodeSystem/immunization-funding-source\nhttp://terminology.kemkes.go.id/CodeSystem/immunization-fundingSource", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Immunization.fundingSource", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10427: {Path: "ServiceRequest.orderDetail", Terminology: "http://dicom.nema.org/resources/ontology/DCM\nhttp://sys-ids.kemkes.go.id/ae-title\nhttp://snomed.info/sct", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ServiceRequest.orderDetail", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10428: {Path: "Observation.derivedFrom", Terminology: "Reference to DocumentReference | ImagingStudy | Media | QuestionnaireResponse | Observation | MolecularSequence", Description: "Path Observation.derivedFrom diisi dengan Reference yang merujuk pada data DocumentReference | ImagingStudy | Media | QuestionnaireResponse | Observation | MolecularSequence dengan mengacu pada id data terkait", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10429: {Path: "Observation.valueDateTime", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10430: {Path: "Encounter.episodeOfCare", Terminology: "Reference EpisodeOfCare/{{EpisodeOfCare_uuid}}", Description: "Path Encounter.episodeOfCare diisi dengan Reference yang merujuk pada data EpisodeOfCare dan {{EpisodeOfCare_uuid}} adalah nomor ID Episode of Care yang didapatkan dari server ketika melakukan POST data Episode of Care sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10431: {Path: "ServiceRequest.locationReference", Terminology: "Location/{{location-id}}", Description: "Path ServiceRequest.locationReference diisi dengan Reference merujuk pada data Location dengan Location/{{location-id}} adalah ID lokasi yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10432: {Path: "DiagnosticReport.identifier", Terminology: "http://sys-ids.kemkes.go.id/diagnostic/{{organization-ihs-number}}/lab", Description: "- Data Diagnostic Report Laboratorium harus menggunakan codesystem http://sys-ids.kemkes.go.id/diagnostic/{{organization-ihs-number}}/lab dengan /{{organization-ihs-number}} adalah Organization ID faskes\n- Data Diagnostic Report Radiologi harus menggunakan codesystem http://sys-ids.kemkes.go.id/diagnostic/{{organization-ihs-number}}/rad dengan /{{organization-ihs-number}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10433: {Path: "MedicationStatement.medicationCodeableConcept", Terminology: "http://sys-ids.kemkes.go.id/kfa\nhttp://snomed.info/sct", Description: "Data harus menggunakan Kode MedicationStatement.medicationCodeableConcept dapat diisi dengan kode KFA sesuai dengan use case masing-masing. Untuk keterangan lebih lanjut dapat melihat dalam playbook masing-masing usecase.\nNilai pada elemen MedicationStatement.medicationCodeableConcept harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10434: {Path: "Observation.component.code", Terminology: "http://snomed.info/sct\nhttp://terminology.kemkes.go.id/CodeSystem/examination", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.component.code\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.component.code.coding", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10435: {Path: "Observation.component.valueCodeableConcept", Terminology: "http://terminology.kemkes.go.id\nhttp://terminology.kemkes.go.id/CodeSystem/icd-10-mm\nhttp://terminology.kemkes.go.id/CodeSystem/icd-10-pm\nhttp://snomed.info/sct\nhttp://terminology.kemkes.go.id/CodeSystem/examination\nhttp://terminology.kemkes.go.id/CodeSystem/clinical-term", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Observation.component.valueCodeableConcept", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10436: {Path: "MedicationStatement.category", Terminology: "http://terminology.hl7.org/CodeSystem/medication-statement-category", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationStatement.category", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10437: {Path: "MedicationStatement.dosage.timing.periodUnit", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for Units of Measure (UCUM) untuk satuan waktu.", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10438: {Path: "MedicationStatement.context", Terminology: "Encounter/{{encounter-id}}", Description: "Path MedicationStatement.context wajib terisi dengan reference merujuk pada data Encounter dengan {{encounter-id}} adalah nomor ID encounter yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10439: {Path: "MedicationStatement.informationSource", Terminology: "Reference ke resource Patient | Practitioner | PractitionerRole | RelatedPerson | Organization.", Description: "Path MedicationStatement.informationSource wajib terisi dan merujuk pada resource (Patient | Practitioner | PractitionerRole | RelatedPerson | Organization) yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10440: {Path: "Specimen.request", Terminology: "ServiceRequest", Description: "Path Specimen.request wajib dikirim dan merujuk pada resource ServiceRequest yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10441: {Path: "Observation.identifier", Terminology: "http://sys-ids.kemkes.go.id/observation/{{org-id}}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/observation/{{org-id}} dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10442: {Path: "Procedure.identifier", Terminology: "http://sys-ids.kemkes.go.id/procedure/{{org-id}}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/procedure/{{org-id}} dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10443: {Path: "AllergyIntolerance.identifier", Terminology: "http://sys-ids.kemkes.go.id/allergy/{{org-id}}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/allergy/{{org-id}} dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10444: {Path: "Condition.identifier", Terminology: "http://sys-ids.kemkes.go.id/condition/{{org-id}}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/condition/{{org-id}}  dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10445: {Path: "MedicationStatement.identifier", Terminology: "http://sys-ids.kemkes.go.id/medicationstatement/{{org-id}}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/medicationstatement/{{org-id}}  dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10446: {Path: "Organization.identifier", Terminology: "http://sys-ids.kemkes.go.id/organization/{{org-id}}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/organization/{{org-id}} dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10447: {Path: "Location.identifier", Terminology: "http://sys-ids.kemkes.go.id/location/{{org-id}}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/location/{{org-id}}  dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10450: {Path: "Immunization.protocolApplied", Terminology: "", Description: "Data harus menggunakan codesystem yang sesuai, lihat penggunaan path secara secara lengkap disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10451: {Path: "Immunization.recorded", Terminology: "", Description: "Data harus terisi menggunakan codesystem yang sesuai, lihat penggunaan path secara secara lengkap disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10453: {Path: "MedicationRequest.dispenseRequest.performer", Terminology: "Organization/{{org-id}}", Description: "Data MedicationRequest.dispenseRequest.performer wajib terisi dengan Reference yang merujuk pada data Organization dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10454: {Path: "MedicationRequest.encounter", Terminology: "\"Encounter/{Id-resource-Encounter}\"", Description: "Path MedicationRequest.encounter wajib diisi dan merujuk pada resource Encounter yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10455: {Path: "MedicationRequest.requester", Terminology: "(Practitioner | PractitionerRole | Organization | Patient | Device | RelatedPerson | CareTeam)", Description: "Path MedicationRequest.requester wajib terisi dan merujuk pada nomor ID resoucer (Practitioner | PractitionerRole | Organization | Patient | Device | RelatedPerson | CareTeam) yang sudah dikirim sebelumnya", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10456: {Path: "Immunization.protocolApplied.doseNumberPositiveInt", Terminology: "Integer must be more than 0", Description: "Data wajib terisi menggunakan codesystem yang sesuai, lihat penggunaan path secara secara lengkap disini", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10458: {Path: "EpisodeOfCare.identifier", Terminology: "http://sys-ids.kemkes.go.id/episode-of-care/{{org_id}}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/episode-of-care/{{org_id}}  dengan {{org-id}} adalah Organization ID faskes", Fix: "", Version: "v1.0", LastUpdate: "14 Maret 2024"},
	10462: {Path: "ServiceRequest.category", Terminology: "http://terminology.kemkes.go.id", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ServiceRequest.category", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10464: {Path: "Composition.identifier", Terminology: "http://sys-ids.kemkes.go.id/composition/{orgID}", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/composition/{orgID} dengan {orgID} adalah Organization ID faskes", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10502: {Path: "DocumentReference.date", Terminology: "not null", Description: "Data harus terisi (tidak boleh kosong)", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10513: {Path: "Invoice.lineItem.priceComponent.code", Terminology: "http://terminology.kemkes.go.id/CodeSystem/kptl\nhttp://sys-ids.kemkes.go.id/kfa", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Invoice.lineItem.priceComponent.code\nData harus menggunakan Kode Invoice.lineItem.priceComponent.code dapat diisi dengan kode KFA sesuai dengan use case masing-masing. Untuk keterangan lebih lanjut dapat melihat dalam playbook masing-masing usecase.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10514: {Path: "ClaimResponse.identifier", Terminology: "http://sys-ids.kemkes.go.id/claim-batch-number/{org_id}\nhttp://sys-ids.kemkes.go.id/claim-number/{org_id} --> org id Asuransi", Description: "Data harus menggunakan codesystem http://sys-ids.kemkes.go.id/claim-batch-number/{org_id} dengan {org_id} adalah Organization ID Fasyankes.\nData harus menggunakan codesystem http://sys-ids.kemkes.go.id/claim-number/{org_id} dengan {org_id} adalah Organization ID Asuransi.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10615: {Path: "Procedure.usedReference", Terminology: "Reference ke Medication", Description: "", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10619: {Path: "Claim.item.productOrService", Terminology: "http://terminology.kemkes.go.id/CodeSystem/kptl\nhttp://sys-ids.kemkes.go.id/kfa", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab  Claim.item.productOrService h\nData harus menggunakan Kode Claim.item.productOrService dapat diisi dengan kode KFA sesuai dengan use case masing-masing. Untuk keterangan lebih lanjut dapat melihat dalam playbook masing-masing usecase.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10631: {Path: "PaymentNotice.status", Terminology: "http://hl7.org/fhir/fm-status", Description: "Nilai pada elemen PaymentNotice.status harus sesuai dengan sistem http://hl7.org/fhir/fm-status.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10632: {Path: "PaymentNotice.created", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:03", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10633: {Path: "PaymentNotice.paymentDate", Terminology: "date", Description: "Data harus menggunakan format date YYYY-MM-DD", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10634: {Path: "PaymentNotice.amount.currency", Terminology: "urn:iso:std:iso:4217\ndecimal", Description: "Nilai pada elemen PaymentNotice.amount.currency dapat diisi dengan kode 'IDR' sesuai standar ISO 4217.\nNilai pada elemen PaymentNotice.amount.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10635: {Path: "PaymentNotice.payment", Terminology: "Reference ke PaymentReconciliation", Description: "Referensi pada elemen PaymentNotice.payment harus mengarah ke resource PaymentReconciliation.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10636: {Path: "PaymentNotice.paymentStatus", Terminology: "http://terminology.hl7.org/CodeSystem/payment-status", Description: "Nilai pada elemen PaymentNotice.paymentStatus harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/payment-status.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10637: {Path: "PaymentReconciliation.status", Terminology: "http://hl7.org/fhir/fm-status", Description: "Nilai pada elemen PaymentReconciliation.status harus sesuai dengan sistem http://hl7.org/fhir/fm-status.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10638: {Path: "PaymentReconciliation.period.start", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:00", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10639: {Path: "PaymentReconciliation.period.end", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:01", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10640: {Path: "PaymentReconciliation.created", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:02", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10641: {Path: "PaymentReconciliation.outcome", Terminology: "http://hl7.org/fhir/remittance-outcome", Description: "Nilai pada elemen PaymentReconciliation.outcome harus sesuai dengan sistem http://hl7.org/fhir/remittance-outcome.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10642: {Path: "PaymentReconciliation.disposition", Terminology: "string", Description: "Nilai pada elemen PaymentReconciliation.disposition harus berupa tipe data string.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10643: {Path: "PaymentReconciliation.paymentDate", Terminology: "date", Description: "Data harus menggunakan format date YYYY-MM-DD", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10644: {Path: "PaymentReconciliation.paymentAmount.currency", Terminology: "urn:iso:std:iso:4217\ndecimal", Description: "Nilai pada elemen PaymentReconciliation.paymentAmount.currency dapat diisi dengan kode 'IDR' sesuai standar ISO 4217.\nNilai pada elemen PaymentReconciliation.paymentAmount.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10645: {Path: "PaymentReconciliation.processNote.type", Terminology: "http://hl7.org/fhir/note-type", Description: "Nilai pada elemen PaymentReconciliation.processNote.type harus sesuai dengan sistem http://hl7.org/fhir/note-type.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10646: {Path: "PaymentReconciliation.processNote.text", Terminology: "string", Description: "Nilai pada elemen PaymentReconciliation.processNote.text harus berupa tipe data string.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10652: {Path: "PaymentReconciliation.detail.type", Terminology: "http://terminology.hl7.org/CodeSystem/payment-type", Description: "Nilai pada elemen PaymentReconciliation.detail.type harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/payment-type.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10653: {Path: "PaymentReconciliation.detail.request", Terminology: "Reference ke Claim dan Invoice", Description: "Referensi pada elemen PaymentReconciliation.detail.request harus mengarah ke resource Claim dan Invoice.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10654: {Path: "PaymentReconciliation.detail.response", Terminology: "Reference ke ClaimResponse", Description: "Referensi pada elemen PaymentReconciliation.detail.response harus mengarah ke resource ClaimResponse.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10655: {Path: "PaymentNotice.request", Terminology: "Reference ke Claim dan Invoice", Description: "Referensi pada elemen PaymentNotice.request harus mengarah ke resource Claim dan Invoice.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10656: {Path: "PaymentNotice.response", Terminology: "Reference ke ClaimResponse", Description: "Referensi pada elemen PaymentNotice.response harus mengarah ke resource ClaimResponse.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10657: {Path: "PaymentReconciliation.paymentIssuer", Terminology: "Reference ke Organization", Description: "Referensi pada elemen PaymentReconciliation.paymentIssuer harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10658: {Path: "PaymentReconciliation.requestor", Terminology: "Reference ke Organization", Description: "Referensi pada elemen PaymentReconciliation.requestor harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10659: {Path: "PaymentNotice.payee", Terminology: "Reference ke Organization", Description: "Referensi pada elemen PaymentNotice.payee harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10660: {Path: "PaymentNotice.provider", Terminology: "Reference ke Organization", Description: "Referensi pada elemen PaymentNotice.provider harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10661: {Path: "PaymentNotice.recipient", Terminology: "Reference ke Organization", Description: "Referensi pada elemen PaymentNotice.recipient harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10663: {Path: "Composition.attester.party", Terminology: "Reference ke Practitioner", Description: "Referensi pada elemen Composition.attester.party harus mengarah ke resource Practitioner.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10665: {Path: "QuestionnaireResponse.item.answer.valueCoding", Terminology: "http://terminology.kemkes.go.id/CodeSystem/clinical-term\nhttp://terminology.hl7.org/CodeSystem/v3-ActCode\nhttp://terminology.kemkes.go.id/CodeSystem/clinical-speciality\nhttp://terminology.kemkes.go.id/CodeSystem/organization-type", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab QuestionnaireResponse.item.answer.valueCoding\nNilai pada elemen QuestionnaireResponse.item.answer.valueCoding harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/v3-ActCode.\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab QuestionnaireResponse.item.X.answer.valueCoding", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10666: {Path: "Composition.section.focus", Terminology: "Reference ke Patient\nReference to any", Description: "Referensi pada elemen Composition.section.focus harus mengarah ke resource Patient.\nElemen Composition.section.focus dapat mereferensikan resource FHIR yang valid.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10669: {Path: "Observation.valueTime", Terminology: "Time", Description: "Format waktu pada elemen Observation.valueTime harus berupa time (HH:mm:ss).", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10670: {Path: "Composition.event.code", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen Composition.event.code harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10671: {Path: "Composition.event.period.start", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:04", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10672: {Path: "Composition.event.period.end", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:05", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10673: {Path: "Composition.event.detail.reference", Terminology: "Reference ke Practitioner", Description: "Referensi pada elemen Composition.event.detail.reference harus mengarah ke resource Practitioner.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10674: {Path: "MedicationDispense.substitution.reason.coding", Terminology: "http://terminology.hl7.org/CodeSystem/v3-ActReason\nhttp://terminology.kemkes.go.id", Description: "Data harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationDispense.substitution.reason.coding", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10675: {Path: "MedicationDispense.substitution.type.coding", Terminology: "http://terminology.hl7.org/CodeSystem/v3-substanceAdminSubstitution", Description: "Nilai pada elemen MedicationDispense.substitution.type.coding harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/v3-substanceAdminSubstitution.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10676: {Path: "MedicationDispense.substitution.responsibleParty", Terminology: "Reference ke Practitioner", Description: "Referensi pada elemen MedicationDispense.substitution.responsibleParty harus mengarah ke resource Practitioner.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10677: {Path: "Composition.section.entry", Terminology: "Reference ke Location\nReference ke Observation\nReference ke QuestionnaireResponse\nReference Any", Description: "Referensi pada elemen Composition.section.entry harus mengarah ke resource Location.\nReferensi pada elemen Composition.section.entry harus mengarah ke resource Observation.\nReferensi pada elemen Composition.section.entry harus mengarah ke resource QuestionnaireResponse.\nElemen Composition.section.entry dapat mereferensikan resource FHIR yang valid.\nElemen Composition.section.X.entry dapat mereferensikan resource FHIR yang valid.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10678: {Path: "Condition.recorder", Terminology: "Reference ke Practitioner", Description: "Referensi pada elemen Condition.recorder harus mengarah ke resource Practitioner.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10679: {Path: "Specimen.container.type", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen Specimen.container.type harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10680: {Path: "Specimen.collection.bodySite", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen Specimen.collection.bodySite harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10681: {Path: "ChargeItem.performer.function", Terminology: "http://snomed.info/sct\nhttp://terminology.kemkes.go.id", Description: "Nilai pada elemen ChargeItem.performer.function harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ChargeItem.performer.function", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10682: {Path: "VisionPrescription.status", Terminology: "http://hl7.org/fhir/fm-status", Description: "Nilai pada elemen VisionPrescription.status harus sesuai dengan sistem http://hl7.org/fhir/fm-status.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10683: {Path: "VisionPrescription.dateTime", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:06", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10684: {Path: "VisionPrescription.patient", Terminology: "Reference ke Patient", Description: "Referensi pada elemen VisionPrescription.patient harus mengarah ke resource Patient.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10685: {Path: "VisionPrescription.encounter", Terminology: "Reference ke Encounter dan not null", Description: "Referensi pada elemen VisionPrescription.encounter harus mengarah ke resource Encounter dan not null.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10686: {Path: "VisionPrescription.dateWritten", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:07", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10687: {Path: "VisionPrescription.prescriber", Terminology: "Reference ke Practitioner", Description: "Referensi pada elemen VisionPrescription.prescriber harus mengarah ke resource Practitioner.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10688: {Path: "VisionPrescription.lensSpecification.product", Terminology: "http://terminology.hl7.org/CodeSystem/ex-visionprescriptionproduct", Description: "Nilai pada elemen VisionPrescription.lensSpecification.product harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/ex-visionprescriptionproduct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10689: {Path: "VisionPrescription.lensSpecification.eye", Terminology: "http://hl7.org/fhir/vision-eye-codes", Description: "Nilai pada elemen VisionPrescription.lensSpecification.eye harus sesuai dengan sistem http://hl7.org/fhir/vision-eye-codes.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10690: {Path: "VisionPrescription.lensSpecification.sphere", Terminology: "decimal", Description: "Nilai pada elemen VisionPrescription.lensSpecification.sphere harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10691: {Path: "VisionPrescription.lensSpecification.cylinder", Terminology: "decimal", Description: "Nilai pada elemen VisionPrescription.lensSpecification.cylinder harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10692: {Path: "VisionPrescription.lensSpecification.axis", Terminology: "integer", Description: "Nilai pada elemen VisionPrescription.lensSpecification.axis harus berupa tipe data integer.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10693: {Path: "VisionPrescription.lensSpecification.prism.amount", Terminology: "decimal", Description: "Nilai pada elemen VisionPrescription.lensSpecification.prism.amount harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10694: {Path: "VisionPrescription.lensSpecification.prism.base", Terminology: "http://hl7.org/fhir/vision-base-codes", Description: "Nilai pada elemen VisionPrescription.lensSpecification.prism.base harus sesuai dengan sistem http://hl7.org/fhir/vision-base-codes.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10695: {Path: "FamilyMemberHistory.condition.onsetAge", Terminology: "Age", Description: "Nilai pada elemen FamilyMemberHistory.condition.onsetAge harus berupa tipe data Age.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10696: {Path: "FamilyMemberHistory.condition.onsetPeriod.start", Terminology: "period", Description: "Format tanggal dan waktu pada elemen FamilyMemberHistory.condition.onsetPeriod.start harus berupa dateTime (YYYY-MM-DDThh:mm:ss+zz:zz).", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10697: {Path: "FamilyMemberHistory.condition.onsetPeriod.end", Terminology: "period", Description: "Format tanggal dan waktu pada elemen FamilyMemberHistory.condition.onsetPeriod.end harus berupa dateTime (YYYY-MM-DDThh:mm:ss+zz:zz).", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10698: {Path: "FamilyMemberHistory.condition.onsetRange.low.value", Terminology: "decimal", Description: "Nilai pada elemen FamilyMemberHistory.condition.onsetRange.low.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10699: {Path: "FamilyMemberHistory.condition.onsetRange.high.value", Terminology: "decimal", Description: "Nilai pada elemen FamilyMemberHistory.condition.onsetRange.high.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10700: {Path: "FamilyMemberHistory.condition.onsetString", Terminology: "string", Description: "Nilai pada elemen FamilyMemberHistory.condition.onsetString harus berupa tipe data string.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10701: {Path: "FamilyMemberHistory.deceasedBoolean", Terminology: "boolean", Description: "Nilai pada elemen FamilyMemberHistory.deceasedBoolean harus berupa tipe data boolean.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10702: {Path: "FamilyMemberHistory.deceasedAge", Terminology: "age", Description: "Nilai pada elemen FamilyMemberHistory.deceasedAge harus berupa tipe data Age.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10703: {Path: "FamilyMemberHistory.deceasedRange.low.value", Terminology: "decimal", Description: "Nilai pada elemen FamilyMemberHistory.deceasedRange.low.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10704: {Path: "FamilyMemberHistory.deceasedRange.high.value", Terminology: "decimal", Description: "Nilai pada elemen FamilyMemberHistory.deceasedRange.high.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10705: {Path: "FamilyMemberHistory.deceasedDate", Terminology: "date", Description: "Data harus menggunakan format date YYYY-MM-DD", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10706: {Path: "FamilyMemberHistory.deceasedString", Terminology: "string", Description: "Nilai pada elemen FamilyMemberHistory.deceasedString harus berupa tipe data string.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10707: {Path: "FamilyMemberHistory.condition.code", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen FamilyMemberHistory.condition.code harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10708: {Path: "FamilyMemberHistory.condition.outcome", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen FamilyMemberHistory.condition.outcome harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10709: {Path: "FamilyMemberHistory.condition.contributedToDeath", Terminology: "boolean", Description: "Nilai pada elemen FamilyMemberHistory.condition.contributedToDeath harus berupa tipe data boolean.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10710: {Path: "RiskAssessment.status", Terminology: "http://hl7.org/fhir/observation-status", Description: "Nilai pada elemen RiskAssessment.status harus sesuai dengan sistem http://hl7.org/fhir/observation-status.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10711: {Path: "RiskAssessment.code", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen RiskAssessment.code harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10712: {Path: "RiskAssessment.reasonReference", Terminology: "Reference ke Condition |", Description: "Referensi pada elemen RiskAssessment.reasonReference harus mengarah ke resource Condition |.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10713: {Path: "RiskAssessment.prediction.outcome", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen RiskAssessment.prediction.outcome harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10714: {Path: "RiskAssessment.prediction.probabilityDecimal", Terminology: "decimal", Description: "Nilai pada elemen RiskAssessment.prediction.probabilityDecimal harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10715: {Path: "RiskAssessment.prediction.qualitativeRisk", Terminology: "http://terminology.hl7.org/CodeSystem/risk-probability", Description: "Nilai pada elemen RiskAssessment.prediction.qualitativeRisk harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/risk-probability.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10716: {Path: "RiskAssessment.prediction.relativeRisk", Terminology: "decimal", Description: "Nilai pada elemen RiskAssessment.prediction.relativeRisk harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10717: {Path: "RiskAssessment.prediction.whenPeriod.start", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:08", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10718: {Path: "RiskAssessment.prediction.whenPeriod.end", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:09", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10719: {Path: "RiskAssessment.prediction.whenRange.high.value", Terminology: "decimal", Description: "Nilai pada elemen RiskAssessment.prediction.whenRange.high.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10720: {Path: "RiskAssessment.prediction.whenRange.low.value", Terminology: "decimal", Description: "Nilai pada elemen RiskAssessment.prediction.whenRange.low.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10721: {Path: "RiskAssessment.note.authorReference", Terminology: "Reference to Practitioner | Patient | Organization | RelatedPerson", Description: "Referensi pada elemen RiskAssessment.note.authorReference harus mengarah ke resource Practitioner | Patient | Organization | RelatedPerson.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10722: {Path: "RiskAssessment.note.authorString", Terminology: "string", Description: "Nilai pada elemen RiskAssessment.note.authorString harus berupa tipe data string.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10723: {Path: "RiskAssessment.note.time", Terminology: "datetime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:10", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10724: {Path: "RiskAssessment.note.text", Terminology: "string", Description: "Nilai pada elemen RiskAssessment.note.text harus berupa tipe data string.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10725: {Path: "RiskAssessment.condition", Terminology: "Reference ke Condition", Description: "Referensi pada elemen RiskAssessment.condition harus mengarah ke resource Condition.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10726: {Path: "RiskAssessment.subject", Terminology: "Reference ke Patient |", Description: "Referensi pada elemen RiskAssessment.subject harus mengarah ke resource Patient |.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10727: {Path: "RiskAssessment.encounter", Terminology: "Reference ke Encounter", Description: "Referensi pada elemen RiskAssessment.encounter harus mengarah ke resource Encounter.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10728: {Path: "RiskAssessment.practitioner", Terminology: "Reference ke Practitioner", Description: "Referensi pada elemen RiskAssessment.practitioner harus mengarah ke resource Practitioner.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10729: {Path: "ClinicalImpression.prognosisReference", Terminology: "Reference ke RiskAssessment", Description: "Referensi pada elemen ClinicalImpression.prognosisReference harus mengarah ke resource RiskAssessment.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10730: {Path: "Goal.lifecycleStatus", Terminology: "http://hl7.org/fhir/goal-status", Description: "Nilai pada elemen Goal.lifecycleStatus harus sesuai dengan sistem http://hl7.org/fhir/goal-status.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10731: {Path: "Goal.achievementStatus", Terminology: "http://terminology.hl7.org/CodeSystem/goal-achievement", Description: "Nilai pada elemen Goal.achievementStatus harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/goal-achievement.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10732: {Path: "Goal.category", Terminology: "http://terminology.hl7.org/CodeSystem/goal-category", Description: "Nilai pada elemen Goal.category harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/goal-category.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10733: {Path: "Goal.description", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen Goal.description harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10734: {Path: "Goal.subject", Terminology: "Reference ke Patient", Description: "Referensi pada elemen Goal.subject harus mengarah ke resource Patient.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10735: {Path: "Goal.target.measure", Terminology: "http://loinc.org", Description: "Nilai pada elemen harus diisi menggunakan kode yang valid dari sistem terminologi http://loinc.org.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10736: {Path: "Goal.target.detailQuantity.value", Terminology: "decimal", Description: "Nilai pada elemen Goal.target.detailQuantity.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10737: {Path: "Goal.target.detailQuantity.code", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for\nUnits of Measure (UCUM). Daftar kode satuan dapat diakses  disini", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10738: {Path: "Goal.target.detailRange.high.value", Terminology: "decimal", Description: "Nilai pada elemen Goal.target.detailRange.high.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10739: {Path: "Goal.target.detailRange.low.value", Terminology: "decimal", Description: "Nilai pada elemen Goal.target.detailRange.low.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10740: {Path: "Goal.target.detailCodeableConcept", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen Goal.target.detailCodeableConcept harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10741: {Path: "Goal.target.detailString", Terminology: "string", Description: "Nilai pada elemen Goal.target.detailString harus berupa tipe data string.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10742: {Path: "Goal.target.detailBoolean", Terminology: "boolean", Description: "Nilai pada elemen Goal.target.detailBoolean harus berupa tipe data boolean.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10743: {Path: "Goal.target.detailInteger", Terminology: "integer", Description: "Nilai pada elemen Goal.target.detailInteger harus berupa tipe data integer.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10744: {Path: "Goal.target.detailRatio.numerator.system", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for\nUnits of Measure (UCUM). Daftar kode satuan dapat diakses  disini", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10745: {Path: "Goal.target.detailRatio.numerator.value", Terminology: "decimal", Description: "Nilai pada elemen Goal.target.detailRatio.numerator.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10746: {Path: "Goal.target.detailRatio.denominator.system", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for\nUnits of Measure (UCUM). Daftar kode satuan dapat diakses  disini", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10747: {Path: "Goal.target.detailRatio.denominator.value", Terminology: "decimal", Description: "Nilai pada elemen Goal.target.detailRatio.denominator.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10748: {Path: "Goal.target.dueDate", Terminology: "date", Description: "Data harus menggunakan format date YYYY-MM-DD", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10750: {Path: "Goal.target.dueDuration", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for\nUnits of Measure (UCUM). Daftar kode satuan dapat diakses  disini", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10751: {Path: "Goal.expressedBy", Terminology: "Reference ke Practitioner", Description: "Referensi pada elemen Goal.expressedBy harus mengarah ke resource Practitioner.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10752: {Path: "Goal.outcomeCode", Terminology: "http://snomed.info/sct\nhttp://terminology.kemkes.go.id", Description: "Nilai pada elemen Goal.outcomeCode harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Goal.outcomeCode", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10753: {Path: "Goal.outcomeReference", Terminology: "Reference ke Observation", Description: "Referensi pada elemen Goal.outcomeReference harus mengarah ke resource Observation.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10755: {Path: "MedicationAdministration.status", Terminology: "http://terminology.hl7.org/CodeSystem/medication-admin-status", Description: "Nilai pada elemen MedicationAdministration.status harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/medication-admin-status.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10756: {Path: "MedicationAdministration.partOf", Terminology: "Reference ke MedicationAdministration |", Description: "Referensi pada elemen MedicationAdministration.partOf harus mengarah ke resource MedicationAdministration |.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10757: {Path: "MedicationAdministration.medicationCodeableConcept", Terminology: "http://sys-ids.kemkes.go.id/kfa", Description: "Data harus menggunakan Kode MedicationAdministration.medicationCodeableConcept dapat diisi dengan kode KFA sesuai dengan use case masing-masing. Untuk keterangan lebih lanjut dapat melihat dalam playbook masing-masing usecase.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10758: {Path: "MedicationAdministration.medicationReference", Terminology: "Reference ke Medication", Description: "Referensi pada elemen MedicationAdministration.medicationReference harus mengarah ke resource Medication.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10759: {Path: "MedicationAdministration.subject", Terminology: "Reference ke Patient", Description: "Referensi pada elemen MedicationAdministration.subject harus mengarah ke resource Patient.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10760: {Path: "MedicationAdministration.context", Terminology: "Reference ke Encounter", Description: "Referensi pada elemen MedicationAdministration.context harus mengarah ke resource Encounter.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10761: {Path: "MedicationAdministration.effectiveDateTime", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:11", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10762: {Path: "MedicationAdministration.performer.actor", Terminology: "Reference ke Practitioner", Description: "Referensi pada elemen MedicationAdministration.performer.actor harus mengarah ke resource Practitioner.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10763: {Path: "MedicationAdministration.reasonCode", Terminology: "http://terminology.hl7.org/CodeSystem/reason-medication-given", Description: "Nilai pada elemen MedicationAdministration.reasonCode harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/reason-medication-given.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10764: {Path: "MedicationAdministration.reasonReference", Terminology: "Reference ke Condition | Observation |", Description: "Referensi pada elemen MedicationAdministration.reasonReference harus mengarah ke resource Condition | Observation |.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10765: {Path: "MedicationAdministration.request", Terminology: "Reference ke MedicationRequest", Description: "Referensi pada elemen MedicationAdministration.request harus mengarah ke resource MedicationRequest.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10766: {Path: "MedicationAdministration.dosage.site", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen MedicationAdministration.dosage.site harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10767: {Path: "MedicationAdministration.dosage.route", Terminology: "http://www.whocc.no/atc", Description: "Nilai pada elemen MedicationAdministration.dosage.route harus sesuai dengan sistem http://www.whocc.no/atc.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10768: {Path: "MedicationAdministration.dosage.method", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen MedicationAdministration.dosage.method harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10769: {Path: "MedicationAdministration.dosage.dose.value", Terminology: "decimal", Description: "Nilai pada elemen MedicationAdministration.dosage.dose.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10770: {Path: "MedicationAdministration.dosage.rateRatio.numerator.system", Terminology: "http://unitsofmeasure.org", Description: "Data harus menggunakan Kode satuan atau unit dapat menggunakan kode yang disediakan oleh Unified Code for\nUnits of Measure (UCUM). Daftar kode satuan dapat diakses  disini", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10771: {Path: "MedicationAdministration.dosage.rateRatio.numerator.value", Terminology: "decimal", Description: "Nilai pada elemen MedicationAdministration.dosage.rateRatio.numerator.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10772: {Path: "MedicationAdministration.dosage.rateRatio.denominator.system", Terminology: "http://unitsofmeasure.org", Description: "Nilai pada elemen MedicationAdministration.dosage.rateRatio.denominator.system harus sesuai dengan sistem http://unitsofmeasure.org.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10773: {Path: "MedicationAdministration.dosage.rateRatio.denominator.value", Terminology: "decimal", Description: "Nilai pada elemen MedicationAdministration.dosage.rateRatio.denominator.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10774: {Path: "MedicationAdministration.dosage.rateQuantity.value", Terminology: "decimal", Description: "Nilai pada elemen MedicationAdministration.dosage.rateQuantity.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10775: {Path: "CoverageEligibilityResponse.insurance.item.benefit.type", Terminology: "http://terminology.hl7.org/CodeSystem/benefit-type", Description: "Nilai pada elemen CoverageEligibilityResponse.insurance.item.benefit.type harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/benefit-type.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10776: {Path: "CoverageEligibilityResponse.insurance.item.benefit.allowedMoney.currency", Terminology: "urn:iso:std:iso:4217\ndecimal", Description: "Nilai pada elemen CoverageEligibilityResponse.insurance.item.benefit.allowedMoney.currency dapat diisi dengan kode 'IDR' sesuai standar ISO 4217.\nNilai pada elemen CoverageEligibilityResponse.insurance.item.benefit.allowedMoney.value harus valid sesuai aturan: Tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10777: {Path: "Account.guarantor.party", Terminology: "Reference ke Patient, Related Person, Organization", Description: "Referensi pada elemen Account.guarantor.party harus mengarah ke resource Patient, Related Person, Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10778: {Path: "ClaimResponse.item.adjudication.category", Terminology: "http://terminology.hl7.org/CodeSystem/adjudication", Description: "Nilai pada elemen ClaimResponse.item.adjudication.category harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/adjudication.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10779: {Path: "ClaimResponse.item.adjudication.amount.currency", Terminology: "urn:iso:std:iso:4217\ndecimal", Description: "Nilai pada elemen ClaimResponse.item.adjudication.amount.currency dapat diisi dengan kode 'IDR' sesuai standar ISO 4217.\nNilai pada elemen ClaimResponse.item.adjudication.amount.value harus valid sesuai aturan: Tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10782: {Path: "Encounter.reasonCode.coding", Terminology: "http://snomed.info/sct\nhttp://terminology.kemkes.go.id", Description: "Nilai pada elemen Encounter.reasonCode.coding harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Encounter.reasonCode.coding", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10783: {Path: "MedicationStatement.dosage.site", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen MedicationStatement.dosage.site harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10784: {Path: "MedicationDispense.dosageInstruction.site", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen MedicationDispense.dosageInstruction.site harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10785: {Path: "ChargeItem.performer.actor", Terminology: "Reference to Practitioner", Description: "Referensi pada elemen ChargeItem.performer.actor harus mengarah ke resource Practitioner.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10810: {Path: "Claim.supportingInfo.valueReference", Terminology: "Reference to Any", Description: "Elemen Claim.supportingInfo.valueReference dapat mereferensikan resource FHIR yang valid.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10811: {Path: "Invoice.lineItem.priceComponent.amount.currency", Terminology: "urn:iso:std:iso:4217\nTipe data decimal, contoh: 52000", Description: "Nilai pada elemen Invoice.lineItem.priceComponent.amount.currency dapat diisi dengan kode 'IDR' sesuai standar ISO 4217.\nNilai pada elemen Invoice.lineItem.priceComponent.amount.value harus valid sesuai aturan: Tipe data decimal, contoh: 52000.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10812: {Path: "Invoice.totalNet.currency", Terminology: "urn:iso:std:iso:4217\nTipe data decimal, contoh: 52000", Description: "Nilai pada elemen Invoice.totalNet.currency dapat diisi dengan kode 'IDR' sesuai standar ISO 4217.\nNilai pada elemen Invoice.totalNet.value harus valid sesuai aturan: Tipe data decimal, contoh: 52000.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10819: {Path: "Immunization.reaction.date", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:12", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10820: {Path: "Immunization.reaction.detail", Terminology: "Reference to Observation", Description: "Referensi pada elemen Immunization.reaction.detail harus mengarah ke resource Observation.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10825: {Path: "Condition.stage.assessment", Terminology: "Reference to ClinicalImpression", Description: "Referensi pada elemen Condition.stage.assessment harus mengarah ke resource ClinicalImpression.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10826: {Path: "Goal.addresses", Terminology: "Reference to Condition", Description: "Referensi pada elemen Goal.addresses harus mengarah ke resource Condition.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10827: {Path: "NutritionOrder.foodPreferenceModifier", Terminology: "http://terminology.hl7.org/CodeSystem/diet", Description: "Nilai pada elemen NutritionOrder.foodPreferenceModifier harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/diet.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10828: {Path: "NutritionOrder.excludeFoodModifier", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen NutritionOrder.excludeFoodModifier harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10839: {Path: "Encounter.length.value", Terminology: "decimal", Description: "Nilai pada elemen Encounter.length.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10840: {Path: "Encounter.length.code", Terminology: "http://unitsofmeasure.org ( Ranap = d Rajal/IGD/ODC = h dan min)", Description: "Nilai pada elemen Encounter.length.code harus sesuai dengan sistem http://unitsofmeasure.org ( Ranap = d Rajal/IGD/ODC = h dan min).", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10841: {Path: "Condition.onsetPeriod.start", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:13", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10842: {Path: "Condition.onsetPeriod.end", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:14", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10843: {Path: "Condition.onsetAge.value", Terminology: "decimal", Description: "Nilai pada elemen Condition.onsetAge.value harus berupa tipe data decimal.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10844: {Path: "Condition.onsetAge.code", Terminology: "http://unitsofmeasure.org", Description: "Nilai pada elemen Condition.onsetAge.code harus sesuai dengan sistem http://unitsofmeasure.org.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10845: {Path: "ClinicalImpression.code", Terminology: "http://snomed.info/sct\nhttp://terminology.kemkes.go.id", Description: "Nilai pada elemen ClinicalImpression.code harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab ClinicalImpression.code", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10846: {Path: "Goal.statusDate", Terminology: "date", Description: "Data harus menggunakan format date YYYY-MM-DD", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10847: {Path: "RiskAssessment.performer", Terminology: "Reference", Description: "Elemen RiskAssessment.performer dapat mereferensikan resource FHIR yang valid.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10848: {Path: "MedicationAdministration.manufacturer", Terminology: "Reference ke Organization", Description: "Referensi pada elemen MedicationAdministration.manufacturer harus mengarah ke resource Organization.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10849: {Path: "MedicationAdministration.category", Terminology: "http://terminology.hl7.org/CodeSystem/medication-admin-category", Description: "Nilai pada elemen MedicationAdministration.category harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/medication-admin-category.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10850: {Path: "MedicationAdministration.effectivePeriod.start", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:15", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10851: {Path: "MedicationAdministration.effectivePeriod.end", Terminology: "dateTime", Description: "Data harus menggunakan format datetime YYYY-MM-DDThh:mm:ss+00:16", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10852: {Path: "MedicationAdministration.dosage.dose", Terminology: "http://terminology.hl7.org/CodeSystem/v3-orderableDrugForm\nhttp://snomed.info/sct\nhttps://terminology.kemkes.go.id/CodeSystem/smallest-logistic-unit\nhttp://unitsofmeasure.org", Description: "Nilai pada elemen MedicationAdministration.dosage.dose harus diisi menggunakan kode yang valid dari sistem terminologi http://terminology.hl7.org/CodeSystem/v3-orderableDrugForm.\nNilai pada elemen MedicationAdministration.dosage.dose harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab MedicationAdministration.dosage.dose\nNilai pada elemen MedicationAdministration.dosage.dose.code harus sesuai dengan sistem http://unitsofmeasure.org.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10853: {Path: "Composition.section.X.code.coding", Terminology: "http://loinc.org\nhttp://terminology.kemkes.go.id\nhttp://terminology.kemkes.go.id/CodeSystem/composition-section", Description: "Nilai pada elemen Composition.section.X.code.coding harus diisi menggunakan kode yang valid dari sistem terminologi http://loinc.org.\nData harus menggunakan codesystem yang sesuai, lihat Dokumen Lampiran Standar Terminologi SATUSEHAT Lihat sub-bab Composition.section.X.code.coding", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10854: {Path: "Composition.section.X.text.status", Terminology: "http://hl7.org/fhir/narrative-status", Description: "Nilai pada elemen Composition.section.X.text.status harus sesuai dengan sistem http://hl7.org/fhir/narrative-status.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10861: {Path: "Observation.hasMember", Terminology: "Reference ke Observation | QuestionnaireResponse", Description: "Referensi pada elemen Observation.hasMember harus mengarah ke resource Observation | QuestionnaireResponse.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10862: {Path: "Observation.component.valueQuantity.code", Terminology: "http://unitsofmeasure.org", Description: "Nilai pada elemen Observation.component.valueQuantity.code harus sesuai dengan sistem http://unitsofmeasure.org.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10863: {Path: "QuestionnaireResponse.item.X.answer.item.X.answer.valueReference", Terminology: "Reference Any", Description: "Elemen QuestionnaireResponse.item.X.answer.item.X.answer.valueReference dapat mereferensikan resource FHIR yang valid.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10900: {Path: "Coverage.payor.display", Terminology: "Mandatory", Description: "Elemen Coverage.payor.display wajib diisi.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10924: {Path: "PractitionerRole.identifier.system", Terminology: "https://fhir.kemkes.go.id/id/sip-number", Description: "Identifier pada elemen PractitionerRole.identifier.system harus menggunakan sistem https://fhir.kemkes.go.id/id/sip-number.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10925: {Path: "PractitionerRole.identifier.use", Terminology: "http://hl7.org/fhir/identifier-use", Description: "Nilai pada elemen PractitionerRole.identifier.use harus sesuai dengan sistem http://hl7.org/fhir/identifier-use.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10926: {Path: "PractitionerRole.identifier.value", Terminology: "Tipe data string", Description: "Nilai pada elemen PractitionerRole.identifier.value harus berupa tipe data string.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10927: {Path: "PractitionerRole.identifier.period.start", Terminology: "Tipe data dateTime", Description: "Format tanggal dan waktu pada elemen PractitionerRole.identifier.period.start harus berupa dateTime (YYYY-MM-DDThh:mm:ss+zz:zz).", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10928: {Path: "PractitionerRole.identifier.period.end", Terminology: "Tipe data dateTime, future date berlaku", Description: "Format tanggal dan waktu pada elemen PractitionerRole.identifier.period.end harus berupa dateTime (YYYY-MM-DDThh:mm:ss+zz:zz).", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10929: {Path: "PractitionerRole.identifier.assigner", Terminology: "Reference (Organization)", Description: "Referensi pada elemen PractitionerRole.identifier.assigner harus mengarah ke resource (Organization).", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10930: {Path: "PractitionerRole.identifier.assigner.display", Terminology: "Tipe data string", Description: "Nilai pada elemen PractitionerRole.identifier.assigner.display harus berupa tipe data string.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10931: {Path: "PractitionerRole.practitioner", Terminology: "Reference ke resource Practitioner (ID SATUSEHAT Practitioner)\nMandatory", Description: "Referensi pada elemen PractitionerRole.practitioner harus mengarah ke resource resource Practitioner (ID SATUSEHAT Practitioner).\nElemen PractitionerRole.practitioner wajib diisi.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10932: {Path: "PractitionerRole.organization", Terminology: "Reference ke resource Organization (ID SATUSEHAT Organization - Fasyankes)\nMandatory", Description: "Referensi pada elemen PractitionerRole.organization harus mengarah ke resource resource Organization (ID SATUSEHAT Organization - Fasyankes).\nElemen PractitionerRole.organization wajib diisi.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10933: {Path: "PractitionerRole.code", Terminology: "http://snomed.info/sct\nMandatory", Description: "Nilai pada elemen PractitionerRole.code harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.\nElemen PractitionerRole.code wajib diisi.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10934: {Path: "PractitionerRole.specialty", Terminology: "http://snomed.info/sct", Description: "Nilai pada elemen PractitionerRole.specialty harus diisi menggunakan kode yang valid dari sistem terminologi http://snomed.info/sct.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10935: {Path: "PractitionerRole.telecom.system", Terminology: "http://hl7.org/fhir/contact-point-system", Description: "Nilai pada elemen PractitionerRole.telecom.system harus sesuai dengan sistem http://hl7.org/fhir/contact-point-system.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10936: {Path: "PractitionerRole.telecom.use", Terminology: "http://hl7.org/fhir/contact-point-use", Description: "Nilai pada elemen PractitionerRole.telecom.use harus sesuai dengan sistem http://hl7.org/fhir/contact-point-use.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10937: {Path: "PractitionerRole.telecom.value", Terminology: "Tipe data string", Description: "Nilai pada elemen PractitionerRole.telecom.value harus berupa tipe data string.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10938: {Path: "PractitionerRole.active", Terminology: "Tipe data boolean", Description: "Nilai pada elemen PractitionerRole.active harus berupa tipe data boolean.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10939: {Path: "PractitionerRole.availableTime.daysOfWeek", Terminology: "http://hl7.org/fhir/days-of-week", Description: "Nilai pada elemen PractitionerRole.availableTime.daysOfWeek harus sesuai dengan sistem http://hl7.org/fhir/days-of-week.", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10940: {Path: "PractitionerRole.availableTime.availableStartTime", Terminology: "Tipe data time", Description: "Format waktu pada elemen PractitionerRole.availableTime.availableStartTime harus berupa time (HH:mm:ss).", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	10941: {Path: "PractitionerRole.availableTime.availableEndTime", Terminology: "Tipe data time", Description: "Format waktu pada elemen PractitionerRole.availableTime.availableEndTime harus berupa time (HH:mm:ss).", Fix: "", Version: "v2.0", LastUpdate: "10 Desember 2025"},
	20002: {Path: "", Terminology: "Found duplicate resource: {Nama Resource}", Description: "Terjadi duplikasi data pada saat pengiriman resource.Pastikan setiap resource yang dikirim memiliki identifier unik dan belum pernah diregistrasikan sebelumnya.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20003: {Path: "", Terminology: "Failed to get permission for client for resource type", Description: "Client ID Anda tidak memiliki scope akses untuk resource yang diminta (misalnya read:Patient atau write:Encounter). Periksa kembali konfigurasi token dan scope di OAuth Client Configuration SATUSEHAT Developer Portal.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20004: {Path: "", Terminology: "Resource {Resource Type} is not allowed for current access class", Description: "Jenis organisasi Anda (contoh: Klinik, Laboratorium, RS) tidak diizinkan untuk memproses resource tertentu. Misalnya, Laboratorium tidak dapat membuat Encounter. Sesuaikan penggunaan resource dengan peran organisasi di Master Organization Registry.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20005: {Path: "", Terminology: "Failed to read request payload", Description: "Server gagal membaca payload karena format JSON tidak sesuai. Cek kembali struktur dan encoding file sebelum dikirim melalui endpoint.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20006: {Path: "", Terminology: "Invalid JSON Payload Format", Description: "Struktur JSON tidak valid. Periksa tanda baca (koma, kurung, atau kutip) dan pastikan formatnya sesuai dengan contoh pada Postman Collection SATUSEHAT API.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20007: {Path: "", Terminology: "Resource not found {Resource Type}/{Resource ID}", Description: "ID resource yang diminta tidak ditemukan di FHIR Store SATUSEHAT. Pastikan ID tersebut benar dan sudah tersimpan melalui proses registrasi sebelumnya.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20008: {Path: "", Terminology: "Failed to unmarshal request payload", Description: "Payload valid secara JSON tetapi tidak sesuai dengan struktur FHIR yang diharapkan. Misalnya field valueInteger dikirim dalam format string. Cek kembali tipe data berdasarkan FHIR Resource Definition.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20009: {Path: "", Terminology: "Please provide Client ID", Description: "Client ID tidak ditemukan pada request. Pastikan parameter client_id dikirimkan saat proses autentikasi (/auth/token) atau saat pemanggilan API yang memerlukan otorisasi.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20010: {Path: "", Terminology: "Can't get Client public key", Description: "Public key yang terkait dengan Client ID tidak dapat ditemukan di sistem. Pastikan public key sudah diregistrasikan di halaman konfigurasi Client di Developer Portal agar proses signature verification dapat berjalan.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20011: {Path: "", Terminology: "Forbidden nested bundle", Description: "FHIR tidak mengizinkan adanya Bundle di dalam Bundle. Pastikan struktur pengiriman resource tidak bersarang; setiap Bundle hanya boleh berisi resource-level entries.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20012: {Path: "", Terminology: "Forbidden resource {Resource Type}  due to previous error", Description: "Dalam transaksi Bundle, resource berikutnya tidak dapat diproses karena resource sebelumnya gagal divalidasi. Perbaiki resource awal yang gagal agar dependensi dapat berjalan.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20013: {Path: "", Terminology: "Invalid resource {Resource Type} due to previous error", Description: "Kesalahan lanjutan pada rule Number 20012, di mana resource bergantung pada resource yang gagal sebelumnya dalam satu Bundle.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20014: {Path: "", Terminology: "Everything search is not supported for resource {Resource Type}", Description: "Parameter $everything belum diaktifkan untuk resource tersebut. Saat ini fitur ini hanya tersedia untuk resource tertentu", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20015: {Path: "", Terminology: "Ops Portal Organization can only edit name, birthdate, gender, address, deceasedBoolean. Please remove attribute: {Attribute list}", Description: "Akses Anda terbatas pada atribut tertentu. Misalnya hanya dapat mengubah nama, tanggal lahir, atau alamat pasien. Hapus atribut lain yang tidak diizinkan dari request.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20016: {Path: "", Terminology: "Organization not allowed to edit Patient", Description: "Fasyankes tidak memiliki izin untuk mengubah data pasien. Akses pengeditan hanya diberikan kepada pihak yang berwenang", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20017: {Path: "", Terminology: "Invalid element {Resource Type}.fullUrl due to invalid fullUrl format. It must be in the format: 'urn:uuid:{{uuid_value}}.{UUID}", Description: "Format fullUrl tidak valid. Nilai fullUrl harus mengikuti format yang benar, yaitu urn:uuid:{{uuid_value}}.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20018: {Path: "", Terminology: "Invalid element {Resource Type}.fullUrl due to invalid fullUrl format. It must begin with 'urn:uuid'.{UUID}", Description: "Format fullUrl tidak sesuai. Setiap fullUrl wajib diawali dengan awalan urn:uuid agar dapat direferensikan secara internal antar resource di dalam Bundle.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20019: {Path: "", Terminology: "Invalid element {Resource Type}.fullUrl due to invalid UUID format", Description: "Format elemen fullUrl salah. Gunakan format standar:\n\n\"fullUrl\": \"urn:uuid:{UUID}\"\n\n\nUUID ini digunakan sebagai referensi internal antar-resource dalam satu Bundle.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20020: {Path: "", Terminology: "Invalid element {Resource Type}.fullUrl due to duplicate fullUrl: {Full URL String}", Description: "fullUrl tidak boleh sama dalam satu Bundle. Pastikan setiap resource memiliki urn:uuid unik.", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
	20021: {Path: "", Terminology: "Invalid resource {Resource Type} due to missing fullUrl element", Description: "Setiap resource dalam transaksi Bundle wajib memiliki fullUrl. Tanpa elemen ini, sistem tidak dapat melakukan referensi antar-resource (misal antara Encounter dan Practitioner).", Fix: "", Version: "v1.1", LastUpdate: "31 Oktober 2025"},
}

// ResponseTypes contains all 14 SATUSEHAT validation error response types.
var ResponseTypes = map[string]ResponseTypeInfo{
	"duplicate":                 {Title: "Found duplicate resource", Cause: "Pengiriman payload request lebih dari satu kali.", Fix: "Pastikan setiap resource yang dikirim memiliki identifier unik dan belum pernah diregistrasikan sebelumnya."},
	"code_not_found":            {Title: "Code not found", Cause: "Kode pada payload request tidak tersedia di sistem.", Fix: "Gunakan codesystem/valueset sesuai standar SATUSEHAT dan/atau FHIR HL7."},
	"invalid_datetime_format":   {Title: "Invalid date time format", Cause: "Format tanggal tidak sesuai standar.", Fix: "Gunakan format YYYY-MM-DDThh:mm:ss+00:00 (UTC +00:00). Contoh: WIB 17:35 -> 2023-08-23T10:35:00+00:00"},
	"invalid_datetime_value":    {Title: "Invalid date time value", Cause: "Tanggal tidak sesuai standar (tanggal future atau older than 31st August 2022).", Fix: "WAJIB gunakan UTC +00. Tanggal tidak boleh di masa depan atau sebelum 31 Agustus 2022."},
	"reference_mandatory":       {Title: "Reference is mandatory", Cause: "Referensi resource tidak valid atau tidak diisi.", Fix: "WAJIB menyertakan referensi resource yang VALID."},
	"invalid_coding_system":     {Title: "Invalid Coding System", Cause: "Codesystem tidak diperbolehkan untuk element terkait.", Fix: "Gunakan codesystem/valueset sesuai standar SATUSEHAT dan/atau FHIR HL7."},
	"element_not_found":         {Title: "Element not found", Cause: "Element mandatory tidak terkirim/tidak diisi.", Fix: "Simbol * sebelum nama element = WAJIB ada. Lihat Daftar Pemetaan Nilai."},
	"invalid_identifier_system": {Title: "Invalid identifier System", Cause: "Codesystem pada Identifier tidak mengikuti standar.", Fix: "WAJIB gunakan codesystem VALID pada Identifier."},
	"integer_formatted_string":  {Title: "Integer formated as a string", Cause: "Value tidak bernilai integer atau menggunakan tanda petik.", Fix: "Value harus integer TANPA tanda petik."},
	"invalid_boolean_value":     {Title: "Invalid boolean value", Cause: "Value tidak sesuai format boolean.", Fix: "Gunakan nilai boolean yang valid: true, false, 1, 0."},
	"not_null":                  {Title: "Element not null", Cause: "Data harus terisi (tidak boleh kosong).", Fix: "Pastikan elemen wajib terisi sebelum pengiriman."},
	"invalid_reference":         {Title: "Invalid reference", Cause: "Referensi resource tidak valid atau belum terdaftar.", Fix: "Pastikan referensi mengarah ke resource yang sudah terdaftar di SATUSEHAT."},
	"invalid_decimal_value":     {Title: "Invalid decimal value", Cause: "Nilai bukan tipe data decimal yang valid.", Fix: "Gunakan tipe data decimal yang valid tanpa karakter non-numerik."},
	"invalid_integer_value":     {Title: "Invalid integer value", Cause: "Nilai bukan tipe data integer yang valid.", Fix: "Gunakan tipe data integer yang valid tanpa karakter non-numerik."},
}

var (
	ruleNumberRegex = regexp.MustCompile(`(?i)RuleNumber:\s*(\d+)`)
	exprRuleRegex   = regexp.MustCompile(`^(\d+)\s`)
)

// SSValidationError maps SATUSEHAT Rule Numbers and OperationOutcome issues to human-readable descriptions.
type SSValidationError struct {
	issues []map[string]any
}

// NewSSValidationError creates a new SSValidationError from a slice of issue maps.
func NewSSValidationError(issues []map[string]any) *SSValidationError {
	if issues == nil {
		issues = []map[string]any{}
	}
	return &SSValidationError{issues: issues}
}

// FromResponse creates an SSValidationError from an *SSResponse, map[string]any, or []map[string]any.
func FromResponse(response any) *SSValidationError {
	var issues []map[string]any

	switch v := response.(type) {
	case *SSResponse:
		if v != nil {
			body := v.GetBody()
			issues = extractIssuesFromBody(body)
		}
	case SSResponse:
		body := v.GetBody()
		issues = extractIssuesFromBody(body)
	case map[string]any:
		issues = extractIssuesFromBody(v)
	case []map[string]any:
		issues = v
	case []any:
		issues = toMapSlice(v)
	}

	return NewSSValidationError(issues)
}

func extractIssuesFromBody(body map[string]any) []map[string]any {
	if body == nil {
		return []map[string]any{}
	}
	if iss := toMapSlice(body["issue"]); iss != nil {
		return iss
	}
	if rt, _ := body["resourceType"].(string); rt == "OperationOutcome" {
		return []map[string]any{}
	}
	return []map[string]any{}
}

// GetRuleNumbers extracts all unique rule numbers from issues.
func (v *SSValidationError) GetRuleNumbers() []int {
	seen := map[int]bool{}
	numbers := []int{}

	for _, issue := range v.issues {
		textStr := extractValidationIssueText(issue)
		if m := ruleNumberRegex.FindStringSubmatch(textStr); len(m) > 1 {
			if n, err := strconv.Atoi(m[1]); err == nil && !seen[n] {
				seen[n] = true
				numbers = append(numbers, n)
			}
		}
		for _, expr := range extractExpressions(issue["expression"]) {
			if m := exprRuleRegex.FindStringSubmatch(expr); len(m) > 1 {
				if n, err := strconv.Atoi(m[1]); err == nil && !seen[n] {
					seen[n] = true
					numbers = append(numbers, n)
				}
			}
		}
	}

	return numbers
}

// GetRuleInfo returns the rule metadata map for a given rule number, or nil if not found.
func (v *SSValidationError) GetRuleInfo(ruleNumber int) map[string]string {
	return GetRuleInfo(ruleNumber)
}

// GetRuleInfo returns the rule metadata map for a given rule number, or nil if not found.
func GetRuleInfo(ruleNumber int) map[string]string {
	if info, ok := RuleMap[ruleNumber]; ok {
		return info.ToMap()
	}
	return nil
}

// GetParsedIssues returns all issues enriched with rule number metadata and detected response type.
func (v *SSValidationError) GetParsedIssues() []map[string]any {
	parsed := make([]map[string]any, 0, len(v.issues))

	for _, issue := range v.issues {
		textStr := extractValidationIssueText(issue)
		expressions := extractExpressions(issue["expression"])
		pathStr := strings.Join(expressions, ", ")

		var ruleNumber any
		var ruleInfo *RuleInfo
		if m := ruleNumberRegex.FindStringSubmatch(textStr); len(m) > 1 {
			if n, err := strconv.Atoi(m[1]); err == nil {
				ruleNumber = n
				if ri, ok := RuleMap[n]; ok {
					ruleInfo = &ri
				}
			}
		}

		code, _ := issue["code"].(string)
		severity, _ := issue["severity"].(string)
		if severity == "" {
			severity = "error"
		}

		respType := detectResponseType(textStr, code)
		var respTypeVal any
		if respType != "" {
			respTypeVal = respType
		}

		desc := ""
		term := ""
		fix := ""
		ver := ""
		lastUpdate := ""

		if ruleInfo != nil {
			desc = ruleInfo.Description
			term = ruleInfo.Terminology
			fix = ruleInfo.Fix
			ver = ruleInfo.Version
			lastUpdate = ruleInfo.LastUpdate
		}
		if desc == "" && respType != "" {
			if rt, ok := ResponseTypes[respType]; ok {
				desc = rt.Title
			}
		}
		if fix == "" && respType != "" {
			if rt, ok := ResponseTypes[respType]; ok {
				fix = rt.Fix
			}
		}

		parsed = append(parsed, map[string]any{
			"rule_number":   ruleNumber,
			"path":          pathStr,
			"severity":      severity,
			"code":          code,
			"text":          textStr,
			"description":   desc,
			"terminology":   term,
			"fix":           fix,
			"version":       ver,
			"last_update":   lastUpdate,
			"response_type": respTypeVal,
			"raw_issue":     issue,
		})
	}

	return parsed
}

func detectResponseType(text, code string) string {
	textLower := strings.ToLower(text)

	if strings.Contains(textLower, "duplicate") || code == "duplicate" {
		return "duplicate"
	}
	if strings.Contains(textLower, "code not found") {
		return "code_not_found"
	}
	if strings.Contains(textLower, "invalid date time format") {
		return "invalid_datetime_format"
	}
	if strings.Contains(textLower, "invalid date time value") {
		return "invalid_datetime_value"
	}
	if strings.Contains(textLower, "reference is mandatory") || strings.Contains(textLower, "reference must") {
		return "reference_mandatory"
	}
	if strings.Contains(textLower, "invalid coding system") {
		return "invalid_coding_system"
	}
	if strings.Contains(textLower, "element not found") {
		return "element_not_found"
	}
	if strings.Contains(textLower, "invalid identifier system") {
		return "invalid_identifier_system"
	}
	if strings.Contains(textLower, "integer formated as a string") || strings.Contains(textLower, "integer formatted") {
		return "integer_formatted_string"
	}
	if strings.Contains(textLower, "invalid boolean") {
		return "invalid_boolean_value"
	}
	if strings.Contains(textLower, "not null") || strings.Contains(textLower, "data harus terisi") {
		return "not_null"
	}
	if strings.Contains(textLower, "invalid reference") || (strings.Contains(textLower, "reference to") && strings.Contains(textLower, "not valid")) {
		return "invalid_reference"
	}
	if strings.Contains(textLower, "invalid decimal") || strings.Contains(textLower, "invalid type") {
		return "invalid_decimal_value"
	}
	if strings.Contains(textLower, "invalid integer") {
		return "invalid_integer_value"
	}

	return ""
}

// GetSummary returns a human-readable summary of all errors.
func (v *SSValidationError) GetSummary() string {
	parsed := v.GetParsedIssues()
	lines := make([]string, 0, len(parsed))

	for _, issue := range parsed {
		ruleNum := "N/A"
		if issue["rule_number"] != nil {
			ruleNum = fmt.Sprintf("%v", issue["rule_number"])
		}
		path, _ := issue["path"].(string)
		if path == "" {
			path = "-"
		}
		desc, _ := issue["description"].(string)
		if desc == "" {
			desc, _ = issue["text"].(string)
		}
		if desc == "" {
			desc = "-"
		}
		lines = append(lines, fmt.Sprintf("[%s] %s => %s", ruleNum, path, desc))
	}

	return strings.Join(lines, "\n")
}

// IsDuplicate returns true if any issue indicates a duplicate resource.
func (v *SSValidationError) IsDuplicate() bool {
	for _, issue := range v.issues {
		text := strings.ToLower(extractValidationIssueText(issue))
		if strings.Contains(text, "duplicate") {
			return true
		}
	}
	return false
}

// IsDateTimeError returns true if any issue is a date/time error.
func (v *SSValidationError) IsDateTimeError() bool {
	for _, issue := range v.issues {
		text := strings.ToLower(extractValidationIssueText(issue))
		if strings.Contains(text, "date time") || strings.Contains(text, "datetime") {
			return true
		}
	}
	return false
}

// GetAffectedPaths returns all unique affected FHIR paths from issue expressions.
func (v *SSValidationError) GetAffectedPaths() []string {
	seen := map[string]bool{}
	paths := []string{}

	for _, issue := range v.issues {
		for _, expr := range extractExpressions(issue["expression"]) {
			if expr != "" && !seen[expr] {
				seen[expr] = true
				paths = append(paths, expr)
			}
		}
	}

	return paths
}

// GetIssues returns the raw issues slice.
func (v *SSValidationError) GetIssues() []map[string]any {
	return v.issues
}

// GetAvailableRuleNumbers returns all rule numbers in RuleMap sorted ascending.
func (v *SSValidationError) GetAvailableRuleNumbers() []int {
	return GetAvailableRuleNumbers()
}

// GetAvailableRuleNumbers returns all rule numbers in RuleMap sorted ascending.
func GetAvailableRuleNumbers() []int {
	nums := make([]int, 0, len(RuleMap))
	for k := range RuleMap {
		nums = append(nums, k)
	}
	sort.Ints(nums)
	return nums
}

// GetRulesForResource returns all rules whose path starts with "{resourceType}.".
func (v *SSValidationError) GetRulesForResource(resourceType string) map[int]map[string]string {
	return GetRulesForResource(resourceType)
}

// GetRulesForResource returns all rules whose path starts with "{resourceType}.".
func GetRulesForResource(resourceType string) map[int]map[string]string {
	prefix := resourceType + "."
	result := map[int]map[string]string{}
	for num, rule := range RuleMap {
		if strings.HasPrefix(rule.Path, prefix) {
			result[num] = rule.ToMap()
		}
	}
	return result
}

// SearchRules searches rules by keyword in description, terminology, or path.
func (v *SSValidationError) SearchRules(keyword string) map[int]map[string]string {
	return SearchRules(keyword)
}

// SearchRules searches rules by keyword in description, terminology, or path.
func SearchRules(keyword string) map[int]map[string]string {
	kw := strings.ToLower(keyword)
	result := map[int]map[string]string{}
	for num, rule := range RuleMap {
		if strings.Contains(strings.ToLower(rule.Description), kw) ||
			strings.Contains(strings.ToLower(rule.Terminology), kw) ||
			strings.Contains(strings.ToLower(rule.Path), kw) {
			result[num] = rule.ToMap()
		}
	}
	return result
}

// GetResponseType returns the ResponseType metadata map for a given type key, or nil if not found.
func (v *SSValidationError) GetResponseType(typeName string) map[string]string {
	return GetResponseType(typeName)
}

// GetResponseType returns the ResponseType metadata map for a given type key, or nil if not found.
func GetResponseType(typeName string) map[string]string {
	if rt, ok := ResponseTypes[typeName]; ok {
		return rt.ToMap()
	}
	return nil
}

func extractValidationIssueText(issue map[string]any) string {
	if issue == nil {
		return ""
	}
	if detailsMap, ok := issue["details"].(map[string]any); ok {
		if text, ok := detailsMap["text"].(string); ok {
			return text
		}
	}
	if detailsStr, ok := issue["details"].(string); ok {
		return detailsStr
	}
	return ""
}

func extractExpressions(expr any) []string {
	if expr == nil {
		return nil
	}
	switch v := expr.(type) {
	case string:
		if v == "" {
			return nil
		}
		return []string{v}
	case []string:
		return v
	case []any:
		out := make([]string, 0, len(v))
		for _, item := range v {
			if s, ok := item.(string); ok {
				out = append(out, s)
			} else if item != nil {
				out = append(out, fmt.Sprintf("%v", item))
			}
		}
		return out
	default:
		return []string{fmt.Sprintf("%v", v)}
	}
}
