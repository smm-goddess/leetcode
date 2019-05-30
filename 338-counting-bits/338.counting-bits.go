package countbits

/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 *
 * https://leetcode.com/problems/counting-bits/description/
 *
 * algorithms
 * Medium (64.30%)
 * Total Accepted:    167.7K
 * Total Submissions: 259.3K
 * Testcase Example:  '2'
 *
 * Given a non negative integer number num. For every numbers i in the range 0
 * ≤ i ≤ num calculate the number of 1's in their binary representation and
 * return them as an array.
 *
 * Example 1:
 *
 *
 * Input: 2
 * Output: [0,1,1]
 *
 * Example 2:
 *
 *
 * Input: 5
 * Output: [0,1,1,2,1,2]
 *
 *
 * Follow up:
 *
 *
 * It is very easy to come up with a solution with run time
 * O(n*sizeof(integer)). But can you do it in linear time O(n) /possibly in a
 * single pass?
 * Space complexity should be O(n).
 * Can you do it like a boss? Do it without using any builtin function like
 * __builtin_popcount in c++ or in any other language.
 *
 */
func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	} else if num == 1 {
		return []int{0, 1}
	} else {
		count := make([]int, num+1, num+1)
		count[0], count[1] = 0, 1
		power, nextPower := 2, 4
		for i := 2; i <= num; i++ {
			if i >= nextPower {
				power, nextPower = nextPower, nextPower*2
			}
			count[i] = 1 + count[i-power]
		}
		return count
	}
}
