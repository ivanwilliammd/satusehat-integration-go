# [satusehat-integration-go](https://github.com/ivanwilliammd/satusehat-integration-go/tree/main/v4)

> **Open-source Go SDK for integrating with SATUSEHAT** — Indonesia's national health data platform powered by FHIR R4. Pure Go, no external framework dependency.

[![Go Reference](https://pkg.go.dev/badge/github.com/ivanwilliammd/satusehat-integration-go/v4.svg)](https://pkg.go.dev/github.com/ivanwilliammd/satusehat-integration-go/v4)
[![version](https://img.shields.io/github/v/tag/ivanwilliammd/satusehat-integration-go?label=version)](https://github.com/ivanwilliammd/satusehat-integration-go/tags)
[![FHIR R4](https://img.shields.io/badge/FHIR-R4-orange)](https://hl7.org/fhir/R4/)
[![License](https://img.shields.io/badge/License-MIT-green)](../LICENSE)
[![CI](https://github.com/ivanwilliammd/satusehat-integration-go/actions/workflows/ci.yml/badge.svg)](https://github.com/ivanwilliammd/satusehat-integration-go/actions)

> Download counts do not exist for Go modules: the proxy serves source from the tag, there is no
> per-version download counter. The version badge above is the honest substitute.

## Status and coverage (measured, not claimed)

| Item | Count | Evidence |
| --- | --- | --- |
| Builder | 141 resource | `src/builder` |
| Builder R5 | 141 struct | `src/r5/builder` (embedding builder R4, `scripts/generate_r5_builder.py --check`) |
| Method builder | 1.405 | `src/builder/surface_parity_test.go` (kontrak `contracts/php_method_surface.json`) |
| Parser | 142 resource | `src/parser` |
| Tests | 555 test runs passing (494 functions) | `go test ./...` |
| CI | `go vet` + `go test` + generator `--check` | `.github/workflows/ci.yml` |

---

**Module:** `github.com/ivanwilliammd/satusehat-integration-go/v4` — Go 1.21+

> The earlier badge here linked to `proxy.golang.org/.../v4` without `/@v/`, which the proxy answers
> with `bad request: missing /@v/`. The proxy is a machine interface, not a product page; the
> human-facing link is pkg.go.dev, and the badge above points there.

---

## Overview

`satusehat-integration-go` is an **open-source** Go SDK for integrating with **SATUSEHAT** — Indonesia's national health data platform powered by FHIR R4.

Built on the official [SATUSEHAT Platform Guidelines](https://satusehat.kemkes.go.id/platform/docs). Capability set matches the reference PHP SDK, with the differences that remain written down in [`docs/PARITY.md`](docs/PARITY.md):
- **141 PayloadBuilder** structs (`src/builder/`) — fluent builders for all FHIR R4 resources (Patient, Practitioner, Organization, Encounter, Observation, Procedure, etc.)
- **141 R5 builder structs** (`src/r5/builder/`) — the R5 namespace is a thin layer over the canonical builders: each R5 struct embeds the R4 builder, so switching to R5 means changing the import, not rewriting call sites. Divergence only where an R5 difference is verified; nothing is invented.
- **33 DataType** structs (`src/datatype/`) — composable FHIR R4 value objects with `ToArray()` serialization
- **OAuth2Client & SSRequest** (`src/ssrequest/`) — full SATUSEHAT HTTP client supporting `DEV`/`STG`/`PROD`, automatic token refresh on 401, and `Retry-After` / exponential backoff on 429 and 5xx
- **SSResponse, BundleResponse & SSValidationError** (`src/ssresponse/`) — response wrappers plus all **558 SATUSEHAT validation rules** and **14 error response types**
- **142 resource parsers** (`src/parser/`) — one file per FHIR resource; the Go port is the one that had the complete parser set first
- **Full Queue Architecture** (`src/queue/`) — `QueueInterface`, `SqliteQueue` (thread-safe `:memory:` or file persistence with zero CGo), `Worker`, `QueueMonitor`, `ErrorClassifier`, `RateLimiter`, `QueueResult`, `QueueException`
- **TerminologyResolver & TerminologyClient** (`src/terminology/`) — castable terminology strings (`"ICD10:A00"`, `"LOINC:2951-2"`, `"SNOMED:38341003"`) directly to CodeableConcept and SmarTerm HTTP client
- **3 SATUSEHAT-specific** resources: BillingStatus (NON-FHIR JSON), PurificationDecision (NON-FHIR JSON), Endpoint (FHIR R4)
- **Go test suite** — every package carries unit tests; 494 test functions

Zero external CGo dependencies — pure Go standard library. No framework required.

---

## Requirements

- Go 1.21 or later

---

## Quick Install

```bash
go install github.com/ivanwilliammd/satusehat-integration-go/v4/cmd/satusehat@v4.19.0
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

### TerminologyClient — SmarTerm HTTP client (terminology.ivanmd.id)

Standalone client for SmarTerm, **separate from SATUSEHAT OAuth**. Two authentication modes:

1. **PAT (Personal Access Token)** — static token sent as a `Bearer` header, taken from the
   `token` argument or the `TERMINOLOGY_API_TOKEN` environment variable.
2. **OIDC client_credentials** against `auth.ivanmd.id` (Passport) — used when no token is supplied
   and client credentials are available (`TERMINOLOGY_CLIENT_ID`, `TERMINOLOGY_CLIENT_SECRET`,
   `TERMINOLOGY_TOKEN_URL`); the JWT is fetched lazily, cached until expiry, and refreshed once on
   401.

When both are supplied, **PAT wins**. Default base URL is `https://terminology.ivanmd.id`
(`TERMINOLOGY_BASE_URL`), HTTP timeout 30 seconds.

```go
import "github.com/ivanwilliammd/satusehat-integration-go/v4/src/terminology"

term := terminology.NewTerminologyClient("", "") // baca env, default SmarTerm

// Fluent query builder
term.Query("pemasangan kateter").On("kptl").Limit(10).Get()
term.Query("demam").Category("diagnosis").First()
term.Query("paracetamol").On("kfa").All()
term.Query("93004944").AsCode().Get() // → validate; the system is inferred from the code prefix (93xxxx → kfa, KFA)

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
