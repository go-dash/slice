package tests

import (
	"testing"
	"reflect"
	. "github.com/go-dash/slice/tests/types"
	"github.com/go-dash/slice/_string"
	"github.com/go-dash/slice/_int"
	"github.com/go-dash/slice/_Person" // github.com/go-dash/slice/tests/types
)

var tableDropString = []struct {
	input  []string
	n      int
	output []string
}{
	{nil, 5, []string{}},
	{[]string{}, 3, []string{}},
	{[]string{"aaa", "aaa", "aaa"}, 2, []string{"aaa"}},
	{[]string{"aa", "bb", "aa", "cc", "bb"}, 3, []string{"cc", "bb"}},
}

func TestDropString(t *testing.T) {
	for _, tt := range tableDropString {
		res := _string.Drop(tt.input, tt.n)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _string.Chain(tt.input).Drop(tt.n).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableDropInt = []struct {
	input  []int
	n      int
	output []int
}{
	{nil, 5, []int{}},
	{[]int{}, 3, []int{}},
	{[]int{1, 1, 1}, 2, []int{1}},
	{[]int{1, 2, 1, 3, 2}, 3, []int{3, 2}},
}

func TestDropInt(t *testing.T) {
	for _, tt := range tableDropInt {
		res := _int.Drop(tt.input, tt.n)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _int.Chain(tt.input).Drop(tt.n).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableDropPerson = []struct {
	input  []Person
	n      int
	output []Person
}{
	{nil, 5, []Person{}},
	{[]Person{}, 3, []Person{}},
	{[]Person{Person{"aa", 18}, Person{"aa", 17}, Person{"bb", 18}, Person{"cc", 19}}, 2, []Person{Person{"bb", 18}, Person{"cc", 19}}},
}

func TestDropPerson(t *testing.T) {
	for _, tt := range tableDropPerson {
		res := _Person.Drop(tt.input, tt.n)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _Person.Chain(tt.input).Drop(tt.n).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableDropRightString = []struct {
	input  []string
	n      int
	output []string
}{
	{nil, 5, []string{}},
	{[]string{}, 3, []string{}},
	{[]string{"aaa", "aaa", "aaa"}, 2, []string{"aaa"}},
	{[]string{"aa", "bb", "aa", "cc", "bb"}, 3, []string{"aa", "bb"}},
}

func TestDropRightString(t *testing.T) {
	for _, tt := range tableDropRightString {
		res := _string.DropRight(tt.input, tt.n)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _string.Chain(tt.input).DropRight(tt.n).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}
