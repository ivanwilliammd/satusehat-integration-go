package builder

import "github.com/ivanwilliammd/satusehat-integration-go/src/datatype"

type BiologicallyDerivedProductBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

func NewBiologicallyDerivedProductBuilder() *BiologicallyDerivedProductBuilder {
	b := &BiologicallyDerivedProductBuilder{ResourceType: "BiologicallyDerivedProduct", Data: make(map[string]interface{})}
	return b
}

func (b *BiologicallyDerivedProductBuilder) SetId(id string) *BiologicallyDerivedProductBuilder { b.Data["id"] = id; return b }
func (b *BiologicallyDerivedProductBuilder) SetProductCategory(category string) *BiologicallyDerivedProductBuilder { b.Data["productCategory"] = category; return b }
func (b *BiologicallyDerivedProductBuilder) SetProductCode(code *datatype.CodeableConcept) *BiologicallyDerivedProductBuilder { b.Data["productCode"] = code.ToArray(); return b }
func (b *BiologicallyDerivedProductBuilder) SetStatus(status string) *BiologicallyDerivedProductBuilder { b.Data["status"] = status; return b }
func (b *BiologicallyDerivedProductBuilder) SetExpirationTime(exp string) *BiologicallyDerivedProductBuilder { b.Data["expirationTime"] = exp; return b }
func (b *BiologicallyDerivedProductBuilder) SetCollected(collected string) *BiologicallyDerivedProductBuilder { b.Data["collectedDateTime"] = collected; return b }
func (b *BiologicallyDerivedProductBuilder) SetStorage(handling string) *BiologicallyDerivedProductBuilder {
	b.Data["storage"] = []interface{}{map[string]interface{}{"duration": map[string]interface{}{"text": handling}}}
	return b
}
func (b *BiologicallyDerivedProductBuilder) Build() map[string]interface{} { return b.Data }
