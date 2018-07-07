package tests

import (
	"testing"
	"reflect"
	. "github.com/go-dash/slice/tests/types"
	"github.com/go-dash/slice/_string"
	"github.com/go-dash/slice/_int"
	"github.com/go-dash/slice/_Person" // github.com/go-dash/slice/tests/types
)

var tableReduceString = []struct {
	input  []string
	output string
}{
	{nil, ""},
	{[]string{}, ""},
	{[]string{"aa", "bb", "cc"}, "aabbcc"},
}

func TestReduceString(t *testing.T) {
	for _, tt := range tableReduceString {
		res := _string.Reduce(tt.input, func (acc string, s string, index int) string {
			return acc + s
		}, "")
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res2 := _string.Chain(tt.input).Reduce(func (acc string, s string, index int) string {
			return acc + s
		}, "").Value()
		if !reflect.DeepEqual(res2[0], tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableReduceInt = []struct {
	input  []int
	output int
}{
	{nil, 0},
	{[]int{}, 0},
	{[]int{1, 2, 3}, 6},
}

func TestReduceInt(t *testing.T) {
	for _, tt := range tableReduceInt {
		res := _int.Reduce(tt.input, func (acc int, n int, index int) int {
			return acc + n
		}, 0)
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res2 := _int.Chain(tt.input).Reduce(func (acc int, n int, index int) int {
			return acc + n
		}, 0).Value()
		if !reflect.DeepEqual(res2[0], tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableReducePerson = []struct {
	input  []Person
	output Person
}{
	{nil, Person{}},
	{[]Person{}, Person{}},
	{[]Person{Person{"aa", 18}, Person{"bb", 16}}, Person{"aabb", 34}},
}

func TestReducePerson(t *testing.T) {
	for _, tt := range tableReducePerson {
		res := _Person.Reduce(tt.input, func (acc Person, p Person, index int) Person {
			return Person{acc.Name + p.Name, acc.Age + p.Age}
		}, Person{})
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res2 := _Person.Chain(tt.input).Reduce(func (acc Person, p Person, index int) Person {
			return Person{acc.Name + p.Name, acc.Age + p.Age}
		}, Person{}).Value()
		if !reflect.DeepEqual(res2[0], tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}