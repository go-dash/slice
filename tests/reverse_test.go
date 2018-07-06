package tests

import (
	"testing"
	"reflect"
	. "github.com/go-dash/slice/tests/types"
	"github.com/go-dash/slice/_string"
	"github.com/go-dash/slice/_int"
	"github.com/go-dash/slice/_Person" // github.com/go-dash/slice/tests/types
)

var tableReverseString = []struct {
	input  []string
	output []string
}{
	{nil, []string{}},
	{[]string{}, []string{}},
	{[]string{"aaa", "aaa", "bbb"}, []string{"bbb", "aaa", "aaa"}},
	{[]string{"aa", "bb", "aa", "cc", "bb"}, []string{"bb", "cc", "aa", "bb", "aa"}},
}

func TestReverseString(t *testing.T) {
	for _, tt := range tableReverseString {
		res := _string.Reverse(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _string.Chain(tt.input).Reverse().Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableReverseInt = []struct {
	input  []int
	output []int
}{
	{nil, []int{}},
	{[]int{}, []int{}},
	{[]int{1, 2, 3}, []int{3, 2, 1}},
	{[]int{1, 2, 1, 3, 2}, []int{2, 3, 1, 2, 1}},
}

func TestReverseInt(t *testing.T) {
	for _, tt := range tableReverseInt {
		res := _int.Reverse(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _int.Chain(tt.input).Reverse().Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableReversePerson = []struct {
	input  []Person
	output []Person
}{
	{nil, []Person{}},
	{[]Person{}, []Person{}},
	{[]Person{Person{"aa", 18}, Person{"bb", 19}}, []Person{Person{"bb", 19}, Person{"aa", 18}}},
}

func TestReversePerson(t *testing.T) {
	for _, tt := range tableReversePerson {
		res := _Person.Reverse(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _Person.Chain(tt.input).Reverse().Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}