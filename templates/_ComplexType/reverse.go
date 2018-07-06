package __ComplexType

import . "_ImportLocation"

func Reverse(slice []_ComplexType) (res []_ComplexType) {
	res = make([]_ComplexType, len(slice))
	for index, entry := range slice {
		res[len(slice)-1-index] = entry
	}
	return
}

func (c *chain) Reverse() *chain {
	return &chain{value: Reverse(c.value)}
}