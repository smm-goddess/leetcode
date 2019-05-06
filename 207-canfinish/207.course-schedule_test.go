package canfinish

import "testing"

func TestCanFinish(t *testing.T) {
	if !canFinish(2, [][]int{{0, 1}}) {
		t.Error("0 1 Error")
	}
	if canFinish(2, [][]int{{0, 1}, {1, 0}}) {
		t.Error("0 1, 1 0 Error")
	}
	if canFinish(5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 0}}) {
		t.Error("5 Error")
	}
	if !canFinish(3, [][]int{{2, 1}, {1, 0}}) {
		t.Error("3 Error")
	}
	if !canFinish(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}) {
		t.Error("4 Error")
	}
}
