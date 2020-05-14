//package main
//
//import (
//	"strconv"
//)

/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

// @lc code=start
func maximumSwap(num int) int {
	if num == 0 {
		return num
	}
	s := strconv.Itoa(num)
	sBytes := []byte(s)
	leftHighestIndex := make([]int, len(s), len(s))
	left, leftPair := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		leftHighestIndex[i] = i
		if i != len(s)-1 && sBytes[leftHighestIndex[i+1]] >= sBytes[i] {
			leftHighestIndex[i] = leftHighestIndex[i+1]
			if sBytes[leftHighestIndex[i+1]] != sBytes[i] {
				left, leftPair = i, leftHighestIndex[i]
			}
		}
	}
	if leftPair != 0 {
		sBytes[left], sBytes[leftPair] = sBytes[leftPair], sBytes[left]
		n, _ := strconv.Atoi(string(sBytes))
		return n
	}
	return num
}

// @lc code=end

//func main() {
//	fmt.Println(maximumSwap(98368))
//}
