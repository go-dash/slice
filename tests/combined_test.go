package tests

import (
	"testing"
	. "github.com/go-dash/slice/tests/types"
	"github.com/go-dash/slice/tests/combined"
	"reflect"
)

func TestUniqCombined(t *testing.T) {
	// string
	stringInput := []string{"aa", "bb", "aa"}
	stringOutput := combined.Uniq(stringInput).([]string)
	stringExpected := []string{"aa", "bb"}
	if !reflect.DeepEqual(stringOutput, stringExpected) {
		t.Fatalf("Expected %v received %v", stringExpected, stringOutput)
	}

	// int
	intInput := []int{1, 2, 1}
	intOutput := combined.Uniq(intInput).([]int)
	intExpected := []int{1, 2}
	if !reflect.DeepEqual(intOutput, intExpected) {
		t.Fatalf("Expected %v received %v", intExpected, intOutput)
	}

	// Person
	personInput := []Person{Person{"aa", 18}, Person{"bb", 19}, Person{"aa", 18}}
	personOutput := combined.Uniq(personInput).([]Person)
	personExpected := []Person{Person{"aa", 18}, Person{"bb", 19}}
	if !reflect.DeepEqual(personOutput, personExpected) {
		t.Fatalf("Expected %v received %v", personExpected, personOutput)
	}
}