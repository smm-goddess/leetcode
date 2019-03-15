package main

import "fmt"

func main() {
	colors := []int{2, 0, 2, 1, 1, 0}
	fmt.Println(colors)
	sortColors1(colors)
	fmt.Println(colors)
}

func sortColors(nums []int) {
	count := []int{0, 0, 0}
	for _, v := range nums {
		count[v] = count[v] + 1
	}
	index := 0
	for i, v := range count {
		c := v
		for c > 0 {
			nums[index] = i
			index ++
			c--
		}
	}
}

func sortColors1(nums []int) {
	s1, s2 := -1, -1
	for i := 0; i < len(nums); i++ {
		if s1 > -1 && s2 > -1 {
			if nums[i] == 0 {
				nums[s1], nums[s2], nums[i] = 0, 1, 2
				s1++
				s2++
			} else if nums[i] == 1 {
				nums[s2], nums[i] = 1, 2
				s2 ++
			}
		} else if s1 > -1 {
			if nums[i] == 0 {
				nums[s1], nums[i] = 0, 1
				s1 ++
			} else if nums[i] == 2 {
				s2 = i
			}
		} else if s2 > -1 {
			if nums[i] == 0 {
				nums[s2], nums[i] = 0, 2
				s2 ++
			} else if nums[i] == 1 {
				nums[s2], nums[i] = 1, 2
				s1 = s2
				s2 ++
			}
		} else {
			if nums[i] == 1 {
				s1 = i
			} else if nums[i] == 2 {
				s2 = i
			}
		}
	}
}
