package main

import "fmt"

func main() {
	fmt.Println(myPow(3, -2))
}

func myPow(x float64, n int) float64 {
	var powIter func(base, temp float64, n int) float64
	powIter = func(base, temp float64, n int) float64 {
		if n == 0 {
			return temp
		} else if n%2 == 0 {
			return powIter(base*base, temp, n/2)
		} else {
			return powIter(base, temp*base, n-1)
		}
	}
	if n < 0 {
		x, n = 1/x, -n
	}
	return powIter(x, 1, n)
}

