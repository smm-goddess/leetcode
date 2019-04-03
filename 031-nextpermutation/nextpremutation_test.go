package nextpermutation

import "testing"

func equal(a, b []int) bool {
	if a == nil && b == nil {
		return true
	} else if len(a) != len(b) || a == nil || b == nil {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func TestNextPermutation(t *testing.T) {
	p123 := []int{1, 2, 3}
	nextPermutation(p123)
	if !equal(p123, []int{1, 3, 2}) {
		t.Error("123 fail")
	}
	p321 := []int{3, 2, 1}
	nextPermutation(p321)
	if !equal(p321, []int{1, 2, 3}) {
		t.Error("321 fail")
	}
	p115 := []int{1, 1, 5}
	nextPermutation(p115)
	if !equal(p115, []int{1, 5, 1}) {
		t.Error("115 fail")
	}
	p132 := []int{1, 3, 2}
	nextPermutation(p132)
	if !equal(p132, []int{2, 1, 3}) {
		t.Error("132 fail")
	}
}
