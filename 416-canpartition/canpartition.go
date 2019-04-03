package canpartition

import "fmt"

/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 *
 * https://leetcode.com/problems/partition-equal-subset-sum/description/
 *
 * algorithms
 * Medium (39.96%)
 * Total Accepted:    79.1K
 * Total Submissions: 197.3K
 * Testcase Example:  '[1,5,11,5]'
 *
 * Given a non-empty array containing only positive integers, find if the array
 * can be partitioned into two subsets such that the sum of elements in both
 * subsets is equal.
 *
 * Note:
 *
 *
 * Each of the array element will not exceed 100.
 * The array size will not exceed 200.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1, 5, 11, 5]
 *
 * Output: true
 *
 * Explanation: The array can be partitioned as [1, 5, 5] and [11].
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [1, 2, 3, 5]
 *
 * Output: false
 *
 * Explanation: The array cannot be partitioned into equal sum subsets.
 *
 *
 *
 *
 */
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum = sum + v
	}
	if sum&1 == 1 {
		return false
	}
	target := sum >> 1
	can := make([]bool, target+1)
	can[0] = true
	for _, v := range nums {
		for j := target; j >= v; j-- {
			can[j] = can[j] || can[j-v]
		}
	}
	fmt.Println(can)
	return can[target]
}
