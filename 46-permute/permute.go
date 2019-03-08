package main

import "fmt"

/**
Given a collection of distinct integers, return all possible permutations.
 */

 // TODO
func main() {
	//p := permuteIter([]int{}, []int{1, 2, 3})
	//for _, v := range p {
	//	fmt.Println(v[0])
	//	fmt.Println(v[1])
	//}
	//fmt.Println(permuteIter([]int{}, []int{1, 2, 3}))
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var iter [][][]int
	r := permuteIter([]int{}, nums)
	for len(r) > 0 && len(r[0][1]) > 0 {
		fmt.Println(r)
		for _, v := range r {
			iter = append(permuteIter(v[0], v[1]))
		}
		r = iter
	}
	fmt.Println(r)
	return [][]int{}
}

func permuteIter(head, remain []int) [][][]int {
	result := make([][][]int, 0)
	remainLength := len(remain) - 1
	headLength := len(head) + 1
	for i, v := range remain {
		newHead := make([]int, headLength)
		copy(newHead, head)
		newHead[len(head)] = v
		newRemain := make([]int, 0, remainLength)
		newRemain = append(newRemain, remain[:i]...)
		newRemain = append(newRemain, remain[i+1:]...)
		item := [][]int{newHead, newRemain}
		result = append(result, item)
	}
	return result
}
