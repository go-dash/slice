package __ComplexType

import . "_ImportLocation"

func First(slice []_ComplexType) (res _ComplexType) {
	if len(slice) == 0 {
		return
	}
	res = slice[0]
	return
}

func (c *chain) First() *chain {
	return &chain{value: []_ComplexType{First(c.value)}}
}