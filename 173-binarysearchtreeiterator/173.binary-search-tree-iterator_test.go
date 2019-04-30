package binarysearchtreeiterator

import (
	"testing"
)

func TestBinarySearchTreeIterator(t *testing.T) {
	tree := &TreeNode{7,
		&TreeNode{3, nil, nil},
		&TreeNode{15,
			&TreeNode{9, nil, nil},
			&TreeNode{20, nil, nil}}}
	bstIterator := Constructor(tree)
	if bstIterator.Next() != 3 {
		t.Error("3 Error")
	} else {
		if bstIterator.Next() != 7 {
			t.Error("7 Error")
		} else {
			if !bstIterator.HasNext() {
				t.Error("7 Next Error")
			} else {
				if bstIterator.Next() != 9 {
					t.Error("9 Error")
				} else {
					if !bstIterator.HasNext() {
						t.Error("9 Next Error")
					} else {
						if bstIterator.Next() != 15 {
							t.Error("15 Error")
						} else {
							if !bstIterator.HasNext() {
								t.Error("15 Next Error")
							} else {
								if bstIterator.Next() != 20 {
									t.Error("20 Error")
								} else {
									if bstIterator.HasNext() {
										t.Error("20 Next Error")
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
