package findpeakelement

import "math"

/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 *
 * https://leetcode.com/problems/find-peak-element/description/
 *
 * algorithms
 * Medium (40.96%)
 * Total Accepted:    231.4K
 * Total Submissions: 563.3K
 * Testcase Example:  '[1,2,3,1]'
 *
 * A peak element is an element that is greater than its neighbors.
 * 
 * Given an input array nums, where nums[i] ≠ nums[i+1], find a peak element
 * and return its index.
 * 
 * The array may contain multiple peaks, in that case return the index to any
 * one of the peaks is fine.
 * 
 * You may imagine that nums[-1] = nums[n] = -∞.
 * 
 * Example 1:
 * 
 * 
 * Input: nums = [1,2,3,1]
 * Output: 2
 * Explanation: 3 is a peak element and your function should return the index
 * number 2.
 * 
 * Example 2:
 * 
 * 
 * Input: nums = [1,2,1,3,5,6,4]
 * Output: 1 or 5 
 * Explanation: Your function can return either index number 1 where the peak
 * element is 2, 
 * or index number 5 where the peak element is 6.
 * 
 * 
 * Note:
 * 
 * Your solution should be in logarithmic complexity.
 * 
 */
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	before, value, after := 0, 0, 0
	for i := 0; i < len(nums); {
		before, value, after = getValue(nums, i-1), getValue(nums, i), getValue(nums, i+1)
		if value > before && value > after {
			return i
		} else {
			if value > after {
				i = i + 2
			} else {
				i++
			}
		}
	}
	//for i := range nums {
	//	before, value, after = getValue(nums, i-1), getValue(nums, i), getValue(nums, i+1)
	//	if value > before && value > after {
	//		return i
	//	}
	//}
	return math.MinInt32
}

func getValue(nums []int, pos int) int {
	if pos == -1 || pos == len(nums) {
		return math.MinInt32
	} else {
		return nums[pos]
	}
}
