package mysqrt

import "math"

func mySqrt(x int) int {
	root := float64(1)
	n := (float64(x)/root + root) / 2
	for math.Floor(n) != math.Floor(root) {
		root = n
		n = (float64(x)/root + root) / 2
	}
	return int(math.Floor(n))
}
