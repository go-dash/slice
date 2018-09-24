package tests

import (
	"reflect"
	"testing"

	"github.com/go-dash/slice/_Person" // github.com/go-dash/slice/tests/types
	"github.com/go-dash/slice/_int"
	"github.com/go-dash/slice/_string"
	. "github.com/go-dash/slice/tests/types"
)

var tableSortString = []struct {
	input  []string
	output []string
}{
	{nil, []string{}},
	{[]string{}, []string{}},
	{[]string{"one", "two", "three", "four"}, []string{"four", "one", "three", "two"}},
}

func TestSortString(t *testing.T) {
	for _, tt := range tableSortString {
		res := _string.Sort(tt.input, func(a, b string) bool {
			return a <= b
		})
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _string.Chain(tt.input).Sort(func(a, b string) bool {
			return a <= b
		}).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableSortInt = []struct {
	input  []int
	output []int
}{
	{nil, []int{}},
	{[]int{}, []int{}},
	{[]int{10, 2, 5, 8, 5, 6}, []int{2, 5, 5, 6, 8, 10}},
}

func TestSortInt(t *testing.T) {
	for _, tt := range tableSortInt {
		res := _int.Sort(tt.input, func(a, b int) bool {
			return a <= b
		})
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _int.Chain(tt.input).Sort(func(a, b int) bool {
			return a <= b
		}).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableSortPerson = []struct {
	input  []Person
	output []Person
}{
	{nil, []Person{}},
	{[]Person{}, []Person{}},
	{
		[]Person{Person{"das", 1018}, Person{"sdbb", 109}, Person{"fecc", 122}, Person{"adad", 122}, Person{"rfaa", 122}, Person{"bb", 1209}, Person{"cc", 250}, Person{"dd", 210}},
		[]Person{Person{"sdbb", 109}, Person{"adad", 122}, Person{"fecc", 122}, Person{"rfaa", 122}, Person{"dd", 210}, Person{"cc", 250}, Person{"das", 1018}, Person{"bb", 1209}},
	},
}

func TestSortPerson(t *testing.T) {
	for _, tt := range tableSortPerson {
		res := _Person.Sort(tt.input, func(a, b Person) bool {
			if a.Age < b.Age {
				return true
			} else if a.Age > b.Age {
				return false
			}
			return a.Name < b.Name
		})
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _Person.Chain(tt.input).Sort(func(a, b Person) bool {
			if a.Age < b.Age {
				return true
			} else if a.Age > b.Age {
				return false
			}
			return a.Name < b.Name
		}).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}
