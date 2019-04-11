package getpermutation

import "testing"

func TestGetPermutation(t *testing.T) {
	if getPermutation(3, 3) != "213" {
		t.Error("3,3 Error")
	}
	if getPermutation(4, 9) != "2314" {
		t.Error("4,9 Error")
	}
	if getPermutation(4, 1) != "1234" {
		t.Error("4,1 Error")
	}
}
