package builder

import (
    "encoding/json"
    "github.com/ivanwilliammd/satusehat-integration-go/v4/src/datatype"
)

type BiologicallyDerivedProductBuilder struct {
    ResourceType string
    Data map[string]interface{}
}

func NewBiologicallyDerivedProductBuilder() *BiologicallyDerivedProductBuilder {
    b := &BiologicallyDerivedProductBuilder{ResourceType: "BiologicallyDerivedProduct", Data: make(map[string]interface{})}
    b.Data["resourceType"] = "BiologicallyDerivedProduct"
    return b
}

func (b *BiologicallyDerivedProductBuilder) setId(id string) *BiologicallyDerivedProductBuilder {
    b.Data["id"] = id
    return b
}

func (b *BiologicallyDerivedProductBuilder) addIdentifier(identifier *datatype.Identifier) *BiologicallyDerivedProductBuilder {
    if b.Data["identifier"] == nil {
        b.Data["identifier"] = []*datatype.Identifier{}
    }
    b.Data["identifier"] = append(b.Data["identifier"].([]*datatype.Identifier), identifier)
    return b
}

func (b *BiologicallyDerivedProductBuilder) setStatus(status string) *BiologicallyDerivedProductBuilder {
    b.Data["status"] = status
    return b
}

func (b *BiologicallyDerivedProductBuilder) setSubject(reference string) *BiologicallyDerivedProductBuilder {
    b.Data["subject"] = datatype.Reference{}.ToArray()
    return b
}

func (b *BiologicallyDerivedProductBuilder) Build() map[string]interface{} {
    return b.Data
}

func (b *BiologicallyDerivedProductBuilder) BuildJSON() ([]byte, error) {
    return json.Marshal(b.Data)
}
