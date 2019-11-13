package find_all_duplicates_in_an_array

import (
	"sort"
)

/*
 * @lc app=leetcode id=442 lang=golang
 *
 * [442] Find All Duplicates in an Array
 *
 * https://leetcode.com/problems/find-all-duplicates-in-an-array/description/
 *
 * algorithms
 * Medium (62.87%)
 * Total Accepted:    120.9K
 * Total Submissions: 192.3K
 * Testcase Example:  '[4,3,2,7,8,2,3,1]'
 *
 * Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements
 * appear twice and others appear once.
 *
 * Find all the elements that appear twice in this array.
 *
 * Could you do it without extra space and in O(n) runtime?
 *
 * Example:
 *
 * Input:
 * [4,3,2,7,8,2,3,1]
 *
 * Output:
 * [2,3]
 *
 */

func findDuplicates(nums []int) []int {
	//start := time.Now()
	//ans := make([]int, 0)
	//m := make(map[int]int)
	//for i := 0; i < len(nums); i++ {
	//	if v, ok := m[nums[i]]; ok {
	//		ans = append(ans, v)
	//	} else {
	//		m[nums[i]] = nums[i]
	//	}
	//}
	//fmt.Println(time.Now().Sub(start))
	//sort.Ints(ans)
	//return ans

	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	ans := make([]int, 0)
	for i, v := range nums {
		if v != i+1 {
			ans = append(ans, v)
		}
	}
	sort.Ints(ans)
	return ans
}
