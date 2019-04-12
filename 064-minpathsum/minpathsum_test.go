package minpathsum

import "testing"

func TestMinPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	if minPathSum(grid) != 7 {
		t.Error("Min Path Error")
	}
	grid = [][]int{{3}}
	if minPathSum(grid) != 3 {
		t.Error("Min Path Error 2")
	}
}
