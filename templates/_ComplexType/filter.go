package __ComplexType

import . "_ImportLocation"

func Filter(slice []_ComplexType, fn func(_ComplexType,int)bool) (res []_ComplexType) {
	res = make([]_ComplexType, 0, len(slice))
	for index, entry := range slice {
		if fn(entry, index) {
			res = append(res, entry)
		}
	}
	return
}

func (c *chain) Filter(fn func(_ComplexType,int)bool) *chain {
	return &chain{value: Filter(c.value, fn)}
}