package palindromic_substrings

/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 *
 * https://leetcode.com/problems/palindromic-substrings/description/
 *
 * algorithms
 * Medium (56.36%)
 * Total Accepted:    123.6K
 * Total Submissions: 212.7K
 * Testcase Example:  '"abc"'
 *
 * Given a string, your task is to count how many palindromic substrings in
 * this string.
 *
 * The substrings with different start indexes or end indexes are counted as
 * different substrings even they consist of same characters.
 *
 * Example 1:
 *
 *
 * Input: "abc"
 * Output: 3
 * Explanation: Three palindromic strings: "a", "b", "c".
 *
 *
 *
 *
 * Example 2:
 *
 *
 * Input: "aaa"
 * Output: 6
 * Explanation: Six palindromic strings: "a", "a", "a", "aa", "aa", "aaa".
 *
 *
 *
 *
 * Note:
 *
 *
 * The input string length won't exceed 1000.
 *
 *
 *
 *
 */
func countSubstrings(s string) int {
	count := len(s)
	start := 0
	length := 2
	for start < len(s)-1 {
		for start+length <= len(s) {
			count += isPalindromic(s[start : start+length])
			length++
		}
		start++
		length = 2
	}
	return count
}

func isPalindromic(s string) int {
	for index := 0; index < len(s)>>1; index++ {
		if s[index] != s[len(s)-index-1] {
			return 0
		}
	}
	return 1
}
