package __SimpleType

func Reduce(slice []_SimpleType, fn func(_SimpleType,_SimpleType,int)_SimpleType, initial _SimpleType) (res _SimpleType) {
	res = initial
	for index, entry := range slice {
		res = fn(res, entry, index)
	}
	return
}

func (c *chain) Reduce(fn func(_SimpleType,_SimpleType,int)_SimpleType, initial _SimpleType) *chain {
	return &chain{value: []_SimpleType{Reduce(c.value, fn, initial)}}
}