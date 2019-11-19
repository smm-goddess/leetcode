package license_key_formatting

import (
	"bytes"
	"strings"
)

/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 *
 * https://leetcode.com/problems/license-key-formatting/description/
 *
 * algorithms
 * Easy (42.07%)
 * Total Accepted:    104K
 * Total Submissions: 246.9K
 * Testcase Example:  '"5F3Z-2e-9-w"\n4'
 *
 * You are given a license key represented as a string S which consists only
 * alphanumeric character and dashes. The string is separated into N+1 groups
 * by N dashes.
 *
 * Given a number K, we would want to reformat the strings such that each group
 * contains exactly K characters, except for the first group which could be
 * shorter than K, but still must contain at least one character. Furthermore,
 * there must be a dash inserted between two groups and all lowercase letters
 * should be converted to uppercase.
 *
 * Given a non-empty string S and a number K, format the string according to
 * the rules described above.
 *
 * Example 1:
 *
 * Input: S = "5F3Z-2e-9-w", K = 4
 *
 * Output: "5F3Z-2E9W"
 *
 * Explanation: The string S has been split into two parts, each part has 4
 * characters.
 * Note that the two extra dashes are not needed and can be removed.
 *
 *
 *
 *
 * Example 2:
 *
 * Input: S = "2-5g-3-J", K = 2
 *
 * Output: "2-5G-3J"
 *
 * Explanation: The string S has been split into three parts, each part has 2
 * characters except the first part as it could be shorter as mentioned
 * above.
 *
 *
 *
 * Note:
 *
 * The length of string S will not exceed 12,000, and K is a positive integer.
 * String S consists only of alphanumerical characters (a-z and/or A-Z and/or
 * 0-9) and dashes(-).
 * String S is non-empty.
 *
 *
 */
func toUpper(c byte) byte {
	if c >= 'a' && c <= 'z' {
		return c - 32
	} else {
		return c
	}
}
func licenseKeyFormatting(S string, K int) string {
	splashCount := strings.Count(S, "-")
	var buffer bytes.Buffer
	letterCount := len(S) - splashCount
	index := letterCount % K
	if index != 0 {
		index = K - index
	}
	for _, ch := range []byte(S) {
		if ch == '-' {
			continue
		} else {
			buffer.WriteByte(toUpper(ch))
			index++
			if index%K == 0 {
				buffer.WriteByte('-')
				index = 0
			}
		}
	}
	if buffer.Len() > 0 {
		return string(buffer.Bytes()[:buffer.Len()-1])
	} else {
		return ""
	}
}

//
//func licenseKeyFormatting(S string, K int) string {
//	var buffer, ansBuffer bytes.Buffer
//	for _, ch := range []byte(S) {
//		if ch == '-' {
//			continue
//		} else if ch >= 'a' && ch <= 'z' {
//			buffer.WriteByte(ch - 32)
//		} else {
//			buffer.WriteByte(ch)
//		}
//	}
//	firstLength := buffer.Len() % K
//	if firstLength > 0 {
//		ansBuffer.Write(buffer.Bytes()[0:firstLength])
//		if buffer.Len() > K {
//			ansBuffer.WriteByte('-')
//		}
//	}
//	count := 1
//	for count*K <= buffer.Len() {
//		ansBuffer.Write(buffer.Bytes()[firstLength+(count-1)*K : firstLength+count*K])
//		count++
//		if count*K <= buffer.Len() {
//			ansBuffer.WriteByte('-')
//		}
//	}
//	return ansBuffer.String()
//}

//
//func licenseKeyFormatting(S string, K int) string {
//	tmp, index := make([]byte, K, K), K-1
//	sBytes := []byte(S)
//	r := make([]string, 0)
//	for k := len(S) - 1; k >= 0; k-- {
//		char := sBytes[k]
//		if char == '-' {
//			continue
//		} else {
//			tmp[index] = toUpper(char)
//			index--
//			if index < 0 {
//				r = append([]string{string(tmp)}, r...)
//				index = K - 1
//			}
//		}
//	}
//	if index != K-1 {
//		r = append([]string{string(tmp[index+1:])}, r...)
//	}
//	return strings.Join(r, "-")
//}
