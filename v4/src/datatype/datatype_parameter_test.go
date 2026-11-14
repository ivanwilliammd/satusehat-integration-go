package datatype_test

import (
	"encoding/json"
	"testing"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

func TestParameterComponentUsesValueX(t *testing.T) {
	param := datatype.ParameterComponent{Name: "subject", ValueKey: "valueString", Value: "Patient/1"}
	raw, err := json.Marshal(param.ToArray())
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var hasil map[string]any
	if err := json.Unmarshal(raw, &hasil); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if hasil["valueString"] != "Patient/1" {
		t.Fatalf("valueString = %v, want Patient/1", hasil["valueString"])
	}
	if _, ada := hasil["value"]; ada {
		t.Fatalf("bare value key must not be emitted: %v", hasil)
	}
}

func TestSampledDataRoundTrip(t *testing.T) {
	period := 0.5
	sd := datatype.SampledData{Period: &period, Dimensions: intPtr(1), Data: stringPtr("0,1,2,3")}
	m := sd.ToArray()
	if m["period"] != 0.5 || m["dimensions"] != 1 {
		t.Fatalf("unexpected payload: %v", m)
	}
}

func intPtr(v int) *int          { return &v }
func stringPtr(v string) *string { return &v }
