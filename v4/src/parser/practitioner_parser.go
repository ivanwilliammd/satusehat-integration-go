package parser

import "github.com/ivanwilliammd/satusehat-integration-go/v4/src/ssrequest"

// PractitionerParser extracts fields from a FHIR Practitioner resource map.
type PractitionerParser struct {
	practitioner map[string]any
	oauth2       *ssrequest.OAuth2Client
}

// NewPractitionerParser creates a new PractitionerParser instance.
func NewPractitionerParser(practitioner map[string]any, oauth2 ...*ssrequest.OAuth2Client) *PractitionerParser {
	var client *ssrequest.OAuth2Client
	if len(oauth2) > 0 {
		client = oauth2[0]
	}
	return &PractitionerParser{
		practitioner: practitioner,
		oauth2:       client,
	}
}

// GetSSNik fetches a Practitioner resource by NIK using OAuth2Client and updates the parser's state.
func (p *PractitionerParser) GetSSNik(nik string, oauth2 ...*ssrequest.OAuth2Client) map[string]any {
	client := p.oauth2
	if len(oauth2) > 0 && oauth2[0] != nil {
		client = oauth2[0]
	}
	if client == nil {
		var err error
		client, err = ssrequest.NewOAuth2Client()
		if err != nil {
			return nil
		}
	}

	statusCode, res, err := client.GetByNIK("Practitioner", nik)
	if err != nil || statusCode != 200 || res == nil {
		return nil
	}

	entries := sliceMaps(res["entry"])
	if len(entries) == 0 {
		p.practitioner = nil
		return nil
	}

	p.practitioner = mapVal(entries[0]["resource"])
	return p.practitioner
}

func (p *PractitionerParser) GetId() string {
	if p.practitioner == nil {
		return ""
	}
	return strVal(p.practitioner["id"])
}

func (p *PractitionerParser) GetIdentifiers() any {
	if p.practitioner == nil {
		return nil
	}
	return p.practitioner["identifier"]
}

func (p *PractitionerParser) GetActive() any {
	if p.practitioner == nil {
		return nil
	}
	return p.practitioner["active"]
}

func (p *PractitionerParser) GetName() any {
	if p.practitioner == nil {
		return nil
	}
	return p.practitioner["name"]
}

func (p *PractitionerParser) GetTelecom() any {
	if p.practitioner == nil {
		return nil
	}
	return p.practitioner["telecom"]
}

func (p *PractitionerParser) GetAddress() any {
	if p.practitioner == nil {
		return nil
	}
	return p.practitioner["address"]
}

func (p *PractitionerParser) GetAddressLine() string {
	if p.practitioner == nil {
		return ""
	}
	addrs := sliceMaps(p.practitioner["address"])
	if len(addrs) == 0 {
		return ""
	}
	switch lines := addrs[0]["line"].(type) {
	case []string:
		if len(lines) > 0 {
			return lines[0]
		}
	case []any:
		if len(lines) > 0 {
			return strVal(lines[0])
		}
	}
	return ""
}

func (p *PractitionerParser) GetCity() string {
	return p.extractAdminExtension(1)
}

func (p *PractitionerParser) GetVillage() string {
	return p.extractAdminExtension(3)
}

func (p *PractitionerParser) extractAdminExtension(idx int) string {
	if p.practitioner == nil {
		return ""
	}
	addrs := sliceMaps(p.practitioner["address"])
	if len(addrs) == 0 {
		return ""
	}
	exts := sliceMaps(addrs[0]["extension"])
	if len(exts) == 0 {
		return ""
	}
	innerExts := sliceMaps(exts[0]["extension"])
	if idx < 0 || idx >= len(innerExts) {
		return ""
	}
	return strVal(innerExts[idx]["valueCode"])
}

func (p *PractitionerParser) GetGender() string {
	if p.practitioner == nil {
		return ""
	}
	return strVal(p.practitioner["gender"])
}

func (p *PractitionerParser) GetBirthDate() string {
	if p.practitioner == nil {
		return ""
	}
	return strVal(p.practitioner["birthDate"])
}

func (p *PractitionerParser) GetPhoto() any {
	if p.practitioner == nil {
		return nil
	}
	return p.practitioner["photo"]
}

func (p *PractitionerParser) GetQualifications() string {
	if p.practitioner == nil {
		return ""
	}
	quals := sliceMaps(p.practitioner["qualification"])
	if len(quals) == 0 {
		return ""
	}
	ids := sliceMaps(quals[0]["identifier"])
	if len(ids) == 0 {
		return ""
	}
	return strVal(ids[0]["value"])
}

func (p *PractitionerParser) GetCommunication() any {
	if p.practitioner == nil {
		return nil
	}
	return p.practitioner["communication"]
}
