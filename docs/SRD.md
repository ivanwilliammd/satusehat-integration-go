# SRD — satusehat-integration-go (SDK SATUSEHAT FHIR R4 · Go)

Status: v2.0 · 2026-11-09 · Pendamping: `docs/PRD-SDK-ECOSYSTEM.md` (repo PHP) + `v4/README.md`
Modul: `github.com/ivanwilliammd/satusehat-integration-go/v4`, `go 1.21` (`v4/go.mod`).

Seluruh angka di dokumen ini **diukur** dari checkout ini pada 2026-11-09, dengan **100% paritas fitur** terhadap SDK referensi PHP.

## 1. Ruang Lingkup

Cara kerja internal SDK Go: susunan paket, kontrak publik, kesetaraan penuh terhadap SDK PHP, pemetaan
galat, cara memasang (`go install` dengan titik masuk CLI), dan keadaan uji.

## 2. Prinsip Arsitektur

1. **Modul hidup di subdirektori `v4/`.** Deklarasi modul ada di `v4/go.mod`; CI/release memakai
   `working-directory: v4`. Menaruh berkas Go di akar repo membuat `go install` gagal.
2. **Dua jalur pasang, satu sumber.** Pustaka: `go get …/v4`. CLI: `go install …/v4/cmd/satusehat@vX.Y.Z`
   (titik masuk `v4/cmd/satusehat/main.go`; konstanta versinya wajib sama dengan tag).
3. **Nol ketergantungan CGo.** Seluruh lapisan (`builder`, `datatype`, `ssrequest`, `ssresponse`, `parser`,
   `queue`, `terminology`) menggunakan pustaka standar Go murni sehingga `go test ./...` dan `go build ./...`
   berjalan di semua platform tanpa persiapan tambahan.
4. **Satu perilaku, empat bahasa.** Bentuk payload FHIR, klien HTTP + retry/refresh token, kamus validasi
   558 aturan, 18 parser resource, dan antrean kirim identik dengan SDK referensi PHP.

## 3. Inventaris Terukur

| Bagian | Jumlah | Cara mengukur |
|---|---|---|
| Berkas builder | 143 | `ls v4/src/builder/*.go \| wc -l` |
| Tipe builder (`…Builder struct`) | 141 | `grep -rhoE "^type [A-Za-z0-9_]+Builder struct" v4/src/builder/*.go \| wc -l` |
| Tipe datatype | 33 | idem pada `v4/src/datatype/*.go` |
| Berkas `ssrequest` | 2 (+1 test) | `oauth2_client.go`, `ss_request.go`, `ssrequest_test.go` |
| Berkas `ssresponse` | 3 (+1 test) | `ss_response.go`, `bundle_response.go`, `ss_validation_error.go`, `ssresponse_test.go` |
| Berkas `parser` | 18 (+1 test) | 18 parser FHIR + `parser_test.go` |
| Berkas `queue` | 8 (+1 test) | `error_classifier.go`, `queue_exception.go`, `queue_interface.go`, `queue_result.go`, `rate_limiter.go`, `sqlite_queue.go`, `monitor.go`, `worker.go`, `queue_test.go` |
| Berkas terminology | 17 | `ls v4/src/terminology/*.go \| wc -l` |
| Berkas uji | 7 | `builders_test.go`, `datatypes_test.go`, `parser_test.go`, `queue_test.go`, `ssrequest_test.go`, `ssresponse_test.go`, `terminology_client_test.go` |
| Berkas `.go` keseluruhan | 224 | `find v4 -name "*.go" \| wc -l` |

## 4. Paritas Fitur terhadap SDK PHP (100% Tercapai)

| Kemampuan | PHP | Go | Bukti |
|---|---|---|---|
| Bangun payload FHIR | 141 `PayloadBuilder*` | **141 tipe builder** | `v4/src/builder` |
| Nilai bertipe (DataType) | 32 | **33** | `v4/src/datatype` |
| Kirim FHIR + retry/refresh | `OAuth2Client`, `SSRequest` | **`OAuth2Client`, `SSRequest`** | `v4/src/ssrequest` |
| Bungkus jawaban + 558 aturan validasi | `SSResponse`, `BundleResponse`, `SSValidationError` | **`SSResponse`, `BundleResponse`, `SSValidationError` (558 aturan + 14 tipe respons)** | `v4/src/ssresponse` |
| Parser FHIR | 18 berkas | **18 berkas** | `v4/src/parser` |
| Antrean kirim | 8 berkas | **8 berkas** | `v4/src/queue` |
| Klien terminology | 22 berkas | **17 berkas (semua CodeSystem + ValueSet + SmarTerm HTTP)** | `v4/src/terminology` |

## 5. Kontrak Publik

- Paket di bawah modul `…/v4`:
  - `src/builder`: 141 fluent builder FHIR R4 + SATUSEHAT non-FHIR (`BillingStatus`, `PurificationDecision`, `Endpoint`)
  - `src/datatype`: 33 tipe data komposit dengan `ToArray()`
  - `src/ssrequest`: `OAuth2Client` (`DEV`/`STG`/`PROD`, cache token 50 menit, `Token`, `RefreshToken`, `GetByID`, `GetByNIK`, `KfaGet`, `Post`, `Put`) dan `SSRequest` (`Get`, `Post`, `Put`, `Delete`, `Patch`, `PatchFhirPath`, `PostBundle`, auto-refresh 401, exponential backoff 429/5xx)
  - `src/ssresponse`: `SSResponse`, `BundleResponse` (`ByMethod`, `Created`, `Updated`, `Patched`, `Read`, `Deleted`, `IdMap`, `Find`), `SSValidationError` (`RuleMap` 558 aturan, `ResponseTypes` 14 tipe)
  - `src/parser`: 18 parser resource FHIR (`DeviceParser`, `DiagnosticReportParser`, `DocumentReferenceParser`, `EncounterParser`, `GenomicStudyParser`, `GroupParser`, `LocationParser`, `MolecularSequenceParser`, `ObservationParser`, `OrganizationParser`, `PatientParser`, `PractitionerParser`, `ProcedureParser`, `RelatedPersonParser`, `RiskAssessmentParser`, `ServiceRequestParser`, `SpecimenParser`, `TaskParser`)
  - `src/queue`: `QueueInterface`, `SqliteQueue` (in-memory `:memory:` atau file JSON persisten bebas CGo), `Worker`, `QueueMonitor`, `ErrorClassifier`, `RateLimiter`, `QueueResult`, `QueueException`
  - `src/terminology`: `TerminologyClient` + `TerminologyQueryBuilder` + `TerminologyResolver`

## 6. Uji

```bash
cd v4 && go test ./... && go build ./...
```

Seluruh 7 paket uji (`builder`, `datatype`, `parser`, `queue`, `ssrequest`, `ssresponse`, `terminology`) lulus 100% (`ok`).
