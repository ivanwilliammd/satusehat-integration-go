package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type SupplyRequestBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSupplyRequestBuilder() *SupplyRequestBuilder {
	b := &SupplyRequestBuilder{ResourceType: "SupplyRequest", Data: make(map[string]interface{})}
	return b
}

func (b *SupplyRequestBuilder) SetId(id string) *SupplyRequestBuilder { b.Data["id"] = id; return b }
func (b *SupplyRequestBuilder) SetStatus(status string) *SupplyRequestBuilder { b.Data["status"] = status; return b }
func (b *SupplyRequestBuilder) SetCategory(category string) *SupplyRequestBuilder { b.Data["category"] = category; return b }
func (b *SupplyRequestBuilder) SetPriority(priority string) *SupplyRequestBuilder { b.Data["priority"] = priority; return b }
func (b *SupplyRequestBuilder) SetItem(item *datatype.CodeableConcept) *SupplyRequestBuilder { b.Data["itemCodeableConcept"] = item.ToArray(); return b }
func (b *SupplyRequestBuilder) SetSubject(subjectRef string) *SupplyRequestBuilder { b.Data["subject"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *SupplyRequestBuilder) SetEncounter(encRef string) *SupplyRequestBuilder { b.Data["encounter"] = map[string]interface{}{"reference": encRef}; return b }
func (b *SupplyRequestBuilder) SetOccurrence(occurrence string) *SupplyRequestBuilder { b.Data["occurrenceDateTime"] = occurrence; return b }
func (b *SupplyRequestBuilder) SetRequester(requesterRef string) *SupplyRequestBuilder { b.Data["requester"] = map[string]interface{}{"reference": requesterRef}; return b }
func (b *SupplyRequestBuilder) SetSupplier(supplierRef string) *SupplyRequestBuilder {
	b.Data["supplier"] = []interface{}{map[string]interface{}{"reference": supplierRef}}; return b
}
func (b *SupplyRequestBuilder) Build() map[string]interface{} { return b.Data }
