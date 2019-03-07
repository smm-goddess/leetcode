package main

import "fmt"

func main() {
	fmt.Println(fourSum([]int{-3, -2, -1, 0, 0, 1, 2, 3}, 0))
}

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return [][]int{}
	}
	result := make([][]int, 0)
	for i := 0; i < length-3; i ++ {
		for j := i + 1; j < length-2; j++ {
			for k := j + 1; k < length-1; k++ {
				for l := k + 1; l < length; l ++ {
					a, b, c, d := nums[i], nums[j], nums[k], nums[l]
					if a+b+c+d == target {
						if a < b {
							b, a = a, b
						}
						if c < d {
							c, d = d, c
						}
						if a < c {
							a, c = c, a
						}
						if d > b {
							b, d = d, b
						}
						if b < c {
							b, c = c, b
						}
						repeat := false
						for _, s := range result {
							if s[0] == a && s[1] == b && s[2] == c && s[3] == d {
								repeat = true
								break
							}
						}
						if !repeat {
							result = append(result, []int{a, b, c, d})
						}
					}
				}
			}
		}
	}
	return result
}
