package threesumclosest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	if threeSumClosest([]int{-1, 2, 1, 4}, 1) != 2 {
		t.Error("fail")
	}
}
