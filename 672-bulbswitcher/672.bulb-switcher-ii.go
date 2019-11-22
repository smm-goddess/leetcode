package fliplight

/*
 * @lc app=leetcode id=672 lang=golang
 *
 * [672] Bulb Switcher II
 *
 * https://leetcode.com/problems/bulb-switcher-ii/description/
 *
 * algorithms
 * Medium (49.66%)
 * Total Accepted:    9.2K
 * Total Submissions: 18.5K
 * Testcase Example:  '1\n1'
 *
 * There is a room with n lights which are turned on initially and 4 buttons on
 * the wall. After performing exactly m unknown operations towards buttons, you
 * need to return how many different kinds of status of the n lights could be.
 *
 * Suppose n lights are labeled as number [1, 2, 3 ..., n], function of these 4
 * buttons are given below:
 *
 *
 * Flip all the lights.
 * Flip lights with even numbers.
 * Flip lights with odd numbers.
 * Flip lights with (3k + 1) numbers, k = 0, 1, 2, ...
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: n = 1, m = 1.
 * Output: 2
 * Explanation: Status can be: [on], [off]
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: n = 2, m = 1.
 * Output: 3
 * Explanation: Status can be: [on, off], [off, on], [off, off]
 *
 *
 *
 *
 * Example 3:
 *
 *
 * Input: n = 3, m = 1.
 * Output: 4
 * Explanation: Status can be: [off, on, off], [on, off, on], [off, off, off],
 * [off, on, on].
 *
 *
 *
 *
 * Note: n and m both fit in range [0, 1000].
 *
 */
/**
1 => [1 2 3 4]
2 => [11 12 13 14]
3 => [1 2 3 4 123 124 134 234]
4 => [11 12 13 14 23 24 34 1234]
5 => 3
6 => 4
 */
func flipLights(n int, m int) int {
	if n == 0 || m == 0 {
		return 1
	} else if m == 1 {
		if n <= 2 {
			return n + 1
		} else {
			return 4
		}
	} else if m == 2 {
		if n < 3 {
			return 2 * n
		} else {
			return 7
		}
	} else {
		if n < 3 {
			return 2 * n
		} else {
			return 8
		}
	}
}
