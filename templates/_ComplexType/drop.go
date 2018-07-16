package __ComplexType

import . "_ImportLocation"

func Drop(slice []_ComplexType, n int) (res []_ComplexType) {
	l := len(slice) - n
	if l < 0 {
		l = 0
	}
	res = make([]_ComplexType, 0, l)
	for _, entry := range slice[len(slice) - l:] {
		res = append(res, entry)
	}
	return
}

func (c *chain) Drop(n int) *chain {
	return &chain{value: Drop(c.value, n)}
}

func DropRight(slice []_ComplexType, n int) (res []_ComplexType) {
	l := len(slice) - n
	if l < 0 {
		l = 0
	}
	res = make([]_ComplexType, 0, l)
	for _, entry := range slice[:l] {
		res = append(res, entry)
	}
	return
}

func (c *chain) DropRight(n int) *chain {
	return &chain{value: DropRight(c.value, n)}
}