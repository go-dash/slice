package __SimpleType

func Reverse(slice []_SimpleType) (res []_SimpleType) {
	res = make([]_SimpleType, len(slice))
	for index, entry := range slice {
		res[len(slice)-1-index] = entry
	}
	return
}

func (c *chain) Reverse() *chain {
	return &chain{value: Reverse(c.value)}
}