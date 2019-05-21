package wordpattern

import "testing"

func TestWordPattern(t *testing.T) {
	if !wordPattern("abba", "dog cat cat dog") {
		t.Error("abba first")
	}
	if wordPattern("abba", "dog cat cat fish") {
		t.Error("abba second")
	}
	if wordPattern("aaaa", "dog cat cat dog") {
		t.Error("aaaa error")
	}
	if wordPattern("abba", "dog dog dog dog") {
		t.Error("abba third")
	}
}
