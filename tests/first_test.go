package tests

import (
	"testing"
	"reflect"
	. "github.com/go-dash/slice/tests/types"
	"github.com/go-dash/slice/_string"
	"github.com/go-dash/slice/_int"
	"github.com/go-dash/slice/_Person" // github.com/go-dash/slice/tests/types
)

var tableFirstString = []struct {
	input  []string
	output string
}{
	{nil, ""},
	{[]string{}, ""},
	{[]string{"aa", "bb", "cc"}, "aa"},
}

func TestFirstString(t *testing.T) {
	for _, tt := range tableFirstString {
		res := _string.First(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res2 := _string.Chain(tt.input).First().Value()
		if !reflect.DeepEqual(res2[0], tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableFirstInt = []struct {
	input  []int
	output int
}{
	{nil, 0},
	{[]int{}, 0},
	{[]int{1, 2, 3}, 1},
}

func TestFirstInt(t *testing.T) {
	for _, tt := range tableFirstInt {
		res := _int.First(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res2 := _int.Chain(tt.input).First().Value()
		if !reflect.DeepEqual(res2[0], tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableFirstPerson = []struct {
	input  []Person
	output Person
}{
	{nil, Person{}},
	{[]Person{}, Person{}},
	{[]Person{Person{"aa", 18}, Person{"bb", 16}}, Person{"aa", 18}},
}

func TestFirstPerson(t *testing.T) {
	for _, tt := range tableFirstPerson {
		res := _Person.First(tt.input)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res2 := _Person.Chain(tt.input).First().Value()
		if !reflect.DeepEqual(res2[0], tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}