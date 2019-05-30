package reversstring

import "testing"

func listEquals(a, b []byte) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil || len(a) != len(b) {
		return false
	} else {
		for i, v := range a {
			if b[i] != v {
				return false
			}
		}
		return true
	}
}
func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(s)
	if !listEquals(s, []byte{'o', 'l', 'l', 'e', 'h'}) {
		t.Error("error")
	}
	s = []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(s)
	if !listEquals(s, []byte{'h', 'a', 'n', 'n', 'a', 'H'}) {
		t.Error("error")
	}
}
