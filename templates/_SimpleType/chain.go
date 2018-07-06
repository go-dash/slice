package __SimpleType

type chain struct {
	value []_SimpleType
}

func Chain(slice []_SimpleType) *chain {
	return &chain{value: slice}
}

func (c *chain) Value() []_SimpleType {
	return c.value
}