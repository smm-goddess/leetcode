package sortlist

import (
	"math"
)

/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 *
 * https://leetcode.com/problems/sort-list/description/
 *
 * algorithms
 * Medium (34.54%)
 * Total Accepted:    180.3K
 * Total Submissions: 516.5K
 * Testcase Example:  '[4,2,1,3]'
 *
 * Sort a linked list in O(n log n) time using constant space complexity.
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
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := findMid(head)
	t := sortList(mid.Next)
	mid.Next = nil
	h := sortList(head)
	return merge(h, t)
}

func findMid(head *ListNode) *ListNode {
	slow, quick := head, head.Next
	for quick != nil && quick.Next != nil {
		slow = slow.Next
		quick = quick.Next.Next
	}
	return slow
}

func merge(l1, l2 *ListNode) *ListNode {
	fakeHead := &ListNode{math.MinInt32, nil}
	tail := fakeHead
	for l1 != nil && l2 != nil {
		var node *ListNode
		if l1.Val < l2.Val {
			node = l1
			l1 = l1.Next
		} else {
			node = l2
			l2 = l2.Next
		}
		node.Next = nil
		tail.Next = node
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	} else if l2 != nil {
		tail.Next = l2
	}
	return fakeHead.Next
}

// quick sort
func quickSort(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pivot := head
	head = head.Next
	pivot.Next = nil
	h, t := partition(pivot, head)
	h = quickSort(h)
	t = quickSort(t)
	pivot.Next = t
	if h != nil {
		ht := h
		for ht.Next != nil {
			ht = ht.Next
		}
		ht.Next = pivot
		return h
	} else {
		return pivot
	}
}

func partition(pivot, list *ListNode) (head, tail *ListNode) {
	if list.Next == nil {
		if list.Val > pivot.Val {
			tail = list
		} else {
			head = list
		}
	} else {
		var tailTail, headTail *ListNode
		for list != nil {
			node := list
			list = list.Next
			node.Next = nil
			if node.Val > pivot.Val {
				if tailTail == nil {
					tail = node
				} else {
					tailTail.Next = node
				}
				tailTail = node
			} else {
				if headTail == nil {
					head = node
				} else {
					headTail.Next = node
				}
				headTail = node
			}
		}
	}
	return
}
