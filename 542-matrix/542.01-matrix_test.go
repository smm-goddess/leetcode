package matrix

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

func listOfListEquals(a, b [][]int) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if !listEquals(b[i], v) {
				return false
			}
		}
		return true
	}
}

func TestUpdateMatrix(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	answer := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	//if !listOfListEquals(updateMatrix(matrix), answer) {
	//	t.Error("First Error")
	//}
	matrix = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
	}
	answer = [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 2, 1},
	}
	if !listOfListEquals(updateMatrix(matrix), answer) {
		t.Error("Second Error")
	}
}
