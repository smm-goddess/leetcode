package search

import (
	"testing"
)

func TestSearch(t *testing.T) {
	//nums := []int{2, 5, 6, 0, 0, 1, 2}
	//if !search(nums, 0) {
	//	t.Error("0 Error")
	//}
	//if search(nums, 3) {
	//	t.Error("3 Error")
	//}
	//if !search([]int{1, 3, 1, 1, 1}, 3) {
	//	t.Error("3 Second Error")
	//}
	//if !search([]int{1, 1, 3, 1}, 3) {
	//	t.Error("3 Third Error")
	//}
	if !search([]int{4, 5, 6, 7, 0, 1, 2}, 0) {
		t.Error("0 Second Error")
	}
}
