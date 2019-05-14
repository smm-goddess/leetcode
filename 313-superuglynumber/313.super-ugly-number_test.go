package uglynumber

import "testing"

func TestUglyNumber(t *testing.T) {
	if nthSuperUglyNumber(12, []int{2, 7, 13, 19}) != 32 {
		t.Error("32 Error")
	}
	if nthSuperUglyNumber(1, []int{2, 7, 13, 19}) != 1 {
		t.Error("1 Error")
	}
	if nthSuperUglyNumber(2, []int{2, 7, 13, 19}) != 2 {
		t.Error("2 Error")
	}
	if nthSuperUglyNumber(100000, []int{7, 19, 29, 37, 41, 47, 53, 59, 61, 79, 83, 89, 101, 103, 109, 127, 131, 137, 139, 157, 167, 179, 181, 199, 211, 229, 233, 239, 241, 251}) != 1092889481 {
		t.Error("long Error")
	}
}
