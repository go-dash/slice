package __SimpleType

func Last(slice []_SimpleType) (res _SimpleType) {
	if len(slice) == 0 {
		return
	}
	res = slice[len(slice) - 1]
	return
}

func (c *chain) Last() *chain {
	return &chain{value: []_SimpleType{Last(c.value)}}
}