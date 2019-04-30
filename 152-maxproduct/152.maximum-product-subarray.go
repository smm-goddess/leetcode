package maxproduct

import (
	"math"
)

/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 *
 * https://leetcode.com/problems/maximum-product-subarray/description/
 *
 * algorithms
 * Medium (28.83%)
 * Total Accepted:    203.8K
 * Total Submissions: 702.4K
 * Testcase Example:  '[2,3,-2,4]'
 *
 * Given an integer array nums, find the contiguous subarray within an array
 * (containing at least one number) which has the largest product.
 * 
 * Example 1:
 * 
 * 
 * Input: [2,3,-2,4]
 * Output: 6
 * Explanation: [2,3] has the largest product 6.
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: [-2,0,-1]
 * Output: 0
 * Explanation: The result cannot be 2, because [-2,-1] is not a subarray.
 * 
 */
func maxProduct(nums []int) int {
	if nums == nil {
		return math.MinInt32
	} else if len(nums) == 1 {
		return nums[0]
	} else {
		product := calculateProduct(nums)
		if product > 0 {
			return product
		} else if product == 0 {
			zeros := findZeros(nums)
			start, end := 0, 0
			for i := range zeros {
				end = zeros[i]
				max := maxProduct(nums[start:end])
				start = end + 1
				if max > product {
					product = max
				}
			}
			if start < len(nums) {
				max := maxProduct(nums[start:])
				if max > product {
					product = max
				}
			}
			return product
		} else {
			negatives := findNegatives(nums)
			for _, i := range negatives {
				max := maxProduct(nums[0:i])
				if max > product {
					product = max
				}
				max = maxProduct(nums[i+1:])
				if max > product {
					product = max
				}
			}
			return product
		}
	}
}

func calculateProduct(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return math.MinInt32
	}
	product := 1
	for i := range nums {
		product = product * nums[i]
	}
	return product
}

func findZeros(nums []int) (places []int) {
	for i := range nums {
		if nums[i] == 0 {
			places = append(places, i)
		}
	}
	return
}

func findNegatives(nums []int) (places []int) {
	for i := range nums {
		if nums[i] < 0 {
			places = append(places, i)
		}
	}
	return
}
