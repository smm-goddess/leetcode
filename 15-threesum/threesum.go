package main

import "fmt"

func main() {
  nums := []int{-1, 0, 1, 2, -1, -4}
  fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
  result := make([][]int, 0, 8)
  length := len(nums)
  for i := 0; i < length-2; i++ {
    for j := i+1; j < length-1; j++ {
      sum1 := nums[i] + nums[j]
      for k := j+1; k < length; k++ {
        if sum1 + nums[k] == 0 {
          val := []int{nums[i],nums[j],nums[k]}
          result = append(result, val)
        }
      }
    }
  }
  return result
}
