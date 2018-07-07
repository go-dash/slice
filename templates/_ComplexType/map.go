package __ComplexType

import . "_ImportLocation"

func Map(slice []_ComplexType, fn func(_ComplexType,int)_ComplexType) (res []_ComplexType) {
	res = make([]_ComplexType, 0, len(slice))
	for index, entry := range slice {
		res = append(res, fn(entry, index))
	}
	return
}

func (c *chain) Map(fn func(_ComplexType,int)_ComplexType) *chain {
	return &chain{value: Map(c.value, fn)}
}