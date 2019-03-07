package main

import (
	"fmt"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) String() string {
	return strconv.Itoa(node.Val)
}

func printAll(node *ListNode) {
	head := node
	for head != nil {
		fmt.Print("--->", head)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	head := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					nil,
				},
			},
		},
	}
	printAll(head)
	printAll(swapPairs(head))
}

func swapPairs(head *ListNode) *ListNode {
	fakeHead := &ListNode{
		0,
		head,
	}

	header := fakeHead
	for header != nil {
		first := header.Next
		if first == nil {
			break
		}
		second := first.Next
		if second == nil {
			break
		}
		first.Next = second.Next
		second.Next = first
		header.Next = second
		header = second.Next
	}
	return fakeHead.Next
}

