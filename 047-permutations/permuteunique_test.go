package permuteunique

import "testing"

func equals(answer, supposed [][]int) bool {
	if answer == nil && supposed == nil {
		return true
	} else if answer == nil || supposed == nil || len(answer) != len(supposed) {
		return false
	} else {
		for i, v := range answer {
			if !equal(supposed[i], v) {
				return false
			}
		}
		return true
	}
}

func TestPermuteUnique(t *testing.T) {
	array := []int{1, 1, 2}
	answer, supposed := permuteUnique(array), [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}
	t.Log(answer)
	if !equals(answer, supposed) {
		t.Error("Not Right")
	}
	array = []int{1, 1, 2, 2}
	answer, supposed = permuteUnique(array), [][]int{{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1}}
	t.Log(len(answer))
	if !equals(answer, supposed) {
		t.Error("Not Right.")
	}
}
