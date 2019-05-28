package validpreorder

import (
	"strings"
)

/*
 * @lc app=leetcode id=331 lang=golang
 *
 * [331] Verify Preorder Serialization of a Binary Tree
 *
 * https://leetcode.com/problems/verify-preorder-serialization-of-a-binary-tree/description/
 *
 * algorithms
 * Medium (38.33%)
 * Total Accepted:    58K
 * Total Submissions: 150.3K
 * Testcase Example:  '"9,3,4,#,#,1,#,#,2,#,6,#,#"'
 *
 * One way to serialize a binary tree is to use pre-order traversal. When we
 * encounter a non-null node, we record the node's value. If it is a null node,
 * we record using a sentinel value such as #.
 *
 *
 * ⁠    _9_
 * ⁠   /   \
 * ⁠  3     2
 * ⁠ / \   / \
 * ⁠4   1  #  6
 * / \ / \   / \
 * # # # #   # #
 *
 *
 * For example, the above binary tree can be serialized to the string
 * "9,3,4,#,#,1,#,#,2,#,6,#,#", where # represents a null node.
 *
 * Given a string of comma separated values, verify whether it is a correct
 * preorder traversal serialization of a binary tree. Find an algorithm without
 * reconstructing the tree.
 *
 * Each comma separated value in the string must be either an integer or a
 * character '#' representing null pointer.
 *
 * You may assume that the input format is always valid, for example it could
 * never contain two consecutive commas such as "1,,3".
 *
 * Example 1:
 *
 *
 * Input: "9,3,4,#,#,1,#,#,2,#,6,#,#"
 * Output: true
 *
 * Example 2:
 *
 *
 * Input: "1,#"
 * Output: false
 *
 *
 * Example 3:
 *
 *
 * Input: "9,#,#,1"
 * Output: false
 */
func isValidSerialization(preorder string) bool {
	return isValidTree(strings.Split(preorder, ","))
}

func isValidTree(preorder []string) bool {
	if len(preorder) == 0 {
		return false
	} else if len(preorder) == 1 {
		return preorder[0] == "#"
	} else {
		if preorder[0] == "#" {
			return false
		} else {
			numberCount, sharpCount, childTreeSep := 0, 0, 0
			preorder = preorder[1:]
			for i := range preorder {
				if preorder[i] == "#" {
					sharpCount++
					if numberCount+1 == sharpCount {
						childTreeSep = i
						break
					}
				} else {
					numberCount++
				}
			}
			return isValidTree(preorder[:childTreeSep+1]) && isValidTree(preorder[childTreeSep+1:])
		}
	}
}
