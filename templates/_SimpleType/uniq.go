package __SimpleType

func Uniq(slice []_SimpleType) (res []_SimpleType) {
	seen := make(map[_SimpleType]bool)
	res = []_SimpleType{}
	for _, entry := range slice {
		if _, found := seen[entry]; !found {
			seen[entry] = true
			res = append(res, entry)
		}
	}
	return
}

func (c *chain) Uniq() *chain {
	return &chain{value: Uniq(c.value)}
}