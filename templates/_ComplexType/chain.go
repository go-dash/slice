package __ComplexType

import . "_ImportLocation"

type chain struct {
	value []_ComplexType
}

func Chain(slice []_ComplexType) *chain {
	return &chain{value: slice}
}

func (c *chain) Value() []_ComplexType {
	return c.value
}