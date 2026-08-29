package builder

import (
	"strings"

	"github.com/ivanwilliammd/satusehat-integration-go/src/datatype"
)

func addExtensionTo(data *map[string]interface{}, url string, value interface{}, valueType string) {
	ext := map[string]interface{}{"url": url}
	if valueType != "" {
		capitalized := strings.ToUpper(valueType[:1]) + valueType[1:]
		ext["value"+capitalized] = value
	} else {
		ext["valueString"] = value
	}
	if _, ok := (*data)["extension"]; !ok {
		(*data)["extension"] = make([]interface{}, 0)
	}
	(*data)["extension"] = append((*data)["extension"].([]interface{}), ext)
}

func ensureArray(data *map[string]interface{}, key string) {
	if _, ok := (*data)[key]; !ok {
		(*data)[key] = make([]interface{}, 0)
	}
}

func appendToArray(data *map[string]interface{}, key string, item interface{}) {
	ensureArray(data, key)
	(*data)[key] = append((*data)[key].([]interface{}), item)
}

func makeCodeableConcept(system, code, display string) *datatype.CodeableConcept {
	cc := &datatype.CodeableConcept{Text: display}
	if system != "" || code != "" || display != "" {
		cc.Coding = []datatype.Coding{{System: system, Code: code, Display: display}}
	}
	return cc
}

func makeReference(reference, display string) *datatype.Reference {
	r := &datatype.Reference{}
	if reference != "" {
		r.Reference = reference
	}
	if display != "" {
		r.Display = display
	}
	return r
}

func makeQuantity(value float64, currency string) *datatype.Quantity {
	q := &datatype.Quantity{}
	if value != 0 {
		q.Value = &value
	}
	if currency != "" {
		q.Code = &currency
	}
	return q
}

func makePeriod(start, end string) *datatype.Period {
	p := &datatype.Period{}
	if start != "" {
		p.Start = &start
	}
	if end != "" {
		p.End = &end
	}
	return p
}
