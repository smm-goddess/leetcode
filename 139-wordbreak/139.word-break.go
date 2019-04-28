package wordbreak

import (
	"strings"
)

// import "strings"

/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 *
 * https://leetcode.com/problems/word-break/description/
 *
 * algorithms
 * Medium (34.78%)
 * Total Accepted:    325.6K
 * Total Submissions: 930.4K
 * Testcase Example:  '"leetcode"\n["leet","code"]'
 *
 * Given a non-empty string s and a dictionary wordDict containing a list of
 * non-empty words, determine if s can be segmented into a space-separated
 * sequence of one or more dictionary words.
 *
 * Note:
 *
 *
 * The same word in the dictionary may be reused multiple times in the
 * segmentation.
 * You may assume the dictionary does not contain duplicate words.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "leetcode", wordDict = ["leet", "code"]
 * Output: true
 * Explanation: Return true because "leetcode" can be segmented as "leet
 * code".
 *
 *
 * Example 2:
 *
 *
 * Input: s = "applepenapple", wordDict = ["apple", "pen"]
 * Output: true
 * Explanation: Return true because "applepenapple" can be segmented as "apple
 * pen apple".
 * Note that you are allowed to reuse a dictionary word.
 *
 *
 * Example 3:
 *
 *
 * Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
 * Output: false
 *
 *
 */
var capable map[string]bool

func wordBreak(s string, wordDict []string) bool {
	capable = make(map[string]bool)
	return wordBreakIter(s, wordDict)
}

func wordBreakIter(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	for _, word := range wordDict {
		if strings.HasPrefix(s, word) {
			newString := s[len(word):]
			if _, ok := capable[newString]; ok {
			} else {
				capa := wordBreakIter(newString, wordDict)
				if capa {
					return true
				} else {
					capable[newString] = capa
				}
			}
		}
	}
	return false
}
