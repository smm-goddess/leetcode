package groupanagrams

import (
	"sort"
)

/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 *
 * https://leetcode.com/problems/group-anagrams/description/
 *
 * algorithms
 * Medium (45.67%)
 * Total Accepted:    314.9K
 * Total Submissions: 689K
 * Testcase Example:  '["eat","tea","tan","ate","nat","bat"]'
 *
 * Given an array of strings, group anagrams together.
 *
 * Example:
 *
 *
 * Input: ["eat", "tea", "tan", "ate", "nat", "bat"],
 * Output:
 * [
 * ⁠ ["ate","eat","tea"],
 * ⁠ ["nat","tan"],
 * ⁠ ["bat"]
 * ]
 *
 * Note:
 *
 *
 * All inputs will be in lowercase.
 * The order of your output does not matter.
 *
 *
 */

func groupAnagrams(strs []string) [][]string {
	ans := make(map[string]int)
	var answer [][]string
	for _, s := range strs {
		bytes := []byte(s)
		ints := make([]int, len(bytes))
		for i, v := range bytes {
			ints[i] = int(v)
		}
		sort.Ints(ints)
		for i, v := range ints {
			bytes[i] = byte(v)
		}
		key := string(bytes)
		if v, ok := ans[key]; ok {
			answer[v] = append(answer[v], s)
		} else {
			answer = append(answer, []string{s})
			ans[key] = len(answer) - 1
		}
	}
	return answer
}
