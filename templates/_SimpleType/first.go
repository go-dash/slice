package __SimpleType

func First(slice []_SimpleType) (res _SimpleType) {
	if len(slice) == 0 {
		return
	}
	res = slice[0]
	return
}

func (c *chain) First() *chain {
	return &chain{value: []_SimpleType{First(c.value)}}
}