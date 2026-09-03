package builder

import (
	"encoding/json"
)

type RelatedPersonBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewRelatedPersonBuilder() *RelatedPersonBuilder {
	b := &RelatedPersonBuilder{ResourceType: "RelatedPerson", Data: make(map[string]interface{})}
	b.Data["resourceType"] = "RelatedPerson"
	return b
}

func (b *RelatedPersonBuilder) setId(id string) *RelatedPersonBuilder {
	b.Data["id"] = id
	return b
}

func (b *RelatedPersonBuilder) addIdentifier(system string, value ...string) *RelatedPersonBuilder {
	ids, _ := b.Data["identifier"].([]map[string]interface{})
	entry := map[string]interface{}{"system": system}
	if len(value) > 0 {
		entry["value"] = value[0]
	}
	ids = append(ids, entry)
	b.Data["identifier"] = ids
	return b
}

func (b *RelatedPersonBuilder) setActive(active bool) *RelatedPersonBuilder {
	b.Data["active"] = active
	return b
}

func (b *RelatedPersonBuilder) setPatient(reference string, display ...string) *RelatedPersonBuilder {
	ref := reference
	if !hasPrefix(ref) && !containsSlash(ref) {
		ref = "Patient/" + ref
	}
	patient := map[string]interface{}{"reference": ref}
	if len(display) > 0 && display[0] != "" {
		patient["display"] = display[0]
	}
	b.Data["patient"] = patient
	return b
}

func (b *RelatedPersonBuilder) addRelationship(code string, display ...string) *RelatedPersonBuilder {
	rels, _ := b.Data["relationship"].([]map[string]interface{})
	disp := code
	if len(display) > 0 && display[0] != "" {
		disp = display[0]
	}
	rels = append(rels, map[string]interface{}{
		"coding": []map[string]interface{}{{
			"system":  "http://terminology.hl7.org/CodeSystem/v2-0131",
			"code":    code,
			"display": disp,
		}},
	})
	b.Data["relationship"] = rels
	return b
}

func (b *RelatedPersonBuilder) addName(text string) *RelatedPersonBuilder {
	names, _ := b.Data["name"].([]map[string]interface{})
	names = append(names, map[string]interface{}{"text": text})
	b.Data["name"] = names
	return b
}

func (b *RelatedPersonBuilder) addTelecom(system, value string, use ...string) *RelatedPersonBuilder {
	tels, _ := b.Data["telecom"].([]map[string]interface{})
	u := "home"
	if len(use) > 0 && use[0] != "" {
		u = use[0]
	}
	tels = append(tels, map[string]interface{}{"system": system, "value": value, "use": u})
	b.Data["telecom"] = tels
	return b
}

func (b *RelatedPersonBuilder) setGender(gender string) *RelatedPersonBuilder {
	b.Data["gender"] = gender
	return b
}

func (b *RelatedPersonBuilder) setBirthDate(birthDate string) *RelatedPersonBuilder {
	b.Data["birthDate"] = birthDate
	return b
}

func (b *RelatedPersonBuilder) addAddress(address map[string]interface{}) *RelatedPersonBuilder {
	addrs, _ := b.Data["address"].([]map[string]interface{})
	addrs = append(addrs, address)
	b.Data["address"] = addrs
	return b
}

func (b *RelatedPersonBuilder) addCommunication(languageCode string, preferred ...bool) *RelatedPersonBuilder {
	comms, _ := b.Data["communication"].([]map[string]interface{})
	p := true
	if len(preferred) > 0 {
		p = preferred[0]
	}
	comms = append(comms, map[string]interface{}{
		"language": map[string]interface{}{
			"coding": []map[string]interface{}{{
				"system":  "http://terminology.hl7.org/CodeSystem/v3-Language",
				"code":    languageCode,
				"display": languageCode,
			}},
		},
		"preferred": p,
	})
	b.Data["communication"] = comms
	return b
}

func (b *RelatedPersonBuilder) addExtension(url string, value interface{}) *RelatedPersonBuilder {
	extensions, _ := b.Data["extension"].([]map[string]interface{})
	ext := map[string]interface{}{"url": url}
	switch v := value.(type) {
	case bool:
		ext["valueBoolean"] = v
	case string:
		ext["valueString"] = v
	case int:
		ext["valueInteger"] = v
	case map[string]interface{}:
		for k, val := range v {
			ext[k] = val
		}
	}
	extensions = append(extensions, ext)
	b.Data["extension"] = extensions
	return b
}

func (b *RelatedPersonBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.Data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}

func (b *RelatedPersonBuilder) BuildJSON() ([]byte, error) {
	return json.Marshal(b.Build())
}
