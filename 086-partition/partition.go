package partition

/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 *
 * https://leetcode.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (36.73%)
 * Total Accepted:    158.8K
 * Total Submissions: 431.4K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * Given a linked list and a value x, partition it such that all nodes less
 * than x come before nodes greater than or equal to x.
 *
 * You should preserve the original relative order of the nodes in each of the
 * two partitions.
 *
 * Example:
 *
 *
 * Input: head = 1->4->3->2->5->2, x = 3
 * Output: 1->2->2->4->3->5
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
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	var result, smaller, stail, bigger, btail *ListNode
	for head != nil {
		if head.Val < x {
			if stail == nil {
				smaller, stail = head, head
			} else {
				stail.Next = head
				stail = head
			}
		} else {
			if btail == nil {
				bigger, btail = head, head
			} else {
				btail.Next = head
				btail = head
			}
		}
		head = head.Next
	}
	if stail == nil {
		btail.Next = nil
		result = bigger
	} else if btail == nil {
		stail.Next = nil
		result = smaller
	} else {
		btail.Next = nil
		stail.Next = bigger
		result = smaller
	}
	return result
}
