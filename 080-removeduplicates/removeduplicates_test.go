package removeduplicates

import "testing"

func equals(nums, removed []int, same int) bool {
	for i := 0; i < same; i++ {
		if nums[i] != removed[i] {
			return false
		}
	}
	return true
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	num := removeDuplicates(nums)
	if num == 5 && equals(nums, []int{1, 1, 2, 2, 3}, num) {
		t.Log("First Pass0")
	} else {
		t.Error("First Error")
	}
	nums = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	num = removeDuplicates(nums)
	if num == 7 && equals(nums, []int{0, 0, 1, 1, 2, 3, 3}, num) {
		t.Log("Second Pass")
	} else {
		t.Error("Second Error")
	}
}
