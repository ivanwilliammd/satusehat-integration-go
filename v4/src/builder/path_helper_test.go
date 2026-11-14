package builder

// A numeric segment in a path points at an existing list element, exactly as the PHP builder's set()
// does. Without that rule the payload gets a literal map key "1" where FHIR expects a list, which is
// valid JSON and completely wrong on the wire.

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSetPathIndexesIntoTheList(t *testing.T) {
	data := map[string]interface{}{}
	pushPath(data, "item", map[string]interface{}{"sequence": 1})
	setPath(data, "item/0/quantity/value", 10.0)

	harapan := map[string]interface{}{
		"item": []interface{}{
			map[string]interface{}{
				"sequence": 1,
				"quantity": map[string]interface{}{"value": 10.0},
			},
		},
	}
	h, _ := json.Marshal(harapan)
	d, _ := json.Marshal(data)
	var hn, dn interface{}
	_ = json.Unmarshal(h, &hn)
	_ = json.Unmarshal(d, &dn)
	if !reflect.DeepEqual(hn, dn) {
		t.Fatalf("jalur bernomor tidak masuk ke elemen daftar\n--- harapan ---\n%s\n--- dapat ---\n%s", h, d)
	}
}

func TestSetPathOnAMissingIndexIsRefused(t *testing.T) {
	data := map[string]interface{}{}
	pushPath(data, "item", map[string]interface{}{"sequence": 1})
	setPath(data, "item/3/quantity/value", 10.0)

	item, _ := data["item"].([]interface{})
	pertama, _ := item[0].(map[string]interface{})
	if _, ada := pertama["quantity"]; ada {
		t.Fatal("indeks yang tidak menunjuk elemen mana pun tidak boleh menulis apa pun")
	}
	if _, ada := data["item/3/quantity/value"]; ada {
		t.Fatal("jalur tidak boleh tersimpan sebagai kunci literal")
	}
}
