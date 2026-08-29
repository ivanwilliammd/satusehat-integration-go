# satusehat-integration-go

> **Open-source Go SDK for integrating with SATUSEHAT** — Indonesia's national health data platform powered by FHIR R4. Pure Go, no external framework dependency.

[![Go](https://img.shields.io/badge/Go-1.21%2B-blue)](https://go.dev)
[![FHIR R4](https://img.shields.io/badge/FHIR-R4-orange)](https://hl7.org/fhir/R4/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)
[![CI](https://github.com/ivanwilliammd/satusehat-integration-go/actions/workflows/ci.yml/badge.svg)](https://github.com/ivanwilliammd/satusehat-integration-go/actions)

---

## Overview

`satusehat-integration-go` is an **open-source** Go SDK for integrating with **SATUSEHAT** — Indonesia's national health data platform powered by FHIR R4.

Built on the official [SATUSEHAT Platform Guidelines](https://satusehat.kemkes.go.id/platform/docs). Ships with:
- **50 DataType** structs — composable FHIR R4 value objects with `ToArray()` serialization
- **50 PayloadBuilder** classes — fluent builders for all FHIR resources (Patient, Practitioner, Organization, etc.)
- **Queue + Rate Limiter** — in-memory queue with configurable RPM rate limiting

Minimal dependencies: only `github.com/google/uuid`. No framework required.

---

## Requirements

- Go 1.21 or later

---

## Quick Install

```bash
go get github.com/ivanwilliammd/satusehat-integration-go
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

All 51 resources fully implemented via PayloadBuilder classes. Core (✅) + Non-Core (💼):

| # | Resource | Notes |
|---|----------|-------|
| 1 | Patient | ✅ MPI |
| 2 | Practitioner | ✅ SDMK |
| 3 | PractitionerRole | ✅ |
| 4 | Organization | ✅ MSI |
| 5 | Location | ✅ |
| 6 | Encounter | ✅ |
| 7 | Condition | ✅ |
| 8 | Observation | ✅ |
| 9 | Procedure | ✅ |
| 10 | MedicationRequest | ✅ |
| 11 | Bundle | ✅ batch/transaction |
| 12–37 | CarePlan through Task | ✅ |
| 38–50 | Account through Invoice | 💼 Billing/Claim resources |

---

## Usage Examples

### Patient

```go
import (
    "fmt"
    "github.com/ivanwilliammd/satusehat-integration-go/src/builder"
    "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
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
| [Wiki Home](https://github.com/ivanwilliammd/satusehat-integration-go.wiki.git) | Full documentation |
| [Getting Started](https://github.com/ivanwilliammd/satusehat-integration-go/wiki/Getting-Started) | Installation, configuration |
| [DataTypes](https://github.com/ivanwilliammd/satusehat-integration-go/wiki/DataTypes) | Complete type reference |
| [Builders](https://github.com/ivanwilliammd/satusehat-integration-go/wiki/Builders) | Builder usage guide |
| [Resources](https://github.com/ivanwilliammd/satusehat-integration-go/wiki/Resources) | All FHIR resources |
| [Claim Module](https://github.com/ivanwilliammd/satusehat-integration-go/wiki/Claim-Module) | BPJS Klaim integration |

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
