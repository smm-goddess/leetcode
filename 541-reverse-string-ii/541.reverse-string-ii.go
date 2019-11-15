package reverse_str

import (
	"bytes"
)

/*
 * @lc app=leetcode id=541 lang=golang
 *
 * [541] Reverse String II
 *
 * https://leetcode.com/problems/reverse-string-ii/description/
 *
 * algorithms
 * Easy (46.56%)
 * Total Accepted:    71.5K
 * Total Submissions: 153.4K
 * Testcase Example:  '"abcdefg"\n2'
 *
 *
 * Given a string and an integer k, you need to reverse the first k characters
 * for every 2k characters counting from the start of the string. If there are
 * less than k characters left, reverse all of them. If there are less than 2k
 * but greater than or equal to k characters, then reverse the first k
 * characters and left the other as original.
 *
 *
 * Example:
 *
 * Input: s = "abcdefg", k = 2
 * Output: "bacdfeg"
 *
 *
 *
 * Restrictions:
 *
 * ⁠The string consists of lower English letters only.
 * ⁠Length of the given string and k will in the range [1, 10000]
 *
 */

func reverseBytes(bs []byte, k int) {
	if len(bs) <= k {
		for i := 0; i < len(bs)>>1; i++ {
			bs[i], bs[len(bs)-1-i] = bs[len(bs)-1-i], bs[i]
		}
	} else {
		for i := 0; i < k>>1; i++ {
			bs[i], bs[k-1-i] = bs[k-1-i], bs[i]
		}
	}
}

func reverseStr(s string, k int) string {
	var ans bytes.Buffer
	tmp := make([]byte, 0)
	for i, ch := range []byte(s) {
		if i%(2*k) == 0 && len(tmp) > 0 {
			reverseBytes(tmp, k)
			ans.Write(tmp)
			tmp[0] = ch
			tmp = tmp[0:1]
		} else {
			tmp = append(tmp, ch)
		}
	}
	if len(tmp) > 0 {
		reverseBytes(tmp, k)
		ans.Write(tmp)
	}
	return ans.String()
}
