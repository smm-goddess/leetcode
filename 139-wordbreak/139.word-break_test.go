package wordbreak

import "testing"

func TestWordBreak(t *testing.T) {
	if !wordBreak("leetcode", []string{"leet", "code"}) {
		t.Error("leetcode Error")
	}
	if !wordBreak("applepenapple", []string{"apple", "pen"}) {
		t.Error("applepenapple Error")
	}
	if wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) {
		t.Error("catsandog Error")
	}
	if !wordBreak("aaaaaaa", []string{"aaaa", "aaa"}) {
		t.Error("aaaaaaaa Error")
	}
	if wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}) {
		t.Error("very long error")
	}
	if !wordBreak("cars", []string{"car", "ca", "rs"}) {
		t.Error("cars Error")
	}
}
