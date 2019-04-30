package compareversion

import "testing"

func TestCompareVersion(t *testing.T) {
	if compareVersion("0.1", "1.1") != -1 {
		t.Error("0.1 1.1 Error")
	}
	if compareVersion("1.0.1", "1") != 1 {
		t.Error("1.0.1 1 Error")
	}
	if compareVersion("7.5.2.4", "7.5.3") != -1 {
		t.Error("7.5.2.4 7.5.3 Error")
	}
	if compareVersion("1.01", "1.001") != 0 {
		t.Error("1.01 1.001 Error")
	}
}
