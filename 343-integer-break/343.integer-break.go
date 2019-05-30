package integerbreak

/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 *
 * https://leetcode.com/problems/integer-break/description/
 *
 * algorithms
 * Medium (47.52%)
 * Total Accepted:    79.2K
 * Total Submissions: 165.7K
 * Testcase Example:  '2'
 *
 * Given a positive integer n, break it into the sum of at least two positive
 * integers and maximize the product of those integers. Return the maximum
 * product you can get.
 *
 * Example 1:
 *
 *
 *
 * Input: 2
 * Output: 1
 * Explanation: 2 = 1 + 1, 1 × 1 = 1.
 *
 *
 * Example 2:
 *
 *
 * Input: 10
 * Output: 36
 * Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
 *
 * Note: You may assume that n is not less than 2 and not larger than 58.
 *
 *
 */
func integerBreak(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 4
	default:
		numberOf3 := n / 3
		remains := n % 3
		if remains == 0 {
			return pow3(numberOf3)
		} else if remains == 1 {
			return pow3(numberOf3-1) * 4
		} else {
			return pow3(numberOf3) * 2
		}
	}
}

func pow3(n int) int {
	var powIter func(base, temp int, n int) int
	powIter = func(base, temp int, n int) int {
		if n == 0 {
			return temp
		} else if n%2 == 0 {
			return powIter(base*base, temp, n/2)
		} else {
			return powIter(base, temp*base, n-1)
		}
	}
	return powIter(3, 1, n)
}
