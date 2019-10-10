package shortest_unsorted_continuous_subarray

/*
 * @lc app=leetcode id=581 lang=golang
 *
 * [581] Shortest Unsorted Continuous Subarray
 *
 * https://leetcode.com/problems/shortest-unsorted-continuous-subarray/description/
 *
 * algorithms
 * Easy (29.80%)
 * Total Accepted:    83.5K
 * Total Submissions: 273.4K
 * Testcase Example:  '[2,6,4,8,10,9,15]'
 *
 * Given an integer array, you need to find one continuous subarray that if you
 * only sort this subarray in ascending order, then the whole array will be
 * sorted in ascending order, too.
 *
 * You need to find the shortest such subarray and output its length.
 *
 * Example 1:
 *
 * Input: [2, 6, 4, 8, 10, 9, 15]
 * Output: 5
 * Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make
 * the whole array sorted in ascending order.
 *
 *
 *
 * Note:
 *
 * Then length of the input array is in range [1, 10,000].
 * The input array may contain duplicates, so ascending order here means .
 *
 *
 */
func findUnsortedSubarray(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	minIndex, maxIndex := 0, len(nums)-1
	for index := 0; index < len(nums); index++ {
		if nums[index] < nums[minIndex] {
			minIndex = index
		}
		if nums[index] > nums[maxIndex] {
			maxIndex = index
		}
	}
	if minIndex == 0 && maxIndex == len(nums)-1 {
		return findUnsortedSubarray(nums[1 : len(nums)-1])
	} else if minIndex == 0 {
		return findUnsortedSubarray(nums[1:])
	} else if maxIndex == len(nums)-1 {
		return findUnsortedSubarray(nums[:len(nums)-1])
	} else {
		return len(nums)
	}
}
