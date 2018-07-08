package tests

import (
	"testing"
	"reflect"
	. "github.com/go-dash/slice/tests/types"
	"github.com/go-dash/slice/_string"
	"github.com/go-dash/slice/_int"
	"github.com/go-dash/slice/_Person" // github.com/go-dash/slice/tests/types
)

var tableConcatString = []struct {
	input  []string
	arg 	 []string
	output []string
}{
	{nil, nil, []string{}},
	{[]string{}, []string{}, []string{}},
	{[]string{"aaa", "aaa"}, []string{"aaa"}, []string{"aaa", "aaa", "aaa"}},
	{[]string{"aa"}, []string{"cc", "bb"}, []string{"aa", "cc", "bb"}},
}

func TestConcatString(t *testing.T) {
	for _, tt := range tableConcatString {
		res := _string.Concat(tt.input, tt.arg)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _string.Chain(tt.input).Concat(tt.arg).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableConcatInt = []struct {
	input  []int
	arg    []int
	output []int
}{
	{nil, nil, []int{}},
	{[]int{}, []int{}, []int{}},
	{[]int{111, 111}, []int{111}, []int{111, 111, 111}},
	{[]int{11}, []int{33, 22}, []int{11, 33, 22}},
}

func TestConcatInt(t *testing.T) {
	for _, tt := range tableConcatInt {
		res := _int.Concat(tt.input, tt.arg)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _int.Chain(tt.input).Concat(tt.arg).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableConcatPerson = []struct {
	input  []Person
	arg    []Person
	output []Person
}{
	{nil, nil, []Person{}},
	{[]Person{}, []Person{}, []Person{}},
	{[]Person{Person{"aa", 18}, Person{"bb", 19}}, []Person{Person{"cc", 20}, Person{"dd", 21}}, []Person{Person{"aa", 18}, Person{"bb", 19}, Person{"cc", 20}, Person{"dd", 21}}},
}

func TestConcatPerson(t *testing.T) {
	for _, tt := range tableConcatPerson {
		res := _Person.Concat(tt.input, tt.arg)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _Person.Chain(tt.input).Concat(tt.arg).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}