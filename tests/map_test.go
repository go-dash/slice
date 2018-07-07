package tests

import (
	"testing"
	"reflect"
	. "github.com/go-dash/slice/tests/types"
	"github.com/go-dash/slice/_string"
	"github.com/go-dash/slice/_int"
	"github.com/go-dash/slice/_Person" // github.com/go-dash/slice/tests/types
	"strings"
	"strconv"
)

var tableMapString = []struct {
	input  []string
	output []string
}{
	{nil, []string{}},
	{[]string{}, []string{}},
	{[]string{"aa", "bb", "cc"}, []string{"AA", "BB", "CC"}},
}

func TestMapString(t *testing.T) {
	for _, tt := range tableMapString {
		res := _string.Map(tt.input, func (s string, index int) string {
			return strings.ToUpper(s)
		})
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _string.Chain(tt.input).Map(func (s string, index int) string {
			return strings.ToUpper(s)
		}).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableMapInt = []struct {
	input  []int
	output []int
}{
	{nil, []int{}},
	{[]int{}, []int{}},
	{[]int{1, 2, 3}, []int{2, 4, 6}},
}

func TestMapInt(t *testing.T) {
	for _, tt := range tableMapInt {
		res := _int.Map(tt.input, func (n int, index int) int {
			return 2*n
		})
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _int.Chain(tt.input).Map(func (n int, index int) int {
			return 2*n
		}).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}

var tableMapPerson = []struct {
	input  []Person
	output []Person
}{
	{nil, []Person{}},
	{[]Person{}, []Person{}},
	{[]Person{Person{"aa", 18}, Person{"bb", 16}}, []Person{Person{"aa18", 19}, Person{"bb16", 17}}},
}

func TestMapPerson(t *testing.T) {
	for _, tt := range tableMapPerson {
		res := _Person.Map(tt.input, func (p Person, index int) Person {
			return Person{p.Name + strconv.Itoa(p.Age), p.Age + 1}
		})
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
		res = _Person.Chain(tt.input).Map(func (p Person, index int) Person {
			return Person{p.Name + strconv.Itoa(p.Age), p.Age + 1}
		}).Value()
		if !reflect.DeepEqual(res, tt.output) {
			t.Fatalf("Expected %v received %v", tt.output, res)
		}
	}
}