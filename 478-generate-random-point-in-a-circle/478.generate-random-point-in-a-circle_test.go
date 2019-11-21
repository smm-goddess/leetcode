package random_point

import (
	"math"
	"testing"
)

var solution Solution

func dist(solution Solution, b []float64) float64 {
	return math.Sqrt(math.Pow(solution.x-b[0], 2) + math.Pow(solution.y-b[1], 2))
}

func TestRandomPoint(t *testing.T) {
	solution = Constructor(1, 0, 0)
	inCircle(solution, t)
	solution = Constructor(10, 5, -7.5)
	inCircle(solution, t)
	solution = Constructor(0.01, -73839.1, -3289891.3)
	inCircle(solution, t)
}

func inCircle(solution Solution, t *testing.T) {
	var point []float64
	for i := 0; i < 100000; i++ {
		point = solution.RandPoint()
		if dist(solution, point) > solution.radius {
			t.Error(point, " out of the circle ", solution.x, " ", solution.y, " ", solution.radius)
		}
	}
}
