package canpartition

import "testing"

func TestCanPartition(t *testing.T) {
	t.Log("start testing")
	nums := []int{1, 5, 11, 5}
	if !canPartition(nums) {
		t.Error("1 5 11 5 error")
	}
	nums = []int{1, 2, 3, 5}
	if canPartition(nums) {
		t.Error("1 2 3 5 error")
	}
	t.Log("end testing")
}
