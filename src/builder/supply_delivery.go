package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type SupplyDeliveryBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewSupplyDeliveryBuilder() *SupplyDeliveryBuilder {
	b := &SupplyDeliveryBuilder{ResourceType: "SupplyDelivery", Data: make(map[string]interface{})}
	return b
}

func (b *SupplyDeliveryBuilder) SetId(id string) *SupplyDeliveryBuilder { b.Data["id"] = id; return b }
func (b *SupplyDeliveryBuilder) SetStatus(status string) *SupplyDeliveryBuilder { b.Data["status"] = status; return b }
func (b *SupplyDeliveryBuilder) SetType(type_ *datatype.CodeableConcept) *SupplyDeliveryBuilder { b.Data["type"] = type_.ToArray(); return b }
func (b *SupplyDeliveryBuilder) SetSubject(subjectRef string) *SupplyDeliveryBuilder { b.Data["patient"] = map[string]interface{}{"reference": subjectRef}; return b }
func (b *SupplyDeliveryBuilder) SetEncounter(encRef string) *SupplyDeliveryBuilder { b.Data["encounter"] = map[string]interface{}{"reference": encRef}; return b }
func (b *SupplyDeliveryBuilder) SetOccurrence(occurrence string) *SupplyDeliveryBuilder { b.Data["occurrenceDateTime"] = occurrence; return b }
func (b *SupplyDeliveryBuilder) SetSupplier(supplierRef string) *SupplyDeliveryBuilder { b.Data["supplier"] = map[string]interface{}{"reference": supplierRef}; return b }
func (b *SupplyDeliveryBuilder) Build() map[string]interface{} { return b.Data }
