package __SimpleType

func Map(slice []_SimpleType, fn func(_SimpleType,int)_SimpleType) (res []_SimpleType) {
	res = make([]_SimpleType, 0, len(slice))
	for index, entry := range slice {
		res = append(res, fn(entry, index))
	}
	return
}

func (c *chain) Map(fn func(_SimpleType,int)_SimpleType) *chain {
	return &chain{value: Map(c.value, fn)}
}