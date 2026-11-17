# SRD — satusehat-integration-go (SATUSEHAT FHIR R4 SDK · Go)

Status: v2.1 · 2026-11-17 · Companions: `docs/PRD-SDK-ECOSYSTEM.md` (PHP repository) + `v4/README.md`
Module: `github.com/ivanwilliammd/satusehat-integration-go/v4`, `go 1.21` (`v4/go.mod`).

Every number in this document is **measured** from this checkout on 2026-11-17 with
`scripts/lang_lint.py`-independent shell counts quoted next to each row. The previous revision
(v2.0) reported 18 parsers and 7 test files; both were stale by a wide margin, which is why the
table below quotes the command that produces each number.

## 1. Scope

The internal shape of the Go SDK: package layout, public contract, the capability table against the
other ports, the error surface, installation (library and CLI), and test state.

## 2. Architecture principles

1. **The module lives in `v4/`.** The module declaration is in `v4/go.mod`; CI and release jobs use
   `working-directory: v4`. Go files at the repository root break `go install`.
2. **Two install paths, one source.** Library: `go get …/v4`. CLI:
   `go install …/v4/cmd/satusehat@vX.Y.Z` (`v4/cmd/satusehat/main.go`; its version constant must equal
   the tag).
3. **No CGo.** Every layer (`builder`, `datatype`, `ssrequest`, `ssresponse`, `parser`, `queue`,
   `terminology`) uses the Go standard library only, so `go test ./...` and `go build ./...` run on any
   platform without extra setup.
4. **Rules are data.** SATUSEHAT required rules live in `required-rules.json`, embedded with
   `go:embed`, and their fingerprint is asserted in tests. Adding a rule is a data change.
5. **Non-standard SATUSEHAT objects are constants, not literals.** `src/terminology/nonstandard.go`
   carries the four identifier systems and seven patient extensions with canonical URLs taken from the
   Implementation Guide.

## 3. Measured inventory

| Item | Count | Measurement |
|---|---|---|
| Builder files | 144 | `ls src/builder/*.go \| grep -v _test \| wc -l` |
| Builder types (`…Builder struct`) | 141 | `grep -rhoE "^type [A-Za-z0-9_]+Builder struct" src/builder/*.go \| wc -l` |
| Datatype types | 33 | `grep -rhoE "^type [A-Za-z0-9_]+ struct" src/datatype/*.go \| wc -l` |
| R5 builder files | 141 | `ls src/r5/builder/*.go \| grep -v _test \| wc -l` |
| Parser files | 142 | `ls src/parser/*.go \| grep -v _test \| wc -l` |
| Queue files | 8 | `ls src/queue/*.go \| grep -v _test \| wc -l` |
| Terminology files | 27 | `ls src/terminology/*.go \| grep -v _test \| wc -l` |
| Test functions | 494 | `grep -rho "^func Test" --include="*_test.go" . \| wc -l` |
| Go files, whole module | 748 | `find . -name "*.go" \| wc -l` |

## 4. Capability table against the other ports

| Capability | PHP | Node | Python | Go | Evidence |
|---|---|---|---|---|---|
| Payload builders | 141 | 141 | 141 | 141 | `src/builder` |
| Datatypes | 33 | 33 | 33 | 33 | `src/datatype` |
| FHIR parsers | 142 | 142 | 141 | 142 | `src/parser` |
| R5 namespace | 142 | 141 | 144 | 141 | `src/r5/builder` |
| Queue modules | 8 | 8 | 8 | 8 | `src/queue` |
| Terminology modules | 23 | 22 | 23 | 22 | `src/terminology` |
| SATUSEHAT required rules | yes | yes | yes | yes | `required-rules.json` + fingerprint test |
| Non-standard SATUSEHAT objects | yes | yes | yes | yes | `src/terminology/nonstandard.go` |
| R4 element guard (opt-in) | yes | yes | yes | yes | `src/conformance/guard.go` |

Module counts differ between languages because a module is a file-layout choice; see `docs/PARITY.md`
for what each number actually contains.

## 5. Public contract

- `src/builder`: 141 fluent FHIR R4 builders plus the SATUSEHAT non-FHIR objects (`BillingStatus`,
  `PurificationDecision`, `Endpoint`).
- `src/datatype`: 33 composite types with `ToArray()`.
- `src/ssrequest`: `OAuth2Client` (`DEV`/`STG`/`PROD`, 50-minute token cache, `Token`, `RefreshToken`,
  `GetByID`, `GetByNIK`, `KfaGet`, `Post`, `Put`) and `SSRequest` (`Get`, `Post`, `Put`, `Delete`,
  `Patch`, `PatchFhirPath`, `PostBundle`, automatic refresh on 401, exponential backoff on 429/5xx).
- `src/ssresponse`: `SSResponse`, `BundleResponse` (`ByMethod`, `Created`, `Updated`, `Patched`,
  `Read`, `Deleted`, `IdMap`, `Find`), `SSValidationError` (rule map, 558 rules).
- `src/parser`: one parser per resource, 142 files.
- `src/queue`: `QueueInterface`, `SqliteQueue` (in-memory `:memory:` or persistent file, no CGo),
  `Worker`, `QueueMonitor`, `ErrorClassifier`, `RateLimiter`, `QueueResult`, `QueueException`.
- `src/terminology`: `TerminologyClient`, `TerminologyQueryBuilder`, `TerminologyResolver`, catalog
  readers, and the non-standard constants.
- `src/conformance`: `UnknownElements`, `AssertR4`, `ElementError`, `StrictR4Enabled`.

## 6. Tests

```bash
cd v4 && go test ./... && go build ./...
```

All packages pass. `go test ./... -v` runs 555 cases (494 test functions plus subtests).
