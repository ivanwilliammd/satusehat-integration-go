package builder_test

// GoalBuilder must produce the same payload as PayloadBuilderGoal in the PHP SDK, including the two
// differently shaped target entries.

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/builder"
	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestGoalMatchesPhpBuilderPayload(t *testing.T) {
	isi, err := os.ReadFile(filepath.Join("testdata", "goal_php_golden.json"))
	if err != nil {
		t.Fatalf("golden hilang: %v", err)
	}

	var harapan map[string]interface{}
	if err := json.Unmarshal(isi, &harapan); err != nil {
		t.Fatalf("golden tidak sah: %v", err)
	}

	b := builder.NewGoalBuilder()
	b.SetId("goal-1")
	b.AddIdentifier(&datatype.Identifier{System: "http://sys-ids.kemkes.go.id/goal/1", Value: "goal-1"})
	b.SetLifecycleStatus("active")
	b.SetDescription(concept("http://snomed.info/sct", "386661006", "Fever"))
	b.SetSubject(&datatype.Reference{Reference: "Patient/1000000001", Display: "Budi Santoso"})
	b.SetStartDateTime("2026-01-05")
	b.AddTargetDetailQuantity(&datatype.Quantity{
		Value: ptr(120.0), Unit: str("mmHg"), System: str("http://unitsofmeasure.org"), Code: str("mm[Hg]"),
	})
	b.AddTargetRange(&datatype.Range{
		Low:  &datatype.Quantity{Value: ptr(60.0), Unit: str("bpm")},
		High: &datatype.Quantity{Value: ptr(100.0), Unit: str("bpm")},
	})
	b.SetStatusReason("Patient declined further follow-up")
	b.AddNote("Target tekanan darah < 140/90")
	b.AddOutcomeReference(&datatype.Reference{Reference: "Observation/3000000001"})

	mentah, err := b.BuildJSON()
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}

	var dapat map[string]interface{}
	if err := json.Unmarshal(mentah, &dapat); err != nil {
		t.Fatalf("hasil bukan JSON sah: %v", err)
	}

	if !reflect.DeepEqual(harapan, dapat) {
		h, _ := json.MarshalIndent(harapan, "", "  ")
		d, _ := json.MarshalIndent(dapat, "", "  ")
		t.Fatalf("payload Go berbeda dari PHP\n--- PHP ---\n%s\n--- Go ---\n%s", h, d)
	}
}
