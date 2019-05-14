package computearea

import "testing"

func TestComputeArea(t *testing.T) {
	if computeArea(-3, 0, 3, 4, 0, -1, 9, 2) != 45 {
		t.Error("45 Fail")
	}
	if computeArea(0, -1, 9, 2, -3, 0, 3, 4) != 45 {
		t.Error("45 Fail")
	}
	if computeArea(0, 0, 2, 3, 4, 3, 8, 8) != 26 {
		t.Error("26 Fail")
	}
	if computeArea(0, 0, 8, 8, 2, 2, 4, 4, ) != 64 {
		t.Error("64 Error")
	}
	if computeArea(0, 0, 8, 8, 2, 2, 4, 10) != 68 {
		t.Error("68 Error")
	}
}
