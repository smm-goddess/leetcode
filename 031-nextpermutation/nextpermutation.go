package nextpermutation

/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 *
 * https://leetcode.com/problems/next-permutation/description/
 *
 * algorithms
 * Medium (30.07%)
 * Total Accepted:    223.1K
 * Total Submissions: 738.9K
 * Testcase Example:  '[1,2,3]'
 *
 * Implement next permutation, which rearranges numbers into the
 * lexicographically next greater permutation of numbers.
 *
 * If such arrangement is not possible, it must rearrange it as the lowest
 * possible order (ie, sorted in ascending order).
 *
 * The replacement must be in-place and use only constant extra memory.
 *
 * Here are some examples. Inputs are in the left-hand column and its
 * corresponding outputs are in the right-hand column.
 *
 * 1,2,3 → 1,3,2
 * 3,2,1 → 1,2,3
 * 1,1,5 → 1,5,1
 *
 */
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
