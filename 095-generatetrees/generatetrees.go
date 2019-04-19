package generatetrees

/*
 * @lc app=leetcode id=95 lang=golang
 *
 * [95] Unique Binary Search Trees II
 *
 * https://leetcode.com/problems/unique-binary-search-trees-ii/description/
 *
 * algorithms
 * Medium (35.18%)
 * Total Accepted:    134.1K
 * Total Submissions: 379.8K
 * Testcase Example:  '3'
 *
 * Given an integer n, generate all structurally unique BST's (binary search
 * trees) that store values 1 ... n.
 *
 * Example:
 *
 *
 * Input: 3
 * Output:
 * [
 * [1,null,3,2],
 * [3,2,null,1],
 * [3,1,null,null,2],
 * [2,1,3],
 * [1,null,2,null,3]
 * ]
 * Explanation:
 * The above output corresponds to the 5 unique BST's shown below:
 *
 * ⁠  1         3     3      2      1
 * ⁠   \       /     /      / \      \
 * ⁠    3     2     1      1   3      2
 * ⁠   /     /       \                 \
 * ⁠  2     1         2                 3
 *
 *
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	permutations := generatePermutation(n)
	var trees []*TreeNode
	for _, permutation := range permutations {
		t := permutationToTree(permutation)
		newTree := true
		for _, tree := range trees {
			if treeEquals(t, tree) {
				newTree = false
				break
			}
		}
		if newTree {
			trees = append(trees, t)
		}
	}
	return trees
}

func treeEquals(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	} else if t1 == nil || t2 == nil {
		return false
	} else if t1.Val == t2.Val {
		return treeEquals(t1.Left, t2.Left) && treeEquals(t1.Right, t2.Right)
	} else {
		return false
	}
}

func permutationToTree(permutation []int) *TreeNode {
	tree := &TreeNode{permutation[0], nil, nil}
	for _, n := range permutation[1:] {
		addTreeNode(tree, n)
	}
	return tree
}

func addTreeNode(tree *TreeNode, n int) {
	if tree.Val > n {
		if tree.Left == nil {
			tree.Left = &TreeNode{n, nil, nil}
		} else {
			addTreeNode(tree.Left, n)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &TreeNode{n, nil, nil}
		} else {
			addTreeNode(tree.Right, n)
		}
	}
}

func generatePermutation(n int) [][]int {
	permutations := make([][]int, 0)
	remains := make([]int, n, n)
	for i := range remains {
		remains[i] = i + 1
	}
	generatePermutationIter(&permutations, []int{}, remains)
	return permutations
}

func generatePermutationIter(permutations *[][]int, alreadyIn, remains []int) {
	if len(remains) == 0 {
		*permutations = append(*permutations, append([]int{}, alreadyIn...))
	} else {
		for i, remain := range remains {
			generatePermutationIter(permutations, append(alreadyIn, remain), append(append([]int{}, remains[0:i]...), remains[i+1:]...))
		}
	}
}
