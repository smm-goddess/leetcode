package singlenumber

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
func TestSingleNumber(t *testing.T) {
	ans := singleNumber([]int{1, 2, 1, 3, 2, 5})
	if !listEquals(singleNumber(ans), []int{3, 5}) && !listEquals(ans, []int{5, 3}) {
		t.Error("Error")
	}
}
