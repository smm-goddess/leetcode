package num_special_equiv_groups

import (
	"sort"
)

/*
 * @lc app=leetcode id=893 lang=golang
 *
 * [893] Groups of Special-Equivalent Strings
 *
 * https://leetcode.com/problems/groups-of-special-equivalent-strings/description/
 *
 * algorithms
 * Easy (63.99%)
 * Total Accepted:    19.6K
 * Total Submissions: 30.7K
 * Testcase Example:  '["abcd","cdab","cbad","xyzz","zzxy","zzyx"]'
 *
 * You are given an array A of strings.
 *
 * A move onto S consists of swapping any two even indexed characters of S, or
 * any two odd indexed characters of S.
 *
 * Two strings S and T are special-equivalent if after any number of moves onto
 * S, S == T.
 *
 * For example, S = "zzxy" and T = "xyzz" are special-equivalent because we may
 * make the moves "zzxy" -> "xzzy" -> "xyzz" that swap S[0] and S[2], then S[1]
 * and S[3].
 *
 * Now, a group of special-equivalent strings from A is a non-empty subset of A
 * such that:
 *
 *
 * Every pair of strings in the group are special equivalent, and;
 * The group is the largest size possible (ie., there isn't a string S not in
 * the group such that S is special equivalent to every string in the group)
 *
 *
 * Return the number of groups of special-equivalent strings from A.
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: ["abcd","cdab","cbad","xyzz","zzxy","zzyx"]
 * Output: 3
 * Explanation:
 * One group is ["abcd", "cdab", "cbad"], since they are all pairwise special
 * equivalent, and none of the other strings are all pairwise special
 * equivalent to these.
 *
 * The other two groups are ["xyzz", "zzxy"] and ["zzyx"].  Note that in
 * particular, "zzxy" is not special equivalent to "zzyx".
 *
 *
 *
 * Example 2:
 *
 *
 * Input: ["abc","acb","bac","bca","cab","cba"]
 * Output: 3
 *
 *
 *
 *
 *
 *
 *
 *
 *
 * Note:
 *
 *
 * 1 <= A.length <= 1000
 * 1 <= A[i].length <= 20
 * All A[i] have the same length.
 * All A[i] consist of only lowercase letters.
 *
 *
 *
 *
 *
 *
 */

type byteArray []byte

func (self byteArray) Equals(other byteArray) bool {
	for i := 0; i < len(self); i++ {
		if self[i] != other[i] {
			return false
		}
	}
	return true
}

func (self byteArray) Len() int {
	return len(self)
}

func (self byteArray) Less(i, j int) bool {
	return self[i] < self[j]
}

func (self byteArray) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func numSpecialEquivGroups(A []string) int {
	keys := make([]byteArray, 0)
	length := len(A[0])
	evenLength, oddLength := length>>1, length>>1+length%2
	oddNum := make([]byte, oddLength, oddLength)
	evenNum := make([]byte, evenLength, evenLength)
	for _, s := range A {
		for index, ch := range []byte(s) { // index starts from 0
			if index&1 == 1 { // even
				evenNum[index/2] = ch
			} else { // odd
				oddNum[index/2] = ch
			}
		}
		sort.Sort(byteArray(oddNum))
		sort.Sort(byteArray(evenNum))
		key := append(append([]byte{}, oddNum...), evenNum...)
		contains := false
		for _, k := range keys {
			if k.Equals(key) {
				contains = true
			}
		}
		if !contains {
			keys = append(keys, key)
		}
	}
	return len(keys)
}
