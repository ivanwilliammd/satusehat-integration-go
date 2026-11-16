# [satusehat-integration-go](https://github.com/ivanwilliammd/satusehat-integration-go/tree/main/v4)

> **Open-source Go SDK for integrating with SATUSEHAT** — Indonesia's national health data platform powered by FHIR R4. Pure Go, no external framework dependency.

## Status & cakupan (diukur, bukan diklaim)

| Bagian | Jumlah | Bukti |
| --- | --- | --- |
| Builder | 141 resource | `src/builder` |
| Builder R5 | 141 struct | `src/r5/builder` (embedding builder kanon, `scripts/generate_r5_builder.py --check`) |
| Method kanon | 1.405 | `src/builder/surface_parity_test.go` (kontrak `contracts/php_method_surface.json`) |
| Parser | 142 resource | `src/parser` |
| Tes | 470 lulus | `go test ./...` |
| CI | `go vet` + `go test` | `.github/workflows/ci.yml` |

Go satu-satunya port dengan parser lengkap sejak awal; PHP, Node, dan Python menyusul.

[![Go](https://img.shields.io/badge/Go-1.21%2B-blue)](https://go.dev)
[![FHIR R4](https://img.shields.io/badge/FHIR-R4-orange)](https://hl7.org/fhir/R4/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![CI](https://github.com/ivanwilliammd/satusehat-integration-go/actions/workflows/ci.yml/badge.svg)](https://github.com/ivanwilliammd/satusehat-integration-go/actions)
[![Go](https://img.shields.io/badge/dynamic/go?label=module&query=import&url=https%3A%2F%2Fproxy.golang.org%2Fshield%2Fgithub.com%252Fivanwilliammd%252Fsatusehat-integration-go%252Fv4%2Fmodule-version)](https://proxy.golang.org/github.com/ivanwilliammd/satusehat-integration-go/v4)

---

## Overview

`satusehat-integration-go` is an **open-source** Go SDK for integrating with **SATUSEHAT** — Indonesia's national health data platform powered by FHIR R4.

Built on the official [SATUSEHAT Platform Guidelines](https://satusehat.kemkes.go.id/platform/docs). Ships with **100% feature parity** with the reference PHP SDK:
- **141 PayloadBuilder** structs (`src/builder/`) — fluent builders for all FHIR R4 resources (Patient, Practitioner, Organization, Encounter, Observation, Procedure, etc.)
- **141 R5 builder structs** (`src/r5/builder/`) — the R5 namespace is a thin layer over the canonical builders: each R5 struct embeds the R4 builder, so switching to R5 means changing the import, not rewriting call sites. Divergence only where an R5 difference is verified; nothing is invented.
- **33 DataType** structs (`src/datatype/`) — composable FHIR R4 value objects with `ToArray()` serialization
- **OAuth2Client & SSRequest** (`src/ssrequest/`) — full SATUSEHAT HTTP client supporting `DEV`/`STG`/`PROD`, automatic token refresh on 401, and `Retry-After` / exponential backoff on 429 and 5xx
- **SSResponse, BundleResponse & SSValidationError** (`src/ssresponse/`) — response wrappers plus all **558 SATUSEHAT validation rules** and **14 error response types**
- **18 FHIR Resource Parsers** (`src/parser/`) — `DeviceParser`, `DiagnosticReportParser`, `DocumentReferenceParser`, `EncounterParser`, `GenomicStudyParser`, `GroupParser`, `LocationParser`, `MolecularSequenceParser`, `ObservationParser`, `OrganizationParser`, `PatientParser`, `PractitionerParser`, `ProcedureParser`, `RelatedPersonParser`, `RiskAssessmentParser`, `ServiceRequestParser`, `SpecimenParser`, `TaskParser`
- **Full Queue Architecture** (`src/queue/`) — `QueueInterface`, `SqliteQueue` (thread-safe `:memory:` or file persistence with zero CGo), `Worker`, `QueueMonitor`, `ErrorClassifier`, `RateLimiter`, `QueueResult`, `QueueException`
- **TerminologyResolver & TerminologyClient** (`src/terminology/`) — castable terminology strings (`"ICD10:A00"`, `"LOINC:2951-2"`, `"SNOMED:38341003"`) directly to CodeableConcept and SmarTerm HTTP client
- **3 SATUSEHAT-specific** resources: BillingStatus (NON-FHIR JSON), PurificationDecision (NON-FHIR JSON), Endpoint (FHIR R4)
- **Go test suite** — all 7 packages have comprehensive unit tests

Zero external CGo dependencies — pure Go standard library. No framework required.

---

## Requirements

- Go 1.21 or later

---

## Quick Install

```bash
go install github.com/ivanwilliammd/satusehat-integration-go/v4/cmd/satusehat@v4.11.2
```

```env
# .env or config
SATUSEHAT_ENV=DEV          # DEV | STG | PROD
SATUSEHAT_BASE_URL_DEV=https://api-satusehat-dev.dto.kemkes.go.id
CLIENTID_DEV=your_client_id
CLIENTSECRET_DEV=your_client_secret
ORGID_DEV=your_org_id
```

---

## Architecture

### DataType Structs (`src/datatype/`)

Atomic FHIR R4 value objects. All provide a `ToArray() map[string]interface{}` method — nested types serialize to clean FHIR JSON automatically.

| Category | Types |
|----------|-------|
| Core | `Coding`, `CodeableConcept`, `Identifier`, `ContactPoint`, `Address`, `HumanName`, `Reference` |
| Quantity | `Age`, `Quantity` |
| Utility | `Period`, `ParameterComponent` |

Example — `HumanName`:

```go
name := datatype.HumanName{
    Family: "Doe",
    Given:  []string{"John", "Michael"},
    Use:    "official",
}
// name.ToArray() → {"family": "Doe", "given": ["John", "Michael"], "use": "official"}
```

### PayloadBuilder Pattern (`src/builder/`)

Fluent builder for each FHIR resource. Each builder exposes chainable methods and returns the resource payload via `ToJSON()`.

```go
patient := builder.NewPatientBuilder().
    SetID("12345678-1234-1234-1234-123456789012").
    AddIdentifier(map[string]interface{}{"system": "https://fhir.kemkes.go.id/id/NIK", "value": "3312345678901234"}).
    AddName(name.ToArray()).
    Build()

payload := patient.ToJSON()
```

### HTTP Client & Validation (`src/ssrequest/` & `src/ssresponse/`)

```go
oauth, _ := ssrequest.NewOAuth2Client()
client := ssrequest.NewSSRequest(oauth, "", "")

res := client.Post("Encounter", encounterPayload)
if !res.IsSuccess() {
    valErr := ssresponse.FromResponse(res)
    fmt.Println(valErr.GetSummary()) // Maps all 558 SATUSEHAT RuleNumbers
}
```

### FHIR Resource Parsers (`src/parser/`)

```go
encParser := parser.NewEncounterParser(res.GetBody())
patientID := encParser.GetSubjectReference()
```

### Queue, Worker & Monitor (`src/queue/`)

```go
q := queue.NewSqliteQueue("satusehat_queue.json") // or ":memory:"
q.EnqueueBundle(bundlePayload, "transaction", "system", 5, nil)

worker := queue.NewWorker(q, map[string]string{
    "base_url":      oauth.BaseURL,
    "fhir_url":      oauth.FhirURL,
    "client_id":     oauth.ClientID,
    "client_secret": oauth.ClientSecret,
}, queue.NewRateLimiter(300))
worker.Process(50)

monitor := queue.NewQueueMonitor(q)
fmt.Println(monitor.HealthCheck())
```

---

## Supported FHIR Resources

**141 PayloadBuilder structs** covering all FHIR R4 resources used in SATUSEHAT interoperability, plus 3 SATUSEHAT-specific resources.

### SATUSEHAT Interoperability Resources (47)

| # | Resource | Builder |
|---|----------|---------|
| 1 | Account | `AccountBuilder` |
| 2 | AllergyIntolerance | `AllergyIntoleranceBuilder` |
| 3 | BillingStatus ⚡NON-FHIR | `BillingStatusBuilder` |
| 4–37 | CarePlan, Condition, Encounter, Goal, Immunization, Location, Medication*, Observation, Organization, Patient, Practitioner, Procedure, ServiceRequest, Specimen, Substance, Task | see `src/builder/` |
| 38 | **Endpoint** | `EndpointBuilder` |
| 39 | **MedicationStatement** | `MedicationStatementBuilder` |
| 40 | **Task** | `TaskBuilder` |
| 41 | **PurificationDecision** ⚡NON-FHIR | `PurificationDecisionBuilder` |
| 42–47 | Claim, ClaimResponse, CoverageEligibilityRequest/Response, DocumentReference, QuestionnaireResponse | see `src/builder/` |

⚡ = NON-FHIR JSON (SATUSEHAT-specific extension)

### BillingStatus (NON-FHIR JSON)
```go
billing := builder.NewBillingStatusBuilder().
    SetID("bs-001").
    AddIdentifier("http://sys-ids.kemkes.go.id/billing/org-001", "BILL-12345").
    SetStatus("active").
    SetInsurer("Organization/org-bpjs", "BPJS Kesehatan").
    SetSubject("100000030009", "Budi Santoso").
    SetRequest("cer-001").
    Build()
```

### Endpoint (FHIR R4)
```go
endpoint := builder.NewEndpointBuilder().
    SetID("ep-001").
    SetStatus("active").
    SetConnectionType("ihe-xcpd", "IHE XCPD").
    SetName("SATUSEHAT FHIR Endpoint").
    SetManagingOrganization("Organization/org-ihs").
    SetAddress("https://satusehat-api.example.com/fhir/r4").
    Build()
```

### PurificationDecision (NON-FHIR JSON)
```go
pd := builder.NewPurificationDecisionBuilder().
    SetID("pd-001").
    AddIdentifier("http://sys-ids.kemkes.go.id/purification/org-001", "PD-12345").
    SetStatus("approved", "Approved").
    SetInsurer("Organization/org-bpjs", "BPJS Kesehatan").
    SetProvider("Organization/hos-001", "Rumah Sakit Sehat").
    SetClaimResponse("cr-001").
    SetCreated("2024-01-15T10:35:00+00:00").
    Build()
```

### TerminologyResolver — castable codes
```go
import "github.com/ivanwilliammd/satusehat-integration-go/v4/src/terminology"

// Cast terminology strings directly to CodeableConcept
terminology.Resolve("ICD10:A00")
// → map[string]interface{}{"coding": []map[string]string{{"system": "http://hl7.org/fhir/sid/icd-10", "code": "A00"}}, "text": "A00"}

terminology.Resolve("LOINC:2951-2")
// → map with loinc.org system

// Batch expand
codes := []string{"ICD10:A00", "ICD10:J18.9"}
resolved := terminology.ExpandArray(codes)
```

### TerminologyClient — client HTTP SmarTerm (terminology.ivanmd.id)

Client mandiri untuk SmarTerm, **terpisah dari OAuth SATUSEHAT**. Auth dual-mode:

1. **PAT (Personal Access Token)** — token statis, header `Bearer`. Dari argumen
   `token` atau env `TERMINOLOGY_API_TOKEN`.
2. **OIDC client_credentials** ke `auth.ivanmd.id` (Passport) — dipakai saat token
   kosong dan kredensial client tersedia (env `TERMINOLOGY_CLIENT_ID`,
   `TERMINOLOGY_CLIENT_SECRET`, `TERMINOLOGY_TOKEN_URL`); fetch JWT lazy, cache
   sampai expire, auto-refresh sekali saat dapat 401.

Jika keduanya diberikan, **PAT menang**. Base URL default
`https://terminology.ivanmd.id` (env `TERMINOLOGY_BASE_URL`), timeout HTTP 30 detik.

```go
import "github.com/ivanwilliammd/satusehat-integration-go/v4/src/terminology"

term := terminology.NewTerminologyClient("", "") // baca env, default SmarTerm

// Fluent query builder
term.Query("pemasangan kateter").On("kptl").Limit(10).Get()
term.Query("demam").Category("diagnosis").First()
term.Query("paracetamol").On("kfa").All()
term.Query("93004944").AsCode().Get() // → validate, sistem ditebak dari prefix kode (93xxxx → kfa)

// Unified & per-system search
term.Search("paracetamol", "all", 20)
term.SearchSystem("icd10", "demam", 20)

// KPTL multi-layer, mapping lintas terminologi
term.KptlCombine("1.02.01", []string{"1.02.01.01", "1.02.01.02"})
term.Map("icd10", "snomed", "A00")

// Validasi & auto-map
term.Validate("kfa", "93004944")
term.AutoMap("icd10", "A00")

// Resolusi kontekstual FHIR
term.ResolveCondition("J18.9", "icd10", "encounter-diagnosis")
term.ResolveChiefComplaint("demam")
term.ResolveObservation("2951-2", "loinc")
term.ResolveMedication("paracetamol")
term.ResolveServiceRequest("LAB", "lab")

// Value sets
term.ValueSets()
term.ValueSet("vs-bpjs-icd10")
```

Semua method mengembalikan `map[string]any` envelope + `error`; HTTP status >= 400
dikembalikan sebagai error `TerminologyClient HTTP {status}`.

---

## Usage Examples

### Patient

```go
import (
    "fmt"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

name := datatype.HumanName{
    Family: "Doe",
    Given:  []string{"John"},
    Use:    "official",
}

patient := builder.NewPatientBuilder().
    SetID("12345678-1234-1234-1234-123456789012").
    AddName(name.ToArray()).
    SetGender("male").
    SetBirthDate("1990-01-15").
    Build()

payload := patient.ToJSON()
fmt.Printf("%+v\n", payload)
```

### Claim (BPJS Klaim)

```go
claim := builder.NewClaimBuilder().
    SetStatus("active").
    SetUse("claim").
    SetType("institutional").
    SetPatient("pat-123", "enc-456").
    AddItem(1, "PROCID001", 150000, "IDR").
    SetTotal(150000, "IDR")

payload := claim.ToJSON()
```

---

## Documentation

| Page | Description |
|------|-------------|
| [Wiki Home](https://github.com/ivanwilliammd/satusehat-integration-go/v4.wiki.git) | Full documentation |
| [Getting Started](https://github.com/ivanwilliammd/satusehat-integration-go/v4/wiki/Getting-Started) | Installation, configuration |
| [DataTypes](https://github.com/ivanwilliammd/satusehat-integration-go/v4/wiki/DataTypes) | Complete type reference |
| [Builders](https://github.com/ivanwilliammd/satusehat-integration-go/v4/wiki/Builders) | Builder usage guide |
| [Resources](https://github.com/ivanwilliammd/satusehat-integration-go/v4/wiki/Resources) | All FHIR resources |
| [Claim Module](https://github.com/ivanwilliammd/satusehat-integration-go/v4/wiki/Claim-Module) | BPJS Klaim integration |

---

## External Resources

- [HL7 FHIR R4 Specification](https://hl7.org/fhir/R4/)
- [SATUSEHAT Platform Docs](https://satusehat.kemkes.go.id/platform/docs)
- [Main PHP SDK](https://github.com/ivanwilliammd/satusehat-integration)
- [SATUSEHAT Sandbox API](https://api-satusehat-dev.dto.kemkes.go.id)

---

## Contributing

Contributions are welcome. Please ensure tests pass and follow existing code conventions.

---

## License

MIT — see [LICENSE](LICENSE).
