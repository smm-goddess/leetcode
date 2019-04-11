package getpermutation

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode id=60 lang=golang
 *
 * [60] Permutation Sequence
 *
 * https://leetcode.com/problems/permutation-sequence/description/
 *
 * algorithms
 * Medium (32.65%)
 * Total Accepted:    133.4K
 * Total Submissions: 408.1K
 * Testcase Example:  '3\n3'
 *
 * The set [1,2,3,...,n] contains a total of n! unique permutations.
 *
 * By listing and labeling all of the permutations in order, we get the
 * following sequence for n = 3:
 *
 *
 * "123"
 * "132"
 * "213"
 * "231"
 * "312"
 * "321"
 *
 *
 * Given n and k, return the k^th permutation sequence.
 *
 * Note:
 *
 *
 * Given n will be between 1 and 9 inclusive.
 * Given k will be between 1 and n! inclusive.
 *
 *
 * Example 1:
 *
 *
 * Input: n = 3, k = 3
 * Output: "213"
 *
 *
 * Example 2:
 *
 *
 * Input: n = 4, k = 9
 * Output: "2314"
 *
 *
 */
func getPermutation(n int, k int) string {
	nums := make([]int, n)
	for i := 1; i <= n; i++ {
		nums[i-1] = i
	}
	for i := 1; i < k; i++ {
		nextPermutation(nums)
	}
	return strings.Trim(strings.Replace(fmt.Sprint(nums), " ", "", -1), "[]")
}

func nextPermutation(nums []int) {
	length := len(nums)
	i := length - 1
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			i--
			break
		}
	}
	start := i
	for j := length - 1; j > 0; j-- {
		if nums[j] > nums[i] {
			start = i + 1
			nums[i], nums[j] = nums[j], nums[i]
			break
		}
	}
	end := length - 1
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
