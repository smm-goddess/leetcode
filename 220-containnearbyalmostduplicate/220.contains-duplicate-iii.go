package containsnearbyalmostduplicate

/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 *
 * https://leetcode.com/problems/contains-duplicate-iii/description/
 *
 * algorithms
 * Medium (19.55%)
 * Total Accepted:    90.7K
 * Total Submissions: 461.5K
 * Testcase Example:  '[1,2,3,1]\n3\n0'
 *
 * Given an array of integers, find out whether there are two distinct indices
 * i and j in the array such that the absolute difference between nums[i] and
 * nums[j] is at most t and the absolute difference between i and j is at most
 * k.
 *
 *
 * Example 1:
 *
 *
 * Input: nums = [1,2,3,1], k = 3, t = 0
 * Output: true
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1,0,1,1], k = 1, t = 2
 * Output: true
 *
 *
 *
 * Example 3:
 *
 *
 * Input: nums = [1,5,9,1,5,9], k = 2, t = 3
 * Output: false
 *
 *
 *
 *
 */
func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if nums == nil || len(nums) == 0 || k <= 0 || t < 0 {
		return false
	} else {
		for i := 0; i < len(nums)-1; i++ {
			for j := 1; j <= k; j++ {
				if i+j < len(nums) && abs(nums[i], nums[i+j]) <= t {
					return true
				}
			}
		}
		return false
	}
}

func abs(a, b int) (result int) {
	result = a - b
	if result < 0 {
		result = -result
	}
	return
}
