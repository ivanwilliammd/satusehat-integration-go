package builder_test

// Task has the largest surface of the eleven builders that had no golden: twelve setters, five of them
// nested. The golden carries every one of them at once.

import (
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestTaskMatchesPHPPayload(t *testing.T) {
	b := builder.NewTaskBuilder()
	b.SetId("task-1")
	b.AddIdentifier("http://sys.id/ni", "NI-70")
	b.SetStatus("requested")
	b.SetIntent("order")
	b.SetPriority("urgent")
	b.SetDescription("Ambil sampel darah")
	b.SetBusinessStatus("in-progress", "In progress", "http://terminology.hl7.org/CodeSystem/task-business-status")
	b.SetStatusReason("waiting", "Menunggu pasien", "http://terminology.hl7.org/CodeSystem/task-status-reason")
	b.SetCode("LAB-1", "Laboratorium", "http://terminology.kemkes.go.id/CodeSystem/task-code")
	b.SetFocus("ServiceRequest/sr-1", "Darah Lengkap")
	b.SetFor("Patient/1000000001", "Budi Santoso")
	b.SetEncounter("Encounter/enc-1", "Kunjungan")
	b.SetRequester("Practitioner/1000000002", "dr. Budi")
	b.SetOwner("Practitioner/1000000003", "Perawat Ani")
	b.SetLocation("Location/3000000001", "Lab")
	b.SetExecutionPeriod("2026-01-05T08:00:00+07:00", "2026-01-05T10:00:00+07:00")
	b.SetAuthoredOn("2026-01-05T07:00:00+07:00")
	b.SetLastModified("2026-01-05T08:30:00+07:00")
	b.SetReasonCode(konsep("http://snomed.info/sct", "386053000", "Evaluation procedure"))
	b.SetReasonReference(&datatype.Reference{Reference: "Observation/obs-1"})
	b.AddInput("", "Darah", konsep("http://terminology.kemkes.go.id/CodeSystem/input", "IN-1", "Jenis sampel"))
	b.AddInput("Catatan", "Puasa 8 jam")
	b.AddOutput("", "Normal", konsep("http://terminology.kemkes.go.id/CodeSystem/output", "OUT-1", "Hasil"))
	b.AddRestrictionPeriod("Practitioner/1000000002", &datatype.Period{Start: ptr("2026-01-05"), End: ptr("2026-01-06")}, 2)
	b.AddNote("Sudah dikonfirmasi pasien")
	b.AddExtension("http://example.org/ext", "nilai")
	samakan(t, bacaGolden(t, "task_full"), b)
}

func TestTaskRefusesUnknownStatusAndIntent(t *testing.T) {
	b := builder.NewTaskBuilder()
	b.SetStatus("invalid")
	if b.Err() == nil {
		t.Fatal("status invalid diterima")
	}
	c := builder.NewTaskBuilder()
	c.SetIntent("invalid")
	if c.Err() == nil {
		t.Fatal("intent invalid diterima")
	}
}
