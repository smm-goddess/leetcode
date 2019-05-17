package producearray

/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 *
 * https://leetcode.com/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (54.41%)
 * Total Accepted:    252.9K
 * Total Submissions: 461.2K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an array nums of n integers where n > 1,  return an array output such
 * that output[i] is equal to the product of all the elements of nums except
 * nums[i].
 *
 * Example:
 *
 *
 * Input:  [1,2,3,4]
 * Output: [24,12,8,6]
 *
 *
 * Note: Please solve it without division and in O(n).
 *
 * Follow up:
 * Could you solve it with constant space complexity? (The output array does
 * not count as extra space for the purpose of space complexity analysis.)
 *
 */
func productExceptSelf(nums []int) []int {
	length := len(nums)
	output := make([]int, length, length)
	for i := range output {
		output[i] = 1
		if i >= 1 {
			output[i] = output[i-1] * nums[i-1]
		}
	}
	mul := 1
	for i := length - 2; i >= 0; i-- {
		mul *= nums[i+1]
		output[i] *= mul
	}
	return output
}
