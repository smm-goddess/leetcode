package frequency_sort

import (
	"bytes"
	"sort"
)

/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 *
 * https://leetcode.com/problems/sort-characters-by-frequency/description/
 *
 * algorithms
 * Medium (57.81%)
 * Total Accepted:    116.2K
 * Total Submissions: 200.7K
 * Testcase Example:  '"tree"'
 *
 * Given a string, sort it in decreasing order based on the frequency of
 * characters.
 *
 * Example 1:
 *
 * Input:
 * "tree"
 *
 * Output:
 * "eert"
 *
 * Explanation:
 * 'e' appears twice while 'r' and 't' both appear once.
 * So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid
 * answer.
 *
 *
 *
 * Example 2:
 *
 * Input:
 * "cccaaa"
 *
 * Output:
 * "cccaaa"
 *
 * Explanation:
 * Both 'c' and 'a' appear three times, so "aaaccc" is also a valid answer.
 * Note that "cacaca" is incorrect, as the same characters must be together.
 *
 *
 *
 * Example 3:
 *
 * Input:
 * "Aabb"
 *
 * Output:
 * "bbAa"
 *
 * Explanation:
 * "bbaA" is also a valid answer, but "Aabb" is incorrect.
 * Note that 'A' and 'a' are treated as two different characters.
 *
 *
 */

type freqList []*freqKey

func (list *freqList) Len() int {
	return len(*list)
}

func (list *freqList) Less(i, j int) bool {
	return (*list)[i].freq > (*list)[j].freq
}

func (list *freqList) Swap(i, j int) {
	(*list)[i], (*list)[j] = (*list)[j], (*list)[i]
}

func (list *freqList) add(b byte) {
	contains := false
	for index, fk := range *list {
		if fk.key == b {
			contains = true
			(*list)[index].freq = (*list)[index].freq + 1
		}
	}
	if !contains {
		*list = append(*list, &freqKey{
			b,
			1,
		})
	}
}

type freqKey struct {
	key  byte
	freq int
}

func frequencySort(s string) string {
	list := &freqList{}
	for index := range s {
		list.add(s[index])
	}
	var buffer bytes.Buffer
	sort.Sort(list)
	for _, fk := range *list {
		for i := 0; i < fk.freq; i++ {
			buffer.WriteByte(fk.key)
		}
	}
	return buffer.String()
}
