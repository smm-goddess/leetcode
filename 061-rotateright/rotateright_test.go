package rotateright

import (
	"fmt"
	"testing"
)

func printAll(l *ListNode) {
	h := l
	for h != nil {
		fmt.Print("->", h)
		h = h.Next
	}
}

func equal(a, b *ListNode) bool {
	for a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
	return a == nil && b == nil
}

func TestRotateRight(t *testing.T) {
	l := &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4,
					&ListNode{5, nil}},
			},
		},
	}
	rotated := rotateRight(l, 2)
	answer := &ListNode{4,
		&ListNode{5,
			&ListNode{1,
				&ListNode{2,
					&ListNode{3, nil}},
			},
		},
	}
	if !equal(rotated, answer) {
		t.Error("Rotate fail")
	}

	l = &ListNode{0,
		&ListNode{1,
			&ListNode{2, nil}},
	}
	rotated = rotateRight(l, 4)
	answer = &ListNode{2,
		&ListNode{0,
			&ListNode{1, nil}},
	}
	if !equal(rotated, answer) {
		t.Error("Rotate 2 fail")
	}
	l = &ListNode{1, nil}
	if !equal(rotateRight(l, 1), &ListNode{1, nil}) {
		t.Error("Rotate 3 fail")
	}
	l := &ListNode{1, &ListNode{2, nil}}
	if !equal(rotateRight(l, 3), &ListNode{2, &ListNode{1, nil}}) {
		t.Error()
	}
}
