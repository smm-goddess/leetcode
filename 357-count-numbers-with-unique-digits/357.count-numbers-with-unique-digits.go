package countdigits

/*
 * @lc app=leetcode id=357 lang=golang
 *
 * [357] Count Numbers with Unique Digits
 *
 * https://leetcode.com/problems/count-numbers-with-unique-digits/description/
 *
 * algorithms
 * Medium (46.73%)
 * Total Accepted:    61.5K
 * Total Submissions: 131.1K
 * Testcase Example:  '2'
 *
 * Given a non-negative integer n, count all numbers with unique digits, x,
 * where 0 ≤ x < 10n.
 *
 *
 * Example:
 *
 *
 * Input: 2
 * Output: 91
 * Explanation: The answer should be the total numbers in the range of 0 ≤ x <
 * 100,
 * excluding 11,22,33,44,55,66,77,88,99
 *
 *
 */
var array = []int{9, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

func countNumbersWithUniqueDigits(n int) int {
	return countUniqueDigits(0, n)
}

func countUniqueDigits(cnt int, n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return cnt + 10
	} else if n > 10 {
		return countUniqueDigits(cnt, 10)
	} else {
		count := 1
		for i := 0; i < n; i++ {
			count *= array[i]
		}
		return countUniqueDigits(cnt+count, n-1)
	}
}
