package rob

import "testing"

func TestRob(t *testing.T) {
	if rob([]int{2, 3, 2}) != 3 {
		t.Error("3 Error1")
	}
	if rob([]int{1, 2, 3, 1}) != 4 {
		t.Error("4 Error")
	}
	if rob([]int{1, 2, 3, 4, 5}) != 8 {
		t.Error("8 Error")
	}
	if rob([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}) != 16 {
		t.Error("16 Error")
	}
}
