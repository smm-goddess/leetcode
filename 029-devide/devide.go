package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(maxNNAnswer(-10, -3))
	//fmt.Println(maxNNAnswer(-10, -1))
	//fmt.Println(maxNNAnswer(-10, -2))
	//fmt.Println(maxNNAnswer(-10, -4))
	//fmt.Println(devide(-10, -3))
	//fmt.Println(devide(1, 2))
	fmt.Println(devide(-2147483648, -1))
	fmt.Println(devide(-2147483648, 1))
	//fmt.Println(devide(-2147483648, 2))
	//fmt.Println(devide(-2147483648, -2147483648))
	//fmt.Println(devide(-2147483648, -2))
}

func devide(dividend, divisor int) int {
	if dividend == 0 {
		return 0
	} else {
		var symbol, result int32
		symbol = 1
		if divisor > 0 && dividend > 0 {
			dividend = -dividend
			divisor = -divisor
		} else if divisor > 0 && dividend < 0 {
			divisor = -divisor
			symbol = -1
		} else if divisor < 0 && dividend > 0 {
			dividend = -dividend
			symbol = -1
		}
		if divisor < dividend {
			return 0
		}
		result = -1
		ans, rem := maxNNAnswer(int32(dividend-divisor), int32(divisor))
		for ans != 0 {
			if math.MinInt32+ans > result {
				if symbol == 1 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
			result = result - ans
			ans, rem = maxNNAnswer(rem, int32(divisor))
		}
		if symbol == 1 && result == math.MinInt32 {
			return math.MaxInt32
		}
		return int(-1 * result * symbol)
	}
}

/**
 * dividend < 0 && divisor < 0
 */
func maxNNAnswer(dividend, divisor int32) (int32, int32) {
	if divisor < dividend {
		return 0, dividend
	} else {
		ans := int32(1)
		for dividend-divisor-divisor <= 0 {
			divisor = divisor + divisor
			ans = ans + ans
		}
		return ans, dividend - divisor
	}
}
