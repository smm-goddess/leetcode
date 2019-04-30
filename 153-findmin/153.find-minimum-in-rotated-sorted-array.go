package findmin

/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 *
 * https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/
 *
 * algorithms
 * Medium (42.70%)
 * Total Accepted:    276.1K
 * Total Submissions: 644.9K
 * Testcase Example:  '[3,4,5,1,2]'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 * 
 * (i.e.,  [0,1,2,4,5,6,7] might become  [4,5,6,7,0,1,2]).
 * 
 * Find the minimum element.
 * 
 * You may assume no duplicate exists in the array.
 * 
 * Example 1:
 * 
 * 
 * Input: [3,4,5,1,2] 
 * Output: 1
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [4,5,6,7,0,1,2]
 * Output: 0
 * 
 * 
 */
func findMin(nums []int) int {
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	} else if len(nums) == 2 {
		return nums[1]
	} else {
		mid := len(nums) >> 1
		if nums[mid] > nums[len(nums)-1] {
			return findMin(nums[mid:])
		} else {
			return findMin(nums[:mid+1])
		}
	}
}
