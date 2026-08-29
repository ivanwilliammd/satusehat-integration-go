package datatype

type SimpleQuantity struct {
	Quantity
}

func (q *SimpleQuantity) ToArray() map[string]interface{} {
	return q.Quantity.ToArray()
}
