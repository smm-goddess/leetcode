package search

/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 *
 * https://leetcode.com/problems/search-in-rotated-sorted-array-ii/description/
 *
 * algorithms
 * Medium (32.52%)
 * Total Accepted:    165K
 * Total Submissions: 507.1K
 * Testcase Example:  '[2,5,6,0,0,1,2]\n0'
 *
 * Suppose an array sorted in ascending order is rotated at some pivot unknown
 * to you beforehand.
 *
 * (i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).
 *
 * You are given a target value to search. If found in the array return true,
 * otherwise return false.
 *
 * Example 1:
 *
 *
 * Input: nums = [2,5,6,0,0,1,2], target = 0
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [2,5,6,0,0,1,2], target = 3
 * Output: false
 *
 * Follow up:
 *
 *
 * This is a follow up problem to Search in Rotated Sorted Array, where nums
 * may contain duplicates.
 * Would this affect the run-time complexity? How and why?
 *
 *
 */
func search(nums []int, target int) bool {
	if nums == nil || len(nums) == 0 {
		return false
	}
	start, end := 0, len(nums)-1
	middle := start + (end-start)/2
	if nums[start] == target || nums[end] == target || nums[middle] == target {
		return true
	} else {
		if nums[middle] > nums[start] {
			if nums[start] < target && nums[middle] > target {
				return binarySearch(nums[start:middle], target)
			} else {
				return search(nums[middle:end], target)
			}
		} else if nums[middle] < nums[start] {
			if nums[middle] < target && nums[end] > target {
				return binarySearch(nums[middle:end], target)
			} else {
				return search(nums[start:middle], target)
			}
		} else {
			return search(nums[start:middle], target) || search(nums[middle:end], target)
		}
	}
}

func binarySearch(nums []int, target int) bool {
	start, end := 0, len(nums)-1
	middle := start + (end-start)/2
	for start < end {
		if nums[start] == target || nums[middle] == target || nums[end] == target {
			return true
		} else {
			if nums[middle] > target {
				start, end = start+1, middle-1
			} else {
				start, end = middle+1, end-1
			}
		}
	}
	return false
}
