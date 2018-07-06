package __ComplexType

import . "_ImportLocation"

func Uniq(slice []_ComplexType) (res []_ComplexType) {
	seen := make(map[_ComplexType]bool)
	res = []_ComplexType{}
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