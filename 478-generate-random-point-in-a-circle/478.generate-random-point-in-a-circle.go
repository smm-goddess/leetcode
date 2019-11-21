package random_point

import (
	"math"
	"math/rand"
	"time"
)

/*
 * @lc app=leetcode id=478 lang=golang
 *
 * [478] Generate Random Point in a Circle
 *
 * https://leetcode.com/problems/generate-random-point-in-a-circle/description/
 *
 * algorithms
 * Medium (37.89%)
 * Total Accepted:    7.7K
 * Total Submissions: 20.1K
 * Testcase Example:  '["Solution", "randPoint", "randPoint", "randPoint"]\n[[1.0, 0.0, 0.0], [], [], []]'
 *
 * Given the radius and x-y positions of the center of a circle, write a
 * function randPoint which generates a uniform random point in the circle.
 *
 * Note:
 *
 *
 * input and output values are in floating-point.
 * radius and x-y position of the center of the circle is passed into the class
 * constructor.
 * a point on the circumference of the circle is considered to be in the
 * circle.
 * randPoint returns a size 2 array containing x-position and y-position of the
 * random point, in that order.
 *
 *
 *
 * Example 1:
 *
 *
 * Input:
 * ["Solution","randPoint","randPoint","randPoint"]
 * [[1,0,0],[],[],[]]
 * Output:
 * [null,[-0.72939,-0.65505],[-0.78502,-0.28626],[-0.83119,-0.19803]]
 *
 *
 *
 * Example 2:
 *
 *
 * Input:
 * ["Solution","randPoint","randPoint","randPoint"]
 * [[10,5,-7.5],[],[],[]]
 * Output:
 * [null,[11.52438,-8.33273],[2.46992,-16.21705],[11.13430,-12.42337]]
 *
 *
 * Explanation of Input Syntax:
 *
 * The input is two lists: the subroutines called and their arguments.
 * Solution's constructor has three arguments, the radius, x-position of the
 * center, and y-position of the center of the circle. randPoint has no
 * arguments. Arguments are always wrapped with a list, even if there aren't
 * any.
 *
 *
 */
type Solution struct {
	x      float64
	y      float64
	radius float64
}

func Constructor(radius, x_center, y_center float64) Solution {
	return Solution{
		x_center,
		y_center,
		radius,
	}
}

func (this *Solution) RandPoint() []float64 {
	rand.Seed(time.Now().UnixNano())
	alpha := 2 * math.Pi * rand.Float64()
	dist := math.Sqrt(rand.Float64()) * this.radius
	return []float64{this.x + dist*math.Cos(alpha), this.y + dist*math.Sin(alpha)}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
