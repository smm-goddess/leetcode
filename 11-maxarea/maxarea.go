package main

import "fmt"

func main() {
  fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
    max := 0
    length := len(height)
    for i := 0; i < length; i++ {
      for j := i + 1; j < length; j++ {
        min := height[i]
        if min > height[j] {
          min = height[j]
        }
        area := min * (j - i)
        if max < area {
          max = area
        }
      }
    }
    return max
}
