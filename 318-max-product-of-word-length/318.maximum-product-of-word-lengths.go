package maxproduct

/*
 * @lc app=leetcode id=318 lang=golang
 *
 * [318] Maximum Product of Word Lengths
 *
 * https://leetcode.com/problems/maximum-product-of-word-lengths/description/
 *
 * algorithms
 * Medium (48.08%)
 * Total Accepted:    79.3K
 * Total Submissions: 163.7K
 * Testcase Example:  '["abcw","baz","foo","bar","xtfn","abcdef"]'
 *
 * Given a string array words, find the maximum value of length(word[i]) *
 * length(word[j]) where the two words do not share common letters. You may
 * assume that each word will contain only lower case letters. If no such two
 * words exist, return 0.
 *
 * Example 1:
 *
 *
 * Input: ["abcw","baz","foo","bar","xtfn","abcdef"]
 * Output: 16
 * Explanation: The two words can be "abcw", "xtfn".
 *
 * Example 2:
 *
 *
 * Input: ["a","ab","abc","d","cd","bcd","abcd"]
 * Output: 4
 * Explanation: The two words can be "ab", "cd".
 *
 * Example 3:
 *
 *
 * Input: ["a","aa","aaa","aaaa"]
 * Output: 0
 * Explanation: No such pair of words.
 *
 *
 */
func maxProduct(words []string) int {
	maxProd := 0
	wordsBytes := make([][]byte, len(words), len(words))
	for i := range words {
		wordsBytes[i] = []byte(words[i])
	}
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if len(wordsBytes[i])*len(wordsBytes[j]) > maxProd && !containsSame(wordsBytes[i], wordsBytes[j]) {
				maxProd = len(wordsBytes[i]) * len(wordsBytes[j])
			}
		}
	}
	return maxProd
}

func containsSame(a, b []byte) bool {
	for _, c := range a {
		for _, c2 := range b {
			if c == c2 {
				return true
			}
		}
	}
	return false
}
