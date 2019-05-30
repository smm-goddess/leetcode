package topk

import "testing"

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
func TestTopK(t *testing.T) {
	if !listEquals(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2), []int{1, 2}) {
		t.Error("1 2 error")
	}
	if !listEquals(topKFrequent([]int{1}, 1), []int{1}) {
		t.Error("1 error")
	}
	if !listEquals(topKFrequent([]int{-1, -1}, 1), []int{-1}) {
		t.Error("-1 error")
	}
	if !listEquals(topKFrequent([]int{3, 0, 0, 1}, 1), []int{0}) {
		t.Error("3 error")
	}
	if !listEquals(topKFrequent([]int{4, 1, -1, 2, -1, 2, 3}, 2), []int{-1, 2}) {
		t.Error("-1 2 error")
	}
}
