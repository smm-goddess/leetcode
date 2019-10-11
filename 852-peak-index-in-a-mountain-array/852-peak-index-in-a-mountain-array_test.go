package peak_index_in_a_mountain_array

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
	if peakIndexInMountainArray([]int{0, 1, 0}) != 1 {
		t.Error("error")
	}
	if peakIndexInMountainArray([]int{0, 2, 1, 0}) != 1 {
		t.Error("error")
	}
}
