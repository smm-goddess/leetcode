package inttoroman

import (
	"fmt"
	"testing"
)

var intRomanMap = map[int]string{3: "III", 4: "IV", 9: "IX", 10: "X", 100: "C", 1000: "M", 58: "LVIII", 1994: "MCMXCIV", 3888: "MMMDCCCLXXXVIII"}

func TestIntToRoman(t *testing.T) {
	for k, v := range intRomanMap {
		if intToRoman2(k) != v {
			t.Error(fmt.Sprintf("%d Error", k))
		}
	}
}
