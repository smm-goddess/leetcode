package singlenumber

/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 *
 * https://leetcode.com/problems/single-number-iii/description/
 *
 * algorithms
 * Medium (56.53%)
 * Total Accepted:    106.1K
 * Total Submissions: 186.4K
 * Testcase Example:  '[1,2,1,3,2,5]'
 *
 * Given an array of numbers nums, in which exactly two elements appear only
 * once and all the other elements appear exactly twice. Find the two elements
 * that appear only once.
 *
 * Example:
 *
 *
 * Input:  [1,2,1,3,2,5]
 * Output: [3,5]
 *
 * Note:
 *
 *
 * The order of the result is not important. So in the above example, [5, 3] is
 * also correct.
 * Your algorithm should run in linear runtime complexity. Could you implement
 * it using only constant space complexity?
 *
 */
func singleNumber(nums []int) []int {
	AxorB := 0
	for _, v := range nums {
		AxorB = AxorB ^ v
	}
	bitFlag := AxorB & (^(AxorB - 1))
	ans := []int{0, 0}
	for _, v := range nums {
		if v&bitFlag == 0 {
			ans[0] = ans[0] ^ v
		} else {
			ans[1] = ans[1] ^ v
		}
	}
	return ans
}
