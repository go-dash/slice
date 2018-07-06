package tests

import (
	"testing"
	"reflect"
	. "github.com/go-dash/slice/tests/types"
	"github.com/go-dash/slice/_string"
	"github.com/go-dash/slice/_int"
	"github.com/go-dash/slice/_Person" // github.com/go-dash/slice/tests/types
)

var tableUniqString = []struct {
	input  []string
	output []string
}{
	{nil, []string{}},
	{[]string{}, []string{}},
	{[]string{"aaa", "aaa", "aaa"}, []string{"aaa"}},
	{[]string{"aa", "bb", "aa", "cc", "bb"}, []string{"aa", "bb", "cc"}},
}

func TestUniqString(t *testing.T) {
	for _, tt := range tableUniqString {
		res := _string.Uniq(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _string.Chain(tt.input).Uniq().Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableUniqInt = []struct {
	input  []int
	output []int
}{
	{nil, []int{}},
	{[]int{}, []int{}},
	{[]int{1, 1, 1}, []int{1}},
	{[]int{1, 2, 1, 3, 2}, []int{1, 2, 3}},
}

func TestUniqInt(t *testing.T) {
	for _, tt := range tableUniqInt {
		res := _int.Uniq(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _int.Chain(tt.input).Uniq().Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableUniqPerson = []struct {
	input  []Person
	output []Person
}{
	{nil, []Person{}},
	{[]Person{}, []Person{}},
	{[]Person{Person{"aa", 18}, Person{"aa", 18}}, []Person{Person{"aa", 18}}},
	{[]Person{Person{"aa", 18}, Person{"aa", 17}, Person{"aa", 18}, Person{"bb", 18}, Person{"aa", 17}}, []Person{Person{"aa", 18}, Person{"aa", 17}, Person{"bb", 18}}},
}

func TestUniqPerson(t *testing.T) {
	for _, tt := range tableUniqPerson {
		res := _Person.Uniq(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _Person.Chain(tt.input).Uniq().Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}