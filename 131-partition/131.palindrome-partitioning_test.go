package partition

import "testing"

func equal(a, b []string) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, s := range a {
			if b[i] != s {
				return false
			}
		}
		return true
	}
}

func equals(a, b [][]string) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, s := range a {
			if !equal(b[i], s) {
				return false
			}
		}
		return true
	}
}

func TestPartition(t *testing.T) {
	// s := "aab"
	// if !equals(partition(s), [][]string{{"a", "a", "b"}, {"aa", "b"}}) {
	// 	t.Error("Error")
	// }
	partition("aaaa")
}
