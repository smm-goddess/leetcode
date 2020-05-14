/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

func trap(height []int) int {
	rightSum, leftSum, rightHighest,leftHeighest, sum := 0, 0, 0, 0, 0
	for index, h := range height {
			rightHighest = max(rightHighest, h)
			rightSum = rightSum + rightHighest
			sum = sum + h
			leftHeighest = max(leftHeighest, height[len(height) - 1 - index])
			leftSum = leftSum + leftHeighest
	}
	return rightSum + leftSum - rightHighest * len(height) - sum 
}
// @lc code=end

