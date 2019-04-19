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
	permutations := generateUniqPermutation(n)
	var trees []*TreeNode
	for _, permutation := range permutations {
		t := permutationToTree(permutation)
		trees = append(trees, t)
	}
	return trees
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

var permutationMap = make(map[int][][]int)

func generateUniqPermutation(n int) [][]int {
	if n == 0 {
		return [][]int{}
	} else if n == 1 {
		return [][]int{{1}}
	} else if n == 2 {
		return [][]int{{2, 1}, {1, 2}}
	} else {
		var permutations [][]int
		for k := 1; k <= n; k++ {
			var before, after [][]int
			if v, ok := permutationMap[k-1]; ok {
				before = v
			} else {
				before = generateUniqPermutation(k - 1)
			}
			if v, ok := permutationMap[n-k]; ok {
				after = v
			} else {
				after = generateUniqPermutation(n - k)
			}
			if len(before) == 0 {
				for _, v := range after {
					for ai := range v {
						v[ai] = v[ai] + k
					}
					permutations = append(permutations, append([]int{k}, v...))
				}
			} else if len(after) == 0 {
				for _, v := range before {
					permutations = append(permutations, append([]int{k}, v...))
				}
			} else {
				for _, v := range after {
					for ai := range v {
						v[ai] = v[ai] + k
					}
					for _, be := range before {
						permutations = append(permutations, append(append([]int{k}, be...), v...))
					}
				}
			}
		}
		return permutations
	}
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
