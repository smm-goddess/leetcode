package insertionsortlist

/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 *
 * https://leetcode.com/problems/insertion-sort-list/description/
 *
 * algorithms
 * Medium (36.76%)
 * Total Accepted:    146.8K
 * Total Submissions: 397K
 * Testcase Example:  '[4,2,1,3]'
 *
 * Sort a linked list using insertion sort.
 * 
 * 
 * 
 * 
 * 
 * A graphical example of insertion sort. The partial sorted list (black)
 * initially contains only the first element in the list.
 * With each iteration one element (red) is removed from the input data and
 * inserted in-place into the sorted list
 * 
 * 
 * 
 * 
 * 
 * Algorithm of Insertion Sort:
 * 
 * 
 * Insertion sort iterates, consuming one input element each repetition, and
 * growing a sorted output list.
 * At each iteration, insertion sort removes one element from the input data,
 * finds the location it belongs within the sorted list, and inserts it
 * there.
 * It repeats until no input elements remain.
 * 
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: 4->2->1->3
 * Output: 1->2->3->4
 * 
 * 
 * Example 2:
 * 
 * 
 * Input: -1->5->3->4->0
 * Output: -1->0->3->4->5
 * 
 * 
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else {
		r := head
		head = head.Next
		r.Next = nil
		for head != nil {
			node := head
			head = head.Next
			node.Next = nil
			r = insert(r, node)
		}
		return r
	}
}

func insert(head, node *ListNode) *ListNode {
	if head.Val > node.Val {
		node.Next = head
		return node
	}
	h := head
	for head != nil && head.Next != nil {
		if head.Val <= node.Val && head.Next.Val > node.Val {
			tail := head.Next
			head.Next = node
			node.Next = tail
			return h
		}
		head = head.Next
	}
	head.Next = node
	return h
}
