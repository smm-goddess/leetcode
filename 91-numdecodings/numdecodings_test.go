package numdecodings

import "testing"

func TestNumDecodings(t *testing.T) {
	if n := numDecodings("12"); n != 2 {
		t.Error("12 error", n)
	}
	if n := numDecodings("227"); n != 2 {
		t.Error("227 error", n)
	}
	if n := numDecodings("101"); n != 1 {
		t.Error("101 error", n)
	}
}
