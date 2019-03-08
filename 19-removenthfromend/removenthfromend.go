package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printAll(node *ListNode) {
	for node != nil {
		fmt.Print("--->", node.Val)
		node = node.Next
	}
	fmt.Println()
}

func main() {
	list := &ListNode{
		1,
		&ListNode{
			2,
			&ListNode{
				3,
				&ListNode{
					4,
					&ListNode{
						5,
						nil,
					},
				},
			},
		},
	}
	printAll(list)
	list = removeNthFromEnd(list, 5)
	printAll(list)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	fakeHead := &ListNode{
		-1,
		head,
	}
	count := 1
	start := fakeHead
	for count < n {
		head = head.Next
		count ++
	}
	for head.Next != nil {
		start = start.Next
		head = head.Next
	}
	start.Next = start.Next.Next
	return fakeHead.Next
}
