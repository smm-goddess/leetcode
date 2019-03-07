package main

import "fmt"

func main() {
	fmt.Println(step(1, 99, 2))
	fmt.Println(step(1, 2, 2))
	fmt.Println(uniquePaths(100, 3))
}

func step(r, s, f int) int {
	if f == 0 {
		return r
	} else {
		return step(r*s, s-1, f-1)
	}
}

func uniquePaths(m int, n int) int {
	if m <= 1 || n <= 1 {
		return 1
	}
	a, b := m-1, n-1
	if a > b {
		b, a = a, b
	}
	return step(1, a+b, a) / step(1, a, a)
}
