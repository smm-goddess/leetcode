package rangebitwiseand

/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 *
 * https://leetcode.com/problems/bitwise-and-of-numbers-range/description/
 *
 * algorithms
 * Medium (35.67%)
 * Total Accepted:    80.3K
 * Total Submissions: 224.5K
 * Testcase Example:  '5\n7'
 *
 * Given a range [m, n] where 0 <= m <= n <= 2147483647, return the bitwise AND
 * of all numbers in this range, inclusive.
 * 
 * Example 1:
 * 
 * 
 * Input: [5,7]
 * Output: 4
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [0,1]
 * Output: 0
 */
func rangeBitwiseAnd(m int, n int) int {
	if m == 0 {
		return 0
	}
	moveFactor := 1
	for m != n {
		m = m >> 1
		n = n >> 1
		moveFactor = moveFactor << 1
	}
	return m * moveFactor
}
