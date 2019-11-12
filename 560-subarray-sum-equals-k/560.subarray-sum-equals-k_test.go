package subarraySum

import "testing"

func TestSubarraySum(t *testing.T) {
	if subarraySum([]int{1, 1, 1}, 2) != 2 {
		t.Error("1 1 1, 2")
	}
}
