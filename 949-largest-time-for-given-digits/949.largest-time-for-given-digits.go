package largest_time_from_digits

import "fmt"

/*
 * @lc app=leetcode id=949 lang=golang
 *
 * [949] Largest Time for Given Digits
 *
 * https://leetcode.com/problems/largest-time-for-given-digits/description/
 *
 * algorithms
 * Easy (34.55%)
 * Total Accepted:    11.6K
 * Total Submissions: 33.4K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given an array of 4 digits, return the largest 24 hour time that can be
 * made.
 *
 * The smallest 24 hour time is 00:00, and the largest is 23:59.  Starting from
 * 00:00, a time is larger if more time has elapsed since midnight.
 *
 * Return the answer as a string of length 5.  If no valid time can be made,
 * return an empty string.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: [1,2,3,4]
 * Output: "23:41"
 *
 *
 *
 * Example 2:
 *
 *
 * Input: [5,5,5,5]
 * Output: ""
 *
 *
 *
 *
 * Note:
 *
 *
 * A.length == 4
 * 0 <= A[i] <= 9
 *
 *
 *
 */

var arrange = [][]int{
	{0, 1, 2, 3},
	{0, 1, 3, 2},
	{0, 2, 1, 3},
	{0, 2, 3, 1},
	{0, 3, 1, 2},
	{0, 3, 2, 1},
	{1, 0, 2, 3},
	{1, 0, 3, 2},
	{1, 2, 3, 0},
	{1, 2, 0, 3},
	{1, 3, 2, 0},
	{1, 3, 0, 2},
	{2, 0, 1, 3},
	{2, 0, 3, 1},
	{2, 1, 3, 0},
	{2, 1, 0, 3},
	{2, 3, 1, 0},
	{2, 3, 0, 1},
	{3, 0, 1, 2},
	{3, 0, 2, 1},
	{3, 1, 2, 0},
	{3, 1, 0, 2},
	{3, 2, 1, 0},
	{3, 2, 0, 1},
}

func largestTimeFromDigits(A []int) string {
	largestTime := -1
	var s string
	var time, hours, minutes int
	for _, locate := range arrange {
		hours = A[locate[0]]*10 + A[locate[1]]
		minutes = A[locate[2]]*10 + A[locate[3]]
		if hours < 24 && minutes < 60 {
			time = hours*60 + minutes
			if time > largestTime {
				largestTime = time
				s = fmt.Sprintf("%d%d:%d%d", A[locate[0]], A[locate[1]], A[locate[2]], A[locate[3]])
			}
		}
	}
	return s
}
