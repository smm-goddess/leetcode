package additive

import (
	"testing"
)

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

func TestIsAdditiveNumber(t *testing.T) {
	if !isAdditiveNumber("112358") {
		t.Error("112358 Error")
	}
	if !isAdditiveNumber("199100199") {
		t.Error("199100199 Error")
	}
	if !isAdditiveNumber("199111992") {
		t.Error("199111992 Error")
	}
	if !isAdditiveNumber("111122335588143") {
		t.Error("111122335588143 error")
	}
}

func TestByteSum(t *testing.T) {
	b1 := []byte{'3', '5', '1', '2'}
	b2 := []byte{'4', '6', '7'}
	if !listEquals(byteSum(b1, b2), []byte{'3', '9', '7', '9'}) {
		t.Error("error")
	}
	if !listEquals(byteSum([]byte{'9', '9', '6', '3', '2'}, []byte{'3', '6', '8'}), []byte{'1', '0', '0', '0', '0', '0'}) {
		t.Error("99999 error")
	}
}

func TestEqual(t *testing.T) {
	array := []byte{'4', '5', '1', '3', '2', '6'}
	target := []byte{'3', '2', '6'}
	if !equals(array, target, 3) {
		t.Error("equal error")
	}
}

func TestSubString(t *testing.T) {
	bs := []byte("111122335588143")
	if !listEquals(subString(bs, 0, 2), []byte{'1', '1'}) {
		t.Error("substring error")
	}
}
