package fizz_buzz

import (
	"strconv"
)

/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 *
 * https://leetcode.com/problems/fizz-buzz/description/
 *
 * algorithms
 * Easy (60.38%)
 * Total Accepted:    252.1K
 * Total Submissions: 417.4K
 * Testcase Example:  '1'
 *
 * Write a program that outputs the string representation of numbers from 1 to
 * n.
 *
 * But for multiples of three it should output “Fizz” instead of the number and
 * for the multiples of five output “Buzz”. For numbers which are multiples of
 * both three and five output “FizzBuzz”.
 *
 * Example:
 *
 * n = 15,
 *
 * Return:
 * [
 * ⁠   "1",
 * ⁠   "2",
 * ⁠   "Fizz",
 * ⁠   "4",
 * ⁠   "Buzz",
 * ⁠   "Fizz",
 * ⁠   "7",
 * ⁠   "8",
 * ⁠   "Fizz",
 * ⁠   "Buzz",
 * ⁠   "11",
 * ⁠   "Fizz",
 * ⁠   "13",
 * ⁠   "14",
 * ⁠   "FizzBuzz"
 * ]
 *
 *
 */
func fizzBuzz(n int) []string {
	ans := make([]string, n, n)
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			ans[i-1] = "FizzBuzz"
		} else if i%5 == 0 {
			ans[i-1] = "Buzz"
		} else if i%3 == 0 {
			ans[i-1] = "Fizz"
		} else {
			ans[i-1] = strconv.FormatInt(int64(i), 10)
		}
	}
	return ans
}
