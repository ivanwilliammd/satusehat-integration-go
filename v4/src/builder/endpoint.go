package builder

type EndpointBuilder struct {
	data map[string]interface{}
}

func NewEndpointBuilder() *EndpointBuilder {
	return &EndpointBuilder{
		data: map[string]interface{}{"resourceType": "Endpoint"},
	}
}

func (b *EndpointBuilder) SetID(id string) *EndpointBuilder {
	b.data["id"] = id
	return b
}

func (b *EndpointBuilder) AddIdentifier(system, value string) *EndpointBuilder {
	ids, _ := b.data["identifier"].([]map[string]string)
	ids = append(ids, map[string]string{"system": system, "value": value})
	b.data["identifier"] = ids
	return b
}

func (b *EndpointBuilder) SetStatus(status string) *EndpointBuilder {
	valid := map[string]bool{"active": true, "suspended": true, "error": true, "off": true, "entered-in-error": true, "test": true}
	if !valid[status] {
		panic("Invalid Endpoint status: " + status)
	}
	b.data["status"] = status
	return b
}

func (b *EndpointBuilder) SetConnectionType(code, display string, system ...string) *EndpointBuilder {
	sys := "http://terminology.hl7.org/CodeSystem/endpoint-connection-type"
	if len(system) > 0 && system[0] != "" {
		sys = system[0]
	}
	b.data["connectionType"] = map[string]interface{}{
		"coding": []map[string]string{
			{"code": code, "display": display, "system": sys},
		},
	}
	return b
}

func (b *EndpointBuilder) SetName(name string) *EndpointBuilder {
	b.data["name"] = name
	return b
}

func (b *EndpointBuilder) SetManagingOrganization(ref string, display ...string) *EndpointBuilder {
	if !hasPrefix(ref) && !containsSlash(ref) {
		ref = "Organization/" + ref
	}
	res := map[string]interface{}{"reference": ref}
	if len(display) > 0 && display[0] != "" {
		res["display"] = display[0]
	}
	b.data["managingOrganization"] = res
	return b
}

func (b *EndpointBuilder) AddContact(system, value string, use ...string) *EndpointBuilder {
	contacts, _ := b.data["contact"].([]map[string]string)
	c := map[string]string{"system": system, "value": value}
	if len(use) > 0 && use[0] != "" {
		c["use"] = use[0]
	}
	contacts = append(contacts, c)
	b.data["contact"] = contacts
	return b
}

func (b *EndpointBuilder) SetPeriod(start string, end ...string) *EndpointBuilder {
	p := map[string]string{"start": start}
	if len(end) > 0 && end[0] != "" {
		p["end"] = end[0]
	}
	b.data["period"] = p
	return b
}

func (b *EndpointBuilder) AddPayloadType(code, display string, system ...string) *EndpointBuilder {
	sys := "http://terminology.hl7.org/CodeSystem/endpoint-payload-type"
	if len(system) > 0 && system[0] != "" {
		sys = system[0]
	}
	pts, _ := b.data["payloadType"].([]map[string]interface{})
	pts = append(pts, map[string]interface{}{
		"coding": []map[string]string{
			{"code": code, "display": display, "system": sys},
		},
	})
	b.data["payloadType"] = pts
	return b
}

func (b *EndpointBuilder) AddPayloadMimeType(mime string) *EndpointBuilder {
	mimes, _ := b.data["payloadMimeType"].([]string)
	mimes = append(mimes, mime)
	b.data["payloadMimeType"] = mimes
	return b
}

func (b *EndpointBuilder) SetAddress(addr string) *EndpointBuilder {
	b.data["address"] = addr
	return b
}

func (b *EndpointBuilder) AddHeader(header string) *EndpointBuilder {
	headers, _ := b.data["header"].([]string)
	headers = append(headers, header)
	b.data["header"] = headers
	return b
}

func (b *EndpointBuilder) Build() map[string]interface{} {
	clean := make(map[string]interface{})
	for k, v := range b.data {
		if v != nil {
			clean[k] = v
		}
	}
	return clean
}
