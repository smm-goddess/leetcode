package majorityelement

import (
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
func TestMajorityElement(t *testing.T) {
	if !listEquals(majorityElement([]int{3, 2, 3}), []int{3}) {
		t.Error("First Error")
	}
	if !listEquals(majorityElement([]int{1, 1, 1, 3, 3, 2, 2, 2}), []int{1, 2}) {
		t.Error("Second Error")
	}
	if !listEquals(majorityElement([]int{2, 2}), []int{2}) {
		t.Error("Third Error")
	}
}
