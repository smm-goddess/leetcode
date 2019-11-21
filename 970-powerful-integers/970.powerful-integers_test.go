package powerful_integers

import (
	"fmt"
	"sort"
	"testing"
)

func listEquals(a, b []int) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
		return true
	}
}

type Case struct {
	X      int
	Y      int
	Bound  int
	Result []int
}

var cases = []*Case{
	{2, 3, 10, []int{2, 3, 4, 5, 7, 9, 10}},
	{3, 5, 15, []int{2, 4, 6, 8, 10, 14}},
	{1, 1, 0, []int{}},
	{1, 1, 2, []int{2}},
	{1, 3, 10, []int{2, 4, 10}},
}

func TestPowerfulIntegers(t *testing.T) {
	for _, cs := range cases {
		r := powerfulIntegers(cs.X, cs.Y, cs.Bound)
		sort.Ints(r)
		fmt.Println(r)
		if !listEquals(r, cs.Result) {
			t.Error(cs.X, " ", cs.Y, " ", cs.Bound, " error")
		}
	}
}
