package findkthlargest

import "testing"

func TestFindKthLargest(t *testing.T) {
	if findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2) != 5 {
		t.Error("5 Error")
	}
	if findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4) != 4 {
		t.Error("4 Error")
	}
}
