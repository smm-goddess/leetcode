package rangesumquery

/*
 * @lc app=leetcode id=307 lang=golang
 *
 * [307] Range Sum Query - Mutable
 *
 * https://leetcode.com/problems/range-sum-query-mutable/description/
 *
 * algorithms
 * Medium (28.00%)
 * Total Accepted:    71K
 * Total Submissions: 248.9K
 * Testcase Example:  '["NumArray","sumRange","update","sumRange"]\n[[[1,3,5]],[0,2],[1,2],[0,2]]'
 *
 * Given an integer array nums, find the sum of the elements between indices i
 * and j (i ≤ j), inclusive.
 *
 * The update(i, val) function modifies nums by updating the element at index i
 * to val.
 *
 * Example:
 *
 *
 * Given nums = [1, 3, 5]
 *
 * sumRange(0, 2) -> 9
 * update(1, 2)
 * sumRange(0, 2) -> 8
 *
 *
 * Note:
 *
 *
 * The array is only modifiable by the update function.
 * You may assume the number of calls to update and sumRange function is
 * distributed evenly.
 *
 *
 */
type NumArray struct {
	sumArray []int
}

func Constructor(nums []int) NumArray {
	var length int
	if nums == nil {
		length = 1
	} else {
		length = len(nums) + 1
	}
	sum := make([]int, length, length)
	for i, v := range nums {
		sum[i+1] = sum[i] + v
	}
	return NumArray{sum}
}

func (this *NumArray) Update(i int, val int) {
	diff := this.sumArray[i+1] - this.sumArray[i] - val
	for k := i + 1; k < len(this.sumArray); k++ {
		this.sumArray[k] = this.sumArray[k] - diff
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.sumArray[j+1] - this.sumArray[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
