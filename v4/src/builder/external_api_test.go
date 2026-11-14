package builder_test

// Uji ini sengaja berada di paket LAIN (builder_test), bukan `builder`.
// Sebelum perbaikan, seluruh metode setter ditulis huruf kecil sehingga hanya bisa
// dipanggil dari dalam paket: uji internal hijau, tetapi pustaka tidak dapat dipakai
// dari luar. Berkas ini menjaga agar hal itu tidak terulang.

import (
	"encoding/json"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestPatientDapatDibangunDariLuarPaket(t *testing.T) {
	patient := builder.NewPatientBuilder().
		SetId("pat-1").
		AddIdentifier(&datatype.Identifier{
			System: "https://fhir.kemkes.go.id/id/nik",
			Value:  "3201234567890123",
		}).
		Build()

	raw, err := json.Marshal(patient)
	if err != nil {
		t.Fatalf("payload tidak dapat diserialisasi: %v", err)
	}
	var hasil map[string]any
	if err := json.Unmarshal(raw, &hasil); err != nil {
		t.Fatalf("JSON tidak sah: %v", err)
	}
	if hasil["resourceType"] != "Patient" {
		t.Fatalf("resourceType = %v, harus Patient", hasil["resourceType"])
	}
	if hasil["id"] != "pat-1" {
		t.Fatalf("id = %v, harus pat-1", hasil["id"])
	}
	ids, ok := hasil["identifier"].([]any)
	if !ok || len(ids) != 1 {
		t.Fatalf("identifier = %v, harus satu elemen", hasil["identifier"])
	}
}

func TestBuilderLainJugaDapatDipanggilDariLuar(t *testing.T) {
	_ = builder.NewEncounterBuilder().SetId("enc-1").SetStatus("finished").Build()
	_ = builder.NewObservationBuilder().SetId("obs-1").SetStatus("final").Build()
	_ = builder.NewBundleBuilder().SetId("bun-1").Build()
}
