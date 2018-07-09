package tests

import (
	"testing"
	"reflect"
	. "github.com/go-dash/slice/tests/types"
	"github.com/go-dash/slice/_string"
	"github.com/go-dash/slice/_int"
	"github.com/go-dash/slice/_Person" // github.com/go-dash/slice/tests/types
)

var tableLastString = []struct {
	input  []string
	output string
}{
	{nil, ""},
	{[]string{}, ""},
	{[]string{"aa", "bb", "cc"}, "cc"},
}

func TestLastString(t *testing.T) {
	for _, tt := range tableLastString {
		res := _string.Last(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res2 := _string.Chain(tt.input).Last().Value()
		if !reflect.DeepEqual(res2[0], tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableLastInt = []struct {
	input  []int
	output int
}{
	{nil, 0},
	{[]int{}, 0},
	{[]int{1, 2, 3}, 3},
}

func TestLastInt(t *testing.T) {
	for _, tt := range tableLastInt {
		res := _int.Last(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res2 := _int.Chain(tt.input).Last().Value()
		if !reflect.DeepEqual(res2[0], tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableLastPerson = []struct {
	input  []Person
	output Person
}{
	{nil, Person{}},
	{[]Person{}, Person{}},
	{[]Person{Person{"aa", 18}, Person{"bb", 16}}, Person{"bb", 16}},
}

func TestLastPerson(t *testing.T) {
	for _, tt := range tableLastPerson {
		res := _Person.Last(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res2 := _Person.Chain(tt.input).Last().Value()
		if !reflect.DeepEqual(res2[0], tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}