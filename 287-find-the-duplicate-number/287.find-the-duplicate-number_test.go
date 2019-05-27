package findduplicate

import "testing"

func TestFindDuplicate(t *testing.T) {
	if findDuplicate([]int{1, 3, 4, 2, 2}) != 2 {
		t.Error("2 error")
	}
	if findDuplicate([]int{3, 1, 3, 4, 2}) != 3 {
		t.Error("3 error")
	}
	if findDuplicate([]int{2, 2, 2, 2, 2}) != 2 {
		t.Error(" 2 error")
	}
}
