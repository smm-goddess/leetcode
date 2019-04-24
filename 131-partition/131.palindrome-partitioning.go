package partition

/*
 * @lc app=leetcode id=131 lang=golang
 *
 * [131] Palindrome Partitioning
 *
 * https://leetcode.com/problems/palindrome-partitioning/description/
 *
 * algorithms
 * Medium (40.16%)
 * Total Accepted:    159.3K
 * Total Submissions: 394.5K
 * Testcase Example:  '"aab"'
 *
 * Given a string s, partition s such that every substring of the partition is
 * a palindrome.
 *
 * Return all possible palindrome partitioning of s.
 *
 * Example:
 *
 *
 * Input: "aab"
 * Output:
 * [
 * ⁠ ["aa","b"],
 * ⁠ ["a","a","b"]
 * ]
 *
 *
 */
var partitionMap = make(map[string][][]string)

func partition(s string) [][]string {
	if len(s) == 0 {
		return [][]string{}
	} else if len(s) == 1 {
		return [][]string{{s}}
	} else {
		result := make([][]string, 0)
		for i := 1; i <= len(s); i++ {
			s1 := s[:i]
			s2 := s[i:]
			if palindrome(s1) {
				if len(s2) == 0 {
					result = append(result, append([]string{s1}))
				} else {
					if partitions, ok := partitionMap[s2]; ok {
						for _, p := range partitions {
							result = append(result, append([]string{s1}, p...))
						}
					} else {
						partitions := partition(s2)
						for _, p := range partitions {
							result = append(result, append([]string{s1}, p...))
						}
						partitionMap[s2] = partitions
					}
				}
			}
		}
		return result
	}
}

func palindrome(s string) bool {
	length := len(s)
	if length == 0 || length == 1 {
		return true
	} else {
		for i := 0; i < length>>1; i++ {
			if s[i] != s[length-1-i] {
				return false
			}
		}
		return true
	}
}
