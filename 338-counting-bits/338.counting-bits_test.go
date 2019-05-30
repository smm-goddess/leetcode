package countbits

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
func TestCountingBits(t *testing.T) {
	if !listEquals(countBits(2), []int{0, 1, 1}) {
		t.Error("error")
	}
	if !listEquals(countBits(5), []int{0, 1, 1, 2, 1, 2}) {
		t.Error("error")
	}
	if !listEquals(countBits(8), []int{0, 1, 1, 2, 1, 2, 2, 3, 1}) {
		t.Error("error")
	}
}
