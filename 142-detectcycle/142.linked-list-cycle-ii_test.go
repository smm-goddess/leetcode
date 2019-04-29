package detectcycle

import (
	"fmt"
	"testing"
)

func printNumbers(l *ListNode, n int) {
	head := l
	for n > 0 {
		fmt.Printf("-->%d", head.Val)
		head = head.Next
		n--
	}
	fmt.Println()
}

func TestDetectCycle(t *testing.T) {
	n1 := &ListNode{3, nil}
	n2 := &ListNode{2, nil}
	n3 := &ListNode{0, nil}
	n4 := &ListNode{-4, nil}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n2
	n := detectCycle(n1)
	if n != n2 {
		t.Error("First Linked List Fail")
	}

	n1 = &ListNode{1, nil}
	n2 = &ListNode{2, nil}
	n1.Next = n2
	n2.Next = n1
	printNumbers(n1, 4)
	n = detectCycle(n1)
	if n != n1 {
		t.Error("Second Fail")
	}
	printNumbers(n, 4)

	n1 = &ListNode{1, nil}
	n = detectCycle(n1)
	if n != nil {
		t.Error("Third Error")
	}

}
