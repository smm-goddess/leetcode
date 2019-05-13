package findkthlargest

import (
	"math"
)

/*
 * @lc app=leetcode id=215 lang=golang
 *
 * [215] Kth Largest Element in an Array
 *
 * https://leetcode.com/problems/kth-largest-element-in-an-array/description/
 *
 * algorithms
 * Medium (46.83%)
 * Total Accepted:    365.3K
 * Total Submissions: 769.1K
 * Testcase Example:  '[3,2,1,5,6,4]\n2'
 *
 * Find the kth largest element in an unsorted array. Note that it is the kth
 * largest element in the sorted order, not the kth distinct element.
 *
 * Example 1:
 *
 *
 * Input: [3,2,1,5,6,4] and k = 2
 * Output: 5
 *
 *
 * Example 2:
 *
 *
 * Input: [3,2,3,1,2,4,5,5,6] and k = 4
 * Output: 4
 *
 * Note:
 * You may assume k is always valid, 1 ≤ k ≤ array's length.
 *
 */
func findKthLargest(nums []int, k int) int {
	minHeap := nums[:k]
	minHeapify(minHeap)
	for i := k; i < len(nums); i++ {
		if nums[i] > minHeap[0] {
			minHeap[0] = nums[i]
			minHeapify(minHeap)
		}
	}
	return minHeap[0]
}

func minHeapify(data []int) {
	if len(data) == 1 {
		return
	} else if len(data) == 2 {
		if data[0] > data[1] {
			data[0], data[1] = data[1], data[0]
		}
		return
	} else if len(data) == 3 {
		data[0], data[1], data[2] = sort(data[0], data[1], data[2])
		return
	}
	length := len(data)
	var rcValue, lcValue, rtValue, rcIndex, lcIndex int
	for root := (length - 2) >> 1; root >= 0; root-- {
		lcIndex = 2*root + 1
		rcIndex = 2*root + 2
		lcValue = data[lcIndex]
		rtValue = data[root]
		if rcIndex < length {
			rcValue = data[rcIndex]
		} else {
			rcValue = math.MaxInt32
		}
		rtValue, lcValue, rcValue = sort(rcValue, lcValue, rtValue)
		data[root] = rtValue
		data[lcIndex] = lcValue
		if rcIndex < length {
			data[rcIndex] = rcValue
		}
	}
}

func sort(a, b, c int) (min, middle, max int) {
	max, middle, min = a, b, c
	if max < middle {
		max, middle = middle, max
	}
	if max < min {
		max, min = min, max
	}
	if min > middle {
		min, middle = middle, min
	}
	return
}
