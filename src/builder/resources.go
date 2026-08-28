package builder

type PatientBuilder struct{ Data map[string]interface{} }
type PractitionerBuilder struct{ Data map[string]interface{} }
type OrganizationBuilder struct{ Data map[string]interface{} }
type LocationBuilder struct{ Data map[string]interface{} }
type EncounterBuilder struct{ Data map[string]interface{} }
type ConditionBuilder struct{ Data map[string]interface{} }

func NewPatient() *PatientBuilder { return &PatientBuilder{Data: make(map[string]interface{})} }
func NewPractitioner() *PractitionerBuilder { return &PractitionerBuilder{Data: make(map[string]interface{})} }
func NewOrganization() *OrganizationBuilder { return &OrganizationBuilder{Data: make(map[string]interface{})} }
func NewLocation() *LocationBuilder { return &LocationBuilder{Data: make(map[string]interface{})} }
func NewEncounter() *EncounterBuilder { return &EncounterBuilder{Data: make(map[string]interface{})} }
func NewCondition() *ConditionBuilder { return &ConditionBuilder{Data: make(map[string]interface{})} }
