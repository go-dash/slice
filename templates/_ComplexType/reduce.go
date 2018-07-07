package __ComplexType

import . "_ImportLocation"

func Reduce(slice []_ComplexType, fn func(_ComplexType,_ComplexType,int)_ComplexType, initial _ComplexType) (res _ComplexType) {
	res = initial
	for index, entry := range slice {
		res = fn(res, entry, index)
	}
	return
}

func (c *chain) Reduce(fn func(_ComplexType,_ComplexType,int)_ComplexType, initial _ComplexType) *chain {
	return &chain{value: []_ComplexType{Reduce(c.value, fn, initial)}}
}