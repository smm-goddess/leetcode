package minheighttree

import (
	"testing"
)

func listEquals(a, b []int) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
		return true
	}
}

func TestMinHeightTree(t *testing.T) {
	if !listEquals(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}), []int{1}) {
		t.Error("error")
	}
	if !listEquals(findMinHeightTrees(6, [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {4, 5}}), []int{4, 3}) {
		t.Error("error")
	}
	if !listEquals(findMinHeightTrees(6, [][]int{{0, 1}, {0, 2}, {0, 3}, {3, 4}, {4, 5}}), []int{3}) {
		t.Error("error")
	}
	if !listEquals(findMinHeightTrees(7, [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}), []int{1, 2}) {
		t.Error("error")
	}
}

//func TestHeightTree(t *testing.T) {
//	if heightTree(4, 3, [][]int{{1, 0}, {1, 2}, {1, 3}}) != 3 {
//		t.Error("error")
//	}
//	if heightTree(6, 4, [][]int{{0, 1}, {0, 2}, {0, 3}, {3, 4}, {4, 5}}) != 4 {
//		t.Error("error")
//	}
//	if heightTree(7, 2, [][]int{{0, 1}, {1, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}}) != 4 {
//		t.Error("error")
//	}
//}
