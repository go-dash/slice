package __ComplexType

import . "_ImportLocation"

func Last(slice []_ComplexType) (res _ComplexType) {
	if len(slice) == 0 {
		return
	}
	res = slice[len(slice) - 1]
	return
}

func (c *chain) Last() *chain {
	return &chain{value: []_ComplexType{Last(c.value)}}
}