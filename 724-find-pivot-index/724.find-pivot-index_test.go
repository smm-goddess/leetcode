package find_pivot_index

import "testing"

func Test_pivotIndex(t *testing.T) {
	//if pivotIndex([]int{1, 7, 3, 6, 5, 6}) != 3 {
	//	t.Error("3 error")
	//}
	//if pivotIndex([]int{1, 2, 3}) != -1 {
	//	t.Error("-1 error")
	//}
	if pivotIndex([]int{-1, -1, 0, 1, 1, 0}) != 5 {
		t.Error("5 error")
	}
}
