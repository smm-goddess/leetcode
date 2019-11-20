package insert_into_max_tree

import "testing"

type Case struct {
	Root   *TreeNode
	Val    int
	Output *TreeNode
}

var cases []*Case

func init() {
	cases = []*Case{
		{
			&TreeNode{4,
				&TreeNode{1, nil, nil},
				&TreeNode{3, &TreeNode{2, nil, nil}, nil}},
			5,
			&TreeNode{5,
				&TreeNode{4,
					&TreeNode{1, nil, nil},
					&TreeNode{3, &TreeNode{2, nil, nil}, nil}},
				nil},
		},
		{
			&TreeNode{5,
				&TreeNode{2, nil, &TreeNode{1, nil, nil}},
				&TreeNode{4, nil, nil}},
			3,
			&TreeNode{5,
				&TreeNode{2, nil, &TreeNode{1, nil, nil}},
				&TreeNode{4, nil, &TreeNode{3, nil, nil}}},
		},
		{
			&TreeNode{5,
				&TreeNode{2, nil, &TreeNode{1, nil, nil}},
				&TreeNode{3, nil, nil}},
			4,
			&TreeNode{5,
				&TreeNode{2, nil, &TreeNode{1, nil, nil}},
				&TreeNode{4, &TreeNode{3, nil, nil}, nil}},
		},
	}
}

func TestInsertIntoMaxTree(t *testing.T) {
	for _, cs := range cases {
		if !tree_equals(insertIntoMaxTree(cs.Root, cs.Val), cs.Output) {
			t.Error(cs.Val, " error")
		}
	}
}
