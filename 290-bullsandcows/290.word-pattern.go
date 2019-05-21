package wordpattern

import (
	"strings"
)

/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 *
 * https://leetcode.com/problems/word-pattern/description/
 *
 * algorithms
 * Easy (34.71%)
 * Total Accepted:    139.5K
 * Total Submissions: 399.3K
 * Testcase Example:  '"abba"\n"dog cat cat dog"'
 *
 * Given a pattern and a string str, find if str follows the same pattern.
 *
 * Here follow means a full match, such that there is a bijection between a
 * letter in pattern and a non-empty word in str.
 *
 * Example 1:
 *
 *
 * Input: pattern = "abba", str = "dog cat cat dog"
 * Output: true
 *
 * Example 2:
 *
 *
 * Input:pattern = "abba", str = "dog cat cat fish"
 * Output: false
 *
 * Example 3:
 *
 *
 * Input: pattern = "aaaa", str = "dog cat cat dog"
 * Output: false
 *
 * Example 4:
 *
 *
 * Input: pattern = "abba", str = "dog dog dog dog"
 * Output: false
 *
 * Notes:
 * You may assume pattern contains only lowercase letters, and str contains
 * lowercase letters that may be separated by a single space.
 *
 */
func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")
	if len(pattern) != len(strs) {
		return false
	}
	byteGroup := make(map[byte][]int)
	for i, v := range []byte(pattern) {
		if group, ok := byteGroup[v]; ok {
			byteGroup[v] = append(group, i)
		} else {
			byteGroup[v] = []int{i}
		}
	}
	for key := range byteGroup {
		for key2 := range byteGroup {
			if key != key2 &&
				strs[byteGroup[key2][0]] == strs[byteGroup[key][0]] {
				return false
			}
		}
	}
	for _, indexs := range byteGroup {
		for _, index := range indexs {
			if strs[index] != strs[indexs[0]] {
				return false
			}
		}
	}
	return true
}
