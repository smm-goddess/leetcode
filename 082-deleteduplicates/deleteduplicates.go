package deleteduplicates

/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 *
 * https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (32.52%)
 * Total Accepted:    175.4K
 * Total Submissions: 538.1K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * Given a sorted linked list, delete all nodes that have duplicate numbers,
 * leaving only distinct numbers from the original list.
 *
 * Example 1:
 *
 *
 * Input: 1->2->3->3->4->4->5
 * Output: 1->2->5
 *
 *
 * Example 2:
 *
 *
 * Input: 1->1->1->2->3
 * Output: 2->3
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

func deleteDuplicatesRec(head *ListNode) *ListNode {
	if head.Next != nil {
		if head.Val == head.Next.Val {
			for head.Next != nil && head.Next.Val == head.Val {
				head = head.Next
			}
			return deleteDuplicates(head.Next)
		} else {
			head.Next = deleteDuplicates(head.Next)
			return head
		}
	} else {
		return head
	}
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else {
		if head.Next != nil {
			var result, tail *ListNode
			head = findNextDifference(head)
			for head != nil {
				if result == nil {
					result = head
				} else {
					tail.Next = head
				}
				tail = head
				head = findNextDifference(head.Next)
				tail.Next = nil
			}
			return result
		} else {
			return head
		}
	}
}

func findNextDifference(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head.Next == nil || (head.Next.Val != head.Val) {
		return head
	} else {
		for head != nil && head.Next != nil && head.Next.Val == head.Val {
			for head != nil && head.Next != nil && head.Next.Val == head.Val {
				head = head.Next
			}
			head = head.Next
		}
		return head
	}
}
