package fizz_buzz

import "testing"

func stringListEquals(a, b []string) bool {
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

func TestFizzbuzz(t *testing.T) {
	if !stringListEquals(fizzBuzz(15), []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}) {
		t.Error("error")
	}
}
