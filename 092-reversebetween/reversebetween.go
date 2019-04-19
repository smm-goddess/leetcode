package reversebetween

/*
 * @lc app=leetcode id=92 lang=golang
 *
 * [92] Reverse Linked List II
 *
 * https://leetcode.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (34.44%)
 * Total Accepted:    187.5K
 * Total Submissions: 542.8K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * Reverse a linked list from position m to n. Do it in one-pass.
 *
 * Note: 1 ≤ m ≤ n ≤ length of list.
 *
 * Example:
 *
 *
 * Input: 1->2->3->4->5->NULL, m = 2, n = 4
 * Output: 1->4->3->2->5->NULL
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
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	fakeHead := &ListNode{0, head}
	startAhead := fakeHead
	for m > 1 {
		startAhead = startAhead.Next
		m--
		n--
	}
	tail := startAhead
	for n >= m {
		tail = tail.Next
		n--
	}
	remains := tail.Next
	tail.Next = nil
	start := startAhead.Next
	startAhead.Next = nil
	reversedStart, reversedEnd := reverse(start)
	startAhead.Next = reversedStart
	reversedEnd.Next = remains
	head = fakeHead.Next
	fakeHead.Next = nil
	fakeHead = nil
	return head
}

func reverse(head *ListNode) (*ListNode, *ListNode) {
	var start, end, tmp *ListNode
	for head != nil {
		if end == nil {
			end = head
		}
		if start == nil {
			start = head
		} else {
			tmp = head
			head = head.Next
			tmp.Next = start
			start = tmp
		}
	}
	return start, end
}
