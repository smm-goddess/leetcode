package superpow

import "testing"

func TestSuperPow(t *testing.T) {
	//if superPow(2, []int{1, 0}) != 1024 {
	//	t.Error("2 [1,0]")
	//}
	//if superPow(2, []int{3}) != 8 {
	//	t.Error("2 [3]")
	//}
	if superPow(2, []int{1, 0, 0, 0}) != 345 {
		t.Error("2 [1,0,0,0]")
	}
}
