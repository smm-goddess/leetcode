package pyramid_transition

import (
	"bytes"
	"strings"
)

/*
 * @lc app=leetcode id=756 lang=golang
 *
 * [756] Pyramid Transition Matrix
 *
 * https://leetcode.com/problems/pyramid-transition-matrix/description/
 *
 * algorithms
 * Medium (53.01%)
 * Total Accepted:    16.8K
 * Total Submissions: 31.7K
 * Testcase Example:  '"ABC"\n["ABD","BCE","DEF","FFF"]'
 *
 * We are stacking blocks to form a pyramid. Each block has a color which is a
 * one letter string.
 *
 * We are allowed to place any color block C on top of two adjacent blocks of
 * colors A and B, if and only if ABC is an allowed triple.
 *
 * We start with a bottom row of bottom, represented as a single string. We
 * also start with a list of allowed triples allowed. Each allowed triple is
 * represented as a string of length 3.
 *
 * Return true if we can build the pyramid all the way to the top, otherwise
 * false.
 *
 * Example 1:
 *
 *
 * Input: bottom = "BCD", allowed = ["BCG", "CDE", "GEA", "FFF"]
 * Output: true
 * Explanation:
 * We can stack the pyramid like this:
 * ⁠   A
 * ⁠  / \
 * ⁠ G   E
 * ⁠/ \ / \
 * B   C   D
 *
 * We are allowed to place G on top of B and C because BCG is an allowed
 * triple.  Similarly, we can place E on top of C and D, then A on top of G and
 * E.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: bottom = "AABA", allowed = ["AAA", "AAB", "ABA", "ABB", "BAC"]
 * Output: false
 * Explanation:
 * We can't stack the pyramid to the top.
 * Note that there could be allowed triples (A, B, C) and (A, B, D) with C !=
 * D.
 *
 *
 *
 *
 * Note:
 *
 *
 * bottom will be a string with length in range [2, 8].
 * allowed will have length in range [0, 200].
 * Letters in all strings will be chosen from the set {'A', 'B', 'C', 'D', 'E',
 * 'F', 'G'}.
 *
 *
 *
 *
 */
var allowedCache map[string][]byte
var allowedPairs []string

func pyramidTransition(bottom string, allowed []string) bool {
	allowedCache = make(map[string][]byte)
	allowedPairs = allowed
	return pyramidTransitionIter(bottom)
}

func pyramidTransitionIter(bottom string) bool {
	allowedByteArray := make([][]byte, 0)
	allowedByteIndex := make([]int, 0)
	for k := 0; k < len(bottom)-1; k++ {
		array := findAvailable(allowedPairs, bottom[k:k+2])
		if len(array) == 0 {
			return false
		} else {
			allowedByteArray = append(allowedByteArray, array)
			allowedByteIndex = append(allowedByteIndex, len(array))
		}
	}
	if len(bottom) == 2 && len(allowedByteArray) > 0 {
		return true
	}
	var index []int
	return allowedPyramid(allowedByteArray, allowedByteIndex, index)
}

func allowedPyramid(allowedByteArray [][]byte, allowedByteIndex []int, index []int) bool {
	if len(allowedByteIndex) == len(index) {
		var bt bytes.Buffer
		for i := 0; i < len(index); i++ {
			bt.WriteByte(allowedByteArray[i][index[i]])
		}
		if pyramidTransitionIter(bt.String()) {
			return true
		}
	} else {
		nextLocation := len(index)
		for i := 0; i < allowedByteIndex[nextLocation]; i++ {
			if allowedPyramid(allowedByteArray, allowedByteIndex, append(index, i)) {
				return true
			}
		}
	}
	return false
}

func findAvailable(allowed []string, bottom string) []byte {
	if cache, ok := allowedCache[bottom]; ok {
		return cache
	} else {
		var ava []byte
		for _, allow := range allowed {
			if strings.HasPrefix(allow, bottom) {
				ava = append(ava, allow[2])
			}
		}
		allowedCache[bottom] = ava
		return ava
	}
}
