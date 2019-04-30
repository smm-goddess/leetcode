package evalrpn

import "testing"

func TestEvalRPN(t *testing.T) {
	if evalRPN([]string{"2", "1", "+", "3", "*"}) != 9 {
		t.Error("First Error")
	}
	if evalRPN([]string{"4", "13", "5", "/", "+"}) != 6 {
		t.Error("Second Error")
	}
	if evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}) != 22 {
		t.Error("Third Error")
	}
}
