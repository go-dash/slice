package __SimpleType

import (
	"sort"
)

// Type ...
type Type struct {
	typ []_SimpleType
	cmp func(a, b _SimpleType) bool
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
func Sort(slice []_SimpleType, cmp func(a, b _SimpleType) bool) (res []_SimpleType) {
	res = make([]_SimpleType, 0, len(slice))
	res = append(res, slice...)
	val := Type{
		typ: res,
		cmp: cmp,
	}
	sort.Sort(val)
	return
}

func (c *chain) Sort(cmp func(a, b _SimpleType) bool) *chain {
	return &chain{value: Sort(c.value, cmp)}
}
