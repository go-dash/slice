package combined

import (
	. "github.com/go-dash/slice/tests/types"
	"github.com/go-dash/slice/_string"
	"github.com/go-dash/slice/_int"
	"github.com/go-dash/slice/_Person"
)

func Uniq(slice interface{}) (res interface{}) {
	switch slice.(type) {
	case []string:
		return _string.Uniq(slice.([]string))
	case []int:
		return _int.Uniq(slice.([]int))
	case []Person:
		return _Person.Uniq(slice.([]Person))
	}
	return nil
}
