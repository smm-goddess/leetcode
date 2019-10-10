package shortest_unsorted_continuous_subarray

import "testing"

func Test_findUnsortedSubarray(t *testing.T) {
	if findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}) != 5 {
		t.Error("5 error")
	}
	if findUnsortedSubarray([]int{1}) != 0 {
		t.Error("0 error")
	}
}
