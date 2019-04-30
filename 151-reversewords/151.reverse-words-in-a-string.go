package reversewords

import (
	"bytes"
)

/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 *
 * https://leetcode.com/problems/reverse-words-in-a-string/description/
 *
 * algorithms
 * Medium (16.29%)
 * Total Accepted:    274K
 * Total Submissions: 1.7M
 * Testcase Example:  '"the sky is blue"'
 *
 * Given an input string, reverse the string word by word.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: "the sky is blue"
 * Output: "blue is sky the"
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: "  hello world!  "
 * Output: "world! hello"
 * Explanation: Your reversed string should not contain leading or trailing
 * spaces.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: "a good   example"
 * Output: "example good a"
 * Explanation: You need to reduce multiple spaces between two words to a
 * single space in the reversed string.
 * 
 * 
 * 
 * 
 * Note:
 * 
 * 
 * A word is defined as a sequence of non-space characters.
 * Input string may contain leading or trailing spaces. However, your reversed
 * string should not contain leading or trailing spaces.
 * You need to reduce multiple spaces between two words to a single space in
 * the reversed string.
 * 
 * 
 * 
 * 
 * Follow up:
 * 
 * For C programmers, try to solve it in-place in O(1) extra space.
 */
func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	ss := split(s)
	var buffer bytes.Buffer
	for i := len(ss) - 1; i >= 0; i-- {
		buffer.WriteString(ss[i])
		buffer.WriteByte(' ')
	}
	if buffer.Len() > 1 {
		buffer.Truncate(buffer.Len() - 1)
	}
	return buffer.String()
}

func split(s string) (result []string) {
	bs := []byte(s)
	var sb []byte
	for _, b := range bs {
		if b == ' ' || b == '\t' {
			if len(sb) > 0 {
				result = append(result, string(sb))
			}
			sb = sb[0:0]
		} else {
			sb = append(sb, b)
		}
	}
	if len(sb) > 0 {
		result = append(result, string(sb))
	}
	return
}
