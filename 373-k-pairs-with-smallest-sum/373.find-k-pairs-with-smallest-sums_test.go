package smallestpair

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
func TestSmallestPair(t *testing.T) {
	if !listOfListEquals(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3), [][]int{{1, 2}, {1, 4}, {1, 6}}) {
		t.Error("error")
	}
	if !listOfListEquals(kSmallestPairs([]int{1, 1, 2}, []int{1, 2, 3}, 2), [][]int{{1, 1}, {1, 1}}) {
		t.Error("error")
	}
	if !listOfListEquals(kSmallestPairs([]int{1, 2}, []int{3}, 3), [][]int{{1, 3}, {2, 3}}) {
		t.Error("error")
	}
}
