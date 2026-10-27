# SRD — satusehat-integration-go (SDK SATUSEHAT FHIR R4 · Go)

Status: v1.0 · 2026-10-27 · Pendamping: `docs/PRD-SDK-ECOSYSTEM.md` (repo PHP) + `v4/README.md`
Modul: `github.com/ivanwilliammd/satusehat-integration-go/v4`, `go 1.21` (`v4/go.mod`).

Seluruh angka di dokumen ini **diukur** dari checkout ini pada 2026-10-27, bukan dari ingatan.

## 1. Ruang Lingkup

Cara kerja internal SDK Go: susunan paket, kontrak publik, selisih nyata terhadap SDK PHP, pemetaan
galat, cara memasang (`go install` dengan titik masuk CLI), dan keadaan uji.

## 2. Prinsip Arsitektur

1. **Modul hidup di subdirektori `v4/`.** Deklarasi modul ada di `v4/go.mod`; CI/release memakai
   `working-directory: v4`. Menaruh berkas Go di akar repo membuat `go install` gagal.
2. **Dua jalur pasang, satu sumber.** Pustaka: `go get …/v4`. CLI: `go install …/v4/cmd/satusehat@vX.Y.Z`
   (titik masuk `v4/cmd/satusehat/main.go`; konstanta versinya wajib sama dengan tag).
3. **Bangun dulu, kirim kemudian.** `src/builder`, `src/datatype` tidak menyentuh jaringan; lapisan
   pengiriman FHIR **belum ada** (lihat §4). Satu-satunya jalur jaringan: `src/terminology`.
4. **Satu perilaku, empat bahasa.** Bentuk payload FHIR harus identik dengan SDK lain pada nomor rilis sama.

## 3. Inventaris Terukur

| Bagian | Jumlah | Cara mengukur |
|---|---|---|
| Berkas builder | 143 | `ls v4/src/builder/*.go \| wc -l` |
| Tipe builder (`…Builder struct`) | 141 | `grep -rhoE "^type [A-Za-z0-9_]+Builder struct" v4/src/builder/*.go \| wc -l` |
| Tipe datatype | 33 | idem pada `v4/src/datatype/*.go` |
| Berkas terminology | 17 | `ls v4/src/terminology/*.go \| wc -l` |
| Berkas queue | 3 | `error_classifier.go`, `monitor.go`, `rate_limiter.go` |
| Berkas uji | 3 | `builders_test.go`, `datatypes_test.go`, `terminology_client_test.go` |
| Fungsi uji (`func Test…`) | 16 | `grep -rhoE "^func Test" v4/src --include="*_test.go" \| wc -l` |
| Berkas `.go` keseluruhan | 192 | `find v4 -name "*.go" \| wc -l` |

## 4. Selisih Nyata terhadap SDK PHP (diukur)

| Kemampuan | PHP | Go | Bukti |
|---|---|---|---|
| Bangun payload FHIR | 141 `PayloadBuilder*` | **141 tipe builder** | `grep -c "^type .*Builder struct"` |
| Nilai bertipe (DataType) | 32 | **33** | `v4/src/datatype` |
| Kirim FHIR + retry/refresh | `SSRequest` | **TIDAK ADA** | `find v4 -iname "*ssrequest*"` → kosong |
| Bungkus jawaban + 558 aturan validasi | `SSResponse`, `SSValidationError` | **TIDAK ADA** | idem |
| Parser FHIR | 18 berkas | **TIDAK ADA** | `find v4 -iname "*parser*"` → kosong |
| Antrean kirim | 8 berkas | 3 penyusun | `ls v4/src/queue` |
| Klien terminology | 22 berkas | 17 berkas | `ls v4/src/terminology` |

**Kesimpulan:** SDK Go adalah **pembangun payload + klien terminology**, bukan SDK pengiriman FHIR.
`SYNC_MANIFEST.json` versi lama menyebut komponen `pending` semua sehingga selisih ini tidak terlihat;
diperbaiki pada commit yang sama.

## 5. Kontrak Publik

- Paket: `src/builder`, `src/datatype`, `src/terminology`, `src/queue` di bawah modul `…/v4`.
- Builder dibuat sebagai nilai (`NewAccountBuilder(...)`) lalu dirangkai; tidak ada `::make()` — sama
  seperti PHP (koreksi klaim README PHP yang menyebut `::make()`).
- `TerminologyClient` menyediakan `Search`, `SearchSystem`, `KptlCombine`, `Map`, `Validate`, `AutoMap`,
  `ResolveCondition`, `ResolveChiefComplaint`, `ResolveObservation`.
- Versi dinaikkan bersama tiga SDK lain; tag Git adalah satu-satunya mekanisme terbit (proxy.golang.org
  mengindeks otomatis, tidak ada langkah publish terpisah).

## 6. Pemetaan Galat (`TerminologyClient`)

- `http.Client` dengan `Timeout: 30 * time.Second`.
- **401** → bila PAT kosong dan kredensial OIDC ada, token diambil lalu permintaan diulang sekali.
- Kegagalan token OIDC → `fmt.Errorf("TerminologyClient OIDC token HTTP %d", status)`.
- Galat HTTP lain diteruskan sebagai `error` dari `doJSON`; tidak ada tipe galat khusus per status.

## 7. Uji

```bash
cd v4 && go test ./...
```

**Keadaan nyata 2026-10-27 — PENGHALANG LINGKUNGAN, bukan kode rusak:** `go version` = go1.24.11,
tetapi `go env GOROOT` = `/usr/lib/golang` dan direktori `src/` di dalamnya **tidak ada**
(`ls /usr/lib/golang/src` kosong). Akibatnya setiap paket gagal sebelum diuji:

```
src/terminology/terminology_client.go:10:2: package strconv is not in std (/usr/lib/golang/src/strconv)
FAIL github.com/ivanwilliammd/satusehat-integration-go/v4/src/terminology [setup failed]
```

Karena itu suite Go **belum pernah lulus/gagal terukur di mesin ini** dan TIDAK boleh diklaim hijau.
Perbaikan lingkungan: pasang distribusi Go lengkap (termasuk pustaka standar) atau setel `GOROOT` ke
instalasi yang punya `src/`. Bukti kode sehat harus datang dari angka `ok …` per paket.

## 8. Utang Teknis & Langkah Berikut

- **T-01** `SYNC_MANIFEST.json` (di `v4/`) menyatakan semua komponen `pending`, total 51 builder;
  nyata 141 tipe builder. Diperbaiki pada commit yang sama.
- **T-02** Tidak ada lapisan pengiriman FHIR maupun Parser. Keputusan pemilik produk: port dari PHP
  atau nyatakan SDK ini "payload builder only".
- **T-03** Uji tidak dapat dijalankan di mesin ini (GOROOT tanpa `src/`) — penghalang lingkungan.
- **T-04** Galat tidak bertipe per status HTTP; pemanggil tidak bisa membedakan 404 vs 401/403.
