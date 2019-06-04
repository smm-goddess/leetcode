package smallest

import "testing"

func TestKthSmallest(t *testing.T) {
	matrix := [][]int{
		{1, 5, 9},
		{10, 11, 13},
		{12, 13, 15},
	}
	if kthSmallest(matrix, 8) != 13 {
		t.Error("error")
	}
}
