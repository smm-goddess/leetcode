package findorder

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
		return listEquals(a[0], b[0]) && listOfListEquals(a[1:], b[1:])
	}
}

func TestFindOrder(t *testing.T) {
	if !listEquals(findOrder(2, [][]int{{1, 0}}), []int{0, 1}) {
		t.Error("2 Error")
	}
	if !listEquals(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}), []int{0, 1, 2, 3}) {
		t.Error("4 Error")
	}
}
