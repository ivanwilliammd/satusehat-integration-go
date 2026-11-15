package builder_test

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestImagingStudyMatchesPHPPayload(t *testing.T) {
	status, deskripsi := "available", "Abdomen"
	nomor := 1

	b := builder.NewImagingStudyBuilder()
	b.SetId("img-1")
	b.SetStatus("available")
	b.SetSubject("Patient/1000000001")
	b.SetStarted("2026-01-05T08:00:00+07:00")
	b.SetEnded("2026-01-05T09:00:00+07:00")
	b.SetLocation("Location/3000000001")
	b.AddReferrer("Practitioner/1000000002")
	b.AddPartOf("ImagingStudy/img-0")
	b.AddModality("http://dicom.nema.org/resources/ontology/DCM", "US", "Ultrasound")
	b.AddSeries("1.2.3", "http://dicom.nema.org/resources/ontology/DCM", "US", "Ultrasound", &status, &deskripsi, &nomor)
	samakan(t, bacaGolden(t, "imaging_study"), b)
}

func TestMolecularSequenceMatchesPHPPayload(t *testing.T) {
	strand, orientasi := "watson", "sense"
	mulai, akhir := 100, 200

	b := builder.NewMolecularSequenceBuilder()
	b.SetId("ms-1")
	b.SetPatient(&datatype.Reference{Reference: "Patient/1000000001"})
	b.SetCoordinateSystem(0)
	b.SetSpecimen(&datatype.Reference{Reference: "Specimen/sp-1"})
	b.SetDevice(&datatype.Reference{Reference: "Device/dev-1"})
	b.SetPerformer(&datatype.Reference{Reference: "Practitioner/1000000002"})
	b.SetReferenceSeq(konsep("http://www.ncbi.nlm.nih.gov/nuccore", "NC_000001.11", "chr1"), &strand, &mulai, &akhir, &orientasi)
	b.AddVariant(100, 101, "A", "G", &datatype.Reference{Reference: "Observation/obs-1"})
	b.SetText("generated", "<div>Urutan</div>")

	if err := b.Err(); err != nil {
		t.Fatalf("coordinateSystem sudah diisi, galat tidak diharapkan: %v", err)
	}
	samakan(t, bacaGolden(t, "molecular_sequence"), b)
}

func TestMolecularSequenceWithoutCoordinateSystemIsAnError(t *testing.T) {
	b := builder.NewMolecularSequenceBuilder()
	b.SetId("ms-2")
	_ = b.Build()

	if b.Err() == nil {
		t.Fatal("coordinateSystem kosong seharusnya dicatat sebagai galat, seperti PHP yang throw")
	}
}

func TestPaymentNoticeMatchesPHPPayload(t *testing.T) {
	b := builder.NewPaymentNoticeBuilder()
	b.SetId("pn-1")
	b.SetStatus("active")
	b.SetPayee("Organization/1000000001")
	b.SetResponse("ClaimResponse/cr-1")
	b.SetPayment(250000, "IDR")
	b.SetPaymentDate("2026-01-10")
	b.SetPaymentStatus("http://terminology.hl7.org/CodeSystem/paymentstatus", "paid", "Paid")
	b.SetPaymentReconciliation("PaymentReconciliation/pr-1")
	b.SetStatusDate("2026-01-10")
	samakan(t, bacaGolden(t, "payment_notice"), b)
}
