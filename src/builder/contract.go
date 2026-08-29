package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type ContractBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewContractBuilder() *ContractBuilder {
	b := &ContractBuilder{ResourceType: "Contract", Data: make(map[string]interface{})}
	return b
}

func (b *ContractBuilder) SetId(id string) *ContractBuilder { b.Data["id"] = id; return b }
func (b *ContractBuilder) SetStatus(status string) *ContractBuilder { b.Data["status"] = status; return b }
func (b *ContractBuilder) SetType(type_ *datatype.CodeableConcept) *ContractBuilder { b.Data["type"] = type_.ToArray(); return b }
func (b *ContractBuilder) SetSubject(subjectRef string) *ContractBuilder {
	b.Data["subject"] = []interface{}{map[string]interface{}{"reference": subjectRef}}; return b
}
func (b *ContractBuilder) SetDate(date string) *ContractBuilder { b.Data["date"] = date; return b }
func (b *ContractBuilder) SetAuthority(authRef string) *ContractBuilder {
	b.Data["authority"] = []interface{}{map[string]interface{}{"reference": authRef}}; return b
}
func (b *ContractBuilder) SetDomain(domainRef string) *ContractBuilder {
	b.Data["domain"] = []interface{}{map[string]interface{}{"reference": domainRef}}; return b
}
func (b *ContractBuilder) SetLegalContent(content string) *ContractBuilder {
	b.Data["legal"] = []interface{}{map[string]interface{}{"contentAttachment": map[string]interface{}{"contentString": content}}}
	return b
}
func (b *ContractBuilder) Build() map[string]interface{} { return b.Data }
