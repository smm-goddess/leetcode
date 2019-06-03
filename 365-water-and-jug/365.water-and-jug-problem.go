package measurewater

/*
 * @lc app=leetcode id=365 lang=golang
 *
 * [365] Water and Jug Problem
 *
 * https://leetcode.com/problems/water-and-jug-problem/description/
 *
 * algorithms
 * Medium (28.81%)
 * Total Accepted:    27.8K
 * Total Submissions: 96.3K
 * Testcase Example:  '3\n5\n4'
 *
 * You are given two jugs with capacities x and y litres. There is an infinite
 * amount of water supply available. You need to determine whether it is
 * possible to measure exactly z litres using these two jugs.
 *
 * If z liters of water is measurable, you must have z liters of water
 * contained within one or both buckets by the end.
 *
 * Operations allowed:
 *
 *
 * Fill any of the jugs completely with water.
 * Empty any of the jugs.
 * Pour water from one jug into another till the other jug is completely full
 * or the first jug itself is empty.
 *
 *
 * Example 1: (From the famous "Die Hard" example)
 *
 *
 * Input: x = 3, y = 5, z = 4
 * Output: True
 *
 *
 * Example 2:
 *
 *
 * Input: x = 2, y = 6, z = 5
 * Output: False
 *
 */
func canMeasureWater(x int, y int, z int) bool {
	if z == 0 || (x == 0 && y == 0) || (z < 0 || z > x+y) {
		return false
	}
	var b int
	if x == 0 {
		b = y
	} else if y == 0 {
		b = x
	} else {
		a := x
		b = y
		c := a % b
		for c != 0 {
			a, b = b, c
			c = a % b
		}
	}
	return z%b == 0
}
