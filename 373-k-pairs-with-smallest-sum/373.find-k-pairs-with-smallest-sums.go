package smallestpair

import (
	"math"
)

/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 *
 * https://leetcode.com/problems/find-k-pairs-with-smallest-sums/description/
 *
 * algorithms
 * Medium (33.48%)
 * Total Accepted:    66.5K
 * Total Submissions: 196.5K
 * Testcase Example:  '[1,7,11]\n[2,4,6]\n3'
 *
 * You are given two integer arrays nums1 and nums2 sorted in ascending order
 * and an integer k.
 *
 * Define a pair (u,v) which consists of one element from the first array and
 * one element from the second array.
 *
 * Find the k pairs (u1,v1),(u2,v2) ...(uk,vk) with the smallest sums.
 *
 * Example 1:
 *
 *
 * Input: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
 * Output: [[1,2],[1,4],[1,6]]
 * Explanation: The first 3 pairs are returned from the sequence:
 * [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
 *
 * Example 2:
 *
 *
 * Input: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
 * Output: [1,1],[1,1]
 * Explanation: The first 2 pairs are returned from the sequence:
 * [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
 *
 * Example 3:
 *
 *
 * Input: nums1 = [1,2], nums2 = [3], k = 3
 * Output: [1,3],[2,3]
 * Explanation: All possible pairs are returned from the sequence: [1,3],[2,3]
 *
 *
 */
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if k == 0 || nums1 == nil || len(nums1) == 0 || nums2 == nil || len(nums2) == 0 {
		return nil
	}
	length1, length2 := len(nums1), len(nums2)
	lastPair := make([]int, length1, length1)
	var smallestIndex, smallestSum int
	var ans [][]int
	for k > 0 && lastPair[length1-1] < length2 {
		smallestIndex, smallestSum = 0, math.MaxInt32
		for i := 0; i < length1; i++ {
			if lastPair[i] < length2 && nums1[i]+nums2[lastPair[i]] < smallestSum {
				smallestSum = nums1[i] + nums2[lastPair[i]]
				smallestIndex = i
			}
		}
		ans = append(ans, []int{nums1[smallestIndex], nums2[lastPair[smallestIndex]]})
		lastPair[smallestIndex]++
		k--
	}
	return ans
}
