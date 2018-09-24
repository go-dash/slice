package __ComplexType

import (
	. "_ImportLocation"
	"sort"
)

// Type ...
type Type struct {
	typ []_ComplexType
	cmp func(a, b _ComplexType) bool
}

func (obj Type) Len() int {
	return len(obj.typ)
}

func (obj Type) Swap(i, j int) {
	obj.typ[i], obj.typ[j] = obj.typ[j], obj.typ[i]
}

func (obj Type) Less(i, j int) bool {
	return obj.cmp(obj.typ[i], obj.typ[j])
}

// Sort ...
func Sort(slice []_ComplexType, cmp func(a, b _ComplexType) bool) (res []_ComplexType) {
	res = make([]_ComplexType, 0, len(slice))
	res = append(res, slice...)
	val := Type{
		typ: res,
		cmp: cmp,
	}
	sort.Sort(val)
	return
}

func (c *chain) Sort(cmp func(a, b _ComplexType) bool) *chain {
	return &chain{value: Sort(c.value, cmp)}
}
