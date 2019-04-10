package generatematrix

import "testing"

func equal(a, b []int) bool {
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
func TestGenerateMatrix(t *testing.T) {
	n := 3
	matrix, m := generateMatrix(n), [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
	if !equals(matrix, m) {
		t.Error("Error 3")
	}
	n = 1
	matrix, m = generateMatrix(n), [][]int{{1}}
	if !equals(matrix, m) {
		t.Error("Error 1")
	}
	generateMatrix(4)
	generateMatrix(5)
	generateMatrix(6)
}
