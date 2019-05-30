package topk

/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 *
 * https://leetcode.com/problems/top-k-frequent-elements/description/
 *
 * algorithms
 * Medium (54.06%)
 * Total Accepted:    204.5K
 * Total Submissions: 372.8K
 * Testcase Example:  '[1,1,1,2,2,3]\n2'
 *
 * Given a non-empty array of integers, return the k most frequent elements.
 *
 * Example 1:
 *
 *
 * Input: nums = [1,1,1,2,2,3], k = 2
 * Output: [1,2]
 *
 *
 *
 * Example 2:
 *
 *
 * Input: nums = [1], k = 1
 * Output: [1]
 *
 *
 * Note:
 *
 *
 * You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
 * Your algorithm's time complexity must be better than O(n log n), where n is
 * the array's size.
 *
 *
 */
func topKFrequent(nums []int, k int) []int {
	frequent := make(map[int]int)
	for _, n := range nums {
		if fre, ok := frequent[n]; ok {
			frequent[n] = fre + 1
		} else {
			frequent[n] = 1
		}
	}
	minHeap, heapData, index := make([]int, k, k), make([]int, k, k), 0
	for num, cnt := range frequent {
		if index < k {
			minHeap[index] = cnt
			heapData[index] = num
			index++
		} else {
			buildHeap(minHeap, heapData)
			if minHeap[0] < cnt {
				minHeap[0] = cnt
				heapData[0] = num
			}
		}
	}
	return heapData
}

func buildHeap(minHeap, heapData []int) {
	length := len(minHeap)
	for i := 0; i < len(minHeap)>>1; i++ {
		if 2*i+1 <= length {
			if minHeap[i] > minHeap[2*i+1] {
				minHeap[i], minHeap[2*i+1] = minHeap[2*i+1], minHeap[i]
				heapData[i], heapData[2*i+1] = heapData[2*i+1], heapData[i]
			}
			if 2*i+2 < length {
				if minHeap[i] > minHeap[2*i+2] {
					minHeap[i], minHeap[2*i+2] = minHeap[2*i+2], minHeap[i]
					heapData[i], heapData[2*i+2] = heapData[2*i+2], heapData[i]
				}
			}
		}
	}
}
