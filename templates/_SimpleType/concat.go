package __SimpleType

func Concat(slice []_SimpleType, slice2 []_SimpleType) (res []_SimpleType) {
	res = make([]_SimpleType, 0, len(slice) + len(slice2))
	for _, entry := range slice {
		res = append(res, entry)
	}
	for _, entry := range slice2 {
		res = append(res, entry)
	}
	return
}

func (c *chain) Concat(slice2 []_SimpleType) *chain {
	return &chain{value: Concat(c.value, slice2)}
}