package tests

import (
	"testing"
	"reflect"
	. "github.com/go-dash/slice/tests/types"
	"github.com/go-dash/slice/_string"
	"github.com/go-dash/slice/_int"
	"github.com/go-dash/slice/_Person" // github.com/go-dash/slice/tests/types
	"strings"
)

var tableFilterString = []struct {
	input  []string
	output []string
}{
	{nil, []string{}},
	{[]string{}, []string{}},
	{[]string{"aaa", "aaa", "aaa"}, []string{}},
	{[]string{"aa", "bb", "aa", "cc", "bb"}, []string{"bb", "cc", "bb"}},
}

func TestFilterString(t *testing.T) {
	for _, tt := range tableFilterString {
		res := _string.Filter(tt.input, func (s string, index int) bool {
			return !strings.HasPrefix(s, "a")
		})
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _string.Chain(tt.input).Filter(func (s string, index int) bool {
			return !strings.HasPrefix(s, "a")
		}).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableFilterInt = []struct {
	input  []int
	output []int
}{
	{nil, []int{}},
	{[]int{}, []int{}},
	{[]int{1, 1, 1}, []int{}},
	{[]int{1, 2, 1, 3, 2}, []int{2, 2}},
}

func TestFilterInt(t *testing.T) {
	for _, tt := range tableFilterInt {
		res := _int.Filter(tt.input, func (n int, index int) bool {
			return n % 2 == 0
		})
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _int.Chain(tt.input).Filter(func (n int, index int) bool {
			return n % 2 == 0
		}).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableFilterPerson = []struct {
	input  []Person
	output []Person
}{
	{nil, []Person{}},
	{[]Person{}, []Person{}},
	{[]Person{Person{"aa", 18}, Person{"aa", 18}}, []Person{Person{"aa", 18}, Person{"aa", 18}}},
	{[]Person{Person{"aa", 18}, Person{"aa", 17}, Person{"aa", 18}, Person{"bb", 16}, Person{"aa", 17}}, []Person{Person{"aa", 18}, Person{"aa", 18}}},
}

func TestFilterPerson(t *testing.T) {
	for _, tt := range tableFilterPerson {
		res := _Person.Filter(tt.input, func (p Person, index int) bool {
			return p.Age >= 18
		})
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _Person.Chain(tt.input).Filter(func (p Person, index int) bool {
			return p.Age >= 18
		}).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}