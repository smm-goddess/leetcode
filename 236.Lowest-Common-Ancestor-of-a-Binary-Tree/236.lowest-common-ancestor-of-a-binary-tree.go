package _36_Lowest_Common_Ancestor_of_a_Binary_Tree

//Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
//
//According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
//
//Given the following binary tree:  root = [3,5,1,6,2,0,8,null,null,7,4]
//

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */

// func locateAncestor(root, p *TreeNode) (ancestors []*TreeNode) {
// 	if root == nil {
// 		return
// 	}
// 	if root.Val == p.Val {
// 		ancestors = []*TreeNode{root}
// 		return
// 	}
// 	if root.Left == nil && root.Right == nil {
// 		return
// 	}
// 	leftAncestors := locateAncestor(root.Left, p)
// 	if len(leftAncestors) > 0 {
// 		ancestors = append(append(ancestors, root), leftAncestors...)
// 	} else {
// 		rightAncestors := locateAncestor(root.Right, p)
// 		if len(rightAncestors) > 0 {
// 			ancestors = append(append(ancestors, root), rightAncestors...)
// 		}
// 	}
// 	return
// }

//func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
//	pAncestors := locateAncestor(root, p)
//	qAncestors := locateAncestor(root, q)
//	for i := len(pAncestors) - 1; i >= 0; i-- {
//		for j := len(qAncestors) - 1; j >= 0; j-- {
//			if pAncestors[i] == qAncestors[j] {
//				return pAncestors[i]
//			}
//		}
//	}
//	return nil
//}

func locateQAndP(root, p, q *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}
	leftNode, leftCount := locateQAndP(root.Left, p, q)
	if leftCount == 2 {
		return leftNode, 2
	}
	rightNode, rightCount := locateQAndP(root.Right, p, q)
	if rightCount == 2{
		return rightNode, 2
	} else if leftCount == 1 && rightCount == 1{
		return root, 2
	} else {
		if root.Val == p.Val || root.Val == q.Val {
			return root, 1+leftCount+rightCount
		} else {
			return nil, leftCount+rightCount
		}
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	node, _ := locateQAndP(root, q, p)
	return node
}
