package palindromic_substrings

import "testing"

func Test_CountSubString(t *testing.T) {
	if countSubstrings("abc") != 3 {
		t.Error("abc error")
	}
	if countSubstrings("aaa") != 6 {
		t.Error("aaa error")
	}
}
