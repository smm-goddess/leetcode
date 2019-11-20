package count_prime_set_bits

import "testing"

type BitsCase struct {
	Num   int
	Count int
}

type CountPrimeCase struct {
	L     int
	R     int
	Count int
}

var bitsCases []BitsCase
var countCases []CountPrimeCase

func init() {
	bitsCases = []BitsCase{
		{10, 2}, {6, 2}, {7, 3,}, {9, 2,},
		{11, 3}, {12, 2}, {13, 3}, {14, 3}, {15, 4},
	}
	countCases = []CountPrimeCase{
		{6, 10, 4}, {10, 15, 5},
	}
}

func TestCountPrime(t *testing.T) {
	for _, cs := range countCases {
		if countPrimeSetBits(cs.L, cs.R) != cs.Count {
			t.Error(cs.L, " ", cs.R, " error")
		}
	}
}

func TestCountBits(t *testing.T) {
	for _, cs := range bitsCases {
		if countBits(cs.Num) != cs.Count {
			t.Error(cs.Num, " error")
		}
	}
}
