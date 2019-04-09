package multiply

import "testing"

func TestMultiply(t *testing.T) {
	s1, s2 := "0", "123"
	if multiply(s1, s2) != "0" {
		t.Error("0 * 123 Error")
	}
	s1, s2 = "123", "456"
	if multiply(s1, s2) != "56088" {
		t.Error("123 * 456 Error")
	}
	s1, s2 = "3242342342", "99898213213"
	if multiply(s1, s2) != "323904206590653764846" {
		t.Error("big Error")
	}
}
