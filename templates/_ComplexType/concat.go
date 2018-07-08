package __ComplexType

import . "_ImportLocation"

func Concat(slice []_ComplexType, slice2 []_ComplexType) (res []_ComplexType) {
	res = make([]_ComplexType, 0, len(slice) + len(slice2))
	for _, entry := range slice {
		res = append(res, entry)
	}
	for _, entry := range slice2 {
		res = append(res, entry)
	}
	return
}

func (c *chain) Concat(slice2 []_ComplexType) *chain {
	return &chain{value: Concat(c.value, slice2)}
}