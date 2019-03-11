package numdecodings

/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 *
 * https://leetcode.com/problems/decode-ways/description/
 *
 * algorithms
 * Medium (21.93%)
 * Total Accepted:    240.5K
 * Total Submissions: 1.1M
 * Testcase Example:  '"12"'
 *
 * A message containing letters from A-Z is being encoded to numbers using the
 * following mapping:
 *
 *
 * 'A' -> 1
 * 'B' -> 2
 * ...
 * 'Z' -> 26
 *
 *
 * Given a non-empty string containing only digits, determine the total number
 * of ways to decode it.
 *
 * Example 1:
 *
 *
 * Input: "12"
 * Output: 2
 * Explanation: It could be decoded as "AB" (1 2) or "L" (12).
 *
 *
 * Example 2:
 *
 *
 * Input: "226"
 * Output: 3
 * Explanation: It could be decoded as "BZ" (2 26), "VF" (22 6), or "BBF" (2 2
 * 6).
 *
 */

func numDecodings(s string) int {
	if len(s) == 0 {
		return 1
	} else if len(s) == 1 {
		if s == "0" {
			return 0
		} else {
			return 1
		}
	} else {
		first, second := s[0], s[1]
		if first == '0' {
			return 0
		} else if first >= '3' && first <= '9' || (first == '2' && second >= '7' && second <= '9') {
			return numDecodings(s[1:])
		} else {
			return numDecodings(s[1:]) + numDecodings(s[2:])
		}
	}
}
