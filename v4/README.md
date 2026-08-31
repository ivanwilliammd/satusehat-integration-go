# [satusehat-integration-go](https://github.com/ivanwilliammd/satusehat-integration-go/tree/main/v4)

> **Open-source Go SDK for integrating with SATUSEHAT** — Indonesia's national health data platform powered by FHIR R4. Pure Go, no external framework dependency.

[![Go](https://img.shields.io/badge/Go-1.21%2B-blue)](https://go.dev)
[![FHIR R4](https://img.shields.io/badge/FHIR-R4-orange)](https://hl7.org/fhir/R4/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![CI](https://github.com/ivanwilliammd/satusehat-integration-go/actions/workflows/ci.yml/badge.svg)](https://github.com/ivanwilliammd/satusehat-integration-go/actions)
[![Go](https://img.shields.io/badge/dynamic/go?label=module&query=import&url=https%3A%2F%2Fproxy.golang.org%2Fshield%2Fgithub.com%252Fivanwilliammd%252Fsatusehat-integration-go%252Fv4%2Fmodule-version)](https://proxy.golang.org/github.com/ivanwilliammd/satusehat-integration-go/v4)

---

## Overview

`satusehat-integration-go` is an **open-source** Go SDK for integrating with **SATUSEHAT** — Indonesia's national health data platform powered by FHIR R4.

Built on the official [SATUSEHAT Platform Guidelines](https://satusehat.kemkes.go.id/platform/docs). Ships with:
- **115+ PayloadBuilder** classes — fluent builders for all FHIR R4 resources (Patient, Practitioner, Organization, Encounter, Observation, Procedure, etc.)
- **50 DataType** structs — composable FHIR R4 value objects with `ToArray()` serialization
- **TerminologyResolver** — castable terminology strings (`"ICD10:A00"`, `"LOINC:2951-2"`, `"SNOMED:38341003"`) directly to CodeableConcept
- **3 SATUSEHAT-specific** resources: BillingStatus (NON-FHIR JSON), PurificationDecision (NON-FHIR JSON), Endpoint (FHIR R4)
- **Queue + Rate Limiter** — in-memory queue with configurable RPM rate limiting
- **Go test suite** — all builders have comprehensive unit tests

Minimal dependencies: only `github.com/google/uuid`. No framework required.

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

---

## Supported FHIR Resources

**115+ PayloadBuilder structs** covering all FHIR R4 resources used in SATUSEHAT interoperability, plus 3 SATUSEHAT-specific resources.

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
