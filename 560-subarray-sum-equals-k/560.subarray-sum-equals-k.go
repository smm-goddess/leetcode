package subarraySum

/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 *
 * https://leetcode.com/problems/subarray-sum-equals-k/description/
 *
 * algorithms
 * Medium (43.10%)
 * Total Accepted:    159K
 * Total Submissions: 369K
 * Testcase Example:  '[1,1,1]\n2'
 *
 * Given an array of integers and an integer k, you need to find the total
 * number of continuous subarrays whose sum equals to k.
 *
 * Example 1:
 *
 * Input:nums = [1,1,1], k = 2
 * Output: 2
 *
 *
 *
 * Note:
 *
 * The length of the array is in range [1, 20,000].
 * The range of numbers in the array is [-1000, 1000] and the range of the
 * integer k is [-1e7, 1e7].
 *
 *
 *
 */
func subarraySum(nums []int, k int) int {
	result := 0
	for length := 1; length <= len(nums); length++ {
		var sum int
		for i := 0; i < length; i++ {
			sum += nums[i]
		}
		if sum == k {
			result++
		}
		for lastIndex := length; lastIndex < len(nums); lastIndex++ {
			sum += nums[lastIndex]
			sum -= nums[lastIndex-length]
			if sum == k {
				result++
			}
		}
	}
	return result
}
