package mysqrt

import "testing"

func TestMySqrt(t *testing.T) {
	if mySqrt(4) != 2 {
		t.Error("4 error")
	}
	if mySqrt(8) != 2 {
		t.Error("4 error")
	}
}
