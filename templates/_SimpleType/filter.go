package __SimpleType

func Filter(slice []_SimpleType, fn func(_SimpleType,int)bool) (res []_SimpleType) {
	res = make([]_SimpleType, 0, len(slice))
	for index, entry := range slice {
		if fn(entry, index) {
			res = append(res, entry)
		}
	}
	return
}

func (c *chain) Filter(fn func(_SimpleType,int)bool) *chain {
	return &chain{value: Filter(c.value, fn)}
}