package reverse_str

import "testing"

func TestReverseStr(t *testing.T) {
	if reverseStr("abcdefg", 2) != "bacdfeg" {
		t.Error("error")
	}
}
