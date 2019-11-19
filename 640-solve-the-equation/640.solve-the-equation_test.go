package solve_equation

import (
	"fmt"
	"testing"
)

var test_cases map[string]string

func init() {
	test_cases = map[string]string{
		"x+5-3+x=6+x-2": "x=2",
		"x=x":           "Infinite solutions",
		"2x=x":          "x=0",
		"2x+3x-6x=x+2":  "x=-1",
		"x=x+2":         "No solution",
		"1+1=x":         "x=2",
		"0x=0":          "Infinite solutions",
	}
}

func TestSolveEquation(t *testing.T) {
	for equation, value := range test_cases {
		if value != solveEquation(equation) {
			t.Error(fmt.Sprintf("equation %s error", equation))
		}
	}
}
