package main

import "fmt"

/**
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].
 */

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	fmt.Println(searchRange([]int{5, 7, 7, 8, 10}, 8))
}

func searchRange(nums []int, target int) []int {
	start, end, index := -1, -1, 0
	for index < len(nums) {
		if nums[index] == target {
			if start == -1 {
				start = index
				end = index
			} else {
				end = index
			}
		}
		index++
	}
	return []int{start, end}
}
