package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(numSquares(2))
	//fmt.Println(numSquares(3))
	//fmt.Println(numSquares(4))
	fmt.Println(numSquares(142))

}

var m = make(map[int]int)

func numSquares(n int) int {
	if n < 1 {
		return 0
	} else if n >= 1 && n <= 3 {
		return n
	}
	squareRoot := int(math.Floor(math.Sqrt(float64(n))))
	result := math.MaxInt32
	for j := squareRoot; j > 1; j-- {
		remain := n - j*j
		var r int
		if val, yes := m[remain]; yes {
			r = 1 + val
		} else {
			resolve := numSquares(remain)
			m[remain] = resolve
			r = 1 + resolve
		}
		if result > r {
			result = r
		}
	}
	return result
}
