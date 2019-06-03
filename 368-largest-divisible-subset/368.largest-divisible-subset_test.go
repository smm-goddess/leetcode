package largestdivisiblesubset

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
func TestLargestDivisibleSubsets(t *testing.T) {
	if !listEquals(largestDivisibleSubset([]int{1, 2, 4, 8}), []int{8, 4, 2, 1}) {
		t.Error("1 2 4 8 error")
	}
	if !listEquals(largestDivisibleSubset([]int{1, 2, 3}), []int{2, 1}) {
		t.Error("1 2 3 error")
	}
	if !listEquals(largestDivisibleSubset([]int{3, 4, 16, 8}), []int{16, 8, 4}) {
		t.Error("3 4 16 8")
	}
	if !listEquals(largestDivisibleSubset([]int{2, 3, 8, 9, 27}), []int{27, 9, 3}) {
		t.Error("2,3,8,9,27")
	}
}
