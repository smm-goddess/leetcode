package main

import "fmt"

func getRightLowestIndex(heights []int, index int) int {
	lowestIndex := -1
	height := heights[index]
	for i := index + 1; i < len(heights); i++ {
		if heights[i] > height {
			break
		}
		if heights[i] < height {
			height = heights[i]
			lowestIndex = i
		}
	}
	return lowestIndex
}

func getLeftLowestIndex(heights []int, index int) int {
	lowestIndex := -1
	height := heights[index]
	for i := index - 1; i >= 0; i-- {
		if heights[i] > height {
			break
		}
		if heights[i] < height {
			height = heights[i]
			lowestIndex = i
		}
	}
	return lowestIndex
}

func pourWater(heights []int, V int, K int) []int {
	if V == 0 {
		return heights
	}
	leftLowestIndex := getLeftLowestIndex(heights, K)
	if leftLowestIndex != -1 {
		heights[leftLowestIndex] = heights[leftLowestIndex] + 1
		return pourWater(heights, V-1, K)
	}
	rightLowestIndex := getRightLowestIndex(heights, K)
	if rightLowestIndex != -1 {
		heights[rightLowestIndex] = heights[rightLowestIndex] + 1
		return pourWater(heights, V-1, K)
	}
	heights[K] = heights[K] + 1
	return pourWater(heights, V-1, K)
}

func main() {
	//fmt.Println(pourWater([]int{3, 1, 3}, 5, 1))
	//fmt.Println(pourWater([]int{2, 1, 1, 2, 1, 2, 2}, 4, 3))
	//fmt.Println(pourWater([]int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1}, 5, 5))
	fmt.Println(pourWater([]int{1, 2, 3, 4, 3, 2, 1, 2, 3, 4, 3, 2, 1}, 10, 2))
}
