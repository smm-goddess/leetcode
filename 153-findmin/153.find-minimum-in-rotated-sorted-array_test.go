package findmin

import "testing"

func TestFindMin(t *testing.T) {
	if findMin([]int{4, 5, 6, 7, 0, 1, 2}) != 0 {
		t.Error("1 Error")
	}
	if findMin([]int{3, 4, 5, 1, 2}) != 1 {
		t.Error("2 Error")
	}
	if findMin([]int{2, 1}) != 1 {
		t.Error("3 Error")
	}
	if findMin([]int{1, 2}) != 1 {
		t.Error("3 Error")
	}
}
