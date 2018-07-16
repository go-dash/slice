package __SimpleType

func Drop(slice []_SimpleType, n int) (res []_SimpleType) {
	l := len(slice) - n
	if l < 0 {
		l = 0
	}
	res = make([]_SimpleType, 0, l)
	for _, entry := range slice[len(slice) - l:] {
		res = append(res, entry)
	}
	return
}

func (c *chain) Drop(n int) *chain {
	return &chain{value: Drop(c.value, n)}
}

func DropRight(slice []_SimpleType, n int) (res []_SimpleType) {
	l := len(slice) - n
	if l < 0 {
		l = 0
	}
	res = make([]_SimpleType, 0, l)
	for _, entry := range slice[:l] {
		res = append(res, entry)
	}
	return
}

func (c *chain) DropRight(n int) *chain {
	return &chain{value: DropRight(c.value, n)}
}