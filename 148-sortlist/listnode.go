package sortlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	return fmt.Sprintf(" %d --> %s", l.Val, l.Next)
}
