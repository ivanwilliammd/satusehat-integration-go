package builder

import (
	"encoding/json"

	"github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

// CoverageBuilder builds a FHIR R4 Coverage resource.
//
// Counterpart of PayloadBuilderCoverage in the PHP SDK, with one structural difference: PHP takes
// bare strings for references and coding pairs for concepts, while this port takes the datatypes and
// writes the same JSON. Callers of the Go SDK therefore pass &datatype.Reference{Reference: "…"}
// where PHP passes "…".
//
// Two PHP quirks are reproduced on purpose:
//   - SetType always writes display, even when the caller passes an empty string;
//   - AddCostToBeneficiary writes value and currency directly instead of a valueMoney object.
//
// @link https://www.hl7.org/fhir/coverage.html
type CoverageBuilder struct {
	ResourceType string
	Data         map[string]interface{}
}

// NewCoverageBuilder creates an empty Coverage with its resourceType set.
func NewCoverageBuilder() *CoverageBuilder {
	b := &CoverageBuilder{ResourceType: "Coverage", Data: map[string]interface{}{}}
	b.Data["resourceType"] = b.ResourceType
	return b
}

func (b *CoverageBuilder) set(k string, v interface{}) *CoverageBuilder {
	setPath(b.Data, k, v)
	return b
}

func (b *CoverageBuilder) push(k string, v interface{}) *CoverageBuilder {
	pushPath(b.Data, k, v)
	return b
}

// SetId sets the logical id.
func (b *CoverageBuilder) SetId(id string) *CoverageBuilder { return b.set("id", id) }

// AddIdentifier appends an identifier.
func (b *CoverageBuilder) AddIdentifier(id *datatype.Identifier) *CoverageBuilder {
	return b.push("identifier", id.ToArray())
}

// SetStatus sets status (active | cancelled | draft | entered-in-error).
func (b *CoverageBuilder) SetStatus(status string) *CoverageBuilder { return b.set("status", status) }

// SetType sets type, writing display even when empty, as PHP does. A concept without display would
// otherwise produce a payload that differs from the PHP output by one empty string.
func (b *CoverageBuilder) SetType(typ *datatype.CodeableConcept) *CoverageBuilder {
	m := typ.ToArray()
	if coding, ok := m["coding"].([]interface{}); ok {
		for _, c := range coding {
			if satu, ok := c.(map[string]interface{}); ok {
				if _, ada := satu["display"]; !ada {
					satu["display"] = ""
				}
			}
		}
	}
	return b.set("type", m)
}

// SetPolicyHolder sets policyHolder.
func (b *CoverageBuilder) SetPolicyHolder(ref *datatype.Reference) *CoverageBuilder {
	return b.set("policyHolder", ref.ToArray())
}

// SetSubscriber sets subscriber.
func (b *CoverageBuilder) SetSubscriber(ref *datatype.Reference) *CoverageBuilder {
	return b.set("subscriber", ref.ToArray())
}

// SetSubscriberId sets subscriberId.
func (b *CoverageBuilder) SetSubscriberId(value string) *CoverageBuilder {
	return b.set("subscriberId", value)
}

// SetBeneficiary sets beneficiary.
func (b *CoverageBuilder) SetBeneficiary(ref *datatype.Reference) *CoverageBuilder {
	return b.set("beneficiary", ref.ToArray())
}

// SetDependent sets dependent.
func (b *CoverageBuilder) SetDependent(value string) *CoverageBuilder {
	return b.set("dependent", value)
}

// SetRelationship sets relationship.
func (b *CoverageBuilder) SetRelationship(rel *datatype.CodeableConcept) *CoverageBuilder {
	return b.set("relationship", rel.ToArray())
}

// AddPayor appends a payor.
func (b *CoverageBuilder) AddPayor(ref *datatype.Reference) *CoverageBuilder {
	return b.push("payor", ref.ToArray())
}

// AddClass appends a class entry. name is optional.
func (b *CoverageBuilder) AddClass(typ *datatype.CodeableConcept, value, name string) *CoverageBuilder {
	class := map[string]interface{}{"type": typ.ToArray(), "value": value}
	if name != "" {
		class["name"] = name
	}
	return b.push("class", class)
}

// SetOrder sets order.
func (b *CoverageBuilder) SetOrder(order int) *CoverageBuilder { return b.set("order", order) }

// SetNetwork sets network.
func (b *CoverageBuilder) SetNetwork(network string) *CoverageBuilder {
	return b.set("network", network)
}

// AddCostToBeneficiary appends a costToBeneficiary entry. PHP writes value and currency at the top
// level of the element rather than inside valueMoney; that shape is kept.
func (b *CoverageBuilder) AddCostToBeneficiary(value float64, currency string) *CoverageBuilder {
	return b.push("costToBeneficiary", map[string]interface{}{
		"value":    value,
		"currency": currency,
	})
}

// SetPeriod sets period.start and, when end is not empty, period.end.
func (b *CoverageBuilder) SetPeriod(start, end string) *CoverageBuilder {
	b.set("period/start", start)
	if end != "" {
		b.set("period/end", end)
	}
	return b
}

// Build returns the payload.
func (b *CoverageBuilder) Build() map[string]interface{} { return b.Data }

// BuildJSON returns the payload as JSON.
func (b *CoverageBuilder) BuildJSON() ([]byte, error) { return json.Marshal(b.Data) }
