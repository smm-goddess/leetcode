//package reorderlist

/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 *
 * https://leetcode.com/problems/reorder-list/description/
 *
 * algorithms
 * Medium (30.25%)
 * Total Accepted:    150.5K
 * Total Submissions: 493.8K
 * Testcase Example:  '[1,2,3,4]'
 *
 * Given a singly linked list L: L0→L1→…→Ln-1→Ln,
 * reorder it to: L0→Ln→L1→Ln-1→L2→Ln-2→…
 *
 * You may not modify the values in the list's nodes, only nodes itself may be
 * changed.
 *
 * Example 1:
 *
 *
 * Given 1->2->3->4, reorder it to 1->4->2->3.
 *
 * Example 2:
 *
 *
 * Given 1->2->3->4->5, reorder it to 1->5->2->4->3.
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
func reorderList(head *ListNode) {
	if head == nil {
		return
	} else {
		t, f, reversed := head, head.Next, true
		t.Next = nil
		f = getHead(f, reversed)
		for f != nil {
			t.Next = f
			t = t.Next
			f = t.Next
			t.Next = nil
			reversed = !reversed
			f = getHead(f, reversed)
		}
	}
}

func getHead(head *ListNode, reversed bool) (first *ListNode) {
	if head == nil || head.Next == nil {
		first = head
	} else {
		if reversed {
			first = head
			for first.Next.Next != nil {
				first = first.Next
			}
			t := first
			first = first.Next
			t.Next = nil
			first.Next = head
		} else {
			first = head
		}
	}
	return
}
