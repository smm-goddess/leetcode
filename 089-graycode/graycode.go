package graycode

/*
 * @lc app=leetcode id=89 lang=golang
 *
 * [89] Gray Code
 *
 * https://leetcode.com/problems/gray-code/description/
 *
 * algorithms
 * Medium (45.29%)
 * Total Accepted:    130.7K
 * Total Submissions: 288K
 * Testcase Example:  '2'
 *
 * The gray code is a binary numeral system where two successive values differ
 * in only one bit.
 *
 * Given a non-negative integer n representing the total number of bits in the
 * code, print the sequence of gray code. A gray code sequence must begin with
 * 0.
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: [0,1,3,2]
 * Explanation:
 * 00 - 0
 * 01 - 1
 * 11 - 3
 * 10 - 2
 *
 * For a given n, a gray code sequence may not be uniquely defined.
 * For example, [0,2,3,1] is also a valid gray code sequence.
 *
 * 00 - 0
 * 10 - 2
 * 11 - 3
 * 01 - 1
 *
 *
 * Example 2:
 *
 *
 * Input: 0
 * Output: [0]
 * Explanation: We define the gray code sequence to begin with 0.
 * A gray code sequence of n has size = 2^n, which for n = 0 the size is 2^0 =
 * 1.
 * Therefore, for n = 0 the gray code sequence is [0].
 *
 *
 */
func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	} else if n == 1 {
		return []int{0, 1}
	} else {
		p := power2(1, 2, n-1)
		g := grayCode(n - 1)
		result := make([]int, len(g)*2)
		copy(result, g)
		for i := len(g) - 1; i >= 0; i-- {
			result[len(g)*2-1-i] = p + g[i]
		}
		return result
	}
}

func power2(answer, base, power int) int {
	if power == 0 {
		return answer
	} else if power&1 == 1 {
		return power2(answer*base, base, power-1)
	} else {
		return power2(answer, base*base, power>>1)
	}
}
