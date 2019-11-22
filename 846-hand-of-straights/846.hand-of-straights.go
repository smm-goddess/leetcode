package is_n_straight_hand

import (
	"math"
)

/*
 * @lc app=leetcode id=846 lang=golang
 *
 * [846] Hand of Straights
 *
 * https://leetcode.com/problems/hand-of-straights/description/
 *
 * algorithms
 * Medium (50.48%)
 * Total Accepted:    25.1K
 * Total Submissions: 49.6K
 * Testcase Example:  '[1,2,3,6,2,3,4,7,8]\n3'
 *
 * Alice has a hand of cards, given as an array of integers.
 *
 * Now she wants to rearrange the cards into groups so that each group is size
 * W, and consists of W consecutive cards.
 *
 * Return true if and only if she can.
 *
 *
 *
 *
 *
 *
 * Example 1:
 *
 *
 * Input: hand = [1,2,3,6,2,3,4,7,8], W = 3
 * Output: true
 * Explanation: Alice's hand can be rearranged as [1,2,3],[2,3,4],[6,7,8].
 *
 * Example 2:
 *
 *
 * Input: hand = [1,2,3,4,5], W = 4
 * Output: false
 * Explanation: Alice's hand can't be rearranged into groups of 4.
 *
 *
 *
 * Note:
 *
 *
 * 1 <= hand.length <= 10000
 * 0 <= hand[i] <= 10^9
 * 1 <= W <= hand.length
 *
 *
 */
func isNStraightHand(hand []int, W int) bool {
	if len(hand)%W != 0 {
		return false
	} else {
		if W == 1 {
			return true
		} else {
			m := make(map[int]int)
			for _, v := range hand {
				if c, ok := m[v]; ok {
					m[v] = c + 1
				} else {
					m[v] = 1
				}
			}
			for len(m) > 0 {
				min, count := math.MaxInt32, 0
				for k, v := range m {
					if k < min {
						min, count = k, v
					}
				}
				for i := min; i < min+W; i++ {
					if c, ok := m[i]; ok {
						if c-count == 0 {
							delete(m, i)
						} else {
							m[i] = c - count
						}
					} else {
						return false
					}
				}
			}
			return true
		}
	}
}
