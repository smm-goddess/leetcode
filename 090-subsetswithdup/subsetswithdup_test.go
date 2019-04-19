package subsetswithdup

import "testing"

func TestSubsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 2}
	subsetsWithDup(nums)
	subsetsWithDup([]int{1, 1, 1, 1})
	subsetsWithDup([]int{})
}
