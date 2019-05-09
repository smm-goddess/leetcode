package combinationSum3

import "testing"

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

func listOfListEquals(a, b [][]int) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if !listEquals(b[i], v) {
				return false
			}
		}
		return true
	}
}

func TestCombinationSum(t *testing.T) {
	if !listOfListEquals(combinationSum3(3, 7), [][]int{{1, 2, 4}}) {
		t.Error("3,7 Error")
	}
	if !listOfListEquals(combinationSum3(3, 9), [][]int{{2, 3, 4}, {1, 3, 5}, {1, 2, 6}}) {
		t.Error("3,9 Error")
	}
}
