package main

import "fmt"

/**
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).
 */
func main() {
	//fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1, 2, 3, 4, 5, 6}, 4))
}

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	middle := low + (high-low)>>1
	for low <= high {
		if nums[low] == target {
			return low
		} else if nums[high] == target {
			return high
		} else if nums[middle] == target {
			return middle
		} else {
			if target > nums[high] && target < nums[low] {
				return -1
			} else if nums[middle] > nums[low] {
				if target > nums[low] && target < nums[middle] {
					low ++
					high = middle - 1
					middle = low + (high-low)>>1
				} else {
					low = middle + 1
					high--
					middle = low + (high-low)>>1
				}
			} else {
				if nums[middle] < target && nums[high] > target {
					low = middle + 1
					high--
					middle = low + (high-low)>>1
				} else {
					low ++
					high = middle - 1
					middle = low + (high-low)>>1
				}
			}
		}
	}
	return -1
}
