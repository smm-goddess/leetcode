package combinationsum

import "testing"

func equal(a, b []int) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func equals(a, b [][]int) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if !equal(b[i], v) {
				return false
			}
		}
	}
	return true
}

func TestCombinationSum(t *testing.T) {
	candidate, target, supposed := []int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}
	combination := combinationSum(candidate, target)
	t.Log(combination)
	if !equals(combination, supposed) {
		t.Error("combination test 1 fail")
	}
	t.Log("first test finished")
	candidate, target, supposed = []int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}
	combination = combinationSum(candidate, target)
	t.Log(combination)
	if !equals(combination, supposed) {
		t.Error("combination test 2 fail")
	}
	candidate, target, supposed = []int{2, 3, 7}, 18, [][]int{{2, 2, 2, 2, 2, 2, 2, 2, 2}, {2, 2, 2, 2, 2, 2, 3, 3}, {2, 2, 2, 2, 3, 7}, {2, 2, 2, 3, 3, 3, 3}, {2, 2, 7, 7}, {2, 3, 3, 3, 7}, {3, 3, 3, 3, 3, 3}}
	combination = combinationSum(candidate, target)
	t.Log(combination)
	if !equals(combination, supposed) {
		t.Error("combination test 3 fail")
	}
}
