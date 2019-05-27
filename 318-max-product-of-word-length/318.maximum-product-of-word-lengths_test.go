package maxproduct

import "testing"

func TestMaxProduct(t *testing.T) {
	if maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}) != 16 {
		t.Error("16 error")
	}
	if maxProduct([]string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"}) != 4 {
		t.Error("4 error")
	}
	if maxProduct([]string{"a", "aa", "aaa", "aaaa"}) != 0 {
		t.Error("0 error")
	}
}
